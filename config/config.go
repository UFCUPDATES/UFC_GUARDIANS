package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var (
	LoggerId int64
	OwnerId  int64

	MongoUri   string
	StartImage string
	Token      string
	StartTime  time.Time
)

func init() {
	godotenv.Load()
	StartTime = time.Now()

	parseToInt64 := func(val string) int64 {
		parsed, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("Error parsing int64: %v", err))
		}
		return parsed
	}

	Token = Getenv[string]("TOKEN", "", nil)
	StartImage = Getenv[string](
		"START_IMG_URL",
		"https://files.catbox.moe/257jgn.jpg",
		nil,
	)
	LoggerId = Getenv("LOGGER_ID", "-1002286471112", parseToInt64)
	MongoUri = Getenv[string]("MONGO_DB_URI", "", nil)
	OwnerId = Getenv("OWNER_ID", "8195333063", parseToInt64)

	if Token == "" {
		log.Panic("TOKEN environment variable is empty")
	}
	if MongoUri == "" {
		log.Panic("MONGO_DB_URI environment variable is empty")
	}
}

func Getenv[T any](key, defaultValue string, convert func(string) T) T {
	value := defaultValue
	if envValue, exists := os.LookupEnv(key); exists {
		value = envValue
	}

	if convert != nil {
		return convert(value)
	}

	return any(value).(T)
}
