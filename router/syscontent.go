package router

import (
"github.com/gin-gonic/gin"
"go-admin/middleware"
"go-admin/apis/syscontent"
jwt "go-admin/pkg/jwtauth"
)

// 需认证的路由代码
func registerSysContentRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

r := v1.Group("/syscontent").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
{
r.GET("/:id", syscontent.GetSysContent)
r.POST("", syscontent.InsertSysContent)
r.PUT("", syscontent.UpdateSysContent)
r.DELETE("/:id", syscontent.DeleteSysContent)
}

l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
{
l.GET("/syscontentList",syscontent.GetSysContentList)
}

}
