package services

import (
	"github.com/gin-gonic/gin"
)

type SystenLoad struct {
	Cpu    string `json:"cpu"`
	Memory string `json:"memory"`
	Disk   string `json:"disk"`
}

func GetAllStatus(c *gin.Context) {
	c.JSON(200, &SystenLoad{
		Cpu:    "foo",
		Memory: "bar",
		Disk:   "rab",
	})
}
