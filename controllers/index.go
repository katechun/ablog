package controllers

import (
	"ablog/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Index(c *gin.Context){
	navlist := NaviList()
	articlelist := Listx()


	c.HTML(http.StatusOK,"index.html",gin.H{
		"navilist":navlist,
		"articlelist":articlelist,
		"gettag":GetTag(),
	})
}


func NaviList() []*models.Navis{

	getnavi := models.Navis{}
	res := getnavi.Get()
	return res
}

func Listx() []*models.Lists{
	getlist := models.Lists{}
	res := getlist.Get()
	return  res
}

func GetTag() string {
	getnavi := models.Navis{}
	res := getnavi.Get()
	for _,v:=range res {
		if strings.Index(v.Css,"layui-this") != -1{
			return v.Button

		}

	}
	return ""
}