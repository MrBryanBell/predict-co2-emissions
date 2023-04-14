package main

import (
	"github.com/MrBryanBell/predict-co2-emissions/backend/src/pkg/car"
	"github.com/gin-gonic/gin"
)

func main() {
	car.InitCo2EmissionsModel()
	r := gin.Default()

	r.GET("/predict/co2-emissions", func(ctx *gin.Context) {
		var car car.Model

		if err := ctx.ShouldBindJSON(&car); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		co2Prediction := car.PredictCO2Emissions()
		car.SetCO2Emissions(co2Prediction)

		ctx.JSON(200, car)
	})

	r.Run()
}
