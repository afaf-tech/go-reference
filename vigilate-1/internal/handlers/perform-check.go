package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
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
	// if the host service status has changed, broadcast to all clients
	// if hostServiceStatusChanged {
	// 	data := make(map[string]string)
	// 	data["message"] = fmt.Sprintf("host service %s on %s has changed to %s", hs.Service.ServiceName, h.HostName, newStatus)
	// 	repo.broadcastMessage("public-channel", "host-service-status-changed", data)
	// 	// if appropriate, send email or SMS message

	// }
	// update hsot service record in db with status (if changed) and last check
	hs.Status = newStatus
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

	// update host service in the database with status (if changed) and last check
	hs.Status = newStatus
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

	}

	// TODO - broadcast to clients if appropriate
	return newStatus, msg
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

func (repo *DBRepo) broadcastMessage(channel, messageType string, data map[string]string) {
	err := app.WsClient.Trigger(channel, messageType, data)
	if err != nil {
		log.Println(err)
	}
}