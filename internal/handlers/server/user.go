package handlers

import (
	"messenger/internal/authentication"
	"messenger/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (hm *handlersManager) CreateUserHandler(c *gin.Context) {
	user := database.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Error().Err(err).Msg("")
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	securedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	user.Id = uuid.New()
	user.Password = string(securedPass)

	err = hm.dm.CreateUser(user)
	if err != nil {
		log.Error().Err(err).Msg("unable to create user")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	token, err := authentication.CreateToken(user.Id.String())
	if err != nil {
		log.Error().Err(err).Msg("unable to create token")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (hm *handlersManager) LoginHandler(c *gin.Context) {
	user := database.User{}
	c.ShouldBindJSON(&user)

	scannedUser, err := hm.dm.GetUserByName(user.Username)
	if err != nil {
		log.Error().Err(err).Msg("user not found")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(scannedUser.Password), []byte(user.Password))
	if err != nil {
		log.Error().Err(err).Msg("non-registered user")
		c.String(http.StatusForbidden, err.Error())
		return
	}

	token, err := authentication.CreateToken(scannedUser.Id.String())
	if err != nil {
		log.Error().Err(err).Msg("cant create token")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (hm *handlersManager) ValidateTokenHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
