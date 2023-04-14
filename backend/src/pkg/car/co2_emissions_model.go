package car

import (
	"github.com/MrBryanBell/predict-co2-emissions/backend/src/pkg/file"
	"github.com/sajari/regression"
)

var r *regression.Regression

func InitCo2EmissionsModel() {
	// Open the file
	filePath := "src/datasets/model_features.csv"
	csvFile := file.OpenCSV(filePath)
	defer csvFile.Close()

	// Read the file
	rows := file.ReadCSV(csvFile)
	cars := GetManyFromRows(rows, true)

	// Create an instance of a new regression.
	r = new(regression.Regression)

	// Set the dependent and independent variables.
	r.SetObserved("CO2_Emissions")
	r.SetVar(0, "Fuel_Consumption")
	r.SetVar(1, "Engine_Size")
	r.SetVar(2, "Cylinders")

	// Train the model.
	for _, car := range cars {
		r.Train(
			regression.DataPoint(
				car.CO2Emissions,
				[]float64{
					car.FuelConsumption,
					car.EngineSize,
					float64(car.Cylinders),
				},
			),
		)
	}

	r.Run()
}
