package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/tsawler/vigilate/internal/certificateutils"
	"github.com/tsawler/vigilate/internal/channeldata"
	"github.com/tsawler/vigilate/internal/helpers"
	"github.com/tsawler/vigilate/internal/models"
)

const (
	HTTP           = 1
	HTTPS          = 2
	SSLCertificate = 3
)

type jsonResp struct {
	OK            bool      `json:"ok"`
	Message       string    `json:"message"`
	ServiceID     int       `json:"service_id"`
	HostServiceID int       `json:"host_service_id"`
	HostID        int       `json:"host_id"`
	OldStatus     string    `json:"old_status"`
	NewStatus     string    `json:"new_status"`
	LastCheck     time.Time `json:"last_check"`
}

// ScheduledChek performs a scheduled check on a host service by id
func (repo *DBRepo) ScheduledCheck(hostServiceID int) {
	log.Println("******** Running check for ", hostServiceID)

	hs, err := repo.DB.GetHostServiceByID(hostServiceID)
	if err != nil {
		log.Println(err)
		return
	}
	h, err := repo.DB.GetHostById(hs.HostID)
	if err != nil {
		log.Println(err)
		return
	}
	// test the service
	newStatus, msg := repo.testServiceForHost(h, hs)

	if newStatus != hs.Status {
		repo.updateHostServiceStatusCount(h, hs, newStatus, msg)
	}

}

func (repo *DBRepo) updateHostServiceStatusCount(h models.Host, hs models.HostService, newStatus, msg string) {
	// update hsot service record in db with status (if changed) and last check
	hs.Status = newStatus
	hs.LastMessage = msg
	hs.LastCheck = time.Now()
	err := repo.DB.UpdateHostService(hs)
	if err != nil {
		log.Println(err)
		return
	}

	pending, healthy, warning, problem, err := repo.DB.GetAllServiceStatusCounts()
	if err != nil {
		log.Println(err)
		return
	}

	data := make(map[string]string)
	data["healthy_count"] = strconv.Itoa(healthy)
	data["pending_count"] = strconv.Itoa(pending)
	data["problem_count"] = strconv.Itoa(problem)
	data["warning_count"] = strconv.Itoa(warning)
	repo.broadcastMessage("public-channel", "host-service-count-changed", data)

	log.Println("New status is", newStatus)
	log.Println("msg is", msg)

}

