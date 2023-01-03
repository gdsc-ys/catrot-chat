package routers

import (
	"catrot-chat/docs"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary norification API
// @Schemes
// @Description notification API
// @Tags index
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router / [get]
func SetRoutes(r *gin.Engine) {

	// 라우트 설정
	setChatRoutes(r.Group("/fi_chat"))

	// notification api 와 스웨거 설정
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "HellGo, World 👋🇺🇸🇰🇷!\n"+time.Now().String()+"\n"+fmt.Sprintf("cores: %d", runtime.GOMAXPROCS(0)))
	})
	r.GET("/docs", func(context *gin.Context) {
		context.Redirect(http.StatusFound, "/docs/index.html")
	})
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	docs.SwaggerInfo.BasePath = "/"
}