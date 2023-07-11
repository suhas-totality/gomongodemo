package env

import (
	"os"
	"strconv"
)

func loadEnvPlain() (*EnvPlain, error) {
	var errSlice []error

	serverLogLevel, errServerLogLevel := strconv.Atoi(os.Getenv("SERVER_LOG_LEVEL"))
	serverGRPCPort, errServerGRPCPort := strconv.Atoi(os.Getenv("SERVER_GRPC_PORT"))
	serverTimeoutSeconds, errServerTimeoutSeconds := strconv.Atoi(os.Getenv("SERVER_TIMEOUT_SECONDS"))
	errSlice = append(errSlice, errServerLogLevel, errServerGRPCPort, errServerTimeoutSeconds)

	mongodbURL := os.Getenv("MONGODB_URL")
	mongodbUserName := os.Getenv("MONGODB_USERNAME")
	mongodbPassword := os.Getenv("MONGODB_PASSWORD")
	mongodbDatabase := os.Getenv("MONGODB_DATABASE")
	mongodbCollection := os.Getenv("MONGODB_COLLECTION")

	if err := checkEnvErrors(errSlice); err != nil {
		return nil, err
	}

	return &EnvPlain{
		SERVER_LOG_LEVEL:       serverLogLevel,
		SERVER_GRPC_PORT:       serverGRPCPort,
		SERVER_TIMEOUT_SECONDS: serverTimeoutSeconds,

		MONGODB_URL:        mongodbURL,
		MONGODB_USERNAME:   mongodbUserName,
		MONGODB_PASSWORD:   mongodbPassword,
		MONGODB_DATABASE:   mongodbDatabase,
		MONGODB_COLLECTION: mongodbCollection,
	}, nil
}
