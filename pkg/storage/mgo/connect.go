package mgo

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConnectInfo struct {
	Client *mongo.Client
	Db     *mongo.Database
	Err    error
}

func (m *ConnectInfo) ConnectMgo() {

	auth := options.Credential{}
	auth.Username = os.Getenv("MGO_USER")
	auth.Password = os.Getenv("MGO_PASS")

	m.Client, m.Err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://172.16.20.12:27017").SetConnectTimeout(5*time.Second).SetAuth(auth))
}
