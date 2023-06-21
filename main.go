package main

import (
	"net/http"
	"fmt"
	//"errors"
	"math/rand"
	"time"
	"github.com/gin-gonic/gin"
)

//change this to a database later
type Mes struct {
    ID	int	`json:"id"`
}
var GID []int
func generateRandomInt() int {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(9000) + 1000
	return randomInt
}
func getUniqueGameID() int{
	//check if serves are full if so return -1
	for true {
		var val int = generateRandomInt()
		if !containsValue(GID, val){
			return val
		}
	}
	return -1
}
func containsValue(slice []int, value int) bool {
    for _, element := range slice {
        if element == value {
            return true
        }
    }
    return false
}
func getGID(c *gin.Context){
	val := Mes{ID:	getUniqueGameID()}
	//if val != -1{
	c.JSON(http.StatusOK,val)
	//}
	
}

func main(){
	fmt.Println(getUniqueGameID())
	router := gin.Default()
	router.GET("/GID",getGID)
	router.Run("localhost:8080")
}