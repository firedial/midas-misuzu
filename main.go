package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"

    "github.com/firedial/midas-go/controller"
    "github.com/firedial/midas-go/config"
    "github.com/firedial/midas-go/entity"
    "github.com/firedial/midas-go/interactor"
)

func main() {

    if config.IS_PRODUCTION {
        gin.SetMode(gin.ReleaseMode)
    }

    r := gin.Default()
    r.Use(cors.Default())

    api := r.Group("/api/v1")
    {
        api.GET("/balance/", func(c *gin.Context) { c.JSON(200, controller.BalanceGet(c.Request.URL.Query())) } )
        api.POST("/balance/", func(c *gin.Context) { 
            var balance entity.Balance
            c.BindJSON(&balance)
            c.String(200, controller.BalancePost(balance)) } )
        api.POST("/move/", func(c *gin.Context) { 
            var move interactor.Move
            c.BindJSON(&move)
            c.JSON(200, controller.MovePost(move)) } )
        api.GET("/kind/", func(c *gin.Context) { c.JSON(200, controller.AttributeGet("kind")) } )
        api.GET("/purpose/", func(c *gin.Context) { c.JSON(200, controller.AttributeGet("purpose")) } )
        api.GET("/place/", func(c *gin.Context) { c.JSON(200, controller.AttributeGet("place")) } )
        api.GET("/sum/", func(c *gin.Context) { c.JSON(200, controller.SumGet(c.Request.URL.Query())) } )
    }

    r.Run()
}
