
package models

import "gorm.io/gorm"

type Issue struct {
    gorm.Model

    Issue_Title           string `json:"issue_title" binding:"required,min=5,max=150"`
    Category             string `json:"category" binding:"required"`
    Priority             string `json:"priority" binding:"required,oneof=Low Medium High Critical"`
    Primary_Language      string `json:"primary_language" binding:"required"`
    Framework_Stack       string `json:"framework_stack" binding:"required"`
    Technology_Tags       string `json:"technology_tags"`
    Detailed_Description  string `json:"detailed_description" binding:"required,min=20"`

    // Environment & Infrastructure Details
    Operating_System      string `json:"operating_system"`
    Language_Version      string `json:"language_version"`
    Database_Cache        string `json:"database_cache"`
    Framework_Version     string `json:"framework_version"`
    Other_Runtimes_Tools   string `json:"other_runtimes_tools"`

    // Error Details
    Error_Message         string `json:"error_message"`
    Stack_Trace           string `json:"stack_trace"`
    Expected_Behavior     string `json:"expected_behavior"`
    Actual_Behavior       string `json:"actual_behavior"`
    Steps_To_Reproduce     string `json:"steps_to_reproduce"`
    Code_Snippets     []snippets `json:"steps_to_reproduce"`
}
