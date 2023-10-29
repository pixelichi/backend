package user

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"shinypothos.com/api/common/request_context"
	"shinypothos.com/api/common/request_util"
	"shinypothos.com/api/data/db_txn"
	"shinypothos.com/api/data/db_txn/db_util"
	db "shinypothos.com/db/sqlc"
	"shinypothos.com/token"
)

type addPlantRequest struct {
	PlantName string `json:"plant_name" binding:"required"`
	Species   string `json:"species" binding:"required"`
}

type addPlantResponse struct {
	ID        int64  `json:"id"`
	PlantName string `json:"plant_name"`
	Species   string `json:"species"`
	CreatedAt string `json:"created_at"`
}

func AddPlant(c *gin.Context) {

	rc, err := request_context.GetReqCtxOrInternalServerError(c)
	if err != nil {
		return
	}

	req, err := request_util.BindJSONOrAbort[addPlantRequest](c, &addPlantRequest{})
	if err != nil {
		return
	}

	tokenPayload, err := token.GetPayloadOrAbort(c)
	if err != nil {
		return
	}

	createPlantParams := db.CreatePlantParams{
		UserID: sql.NullInt64{
			Int64: tokenPayload.UserID,
			Valid: tokenPayload.UserID != 0,
		},

		PlantName: req.PlantName,

		Species: sql.NullString{
			String: req.Species,
			Valid:  len(req.Species) != 0,
		},
	}

	plant, err := db_txn.AddPlantOrAbort(c, rc.DB, &createPlantParams)
	if err != nil {
		return
	}

	response := addPlantResponse{
		ID:        plant.ID,
		PlantName: plant.PlantName,
		Species:   db_util.NullStringToString(plant.Species),
		CreatedAt: db_util.TimeToString(plant.CreatedAt),
	}

	c.JSON(http.StatusOK, response)
}
