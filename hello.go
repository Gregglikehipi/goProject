package main

import (
	"net/http"

	"goProject/postgresql"

	"github.com/gin-gonic/gin"

	"strconv"

	"reflect"
)

type post struct {
	Id     int    `json:"id"`
	Data   string `json:"data"`
	UserId int    `json:"userId"`
	Likes  int    `json:"likes"`
}

func getPost(c *gin.Context) {
	id := c.Query("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		// ... handle error
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, postgresql.GetPost(i))
}

func getPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, postgresql.GetPosts())
}

func postHuman(c *gin.Context) {
	name := c.Query("name")
	pass := c.Query("pass")
	postgresql.PostHuman(name, pass)
	c.IndentedJSON(http.StatusOK, gin.H{"ans": "Yes"})
}

func insertPost(c *gin.Context) {
	data := c.Query("data")
	name := c.Query("name")
	pass := c.Query("pass")
	human := postgresql.GetHumanName(name, pass)
	if reflect.DeepEqual(human, post{}) {
		c.IndentedJSON(http.StatusOK, gin.H{"ans": "No"})
	} else {
		if human.Name == name && human.Pass == pass {
			postgresql.InsertPost(data, human.Id)
			c.IndentedJSON(http.StatusOK, gin.H{"ans": "Yes"})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"ans": "No"})
		}
	}
}

func putPost(c *gin.Context) {
	id := c.Query("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		// ... handle error
		panic(err)
	}
	post := postgresql.GetPost(i)
	postgresql.UpdatePost(i, post.Likes+1)
	c.IndentedJSON(http.StatusOK, gin.H{"ans": "Yes"})
}

func deletePost(c *gin.Context) {
	id := c.Query("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		// ... handle error
		panic(err)
	}
	name := c.Query("name")
	pass := c.Query("pass")
	human := postgresql.GetHumanName(name, pass)
	if reflect.DeepEqual(human, post{}) {
		c.IndentedJSON(http.StatusOK, gin.H{"ans": "No"})
	} else {
		post := postgresql.GetPost(i)
		if human.Id == post.UserId {
			postgresql.DeletePost(i)
			c.IndentedJSON(http.StatusOK, gin.H{"ans": "Yes"})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"ans": "No"})
		}
	}
}

func main() {
	//postgresql.GetPosts()
	//fmt.Println(helper.Factorial(5))
	//fmt.Println("hi")

	router := gin.Default()
	router.GET("/post", getPost)
	router.GET("/posts", getPosts)
	router.POST("/post", insertPost)
	router.PUT("/post", putPost)
	router.DELETE("/post", deletePost)
	router.POST("/user", postHuman)

	router.Run("localhost:8080")
}
