package handler

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google-service/models"
	"time"
)

var AuthMiddleware, _ = jwt.New(&jwt.GinJWTMiddleware{
	Realm:       "Moscow",
	Key:         []byte("SESSION_SECRET"),
	Timeout:     time.Hour,
	MaxRefresh:  time.Hour,
	IdentityKey: "id",
	PayloadFunc: func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*models.User); ok {
			return jwt.MapClaims{
				"id": v.ID,
			}
		}
		return jwt.MapClaims{}
	},
	IdentityHandler: func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		id, _ := claims["id"].(string)
		objectID, _ := primitive.ObjectIDFromHex(id)

		return &models.User{
			ID: objectID,
		}
	},
	Authorizator: func(data interface{}, c *gin.Context) bool {
		if _, ok := data.(*models.User); ok {
			return true
		}

		return false
	},
	Unauthorized: func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	},
	LoginResponse: func(c *gin.Context, code int, message string, time time.Time) {
		RespondWithJSON(c, code, map[string]interface{}{"token": message, "expire": time})
	},
	TokenLookup:   "header: Authorization, query: token, cookie: jwt",
	TokenHeadName: "Bearer",
	TimeFunc:      time.Now,
})
