package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	//1创建路由
	//默认使用了2个中间件 Logger() Recovery()
	r := gin.Default()
	//form绑定
	r.POST("/test",test)
	//多种响应方式 json
	r.GET("json", func(c *gin.Context) {
		c.JSON(200,gin.H{"msg":"someMsg","status":200})
	})
	//结构体
	r.GET("struct", func(c *gin.Context) {
		 var msg struct{
		 	Name string
		 	Age int
		 }
		 msg.Name ="root"
		 msg.Age = 11
		 c.JSON(200,msg)
	})
	//xml 用的不多 会下载
	r.GET("xml", func(c *gin.Context) {
		c.XML(200,gin.H{"meg":"xml"})
	})
	//yaml 会下载
	r.GET("yaml", func(c *gin.Context) {
		c.YAML(200,gin.H{"msg":"yaml"})
	})
	//protobuf格式 谷歌开发的高效读存取工具 会下载
	r.GET("protobuf", func(c *gin.Context) {
		resp:= []int64{int64(1),int64(2)}
		label := "label test"//定义数据
		data:=&protoexample.Test{//格式数据
			Label: &label,
			Reps: resp,
		}
		c.ProtoBuf(200,data)
	})
	//html渲染 本质是字符串替换
	r.LoadHTMLGlob("../../html/*")//加载文件模板，正则匹配路径
	r.GET("index", func(c *gin.Context) {
		//r.LoadHTMLFiles("html/绝对路径")
		//模板那边是 .title
		c.HTML(200,"form.html",gin.H{"title":"我是标题"})
	})
	//3监听端口
	r.Run(":8000")

}
type Form struct {
	//要可以访问且 form绑定设置了,就是表单中的name值
	User string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pwd string `form:"pwd"`
}
func test (c *gin.Context ){
	var form Form
	if err :=c.Bind(&form);err !=err{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	if form.User !="root" || form.Pwd != "123"{
		c.JSON(http.StatusBadRequest,gin.H{"status":"304"})
		c.JSON(http.StatusBadRequest,fmt.Sprintf("user:%v pwd:%v",form.User,form.Pwd))
		return
	}
	c.JSON(http.StatusOK,gin.H{"status":"200"})
}