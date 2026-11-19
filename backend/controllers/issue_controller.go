
package controllers

import (
    "backend/config"
    "backend/models"
    "github.com/gin-gonic/gin"
)

//  Creating any issue 
func CreateIssue (c *gin.Context){
	 var issues models.Issue

	 if err := c.ShouldBindJSON(&issues); err !=nil{
		c.JSON(400, gin.H{"error ": err})
		return 
	 }
	 if err := config.DB.Create(&issues) ; err !=nil{
		c.JSON(400, gin.H{"error ": "Failed to Create data in DB..."})
		return
	 }

	//  Reading issues 
	func ReadIssues (c *gin.Context){
		 var readIssue []models.Issue

		config.DB.Where("priority = ?" , c.Query("priority")).Find(&issues)
		c.JSON(200, gin.H{"message ": issues})
	}  

	// For Updation 
func UpdateIssue(c *gin.Context) {
    var issue models.Issue
    config.DB.First(&issue, c.Param("id"))
    c.ShouldBindJSON(&issue)
    config.DB.Save(&issue)
    c.JSON(200, issue)
}


//  For Deletiom
func DeleteIssue(c *gin.Context) {
    config.DB.Delete(&models.Issue{}, c.Param("id"))
    c.JSON(200, gin.H{"message": "Deleted"})
}

	 c.JSON(200, gin.H{
		"message ": "Issues Created"
	 })



}