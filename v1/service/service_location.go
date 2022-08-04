package service

import (
	"github.com/SushiTuna/go-weather-api/v1/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type LocationDAO struct {
	Location *model.Location
	Client   *mongo.Client
	Context  *gin.Context
}

func (dao *LocationDAO) Add() (*model.Location, error) {
	return nil, nil
}
