package main

import (
	"demo_api/conf"
	"fmt"
	"net/http"
	"time"

	"os"

	"io"

	"github.com/gin-gonic/gin"
)

var picbasepath string

func main() {

	myconf := new(conf.Config)
	myconf.InitConfig("config.ini")
	picbasepath = myconf.Read("path", "picbasepath")
	router := gin.Default()

	router.GET("/hh", func(c *gin.Context) {
		c.JSON(200, gin.H{"key": "value"})
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.POST("/upload", UpLoad)
	router.POST("/uploadmul", UpLoadMul)
	router.Run(":8080")

}

// UpLoad 上传文件接口
func UpLoad(c *gin.Context) {
	file, err1 := c.FormFile("file")
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "上传错误"})
		return
	}

	upfileh, err2 := file.Open()
	defer upfileh.Close()
	if err2 != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "保存错误"})
		return
	}

	// os.Chdir(picbasepath)

	_, err := os.Stat(picbasepath)
	if err != nil {
		os.MkdirAll(picbasepath, 0777)
	}

	createh, err3 := os.Create(picbasepath + "/" + file.Filename)
	defer createh.Close()
	if err3 != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "保存错误"})
		return
	}
	ret, err4 := io.Copy(createh, upfileh)
	if err4 != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "保存错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "保存成功", "filelen": ret})
	return

}

//UpLoadMul  上传多个文件
func UpLoadMul(c *gin.Context) {

	c.Request.ParseMultipartForm(32 << 20) //在使用r.MultipartForm前必须先调用ParseMultipartForm方法，参数为最大缓存
	if c.Request.MultipartForm != nil && c.Request.MultipartForm.File != nil {
		files := c.Request.MultipartForm.File["files"] //获取所有上传文件信息

		for idx, file := range files {
			upFileName := file.Filename
			upHandler, err := file.Open()
			defer upHandler.Close()
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"success": false, "message": "上传错误"})
				return
			}

			saveHanler, err := os.Create(upFileName)
			defer saveHanler.Close()
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"success": false, "message": "保存错误"})
				return
			}

			ret, err := io.Copy(saveHanler, upHandler)

			if err != nil || ret == 0 {
				c.JSON(200, gin.H{"success": false, "message": "保存错误"})
				return
			}
			fileStat, _ := saveHanler.Stat()
			fmt.Printf("%s  NO.: %d  Size: %d KB  Name：%s\n", time.Now().Format("2006-01-02 15:04:05"), idx, fileStat.Size()/1024, upFileName)
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "上传成功"})
		return
	}

}
