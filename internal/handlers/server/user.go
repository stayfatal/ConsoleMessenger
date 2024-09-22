package handlers

import (
	"messenger/internal/authentication"
	"messenger/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (hm *handlersManager) RegisterHandler(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "binding json")).Msg("")
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	securedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "generating hashed password")).Msg("")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	user.Password = string(securedPass)

	userId, err := hm.dm.CreateUser(user)
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "calling CreateUser")).Msg("")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	token, err := authentication.CreateToken(userId)
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "calling CreateToken")).Msg("")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (hm *handlersManager) LoginHandler(c *gin.Context) {
	user := models.User{}
	c.ShouldBindJSON(&user)

	scannedUser, err := hm.dm.GetUserByName(user.Username)
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "calling GetUserByName")).Msg("")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(scannedUser.Password), []byte(user.Password))
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "comparing password to hashed password")).Msg("")
		c.String(http.StatusForbidden, err.Error())
		return
	}

	token, err := authentication.CreateToken(scannedUser.Id)
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "calling CreateToken")).Msg("")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (hm *handlersManager) ValidateTokenHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
