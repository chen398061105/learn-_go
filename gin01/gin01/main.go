package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//1创建路由
	//默认使用了2个中间件 Logger() Recovery()
	r := gin.Default()
	//2封装了request和respone
	r.GET("/test/:name/*action", func(context *gin.Context) {
		//传api参数 通常冒号是写死  *可以任意，可以接收后面所有的query
		name := context.Param("name")
		action := context.Param("action")
		context.String(http.StatusOK,name+"  "+action)
	})
	//url参数
	r.GET("/test2", func(context *gin.Context) {
		//context.String(http.StatusOK,"hello world2s!")
		//传api参数 通常冒号是写死  *可以任意，可以接收后面所有的query
		name:= context.DefaultQuery("name","default")
	    //http://localhost:8000/test2?name=xiao
		context.String(http.StatusOK,name)
	})
	//表单
	r.POST("/form", func(context *gin.Context) {
		//设置表单默认值
		type1:=context.DefaultPostForm("type","alert")
		username:= context.PostForm("username")
		password := context.PostForm("pwd")
		//复选框
		hobbys := context.PostFormArray("hobby")
		context.String(http.StatusOK,
			fmt.Sprintf("type is %s\n,username is%s\n,password is %s\n ,hobbys is %v\n",
				type1,username,password,hobbys))
	})
	//单文件上传
	r.POST("/upload", func(context *gin.Context) {
		file,_:=context.FormFile("file")
		//log.Println(file.Filename)
		//默认传到根目录
		context.SaveUploadedFile(file,file.Filename)
		context.String(200,fmt.Sprintf("%s is upload",file.Filename))
	})
	//多文件上传
	r.POST("/uploads", func(context *gin.Context) {
		//控制上传大小 8mb 默认32
		r.MaxMultipartMemory = 8<<20
		//获取所有文件
		form,err:= context.MultipartForm()
		if err != nil{
			context.String(400,fmt.Sprintf("get file err:%v",err))
			return
		}
		//获取所有上传文件
		files:= form.File["file"]
		for _,file :=range files{
			//保存在根目录并通知
			if err:=context.SaveUploadedFile(file,file.Filename);err !=nil{
				context.String(400,fmt.Sprintf("upload file err:%v",err))
				return
			}
		}
		context.String(200,fmt.Sprintf("%v:upload sucess",len(files)))
	})
	//3监听端口
	r.Run(":8000")

}
func getting (c *gin.Context ){
	c.String(http.StatusOK,"hello world2!")
}