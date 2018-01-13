package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type MetaData struct {
	total int64
}

type Error struct {}

var ORM orm.Ormer


func InitDB() {
	ORM = GetOrmObject()
}

func GetBooks(c *gin.Context) {

	var books []Book

	res, err := ORM.Raw("SELECT * FROM book").QueryRows(&books)

	if err != nil {
		checkErr(err, "Failed to fetch books")
		return
	}

	var metadata MetaData
	metadata.total = res

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": &books,
		"metadata": metadata,
	})
}

func CreateBook(c *gin.Context) {
	var book Book
	c.Bind(&book)

	if book.Title != "" && book.Author != "" {
		_, err := ORM.Insert(&book)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
				"title": book.Title,
				"author": book.Author,
			})
		} else {
			checkErr(err, "Creation failed")
		}
	} else {
		c.JSON(422, gin.H{"error": "fields are empty"})
	}
}

func GetBook(c *gin.Context) {
	if book, err := checkParam(c); err == nil {
		read_error := ORM.Read(&book)
		if read_error == orm.ErrNoRows {
			Error{}.notFound(c)
			return
		} else if read_error != nil {
			c.JSON(404, gin.H{
				"error": "Error while fetching users",
			})
			return
		} else {
			c.JSON(200, gin.H{"book": book})
		}
	} else {
		Error{}.notFound(c)
	}
}

func UpdateBook(c *gin.Context) {
	var body Book
	book, err := checkParam(c); if err == nil {
		c.Bind(&body)
		readError := ORM.Read(&book)
		if readError == nil {
			toUpdate := []string{}
			if len(body.Author) > 0 {
				book.Author = body.Author
				toUpdate = append(toUpdate, "Author")
			}
			if len(body.Title) > 0 {
				book.Title = body.Title
				toUpdate = append(toUpdate, "Title")
			}
			if _, err := ORM.Update(&book, toUpdate...); err == nil {
				c.JSON(204, gin.H{})
			}
		} else if (readError == orm.ErrNoRows) {
			Error{}.notFound(c)
		}
		return
	}
	Error{}.notFound(c)
}

func DeleteBook(c *gin.Context) {
	if book, err := checkParam(c); err == nil {
		readError := ORM.Read(&book)
		if readError == nil {
			if _, err := ORM.Delete(&book); err == nil {
				c.JSON(204, gin.H{})
			} else {
				Error{}.serverError(c)
			}
		} else {
			Error{}.notFound(c)
		}
		return
	}
	Error{}.notFound(c)
}

func checkParam(c *gin.Context) (Book, error) {
	param, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return Book{}, err
	}
	book := Book{Id: param}
	return book, nil
}

func (e Error) notFound(c *gin.Context) {
	c.JSON(404, gin.H{ "error": "resource not found" })
	return
}

func (e Error) serverError(c *gin.Context) {
	c.JSON(500, gin.H{"error": "an unexpected error occurred please try again in a few minute"})
	return
}
