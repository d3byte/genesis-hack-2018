package handler

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google-service/driver"
	"google-service/models"
	"google-service/repository"
	"google-service/repository/orm"
	"net/http"
	"strings"
)

func NewMiptHandler(db *driver.DB) *Mipt {
	return &Mipt{
		repo: orm.MongoMiptRepo{Conn: db.Mgo}.NewSQLRepo(),
	}
}

type Mipt struct {
	repo repository.MiptRepo
}

// Список конфигов godoc
// @Summary Список конфигов
// @Description Получение списка конфигов
// @Tags Mipt
// @Accept json
// @Produce json
// @Success 200 {array} models.HTTPError
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /mipts [get]
func (c *Mipt) Fetch(ctx *gin.Context) {
	payload, err := c.repo.Fetch(ctx, 0)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, err)
	}

	RespondWithJSON(ctx, http.StatusOK, payload)
}

// Создание конфига godoc
// @Summary Создать конфиг
// @Description Создание конфига
// @Tags Mipt
// @Param Mipt body models.ConfigInterface true "Поля необходимые для создания конфига"
// @Accept json
// @Produce json
// @Success 200 {object} models.HTTPError
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /mipt [post]
func (c *Mipt) Create(ctx *gin.Context) {
	var model *models.ConfigInterface

	err := ctx.BindJSON(&model)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

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
			RespondWithError(ctx, http.StatusBadRequest, err)
			return
		}

		model.CreatorID = oid
		model.ID = primitive.NewObjectID()
	} else {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := c.repo.Create(ctx, model)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	RespondWithJSON(ctx, http.StatusOK, map[string]string{"success": "Information added", "id": res})
}

// Поиск конфига по id godoc
// @Summary Поиск конфига по id
// @Description Поиск и вывод конфига по id
// @Tags Mipt
// @Param id path primitive.ObjectID true "ID конфига"
// @Accept json
// @Produce json
// @Success 200 {object} models.HTTPError
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /mipt/{id} [get]
func (c *Mipt) GetMiptById(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))

	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	payload, err := c.repo.GetMiptById(ctx, id)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	RespondWithJSON(ctx, http.StatusOK, payload)
}

func (c *Mipt) Publish(ctx *gin.Context, model *models.ConfigInterface) {
	res, err := c.repo.Create(ctx, model)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	RespondWithJSON(ctx, http.StatusOK, map[string]string{"success": "Information added", "id": res})
}
