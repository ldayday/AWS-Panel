package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yuzuki999/Aws-Panel/data"
)

func AddSecret(c *gin.Context) {
	username := SessionCheck(c)
	name := c.PostForm("name")
	id := c.PostForm("id")
	secret := c.PostForm("secret")
	if name == "" || id == "" || secret == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "信息填写不完整",
		})
	}
	if username == "" {
		return
	}
	addErr := data.AddSecret(username, name, id, secret)
	if addErr == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "添加成功",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  addErr.Error(),
		})
		return
	}
}

func ListSecret(c *gin.Context) {
	username := SessionCheck(c)
	if username == "" {
		return
	}
	secret, listErr := data.ListSecret(username)
	if listErr == nil {
		var tmp []map[string]string
		for _, v := range secret {
			tmp = append(tmp, map[string]string{
				"name":   v.Name,
				"id":     v.SecretId,
				"secret": v.Secret,
			})
		}
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"data": tmp,
		})
	}
}

func GetSecretInfo(c *gin.Context) {
	username := SessionCheck(c)
	if username == "" {
		return
	}
	name := c.PostForm("name")
	if name == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "信息填写不完整",
		})
	}
	secret, getErr := data.GetSecret(username, name)
	if getErr == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"data": secret,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  getErr.Error(),
		})
	}
}

func DelSecret(c *gin.Context) {
	username := SessionCheck(c)
	if username == "" {
		return
	}
	name := c.PostForm("name")
	if name == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "信息填写不完整",
		})
	}
	delErr := data.DelSecret(username, name)
	if delErr == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "删除成功",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  delErr.Error(),
		})
		return
	}
}
