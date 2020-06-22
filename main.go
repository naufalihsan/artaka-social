package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/naufalihsan/artaka-social/wrapper"
	"os"
)

func main() {
	conn := "host=localhost port=5432 user=ihsan dbname=artaka-social sslmode=disable"
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&wrapper.User{})

	err = godotenv.Load()
	if err != nil {
		panic("failed to load env")
	}

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	r := gin.Default()
	instagram := r.Group("/instagram")
	{
		instagram.GET("/auth", func(c *gin.Context) {
			code := c.Query("code")
			if len(code) > 0 {
				code = wrapper.RemoveTrail(code)
				data := wrapper.ExchangeParams{
					ClientID:     clientID,
					ClientSecret: clientSecret,
					GrantType:    "authorization_code",
					RedirectURI:  "https://artaka.ap.ngrok.io/instagram/auth/",
					Code:         code,
				}

				token, err := wrapper.ExchangeCodeForToken(&data)
				if err != nil {
					panic(err)
				}

				db.Create(&wrapper.User{
					UserID:      token.UserID,
					AccessToken: token.AccessToken,
				})
			}

			c.JSON(200, gin.H{
				"message": "success",
			})
		})

		instagram.GET("/photos", func(c *gin.Context) {
			var user wrapper.User
			db.Take(&user)

			response, err := wrapper.QueryUserMedia(user.AccessToken)
			if err != nil {
				panic(err)
			}

			c.JSON(200, gin.H{
				"message": response,
			})
		})

		instagram.GET("/photo/:media", func(c *gin.Context) {
			mediaID := c.Param("media")

			var user wrapper.User
			db.Take(&user)

			response, err := wrapper.QueryMediaNode(mediaID, user.AccessToken)
			if err != nil {
				panic(err)
			}

			c.JSON(200, gin.H{
				"message": response,
			})
		})

	}

	r.Run(":8006")
}
