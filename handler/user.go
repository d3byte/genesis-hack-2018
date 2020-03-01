package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	"golang.org/x/crypto/bcrypt"
	"google-service/driver"
	"google-service/models"
	"google-service/repository"
	"google-service/repository/orm"
	"net/http"
	"time"
)

func NewUserHandler(db *driver.DB) *User {
	return &User{
		repo: orm.MongoUserRepo{Conn: db.Mgo}.NewSQLRepo(),
	}
}

type User struct {
	repo repository.UserRepo
}

// Создание пользователя godoc
// @Summary Создать пользователя
// @Description Создание пользователя
// @Tags Users
// @Param user body models.User true "Поля необходимые для создания аккаунта"
// @Accept json
// @Produce json
// @Success 200 {object} models.UserWithHiddenFields
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /user [post]
func (u *User) Create(ctx *gin.Context) {
	var userModel *models.User

	err := ctx.BindJSON(&userModel)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	_, err = u.repo.GetUserByEmail(ctx, userModel.Email)
	if err == nil {
		RespondWithError(ctx, http.StatusBadRequest, errors.New(400, "Exist email"))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), bcrypt.DefaultCost)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	userModel.Password = string(hashedPassword)
	userModel.Confirmed = new(bool)
	userModel.Time = time.Now()
	id, err := u.repo.Create(ctx, userModel)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	RespondWithJSON(ctx, http.StatusOK, map[string]interface{}{"success": "Information added", "id": id})
}

// Авторизация пользователя godoc
// @Summary Авторизация пользователя
// @Description Авторизация пользователя
// @Tags Users
// @Param user body models.UserLogin true "Поля необходимые для авторизации аккаунта"
// @Accept json
// @Produce json
// @Success 200 {object} models.UserWithHiddenFields
// @Failure 400 {object} models.HTTPError
// @Failure 404 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /login [post]
func (u *User) Login(ctx *gin.Context) {
	var userModel *models.UserLogin

	err := ctx.BindJSON(&userModel)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, err)
		return
	}

	if userModel.Email == "" && userModel.Password == "" {
		token, err := u.repo.GenerateToken(ctx)
		if err != nil {
			RespondWithError(ctx, http.StatusBadRequest, errors.New(400, "Wrong token"))
			return
		}

		RespondWithJSON(ctx, http.StatusOK, map[string]interface{}{"token": token})
		return
	}

	user, err := u.repo.GetUserByEmail(ctx, userModel.Email)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, errors.New(400, "Wrong email"))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userModel.Password)); err != nil {
		RespondWithError(ctx, http.StatusBadRequest, errors.New(400, "Wrong password"))
		return
	}

	RespondWithJSON(ctx, http.StatusOK, map[string]interface{}{"token": user.Token, "confirmed": user.Confirmed})
}
