package models 

import "gorm.io/gorm"

type snippets struct{
     gorm.Model
     
	Title       string `json:"title" binding:"required,min=3,max=100"`
    Stack       string `json:"stack" binding:"required"`
    Description string `json:"description" binding:"required,min=10"`
}