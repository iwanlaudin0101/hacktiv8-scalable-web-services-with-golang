package controllers

import (
	"fmt"
	"net/http"
	"rest-api-postgres/helpers"
	"rest-api-postgres/models"
	"rest-api-postgres/services"

	"github.com/gin-gonic/gin"
)

func AllBook(ctx *gin.Context) {
	books := services.GetAllBook()
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": books,
	})
}

func CreatedBook(ctx *gin.Context) {
	var newBook models.Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid params!",
		})
		return
	}

	book := services.CreatedBook(newBook)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    book,
		"message": "Successfully added data!",
	})
}

func DetailBook(ctx *gin.Context) {
	Id, err := helpers.StrToInt(ctx.Param("Id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid params!",
		})
		return
	}

	book, err := services.GetBookById(Id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": fmt.Sprintf("Book with ID %d not found", Id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": book,
	})
}

func UpdatedBook(ctx *gin.Context) {
	var book models.Book

	Id, err := helpers.StrToInt(ctx.Param("Id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		return
	}

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		return
	}

	_, err = services.GetBookById(Id)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": fmt.Sprintf("Book with ID %d not found", Id),
		})
		return
	}

	rows := services.UpdatedBook(Id, book)
	if rows < 1 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Data failed updated!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Data successfully updated!",
	})
}

func DeletedBook(ctx *gin.Context) {
	Id, err := helpers.StrToInt(ctx.Param("Id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		return
	}

	_, err = services.GetBookById(Id)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": fmt.Sprintf("Book with ID %d not found", Id),
		})
		return
	}

	rows := services.DeletedBook(Id)
	if rows < 1 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Data successfully deleted!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Data successfully deleted!",
	})
}
