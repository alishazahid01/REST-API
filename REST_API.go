package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type identity struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

var identities = []identity{
	{ID: "1", Name: "Alex", Age: 21, Gender: "Female"},
	{ID: "2", Name: "Mark", Age: 22, Gender: "Male"},
	{ID: "3", Name: "Ronaldo", Age: 20, Gender: "Female"},
	{ID: "4", Name: "Messi", Age: 17, Gender: "Male"},
	{ID: "5", Name: "Roy", Age: 14, Gender: "Female"},
}

func getIdentity(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, identities)
}

func returnIdentity(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
		return
	}

	identity, err := getIdentityByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Identity not Found"})
		return
	}

	identity.ID = "updated"
	c.IndentedJSON(http.StatusOK, identity)
}

func identityByID(c *gin.Context) {
	id := c.Param("id")
	identity, err := getIdentityByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Identity not Found"})
		return
	}

	c.IndentedJSON(http.StatusOK, identity)
}

func getIdentityByID(id string) (*identity, error) {
	for a, i := range identities {
		if i.ID == id {
			return &identities[a], nil
		}
	}

	return nil, errors.New("identity not found")
}

func createIdentity(c *gin.Context) {
	var newIdentity identity

	if err := c.BindJSON(&newIdentity); err != nil {
		return
	}

	identities = append(identities, newIdentity)
	c.IndentedJSON(http.StatusCreated, newIdentity)
}

func main() {

	router := gin.Default()
	router.GET("/identities", getIdentity)
	router.GET("/identities/:id", identityByID)
	router.POST("/identities", createIdentity)
	router.PATCH("/return", returnIdentity)
	router.Run("localhost:8080")

}
