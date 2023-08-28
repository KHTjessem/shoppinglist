package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerPages(router *gin.Engine) {
	router.GET("/", homePage)
	router.GET("/lists", getAllLists)
	router.GET("/listitems/:lid", getAllListItems)
	router.POST("/newlist", saveNewList)
	router.POST("/newitem", saveNewItem)
	router.POST("/updateitem", updateItem)
	router.POST("/deleteitem", deleteItem)
	router.POST("/updatelist", updateList)
	router.POST("/deletelist", deleteList)
	router.POST("/completelist", changeListStatus)
	router.POST("/completeitem", changeItemStatus)

}
func homePage(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Welcome to the backend!")
}

func getAllLists(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	lists, _ := DATABASE.getAllLists()
	c.IndentedJSON(http.StatusOK, lists)
}

func getAllListItems(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	asd := c.Param("lid")
	lid, err := strconv.Atoi(asd)
	if err != nil {
		fmt.Fprintf(c.Writer, "Invalid id")
		return
	}

	items, _ := DATABASE.getListItems(lid)

	c.JSON(http.StatusOK, items)
}

func saveNewItem(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	var newItem Item

	if err := c.BindJSON(&newItem); err != nil {
		println(err.Error())
		return
	}
	itID, err := DATABASE.InsertNewItem(&newItem)
	if err != nil {
		return
	}
	DATABASE.InsertListItemRel(newItem.ListID, itID)

	resp, err := DATABASE.getItemByID(itID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Encountered an error.")
		return
	}

	c.JSON(http.StatusOK, resp)
}

func saveNewList(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	var newList List

	if err := c.BindJSON(&newList); err != nil {
		return
	}

	lid, err := DATABASE.InsertNewList(&newList)
	if err != nil {
		return
	}
	retList, err := DATABASE.getListByID(lid)
	if err != nil {

	}
	c.JSON(http.StatusCreated, retList)
}

func updateItem(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	var newItem Item

	if err := c.BindJSON(&newItem); err != nil {
		println(err.Error())
		return
	}

	DATABASE.updateItem(&newItem)
	c.JSON(http.StatusOK, "{'status': true}")
}
func changeItemStatus(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	var it Item

	if err := c.BindJSON(&it); err != nil {
		println(err.Error())
		return
	}

	DATABASE.completeItem(&it)
}

func deleteItem(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	var it Item

	if err := c.BindJSON(&it); err != nil {
		println(err.Error())
		return
	}
	DATABASE.delteItem(&it)
}

func updateList(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	var li List

	if err := c.BindJSON(&li); err != nil {
		println(err.Error())
		return
	}
	DATABASE.updateList(&li)
}

func deleteList(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	var li List

	if err := c.BindJSON(&li); err != nil {
		println(err.Error())
		return
	}
	DATABASE.deleteList(&li)
}

func changeListStatus(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	var li List

	if err := c.BindJSON(&li); err != nil {
		println(err.Error())
		return
	}

	DATABASE.changeListStatus(&li)
}
