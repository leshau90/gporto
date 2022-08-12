package main
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)


func main(){
	r:=gin.Default()
	i :=0
	r.GET("/ping",func(c *gin.Context){
		fmt.Printf("serving the %d th /n",i)
		i++
		c.JSON(http.StatusOK,
		gin.H{"message":"hey there","meaning":"are you okay"})
	})
	r.Run(":3000")
}
