package controllers

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/models"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FileController struct {
	DB *gorm.DB
}

func Index(c *gin.Context) {
	var posts []models.Post

	database.DB.Find(&posts)
	c.JSON(http.StatusOK, gin.H{"posts": posts})

}

func Detail(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	if err := database.DB.First(&post, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
}

func (cf *FileController) Create(c *gin.Context) {

	title := c.PostForm("title")
	content := c.PostForm("content")

	image, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// mengambil ekstensi file
	splitDots := strings.Split(image.Filename, ".")
	ext := splitDots[len(splitDots)-1]
	fmt.Println(ext)
	nameNewImage := fmt.Sprintf("%s.%s", time.Now().Format("20060102150405"), ext)
	fmt.Println(nameNewImage)

	imagePath := filepath.Join("uploads", nameNewImage)
	if err := c.SaveUploadedFile(image, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	fileMetadata := models.Post{
		Title:   title,
		Content: content,
		Image:   nameNewImage,
	}

	if err := cf.DB.Create(&fileMetadata).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file metadata"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "Details": fileMetadata})

}

func (cf *FileController) Update(c *gin.Context) {

	title := c.PostForm("title")
	content := c.PostForm("content")
	id := c.Param("id")

	image, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var post models.Post
	// Retrieve the file metadata from the database
	err = cf.DB.Where("id = ?", id).First(&post).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	// Define the path of the file to be deleted
	filePath := filepath.Join("uploads", post.Image)
	// Delete the file from the server
	err = os.Remove(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file from upload folder"})
		return
	}

	// mengambil ekstensi file
	splitDots := strings.Split(image.Filename, ".")
	ext := splitDots[len(splitDots)-1]
	nameNewImage := fmt.Sprintf("%s.%s", time.Now().Format("20060102150405"), ext)
	imagePath := filepath.Join("uploads", nameNewImage)
	if err := c.SaveUploadedFile(image, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	fileMetadata := models.Post{
		Title:   title,
		Content: content,
		Image:   nameNewImage,
	}

	if database.DB.Model(&fileMetadata).Where("id = ?", id).Updates(&fileMetadata).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

func (cf *FileController) Delete(c *gin.Context) {

	id := c.Param("id")
	var post models.Post
	// Retrieve the file metadata from the database
	err := cf.DB.Where("id = ?", id).First(&post).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	// Define the path of the file to be deleted
	filePath := filepath.Join("uploads", post.Image)
	// Delete the file from the server
	err = os.Remove(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file from upload folder"})
		return
	}

	if database.DB.Delete(&post, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
