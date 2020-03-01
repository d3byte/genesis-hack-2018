package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google-service/docs"
	"google-service/driver"
	"google-service/handler"
	"google-service/models"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

func SetupRouter(db *driver.DB) *gin.Engine {
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.LoadHTMLGlob("static/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	lHandler := handler.GoogleHandler()
	v1 := r.Group("/api/v1/")
	{
		// @tag.name Users
		// @tag.description Работа с пользователями
		userHandler := handler.NewUserHandler(db)
		v1.POST("/login", userHandler.Login)
		v1.POST("/user", userHandler.Create)

		// @tag.name Spreadsheet
		// @tag.description
		lRouter := v1.Group("/spreadsheet")
		{
			lRouter.POST("", lHandler.CreateTable)
			lRouter.POST("/copy", lHandler.CopyTable)
			lRouter.POST("/append", lHandler.AppendData)
			lRouter.POST("/clear", lHandler.ClearData)
		}

		// @tag.name Config
		// @tag.description Config
		comHandler := handler.NewMiptHandler(db)
		v1.GET("/mipt", comHandler.Fetch)
		comRouter := v1.Group("/mipt")
		{
			comRouter.GET("/:id", comHandler.GetMiptById)
			comRouter.POST("", handler.AuthMiddleware.MiddlewareFunc(), comHandler.Create)
			//comRouter.PATCH("/:id", comHandler.Update)
		}

		v1.POST("/excel", handler.AuthMiddleware.MiddlewareFunc(), func(ctx *gin.Context) {
			file, _, err := ctx.Request.FormFile("file")
			if err != nil {
				ctx.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
				return
			}

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			part, err := writer.CreateFormFile("file", "Prepodavatel.xlsm")
			if err != nil {
				handler.RespondWithError(ctx, http.StatusBadRequest, err)
				return
			}
			_, err = io.Copy(part, file)

			err = writer.Close()
			if err != nil {
				handler.RespondWithError(ctx, http.StatusBadRequest, err)
				return
			}

			req, err := http.NewRequest("POST", "http://10.55.124.223:3000", body)
			req.Header.Set("Content-Type", writer.FormDataContentType())

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			} else {
				body := &bytes.Buffer{}
				_, err := body.ReadFrom(resp.Body)
				if err != nil {
					log.Fatal(err)
				}
				resp.Body.Close()

				var bodyConfig models.ConfigInterface

				err = json.Unmarshal(body.Bytes(), &bodyConfig)
				if err != nil {
					handler.RespondWithError(ctx, http.StatusBadRequest, err)
					return
				}

				bodyConfig.ID = primitive.NewObjectID()

				s := strings.Split(ctx.Request.Header["Authorization"][0], " ")
				token, err := jwt.Parse(s[1], func(token *jwt.Token) (interface{}, error) {
					// Don't forget to validate the alg is what you expect:
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}

					// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
					return []byte("SESSION_SECRET"), nil
				})

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					oid, err := primitive.ObjectIDFromHex(claims["id"].(string))
					if err != nil {
						handler.RespondWithError(ctx, http.StatusBadRequest, err)
						return
					}

					bodyConfig.CreatorID = oid
					bodyConfig.Title = ctx.Request.Form["title"][0]
					bodyConfig.PublicToken = ctx.Request.Form["publicToken"][0]

					t, err := time.Parse(time.RFC3339, ctx.Request.Form["expirationDate"][0])
					if err != nil {
						handler.RespondWithError(ctx, http.StatusBadRequest, err)
						return
					}
					bodyConfig.ExpirationDate = t
				} else {
					handler.RespondWithError(ctx, http.StatusBadRequest, err)
					return
				}

				comHandler.Publish(ctx, &bodyConfig)
				return
			}
		})

		// @tag.name State
		// @tag.description State
		stateHandler := handler.NewStateHandler(db)
		stateRouter := v1.Group("/state")
		{
			stateRouter.GET("/:id", handler.AuthMiddleware.MiddlewareFunc(), stateHandler.GetStateById)
			stateRouter.POST("/:id", handler.AuthMiddleware.MiddlewareFunc(), stateHandler.Create)
			//stateRouter.POST("/:id", handler.AuthMiddleware.MiddlewareFunc(), stateHandler.Create)
		}
	}



	// Set swagger info
	docs.SwaggerInfo.Title = "MIPT API"
	docs.SwaggerInfo.Description = "MIPT API description"
	docs.SwaggerInfo.Version = "0.1"
	docs.SwaggerInfo.Host = "auth.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
