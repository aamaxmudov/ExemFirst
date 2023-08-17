package controller

import (
	"exemFirstM/json/jsonDB/config"
	"exemFirstM/json/jsonDB/storage"
)

type Controller struct {
	Cfg  *config.Config
	Strg storage.StorageI
}

func NewController(cfg *config.Config, storage storage.StorageI) *Controller {
	return &Controller{
		Cfg:  cfg,
		Strg: storage,
	}
}