func (repo *DBRepo) TestCheck(w http.ResponseWriter, r *http.Request) {
	hostServiceID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	oldStatus := chi.URLParam(r, "oldstatus")
	okay := true

	log.Println(hostServiceID, oldStatus)

	// get host service
	hs, err := repo.DB.GetHostServiceByID(hostServiceID)
	if err != nil {
		log.Println(err)
		okay = false
	}

	// get host?
	h, err := repo.DB.GetHostById(hs.HostID)
	if err != nil {
		log.Println(err)
		okay = false
	}
	// test the service
	newStatus, msg := repo.testServiceForHost(h, hs)

	// save event
	event := models.Event{
		EventType:     newStatus,
		HostServiceID: hs.ID,
		HostID:        h.ID,
		ServiceName:   hs.Service.ServiceName,
		HostName:      h.HostName,
		Message:       msg,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	err = repo.DB.InsertEvent(event)
	if err != nil {
		log.Println(err)
	}

	// broadcast service status chanded event
	if newStatus != hs.Status {
		repo.pushStatusChangedEvent(h, hs, newStatus)
	}

	// update host service in the database with status (if changed) and last check
	hs.Status = newStatus
	hs.LastMessage = msg
	hs.UpdatedAt = time.Now()

	err = repo.DB.UpdateHostService(hs)
	if err != nil {
		log.Println(err)
		okay = false
	}

	// broadcast service status changed event
	log.Println(newStatus, msg)

	var resp jsonResp
	// create json
	if okay {
		resp = jsonResp{
			OK:            true,
			Message:       msg,
			ServiceID:     hs.ServiceID,
			HostServiceID: hs.ID,
			HostID:        h.ID,
			OldStatus:     oldStatus,
			NewStatus:     newStatus,
			LastCheck:     time.Now(),
		}
	} else {
		resp.OK = false
		resp.Message = "something went wrong"
	}

	// send json to client

	out, _ := json.MarshalIndent(resp, "", "   ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (repo *DBRepo) testServiceForHost(h models.Host, hs models.HostService) (string, string) {
	var msg, newStatus string

	switch hs.ServiceID {
	case HTTP:
		msg, newStatus = testHttpForHost(h.URL)
		break
	case HTTPS:
		msg, newStatus = testHttpsForHost(h.URL)
		break
	case SSLCertificate:
		msg, newStatus = testSSLForHost(h.URL)
		break
	}

	// broadcast to clients if appropriate
	if hs.Status != newStatus {
		repo.pushStatusChangedEvent(h, hs, newStatus)

		// save event
		event := models.Event{
			EventType:     newStatus,
			HostServiceID: hs.ID,
			HostID:        h.ID,
			ServiceName:   hs.Service.ServiceName,
			HostName:      h.HostName,
			Message:       msg,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}
		err := repo.DB.InsertEvent(event)
		if err != nil {
			log.Println(err)
		}
		// - send email if appropriate.
		if repo.App.PreferenceMap["notify_via_email"] == "1" {
			if hs.Status != "pending" {
				mm := channeldata.MailData{
					ToName:    repo.App.PreferenceMap["notify_name"],
					ToAddress: repo.App.PreferenceMap["notify_email"],
				}

				if newStatus == "healthy" {
					mm.Subject = fmt.Sprintf("HEALTHY %s on %s", hs.Service.ServiceName, hs.HostName)
					mm.Content = template.HTML(fmt.Sprintf(`<p>Service on %s on %s reported healthy status</p>
					<p><strong>Message received : %s </strong></p>`, hs.Service.ServiceName, hs.HostName, msg))
				} else if newStatus == "problem" {
					mm.Subject = fmt.Sprintf("PROBLEM %s on %s", hs.Service.ServiceName, hs.HostName)
					mm.Content = template.HTML(fmt.Sprintf(`<p>Service on %s on %s reported problem status</p>
						<p><strong>Message received : %s </strong></p>`, hs.Service.ServiceName, hs.HostName, msg))

				} else if newStatus == "warning" {
					mm.Subject = fmt.Sprintf("WARNING %s on %s", hs.Service.ServiceName, hs.HostName)

				}

				helpers.SendEmail(mm)
			}
		}
		// send sms if appropriate
		// the feature is temporarily turned off.
		// if repo.App.PreferenceMap["notify_via_sms"] == "1" {
		// 	to := repo.App.PreferenceMap["sms_notify_number"]
		// 	smsMessage := ""

		// 	if newStatus == "healthy" {
		// 		smsMessage = fmt.Sprintf("Service %s on %s is healthy", hs.Service.ServiceName, hs.HostName)
		// 	} else if newStatus == "problem" {
		// 		smsMessage = fmt.Sprintf("Service %s on %s reports a problem: %s", hs.Service.ServiceName, hs.HostName, msg)
		// 	} else if newStatus == "warning" {
		// 		smsMessage = fmt.Sprintf("Service %s on %s reports a warning: %s", hs.Service.ServiceName, hs.HostName, msg)
		// 	}

		// 	err := sms.SendTextTwilio(to, smsMessage, repo.App)
		// 	if err != nil {
		// 		log.Println("Error sending sms in peform-checks.go", err)
		// 	}
		// }
	}

	repo.pushScheduleChangeEvent(hs, newStatus)

	return newStatus, msg
}

func (repo *DBRepo) pushStatusChangedEvent(h models.Host, hs models.HostService, newStatus string) {
	data := make(map[string]string)
	data["host_id"] = strconv.Itoa(hs.HostID)
	data["host_service_id"] = strconv.Itoa(hs.ID)
	data["host_name"] = h.HostName
	data["service_name"] = hs.Service.ServiceName
	data["icon"] = hs.Service.Icon
	data["status"] = newStatus
	data["message"] = fmt.Sprintf("%s on %s reports %s", hs.Service.ServiceName, h.HostName, newStatus)
	data["last_check"] = time.Now().Format("2006-01-02 3:04:05 PM")

	repo.broadcastMessage("public-channel", "host-service-status-changed", data)
}

func (repo *DBRepo) pushScheduleChangeEvent(hs models.HostService, newStatus string) {
	// broadcast schedule-changed-event
	yearOne := time.Date(0001, 1, 1, 0, 0, 0, 1, time.UTC)
	data := make(map[string]string)
	data["host_service_id"] = strconv.Itoa(hs.ID)
	data["service_id"] = strconv.Itoa(hs.ServiceID)
	data["host_id"] = strconv.Itoa(hs.HostID)

	if app.Scheduler.Entry(repo.App.MonitorMap[hs.ID]).Next.After(yearOne) {
		data["next_run"] = repo.App.Scheduler.Entry(repo.App.MonitorMap[hs.ID]).Next.Format("2006-01-02 3:04:05 PM")
	} else {
		data["next_run"] = "Pending..."
	}

	data["last_run"] = time.Now().Format("2006-01-02 3:04:05 PM")
	data["host"] = hs.HostName
	data["service"] = hs.Service.ServiceName
	data["schedule"] = fmt.Sprintf("@every %d%s", hs.ScheduleNumber, hs.ScheduleUnit)
	data["status"] = newStatus
	data["icon"] = hs.Service.Icon

	repo.broadcastMessage("public-channel", "schedule-changed-event", data)
}

func testHttpForHost(url string) (string, string) {
	if strings.HasSuffix(url, "/") {
		url = strings.TrimSuffix(url, "/")
	}

	url = strings.Replace(url, "https://", "http://", -1)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("%s = %s", url, "error connecting"), "problem"
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("%s = %s", url, resp.Status), "problem"
	}
	return fmt.Sprintf("%s = %s", url, resp.Status), "healthy"
}

//  test HTTPS service
func testHttpsForHost(url string) (string, string) {
	if strings.HasSuffix(url, "/") {
		url = strings.TrimSuffix(url, "/")
	}

	url = strings.Replace(url, "http://", "https://", -1)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("%s = %s", url, "error connecting"), "problem"
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("%s = %s", url, resp.Status), "problem"
	}
	return fmt.Sprintf("%s = %s", url, resp.Status), "healthy"
}

