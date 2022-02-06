package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type BingoHandler interface {
	GetBingo(c echo.Context) error
	GetTestBingo(c echo.Context) error
	GetBingoById(c echo.Context) error
	PostBingo(c echo.Context) error
	GetSearch(c echo.Context) error
	GetStatistics(c echo.Context) error
	PostCreateIndex(c echo.Context) error
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

func (h *bingoHandler) GetBingo(c echo.Context) error {
	return c.JSON(http.StatusOK, TestData())
}

func (h *bingoHandler) GetTestBingo(c echo.Context) error {
	bingo := ApiData()
	bingo.Fields = *h.service.Shuffle(bingo.Fields)
	return c.JSON(http.StatusOK, bingo)
}

func (h *bingoHandler) GetBingoById(c echo.Context) error {
	id := c.Param("id")
	bingo, err := h.service.GetBingoById(id)
	if err != nil {
		h.logger.Error("bingo not found", zap.String("id", id), zap.Error(err))
		return c.NoContent(http.StatusNotFound)
	}
	bingo.Fields = *h.service.Shuffle(bingo.Fields)
	return c.JSON(http.StatusOK, bingo)
}

func (h *bingoHandler) PostBingo(c echo.Context) error {
	var matrix Bingo
	err := c.Bind(&matrix)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	id, err := h.service.SaveBingo(matrix)
	if err != nil {
		h.logger.Error("bingo not saved", zap.String("id", id), zap.Error(err))
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.String(http.StatusCreated, id)
}

func (h *bingoHandler) GetSearch(c echo.Context) error {
	query := c.Param("query")
	result, err := h.service.SearchBingoByTitle(query)
	if err != nil {
		h.logger.Error("error during bingo search", zap.String("query", query), zap.Error(err))
		return c.NoContent(http.StatusInternalServerError)
	}
	if *result == nil && len(*result) == 0 {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, result)
}

func (h *bingoHandler) GetStatistics(c echo.Context) error {
	count, err := h.service.Count()
	if err != nil {
		h.logger.Error("stats not available", zap.Error(err))
		return c.NoContent(http.StatusInternalServerError)
	}
	stats := Stats{Count: count}
	return c.JSON(http.StatusOK, stats)
}

func (h *bingoHandler) PostCreateIndex(c echo.Context) error {
	err := h.service.CreateIndexOnTitle()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}
