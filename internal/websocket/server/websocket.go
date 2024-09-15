package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
)

func (wm *WebsocketManager) Upgrade(c *gin.Context) error {
	conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
	if err != nil {
		return err
	}

	id := c.GetInt("id")

	wm.AddChatMember(id, conn)

	return nil
}
