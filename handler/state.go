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

func NewStateHandler(db *driver.DB) *State {
	return &State{
		repo: orm.MongoStateRepo{Conn: db.Mgo}.NewSQLRepo(),
	}
}

type State struct {
	repo repository.StateRepo
}

// Создание стейта godoc
// @Summary Создать стейт
// @Description Создание стейта
// @Tags State
// @Param id path string true "ID конфига"
// @Param State body models.StateInterface true "Поля необходимые для создания стейта"
// @Accept json
// @Produce json
// @Success 200 {object} models.HTTPError
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /state/{id} [post]
func (c *State) Create(ctx *gin.Context) {
	var model *models.StateInterface

	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	err = c.repo.CheckMiptById(ctx, id)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	err = ctx.BindJSON(&model)
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

		model.UserID = oid
		model.ConfigID = id
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

// Поиск стейта по id godoc
// @Summary Поиск стейта по id
// @Description Поиск и вывод стейта по id
// @Tags State
// @Param id path string true "ID конфига"
// @Accept json
// @Produce json
// @Success 200 {object} models.HTTPError
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /state/{id} [get]
func (c *State) GetStateById(ctx *gin.Context) {
	var model models.StateInterface

	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	err = c.repo.CheckMiptById(ctx, id)
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

		model.UserID = oid
		model.ConfigID = id
	} else {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := c.repo.GetStateById(ctx, &model)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	RespondWithJSON(ctx, http.StatusOK, res)
}
