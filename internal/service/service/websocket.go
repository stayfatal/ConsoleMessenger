package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
)

func (cs *ChatService) Upgrade(c *gin.Context) error {
	conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
	if err != nil {
		return err
	}

	id := c.GetInt("id")

	cs.addChatMember(id, conn)

	return nil
}
