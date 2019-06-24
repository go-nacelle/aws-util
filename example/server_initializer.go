package main

import (
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/go-nacelle/httpbase"
	"github.com/go-nacelle/nacelle"
)

type ServerInitializer struct {
	Logger    nacelle.Logger            `service:"logger"`
	DynamoDB  dynamodbiface.DynamoDBAPI `service:"dynamodb"`
	tableName string
}

type Config struct {
	TableName string `env:"dynamodb_table_name" required:"true"`
}

func NewServerInitializer() httpbase.ServerInitializer {
	return &ServerInitializer{}
}

func (si *ServerInitializer) Init(config nacelle.Config, server *http.Server) error {
	serverConfig := &Config{}
	if err := config.Load(serverConfig); err != nil {
		return err
	}

	si.tableName = serverConfig.TableName
	server.Handler = http.HandlerFunc(si.handle)
	return nil
}

func (si *ServerInitializer) handle(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Path[1:]; id != "" {
		if r.Method == "GET" {
			si.handleGet(w, r, id)
			return
		}

		if r.Method == "POST" {
			si.handlePost(w, r, id)
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func (si *ServerInitializer) handleGet(w http.ResponseWriter, r *http.Request, id string) {
	out, err := si.DynamoDB.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(si.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"key": &dynamodb.AttributeValue{S: aws.String(id)},
		},
	})

	if err != nil {
		si.error(w, "Failed to get item (%s)", err)
		return
	}

	if out.Item == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	si.Logger.Debug("Retrieved key %s", id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(*out.Item["value"].S))
}

func (si *ServerInitializer) handlePost(w http.ResponseWriter, r *http.Request, id string) {
	defer r.Body.Close()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		si.error(w, "Failed to read request body (%s)", err)
		return
	}

	_, err = si.DynamoDB.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(si.tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"key":   &dynamodb.AttributeValue{S: aws.String(id)},
			"value": &dynamodb.AttributeValue{S: aws.String(string(content))},
		},
	})

	if err != nil {
		si.error(w, "Failed to put item (%s)", err)
		return
	}

	si.Logger.Debug("Set key %s", id)
	w.WriteHeader(http.StatusOK)
}

func (si *ServerInitializer) error(w http.ResponseWriter, format string, err error) {
	si.Logger.Error(format, err.Error())
	w.WriteHeader(http.StatusInternalServerError)
}
