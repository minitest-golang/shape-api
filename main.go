package main

import (
	"minitest/api"
	"minitest/db"
	"minitest/utils"
	"os"
	"sync"
	"time"

	_ "minitest/docs" // Need this for swagger

	"github.com/joho/godotenv"
)

func main() {
	logFile := utils.LoggerInit()
	if logFile != nil {
		defer logFile.Close()
	}

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		godotenv.Load("./.env")
		os.Setenv("DB_HOST", "localhost")
	}

	var wg sync.WaitGroup
	db.DB = db.NewDbDriver("postgresql")
	wg.Add(1)
	go func() {
		for {
			err := db.DB.DbInit()
			if err == nil {
				break
			}
			time.Sleep(time.Second * 5)
			utils.ErrorLog("Failed to connecto to DB, try to connect again!")
		}
		wg.Done()
	}()
	wg.Wait()

	defer db.DB.DbClose()

	// Start REST server
	api.CreateRestApis()
}
