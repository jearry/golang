package main

import (
	//	"github.com/gin-gonic/gin"
	//	"github.com/jung-kurt/gofpdf"
	qrcode "github.com/skip2/go-qrcode"
	"log"
	//"net/http"
)

func main() {
	err := qrcode.WriteFile("http://v9.vc/abcd", qrcode.Low, 800, "qr.png")

	log.Println(err)

	//	pdf := gofpdf.New("P", "mm", "A4", "")
	//	pdf.AddPage()
	//	pdf.Image("qr.png", 50, 50, 100, 100, false, "", 0, "")

	//	err = pdf.OutputFileAndClose("basic.pdf")

	//	log.Println(err)

	//	router := gin.Default()

	//	// This handler will match /user/john but will not match neither /user/ or /user
	//	router.GET("/user/:name", func(c *gin.Context) {
	//		name := c.Param("name")
	//		c.String(http.StatusOK, "Hello %s", name)
	//	})

	//	router.GET("localhost/:name", func(c *gin.Context) {
	//		name := c.Param("name")
	//		c.String(http.StatusOK, "Root Hello %s", name)
	//	})

	//	// However, this one will match /user/john/ and also /user/john/send
	//	// If no other routers match /user/john, it will redirect to /user/join/
	//	router.GET("/user/:name/*action", func(c *gin.Context) {
	//		name := c.Param("name")
	//		action := c.Param("action")
	//		message := name + " is " + action
	//		c.String(http.StatusOK, message)
	//	})

	//	router.Run(":8080")
}
