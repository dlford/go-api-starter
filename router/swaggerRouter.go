package router

import (
	"net/url"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"api/constants"
	docs "api/docs"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Enter "Bearer" followed by a space and then your token, e.g. "Bearer YOUR_TOKEN"
func swaggerRouter(r *gin.Engine) {
	fqdnUrl, err := url.Parse(constants.FQDN)
	if err != nil {
		panic(err)
	}

	docs.SwaggerInfo.Title = "Api Documentation"
	docs.SwaggerInfo.Version = "v1"
	docs.SwaggerInfo.Host = fqdnUrl.Host
	docs.SwaggerInfo.Schemes = []string{fqdnUrl.Scheme}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
