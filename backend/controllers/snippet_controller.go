package controllers

import (
	 "backend/config"   
    "backend/models"
    "net/http"

    "github.com/gin-gonic/gin"
)

func CreateSnippet(c *gin.Context){
    var snippet models.snippets

    if err := c.ShouldBindJSON(&snippet); err !=nil{
        c.JSON(400,  gin.H{"error": err})
        return 
    }
    if err := Config.DB.Create(&snippet); err !=nil{
        c.JSON(400, gin.H{"error ": "Snippet not Created in DataBase...!"})
        return 
    }
     
    c.JSON(200, gin.H{
        "Message ": "Snippet Created  Successfully >>>"
    })

    // For Reading the Snippet
    func ReadSnippets(c *gin.Context){
        var snippets []models.snippet
        config.DB.Where("stack = ?", c.Query("stack")).Find(&snippets)

        c.JSON(200, gin.H{"data": snippets})
    }

    // For Updating Snippet
    func UpdateSnippets(c *gin.Context){
        var snippets  models.snippet

        config.DB.First(&snippets, c.Param("id"))
        c.ShouldBindJSON(&snippets)
        config.DB.Save(&snippets)

        c.JSON(200, gin.H{
            "message": "Snippet Updated Successfully >>>"
        })


    // For Deleting Snippets 
    func DeleteSnippet(c *gin.Context){
      config.DB.Delete(&models.snippets{})
      c.JSON(200, gin.H{"message": "Deleted Successfully >>>"})
    }
    }




}