package config

type Config struct {
	App struct {
		Name string
		Env  string
		Port string
	}

	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}

	Queue struct {
		URL string
	}

	Security struct {
		APIKeySecret string
	}
}
