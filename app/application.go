package app

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	router.Run( addr: "8008")
}
