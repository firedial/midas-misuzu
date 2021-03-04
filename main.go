package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"

    "github.com/firedial/midas-misuzu/controller"
    "github.com/firedial/midas-misuzu/config"
    "github.com/firedial/midas-misuzu/entity"
    "github.com/firedial/midas-misuzu/interactor"
    "github.com/firedial/midas-misuzu/auth"
)

func main() {

    if config.IS_PRODUCTION {
        gin.SetMode(gin.ReleaseMode)
    }

    salt := auth.GetSalt()
    authToken := auth.GetAuthToken(salt, config.ACCESS_PASSWORD)

    r := gin.Default()
    r.Use(cors.Default())
    r.Use(Auth(authToken))

    api := r.Group("/api/v1")
    {
        api.POST("/authentication/", func(c *gin.Context) { 
            c.JSON(200, controller.AuthPost(salt, c.PostForm("pass"))) } )
        api.GET("/authentication/check/", func(c *gin.Context) { 
            c.JSON(200, controller.AuthCheck("Bearer " + authToken == c.Request.Header.Get("Authorization"))) } )
        api.GET("/authentication/refresh/", func(c *gin.Context) { 
            salt = auth.GetSalt()
            authToken = auth.GetAuthToken(salt, config.ACCESS_PASSWORD)
            c.JSON(200, gin.H{
                "status": "OK",
                "message": "",
            })
        })

        api.GET("/balance/", func(c *gin.Context) { c.JSON(200, controller.BalanceGet(c.Request.URL.Query())) } )
        api.POST("/balance/", func(c *gin.Context) { 
            var balance entity.Balance
            c.BindJSON(&balance)
            c.JSON(200, controller.BalancePost(balance)) } )
        api.POST("/move/", func(c *gin.Context) { 
            var move interactor.Move
            c.BindJSON(&move)
            c.JSON(200, controller.MovePost(move)) } )
        api.GET("/sum/", func(c *gin.Context) { c.JSON(200, controller.SumGet(c.Request.URL.Query())) } )
        api.GET("/chart/", func(c *gin.Context) { c.JSON(200, controller.ChartGet(c.Request.URL.Query())) } )

        api.GET("/kind_elements/", func(c *gin.Context) { c.JSON(200, controller.AttributeElementsGet("kind")) } )
        api.GET("/purpose_elements/", func(c *gin.Context) { c.JSON(200, controller.AttributeElementsGet("purpose")) } )
        api.GET("/place_elements/", func(c *gin.Context) { c.JSON(200, controller.AttributeElementsGet("place")) } )
        api.POST("/kind_elements/", func(c *gin.Context) { 
            var attribute entity.AttributeElement
            c.BindJSON(&attribute)
            c.JSON(200, controller.AttributeElementPost("kind", attribute)) } )
        api.POST("/purpose_elements/", func(c *gin.Context) { 
            var attribute entity.AttributeElement
            c.BindJSON(&attribute)
            c.JSON(200, controller.AttributeElementPost("purpose", attribute)) } )
        api.POST("/place_elements/", func(c *gin.Context) { 
            var attribute entity.AttributeElement
            c.BindJSON(&attribute)
            c.JSON(200, controller.AttributeElementPost("place", attribute)) } )

        api.GET("/kind_categories/", func(c *gin.Context) { c.JSON(200, controller.AttributeCategoriesGet("kind")) } )
        api.GET("/purpose_categories/", func(c *gin.Context) { c.JSON(200, controller.AttributeCategoriesGet("purpose")) } )
        api.GET("/place_categories/", func(c *gin.Context) { c.JSON(200, controller.AttributeCategoriesGet("place")) } )
        api.POST("/kind_categories/", func(c *gin.Context) { 
            var attribute entity.AttributeCategory
            c.BindJSON(&attribute)
            c.JSON(200, controller.AttributeCategoryPost("kind", attribute)) } )
        api.POST("/purpose_categories/", func(c *gin.Context) { 
            var attribute entity.AttributeCategory
            c.BindJSON(&attribute)
            c.JSON(200, controller.AttributeCategoryPost("purpose", attribute)) } )
        api.POST("/place_categories/", func(c *gin.Context) { 
            var attribute entity.AttributeCategory
            c.BindJSON(&attribute)
            c.JSON(200, controller.AttributeCategoryPost("place", attribute)) } )
 
        api.GET("/kind_groups/", func(c *gin.Context) { c.JSON(200, controller.AttributeGroupsGet("kind")) } )
        api.GET("/purpose_groups/", func(c *gin.Context) { c.JSON(200, controller.AttributeGroupsGet("purpose")) } )
        api.GET("/place_groups/", func(c *gin.Context) { c.JSON(200, controller.AttributeGroupsGet("place")) } )
        api.POST("/kind_groups/", func(c *gin.Context) { 
            var attribute entity.AttributeGroup
            c.BindJSON(&attribute)
            c.JSON(200, controller.AttributeGroupPost("kind", attribute)) } )
        api.POST("/purpose_groups/", func(c *gin.Context) { 
            var attribute entity.AttributeGroup
            c.BindJSON(&attribute)
            c.JSON(200, controller.AttributeGroupPost("purpose", attribute)) } )
        api.POST("/place_groups/", func(c *gin.Context) { 
            var attribute entity.AttributeGroup
            c.BindJSON(&attribute)
            c.JSON(200, controller.AttributeGroupPost("place", attribute)) } )
 
        api.GET("/kind_group_relations/", func(c *gin.Context) { c.JSON(200, controller.AttributeGroupRelationsGet("kind")) } )
        api.GET("/purpose_group_relations/", func(c *gin.Context) { c.JSON(200, controller.AttributeGroupRelationsGet("purpose")) } )
        api.GET("/place_group_relations/", func(c *gin.Context) { c.JSON(200, controller.AttributeGroupRelationsGet("place")) } )
        api.POST("/kind_group_relations/", func(c *gin.Context) { 
            var attribute entity.AttributeGroupRelation
            c.BindJSON(&attribute)
            c.JSON(200, controller.AttributeGroupRelationPost("kind", attribute)) } )
        api.POST("/purpose_group_relations/", func(c *gin.Context) { 
            var attribute entity.AttributeGroupRelation
            c.BindJSON(&attribute)
            c.JSON(200, controller.AttributeGroupRelationPost("purpose", attribute)) } )
        api.POST("/place_group_relations/", func(c *gin.Context) { 
            var attribute entity.AttributeGroupRelation
            c.BindJSON(&attribute)
            c.JSON(200, controller.AttributeGroupRelationPost("place", attribute)) } )
 
    }
    r.Run()

}

func Auth(authToken string) gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.Request.Header.Get("Authorization")

        if c.Request.URL.Path == "/api/v1/authentication/" {

        } else if c.Request.URL.Path == "/api/v1/authentication/check/" {

        } else if token != "Bearer " + authToken {
            c.AbortWithStatus(401)
        }
        c.Next()
    }
}
