package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"strconv"
)

type persona struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
	Cedula   string `json:"cedula"`
}

var personas = []persona{
	{ID: 1, Nombre: "Juan", Apellido: "Pérez", Edad: 30, Cedula: "123456789"},
	{ID: 2, Nombre: "Ana", Apellido: "Gómez", Edad: 25, Cedula: "987654321"},
	{ID: 3, Nombre: "Luis", Apellido: "Martínez", Edad: 40, Cedula: "456789123"},
}

func getPersonas(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, personas)
}

func createPersona(c *gin.Context) {
	var newPersona persona

	if err := c.BindJSON(&newPersona); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	for i := 0; i < len(personas); i++ {
		if personas[i].ID == newPersona.ID {
			c.IndentedJSON(http.StatusConflict, gin.H{"error": "Persona with ID already exists"})
			return
		}
	}

	newPersona.ID = len(personas) + 1
	personas = append(personas, newPersona)

	c.IndentedJSON(http.StatusCreated, newPersona)
}

func updatePersona(c *gin.Context) {
	var updatedPersona persona

	if err := c.BindJSON(&updatedPersona); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	for i, p := range personas {
		if newid == p.ID {
			personas[i] = updatedPersona
			c.IndentedJSON(http.StatusOK, updatedPersona)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Persona not found"})
}

func updateField(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updated map[string]interface{}

	if err := c.BindJSON(&updated); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	for i, p := range personas {
		if p.ID == newid {
			// Modificar nombre si viene en el JSON
			if nombre, ok := updated["nombre"].(string); ok {
				personas[i].Nombre = nombre
			}
			// Modificar apellido si viene en el JSON
			if apellido, ok := updated["apellido"].(string); ok {
				personas[i].Apellido = apellido
			}
			// Modificar edad si viene en el JSON
			if edad, ok := updated["edad"].(float64); ok {
				personas[i].Edad = int(edad)
			}
			// Modificar cedula si viene en el JSON
			if cedula, ok := updated["cedula"].(string); ok {
				personas[i].Cedula = cedula
			}
			c.IndentedJSON(http.StatusOK, personas[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Persona not found"})
}

func deletePersona(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	for i, p := range personas {
		if newid == p.ID {
			personas = append(personas[:i], personas[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Persona deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Persona not found"})
}

func getPersonaByID(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, p := range personas {
		if newid == p.ID {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Persona not found"})
}

func getPersonaByNombre(c *gin.Context) {
	nombre := c.Param("nombre")

	for _, p := range personas {
		if nombre == p.Nombre {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Persona not found"})
}

func getPersonaByApellido(c *gin.Context) {
	apellido := c.Param("apellido")

	for _, p := range personas {
		if apellido == p.Apellido {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Persona not found"})
}

func getPersonaByEdad(c *gin.Context) {
	edad := c.Param("edad")
	newEdad, err := strconv.Atoi(edad)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid age"})
		return
	}

	for _, p := range personas {
		if newEdad == p.Edad {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Persona not found"})
}

func getPersonaByCedula(c *gin.Context) {
	cedula := c.Param("cedula")

	for _, p := range personas {
		if cedula == p.Cedula {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Persona not found"})
}

func index(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

func main() {

	r := gin.Default()
	r.GET("/", index)
	r.GET("/personas", getPersonas)
	r.GET("/personas/id/:id", getPersonaByID)
	r.GET("/personas/nombre/:nombre", getPersonaByNombre)
	r.GET("/personas/apellido/:apellido", getPersonaByApellido)
	r.GET("/personas/edad/:edad", getPersonaByEdad)
	r.GET("/personas/cedula/:cedula", getPersonaByCedula)
	r.PATCH("personas/:id", updateField)
	r.POST("/personas", createPersona)
	r.PUT("/personas/:id", updatePersona)
	r.DELETE("/personas/:id", deletePersona)

	r.Run("localhost:8080")
}
