package config

//Config struct
type Config struct {

	DatabaseURL     string `envconfig:"MONGO_CONNECTION_STRING" default:""`
	Database        string `envconfig:"MONGO_DATABASE_NAME" default:"test_restapi"`
}