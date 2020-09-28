package app

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BingoHandler struct {
	service *BingoService
	logger  *zap.Logger
}

func (h *BingoHandler) HandleGetBingo(c *gin.Context) {
	c.JSON(200, TestData())
}

func (h *BingoHandler) HandleGetTestBingo(c *gin.Context) {
	bingo := ApiData()
	bingo.Fields = *h.service.Shuffle(bingo.Fields)
	c.JSON(200, bingo)
}

func (h *BingoHandler) HandleGetBingoById(c *gin.Context) {
	id := c.Param("id")
	bingo, err := h.service.GetBingoById(id)
	if err != nil {
		c.Status(500)
	}
	bingo.Fields = *h.service.Shuffle(bingo.Fields)
	c.JSON(200, bingo)
}

func (h *BingoHandler) HandlePostBingo(c *gin.Context) {
	var matrix Bingo
	_ = c.BindJSON(&matrix)
	id, err := h.service.SaveBingo(matrix)
	if err != nil {
		c.Status(500)
	}
	c.String(201, id)
}

func (h *BingoHandler) HandleSearch(c *gin.Context) {
	query := c.Param("query")
	result, err := h.service.SearchBingoByTitle(query)
	if err != nil {
		c.Status(500)
	}
	if *result == nil && len(*result) == 0 {
		c.Status(404)
	}
	c.JSON(200, result)
}

func (h *BingoHandler) HandleGetStatistics(c *gin.Context) {
	count, err := h.service.Count()
	if err != nil {
		c.Status(500)
	}
	stats := Stats{Count: count}
	c.JSON(200, stats)
}

func (h *BingoHandler) HandleCreateIndex(c *gin.Context) {
	err := h.service.CreateIndexOnTitle()
	if err != nil {
		c.Status(500)
	}
	c.Status(200)
}
