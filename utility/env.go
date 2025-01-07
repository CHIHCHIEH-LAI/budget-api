package utility

import (
	"budget_manager/database"
	"os"

	"github.com/joho/godotenv"
)

func LoadDotenvIfExists(dotenv_filepath string) {
	if FileExists(dotenv_filepath) {
		godotenv.Load(dotenv_filepath)
	}
}

func GetDBConfig() database.Config {
	var DBConfig database.Config
	DBConfig.Host = os.Getenv("DB_HOST")
	DBConfig.Port = os.Getenv("DB_PORT")
	DBConfig.User = os.Getenv("DB_USER")
	DBConfig.Password = os.Getenv("DB_PASSWORD")
	DBConfig.Name = os.Getenv("DB_NAME")
	return DBConfig
}
