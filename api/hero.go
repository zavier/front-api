package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Hero struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var heroes = []*Hero{
	{
		Id:   11,
		Name: "Dr Nice",
	},
	{
		Id:   12,
		Name: "Nacro",
	},
	{
		Id:   13,
		Name: "Bombasto",
	},
	{
		Id:   14,
		Name: "Celeritas",
	},
	{
		Id:   15,
		Name: "Magneta",
	},
	{
		Id:   16,
		Name: "RubberMan",
	},
	{
		Id:   17,
		Name: "Dyname",
	},
	{
		Id:   18,
		Name: "Dr IQ",
	},
	{
		Id:   19,
		Name: "Magma",
	},
	{
		Id:   20,
		Name: "Tornado",
	},
}

func ConfigHero(router *gin.Engine) {
	heroRouter := router.Group("/api/hero")

	heroRouter.GET("heroes", list)

	heroRouter.GET("heroes/:id", getById)

	heroRouter.PUT("heroes", update)

	heroRouter.POST("heroes", add)

	heroRouter.DELETE("heroes/:id", deleteById)
}

var deleteById = func(c *gin.Context) {
	idStr := c.Param("id")
	id, e := strconv.Atoi(idStr)
	if e != nil {
		log.Println("id param error, id=", id)
		c.JSON(http.StatusBadRequest, "")
	}
	length := len(heroes)
	for i := 0; i < length; i++ {
		if heroes[i].Id == id {
			heroes = append(heroes[:i], heroes[i+1:]...)
			break
		}
	}
	c.JSON(http.StatusOK, "")
}

var add = func(c *gin.Context) {
	hero := &Hero{}
	err := c.ShouldBind(hero)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	maxId := -1
	for _, h := range heroes {
		if h.Id > maxId {
			maxId = h.Id
		}
	}
	hero.Id = maxId + 1
	heroes = append(heroes, hero)
	c.JSON(http.StatusOK, hero)
}

var update = func(c *gin.Context) {
	hero := &Hero{}
	err := c.ShouldBind(hero)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	id := hero.Id
	for _, h := range heroes {
		if h.Id == id {
			h.Name = hero.Name
		}
	}
	c.JSON(http.StatusOK, "")
}

var getById = func(c *gin.Context) {
	idStr := c.Param("id")
	id, e := strconv.Atoi(idStr)
	if e != nil {
		log.Println("id param error, id=", id)
		c.JSON(http.StatusBadRequest, "")
	}
	for _, hero := range heroes {
		if hero.Id == id {
			c.JSON(http.StatusOK, hero)
			return
		}
	}
	c.JSON(http.StatusBadRequest, "")
}

var list = func(c *gin.Context) {
	queryName := c.Query("name")
	if queryName == "" {
		c.JSON(http.StatusOK, heroes)
		return
	} else {
		heroesList := make([]*Hero, 0)
		for _, h := range heroes {
			contains := strings.Contains(h.Name, queryName)
			if contains {
				heroesList = append(heroesList, h)
			}
		}
		c.JSON(http.StatusOK, heroesList)
	}
}
