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

// Initialize DataBase ORM
func InitDB() {
	ORM = GetOrmObject()
}

// Get All books on the application
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
		"data": &books,
		"metadata": metadata,
	})
}

// Create a single book
func CreateBook(c *gin.Context) {
	var book Book
	c.Bind(&book)

	// Ensure both title and author are passed in
	if book.Title != "" && book.Author != "" {
		_, err := ORM.Insert(&book)
		if err == nil {
			c.JSON(http.StatusCreated, gin.H{
				"title": book.Title,
				"author": book.Author,
			})
		} else {
			Error{}.serverError(c)
		}
	} else {
		c.JSON(422, gin.H{"error": "fields are empty"})
	}
}

// Get a book by its ID
func GetBook(c *gin.Context) {
	// Check if the parameter passed into the URL is valid
	if book, err := checkParam(c); err == nil {
		read_error := ORM.Read(&book)
		// If book does not exist
		if read_error == orm.ErrNoRows {
			Error{}.notFound(c)
			return
		} else if read_error != nil { // If some other sort of unknown error
			Error{}.serverError(c)
			return
		} else {
			c.JSON(200, gin.H{"book": book})
		}
	} else {
		Error{}.notFound(c)
	}
}

// Update the content of a book
func UpdateBook(c *gin.Context) {
	var body Book
	// Check if parameter url is valid
	book, err := checkParam(c); if err == nil {
		c.Bind(&body)
		readError := ORM.Read(&book)
		if readError == nil {
			toUpdate := []string{}
			if len(body.Author) > 0 { // check if user is updating the author field
				book.Author = body.Author
				toUpdate = append(toUpdate, "Author")
			}
			if len(body.Title) > 0 { // check if user is updating the title field
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

// Delete a book from the platform
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

// Check the validity of a URL parameter
func checkParam(c *gin.Context) (Book, error) {
	param, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return Book{}, err
	}
	book := Book{Id: param}
	return book, nil
}

// Handle 404 errors
func (e Error) notFound(c *gin.Context) {
	c.JSON(404, gin.H{ "error": "resource not found" })
	return
}

// Handle internal server errors
func (e Error) serverError(c *gin.Context) {
	c.JSON(500, gin.H{"error": "an unexpected error occurred please try again in a few minute"})
	return
}
