package service

import (
	"math/rand"
	"time"

	"github.com/suhas-totality/gomongodemo/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceRegistry struct {
	LoggerClient  *logger.Logger
	MongoDBClient *mongo.Client
}

func NewServiceRegistry() *ServiceRegistry {
	rand.Seed(time.Now().Unix())
	return &ServiceRegistry{}
}
