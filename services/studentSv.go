package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	"helloworld/struct"
	"net/http"
)

var student = model.Student{FirstName:"Krithcat", LastName: "Rojanaphruk", Age: 26}

func GetStudent(c *gin.Context) {
	c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
	var json model.Student
	if isRequestValid(c, &json){
		err := update(json)
		if err != nil{
			c.JSON(http.StatusOK, generateResponse("400", err.Error()))
		}else{
			c.JSON(http.StatusOK, generateResponse("200", "Update Student Successfully"))
		}
	}
}

func update(st model.Student) error {
	if st.FirstName == student.FirstName {
		return errors.New("FirstName is same with current name")
	}else {
		student.Age = st.Age
		student.FirstName = st.FirstName
		student.LastName = st.LastName
		return nil
	}
}

func generateResponse(statusCode string, statusDesc string)model.AbsResponse{
	return model.AbsResponse{StatusCode:statusCode, StatusDesc:statusDesc}
}

func isRequestValid(c *gin.Context, json *model.Student) bool{
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	return true
}