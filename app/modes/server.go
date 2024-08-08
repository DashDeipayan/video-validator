package modes

import (
	"fmt"

	"github.com/spf13/viper"
	routers "github.com/video-validator/internal/routes"
)

func StartServer() {
	r := routers.SetupRouter()
	port := fmt.Sprintf(":%s", viper.GetString("PORT"))
	r.Run(port)
}
