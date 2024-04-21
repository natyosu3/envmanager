package main

import (
	"github.com/gin-gonic/gin"
	"envmanager/pkg/engine"
	"envmanager/pkg/db/create"
)

func init() {
	create.CreateDefaultTable()
}

func main() {
	r := gin.New()
	r = engine.Engine(r)

	r.Run()
}