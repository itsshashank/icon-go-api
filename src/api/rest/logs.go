package rest

import (
	"encoding/json"

	fiber "github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/sudoblockio/icon-go-api/config"
	"github.com/sudoblockio/icon-go-api/crud"
)

type LogsQuery struct {
	Limit int `query:"limit"`
	Skip  int `query:"skip"`

	BlockNumber     uint32 `query:"block_number"`
	TransactionHash string `query:"transaction_hash"`
	ScoreAddress    string `query:"score_address"`
	Method          string `query:"method"`
}

func LogsAddHandlers(app *fiber.App) {

	prefix := config.Config.RestPrefix + "/logs"

	app.Get(prefix+"/", handlerGetLogs)
}

// Logs
// @Summary Get Logs
// @Description get historical logs
// @Tags Logs
// @BasePath /api/v1
// @Accept */*
// @Produce json
// @Param limit query int false "amount of records"
// @Param skip query int false "skip to a record"
// @Param block_number query int false "skip to a record"
// @Param transaction_hash query string false "find by transaction hash"
// @Param score_address query string false "find by score address"
// @Param method query string false "find by method"
// @Router /api/v1/logs [get]
// @Success 200 {object} []models.Log
// @Failure 422 {object} map[string]interface{}
func handlerGetLogs(c *fiber.Ctx) error {
	params := new(LogsQuery)
	if err := c.QueryParser(params); err != nil {
		zap.S().Warnf("Logs Get Handler ERROR: %s", err.Error())

		c.Status(422)
		return c.SendString(`{"error": "could not parse query parameters"}`)
	}

	// Default Params
	if params.Limit <= 0 {
		params.Limit = 25
	}

	// Check Params
	if params.Limit < 1 || params.Limit > config.Config.MaxPageSize {
		c.Status(422)
		return c.SendString(`{"error": "limit must be greater than 0 and less than 101"}`)
	}
	if params.Skip < 0 || params.Skip > config.Config.MaxPageSkip {
		c.Status(422)
		return c.SendString(`{"error": "invalid skip"}`)
	}

	// Get Logs
	logs, err := crud.GetLogCrud().SelectMany(
		params.Limit,
		params.Skip,
		params.BlockNumber,
		params.TransactionHash,
		params.ScoreAddress,
		params.Method,
	)
	if err != nil {
		zap.S().Warnf("Logs CRUD ERROR: %s", err.Error())
		c.Status(500)
		return c.SendString(`{"error": "could not retrieve logs"}`)
	}

	if len(*logs) == 0 {
		// No Content
		c.Status(204)
	}

	// Set X-TOTAL-COUNT
	// TODO

	body, _ := json.Marshal(logs)
	return c.SendString(string(body))
}
