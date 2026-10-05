package config

type Config struct {
	Database          string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"
