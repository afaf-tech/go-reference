package simple

type Configuration struct {
	Name string `json:"name"`
}

type Application struct {
	*Configuration
}

func NewApplication() *Application {
	return &Application{
		Configuration: &Configuration{
			Name: "Afaf-tech, and Afaf-sitek",
		},
	}
}
