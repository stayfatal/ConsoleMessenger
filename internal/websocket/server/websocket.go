package websocket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/rs/zerolog/log"
)

func (wm *WebsocketManager) Upgrade(c *gin.Context) {
	conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		log.Error().Err(err).Msg("cant upgrade to websocket")
		return
	}

	id := c.GetInt("id")

	wm.AddChatMember(id, conn)
}
