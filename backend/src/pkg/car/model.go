package car

import "strconv"

type Model struct {
	FuelConsumption float64 `json:"fuel_consumption"`
	EngineSize      float64 `json:"engine_size"`
	Cylinders       int     `json:"cylinders"`
	CO2Emissions    float64 `json:"co2_emissions"`
}

func (c Model) PredictCO2Emissions() float64 {
	p, _ := r.Predict(
		[]float64{c.FuelConsumption, c.EngineSize, float64(c.Cylinders)},
	)

	return p
}

func CO2EmissionsFormula() string {
	return r.Formula
}

func (c *Model) SetCO2Emissions(co2Emissions float64) {
	c.CO2Emissions = co2Emissions
}

// index of the columns in the csv file
const (
	engine_size_idx = iota
	cylinders_idx
	fuel_consumption_idx
	co2_emissions_idx
)

func New(fuelConsumption, engineSize float64, cylinders int) Model {
	car := Model{
		FuelConsumption: fuelConsumption,
		EngineSize:      engineSize,
		Cylinders:       cylinders,
		CO2Emissions:    0,
	}

	return car
}

func NewFromRow(row []string) Model {
	fuelConsumption, _ := strconv.ParseFloat(row[fuel_consumption_idx], 64)
	engineSize, _ := strconv.ParseFloat(row[engine_size_idx], 64)
	cylinders, _ := strconv.Atoi(row[cylinders_idx])

	co2Emissions, _ := strconv.ParseFloat(row[co2_emissions_idx], 64)

	car := Model{
		Cylinders:       cylinders,
		EngineSize:      engineSize,
		FuelConsumption: fuelConsumption,
		CO2Emissions:    co2Emissions,
	}

	return car
}

func GetManyFromRows(rows [][]string, isHeaderRowIncluded bool) []Model {
	var cars []Model

	for i, row := range rows {
		if i == 0 && isHeaderRowIncluded {
			continue
		}

		car := NewFromRow(row)
		cars = append(cars, car)
	}

	return cars
}
