package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"sync"
)

type ServerStatus struct {
	Engine     *gin.Engine
	Err        error
	Status     string
	HandleFunc func(c *gin.Context)
}

var ServerContainer sync.Map

func Maker(port string, httpMethod, relativePath string) error {
	var err error
	defer func(err *error) {
		e := recover()
		// fmt.Println(e)
		temp := errors.New(fmt.Sprintf("%v", e))
		err = &temp
	}(&err)

	v, ok := ServerContainer.Load(port)
	if ok {

		v.(*ServerStatus).Engine.Handle(strings.ToTitle(httpMethod), relativePath, func(c *gin.Context) {
			c.JSON(200, Body())
		})
	} else {
		server := new(ServerStatus)
		server.Engine = gin.Default()
		go func(s *ServerStatus) {
			s.Err = s.Engine.Run(":" + port)
		}(server)
		// todo 根据数据库、sql、参数 来生成返回响应体
		server.Engine.Handle(strings.ToTitle(httpMethod), relativePath, func(c *gin.Context) {
			c.JSON(200, Body())
		})
		ServerContainer.Store(port, server)
	}
	return err
}

func Body() interface{} {
	m := make(map[string]interface{})
	m["string"] = "123"
	m["int"] = 10
	return m
}
