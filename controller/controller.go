package controller

import (
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/rofinafiin/hdbackend"
	"github.com/rofinafiin/iteung/config"
	"github.com/whatsauth/whatsauth"
)

var Helpercol = "helperdata"
var Datacomcol = "data_complain"

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
	getstats := hdbackend.GetDataAllbyStats("Aktif", config.MongoConn, Datacomcol)
	return c.JSON(getstats)
}

func GetdataHelper(c *fiber.Ctx) error {
	hp := c.Params("handphone")
	getdata := hdbackend.GetDataHelperFromPhone(hp, config.MongoConn, Helpercol)
	return c.JSON(getdata)
}

func GetDataComplainbyNumber(c *fiber.Ctx) error {
	hp := c.Params("status")
	crot := hdbackend.GetDataCompFromStatus(hp, config.MongoConn, Datacomcol)
	return c.JSON(crot)
}
