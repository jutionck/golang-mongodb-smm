package manager

import (
	"context"
	"enigmacamp.com/golang-with-mongodb/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type InfraManager interface {
	DbConn() *mongo.Database
}

type infraManager struct {
	db  *mongo.Database
	cfg config.Config
}

func (i *infraManager) initDb() {
	credential := options.Credential{
		Username: i.cfg.User,
		Password: i.cfg.Password,
	}
	mongoUri := fmt.Sprintf("mongodb://%s:%s", i.cfg.Host, i.cfg.Port)
	clientOptions := options.Client()
	clientOptions.ApplyURI(mongoUri).SetAuth(credential)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	i.db = client.Database(i.cfg.DbName)
}

func (i *infraManager) DbConn() *mongo.Database {
	return i.db
}

func NewInfraManager(config config.Config) InfraManager {
	infra := infraManager{cfg: config}
	infra.initDb()
	return &infra
}