// scanHost gets cert details from an internet host
func scanHost(hostname string, certDetailsChannel chan certificateutils.CertificateDetails, errorsChannel chan error) {

	res, err := certificateutils.GetCertificateDetails(hostname, 10)
	if err != nil {
		errorsChannel <- err
	} else {
		certDetailsChannel <- res
	}
}

// testSSLForHost tests an ssl certificate for a host
func testSSLForHost(url string) (string, string) {
	if strings.HasPrefix(url, "https://") {
		url = strings.Replace(url, "https://", "", -1)
	}
	if strings.HasPrefix(url, "http://") {
		url = strings.Replace(url, "http://", "", -1)
	}
	var certDetailsChannel chan certificateutils.CertificateDetails
	var errorsChannel chan error
	certDetailsChannel = make(chan certificateutils.CertificateDetails, 1)
	errorsChannel = make(chan error, 1)

	var msg string
	var newStatus string

	scanHost(url, certDetailsChannel, errorsChannel)

	for i, certDetailsInQueue := 0, len(certDetailsChannel); i < certDetailsInQueue; i++ {
		certDetails := <-certDetailsChannel
		certificateutils.CheckExpirationStatus(&certDetails, 30)

		if certDetails.Expired {
			// cert expired
			msg = certDetails.Hostname + " has expired!"

		} else if certDetails.ExpiringSoon {
			// cert expiring sono
			if certDetails.DaysUntilExpiration < 7 {
				msg = certDetails.Hostname + " expiring in " + strconv.Itoa(certDetails.DaysUntilExpiration) + " days"
				newStatus = "problem"
			} else {
				msg = certDetails.Hostname + " expiring in " + strconv.Itoa(certDetails.DaysUntilExpiration) + " days"
				newStatus = "warning"
			}
		} else {
			// cert okay
			msg = certDetails.Hostname + " expiring in " + strconv.Itoa(certDetails.DaysUntilExpiration) + " days"
			newStatus = "healthy"
		}
	}

	if len(errorsChannel) > 0 {
		fmt.Printf("There were %d error(s):\n", len(errorsChannel))
		for i, errorsInChannel := 0, len(errorsChannel); i < errorsInChannel; i++ {
			msg = fmt.Sprintf("%s\n", <-errorsChannel)
		}
		fmt.Printf("\n")
		newStatus = "problem"
	}

	return msg, newStatus
}
func (repo *DBRepo) broadcastMessage(channel, messageType string, data map[string]string) {
	err := app.WsClient.Trigger(channel, messageType, data)
	if err != nil {
		log.Println(err)
	}
}
