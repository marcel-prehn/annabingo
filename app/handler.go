package app

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BingoHandler interface {
	HandleGetBingo(c *gin.Context)
	HandleGetTestBingo(c *gin.Context)
	HandleGetBingoById(c *gin.Context)
	HandlePostBingo(c *gin.Context)
	HandleSearch(c *gin.Context)
	HandleGetStatistics(c *gin.Context)
	HandleCreateIndex(c *gin.Context)
}

type bingoHandler struct {
	service BingoService
	logger  *zap.Logger
}

func NewBingoHandler(service BingoService, logger *zap.Logger) BingoHandler {
	return &bingoHandler{
		service: service,
		logger:  logger,
	}
}

func (h *bingoHandler) HandleGetBingo(c *gin.Context) {
	c.JSON(200, TestData())
}

func (h *bingoHandler) HandleGetTestBingo(c *gin.Context) {
	bingo := ApiData()
	bingo.Fields = *h.service.Shuffle(bingo.Fields)
	c.JSON(200, bingo)
}

func (h *bingoHandler) HandleGetBingoById(c *gin.Context) {
	id := c.Param("id")
	bingo, err := h.service.GetBingoById(id)
	if err != nil {
		h.logger.Error("bingo not found", zap.String("id", id), zap.Error(err))
		c.Status(500)
	}
	bingo.Fields = *h.service.Shuffle(bingo.Fields)
	c.JSON(200, bingo)
}

func (h *bingoHandler) HandlePostBingo(c *gin.Context) {
	var matrix Bingo
	_ = c.BindJSON(&matrix)
	id, err := h.service.SaveBingo(matrix)
	if err != nil {
		h.logger.Error("bingo not saved", zap.String("id", id), zap.Error(err))
		c.Status(500)
	}
	c.String(201, id)
}

func (h *bingoHandler) HandleSearch(c *gin.Context) {
	query := c.Param("query")
	result, err := h.service.SearchBingoByTitle(query)
	if err != nil {
		h.logger.Error("error during bingo search", zap.String("query", query), zap.Error(err))
		c.Status(500)
	}
	if *result == nil && len(*result) == 0 {
		h.logger.Error("bingo not found", zap.String("query", query), zap.Error(err))
		c.Status(404)
	}
	c.JSON(200, result)
}

func (h *bingoHandler) HandleGetStatistics(c *gin.Context) {
	count, err := h.service.Count()
	if err != nil {
		h.logger.Error("stats not available", zap.Error(err))
		c.Status(500)
	}
	stats := Stats{Count: count}
	c.JSON(200, stats)
}

func (h *bingoHandler) HandleCreateIndex(c *gin.Context) {
	err := h.service.CreateIndexOnTitle()
	if err != nil {
		c.Status(500)
	}
	c.Status(200)
}
