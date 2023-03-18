package controller

import (
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/rofinafiin/hdbackend"
	"github.com/rofinafiin/iteung/config"
	"github.com/whatsauth/whatsauth"
)

func WsWhatsAuthQR(c *websocket.Conn) {
	whatsauth.RunSocket(c, config.PublicKey, config.Usertables[:], config.Ulbimariaconn)
}

func PostWhatsAuthRequest(c *fiber.Ctx) error {
	if string(c.Request().Host()) == config.Internalhost {
		var req whatsauth.WhatsauthRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}
		ntfbtn := whatsauth.RunModuleLegacy(req, config.PrivateKey, config.Usertables[:], config.Ulbimariaconn)
		return c.JSON(ntfbtn)
	} else {
		var ws whatsauth.WhatsauthStatus
		ws.Status = string(c.Request().Host())
		return c.JSON(ws)
	}

}

func GetHome(c *fiber.Ctx) error {
	getip := musik.GetIPaddress()
	return c.JSON(getip)
}

func GetdataHD(c *fiber.Ctx) error {
	getstats := hdbackend.GetDataAllbyStats("Aktif", config.MongoConn, "data_complain")
	return c.JSON(getstats)
}

func GetdataHelper(c *fiber.Ctx) error {
	getdata := hdbackend.GetDataHelperFromPhone("085156007137", config.MongoConn, "helperdata")
	return c.JSON(getdata)
}
