package env

type EnvPlain struct {
	SERVER_LOG_LEVEL       int
	SERVER_GRPC_PORT       int
	SERVER_TIMEOUT_SECONDS int

	MONGODB_URL        string
	MONGODB_USERNAME   string
	MONGODB_PASSWORD   string
	MONGODB_DATABASE   string
	MONGODB_COLLECTION string
}
