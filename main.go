package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"time"
	"datingapp/config"
	"datingapp/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"path/filepath"
	limit "github.com/yangxikun/gin-limit-by-key"
	middleware "datingapp/middleware"
	"golang.org/x/time/rate"
)

func goDotEnvVariable(key string) string {

  	// load .env file
	path := os.Getenv("GIN_ENV")
	folder := ""
	if path == "production" {
  		folder = "/var/www/golang/datingapp/"
  	}
    //fmt.Println(filepath.Join(path,folder, ".env"))

    err := godotenv.Load(filepath.Join(folder, ".env"))

  	if err != nil {
    	log.Fatalf("Error loading .env file")
  	}

  return os.Getenv(key)
}

func emptySuccessResponse(c *gin.Context) {
	time.Sleep(200 * time.Microsecond)
	//c.JSON(http.StatusOK, "")
	c.Next()
}

func timeoutResponse(c *gin.Context) {
	result := gin.H{
		"message": "Request timeout, please try again",
		"status":  http.StatusRequestTimeout,
		"data":   nil,
	}
	c.JSON(http.StatusRequestTimeout, result)
	c.Abort()
	return
}

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}
	path := os.Getenv("GIN_ENV")
	folder := ""
	if path == "production" {
  		folder = "/var/www/golang/datingapp/"
  	}
    //fmt.Println(filepath.Join(path,folder, ".env"))

    err := godotenv.Load(filepath.Join(folder, ".env"))
  	if err != nil {
    	fmt.Println("Failed")
  	}else{
  		fmt.Println("OK")
  	}

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1","35.197.139.3"})

	router.Use(limit.NewRateLimiter(func(c *gin.Context) string {
		return c.ClientIP() // limit rate by client ip
	}, func(c *gin.Context) (*rate.Limiter, time.Duration) {
		return rate.NewLimiter(rate.Every(5*time.Second), 60), 1*time.Minute 
	}, func(c *gin.Context) {
		result := gin.H{
			"message": "Too many request, please try again in 30-60 seconds",
			"status":  http.StatusTooManyRequests,
			"data":   nil,
		}
		c.JSON(http.StatusTooManyRequests, result)
		c.AbortWithStatus(429)
		return
	}))

	v1 := router.Group("/api/v1")
	{
		v1.GET("/test-connection", func(c *gin.Context) {
			expirationTime := time.Now().Add(time.Minute * 15)
			result := gin.H{
				"responseCode": 200,
				"responseMsg": "Ok,Dating Connecting looks good!",
				"times": expirationTime,
			}
			c.JSON(http.StatusOK, result)
			c.AbortWithStatus(200)
			return
		})
		v1.POST("/register", inDB.RegisterService)
		v1.POST("/login", inDB.LoginService)
	}

	secured := router.Group("/api/v1/auth").Use(middleware.Auth()) 
	{
		secured.POST("/like", inDB.LikeService)
		secured.POST("/dislike", inDB.DislikeService)
		secured.POST("/match", inDB.MatchService)
		secured.POST("/profil/list", inDB.ProfilService)
		secured.POST("/profil/update", inDB.ProfilUpdateService)
	}

	router.Run(":3000")
}