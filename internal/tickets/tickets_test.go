package tickets

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var Dataemu = []Ticket{
	{Id: "994",
		Nombre:          "Alleen Cramer",
		Email:           "acramerrl@ehow.com",
		País_de_destino: "Argentina",
		Hora_del_vuelo:  "19:38",
		Precio:          "628"},
	{
		Id:              "995",
		Nombre:          "Daphene Snoxill",
		Email:           "dsnoxillrm@macromedia.com",
		País_de_destino: "Belarus",
		Hora_del_vuelo:  "0:01",
		Precio:          "1465",
	},
	{Id: "996",
		Nombre:          "Rolland Milam",
		Email:           "rmilamrn@nifty.com",
		País_de_destino: "Czech Republic",
		Hora_del_vuelo:  "16:02",
		Precio:          "737",
	},
	{Id: "997",
		Nombre:          "Agustin Lyddiatt",
		Email:           "alyddiattro@github.com",
		País_de_destino: "Kazakhstan",
		Hora_del_vuelo:  "8:31",
		Precio:          "569",
	},
	{
		Id:              "998",
		Nombre:          "Gretchen Skurray",
		Email:           "gskurrayrp@ask.com",
		País_de_destino: "Indonesia",
		Hora_del_vuelo:  "0:30",
		Precio:          "834",
	},
	{Id: "999",
		Nombre:          "Stacy Fallon",
		Email:           "sfallonrq@etsy.com",
		País_de_destino: "Indonesia",
		Hora_del_vuelo:  "20:44",
		Precio:          "550",
	},
	{Id: "1000",
		Nombre:          "Magdalen Covelle",
		Email:           "mcovellerr@wired.com",
		País_de_destino: "Jamaica",
		Hora_del_vuelo:  "23:20",
		Precio:          "563"}}

func TestGetTotalTickets(t *testing.T) {
	t.Run("TEST caso 1 camino feliz", func(t *testing.T) {

		//Arrange
		country := "Indonesia"

		valesperado1 := 2
		//Action
		val, _ := GetTotalTickets(Dataemu, country)
		//Assert
		assert.Equal(t, valesperado1, val, "Fallo el test")

	})
	t.Run("TestGetTotalTickets Error", func(t *testing.T) {

		//Arrange
		country := "Peru"

		//valesperado1 := 2
		//Action
		_, err := GetTotalTickets(Dataemu, country)
		//Assert
		assert.Equal(t, errors.New("No found country..."), err, "Fallo el test")
		assert.Error(t, err, "Fallo el error")
	})

}

func TestCalcAverage(t *testing.T) {
	t.Run("TestCalcAverage caso 1 camino feliz", func(t *testing.T) {

		//Arrange
		country := "Indonesia"

		//Action
		val6, _ := AverageDestination(Dataemu, country, len(Dataemu))

		valesperado1 := 2.0 / float64(len(Dataemu))
		//Assert
		assert.Equal(t, valesperado1, val6, "Fallo el test")

	})
	t.Run("TestCalcAverage Error", func(t *testing.T) {

		//Arrange
		country := "Misiones"

		//Action
		_, err := AverageDestination(Dataemu, country, len(Dataemu))

		//Assert
		assert.Equal(t, errors.New("No found country..."), err, "Fallo el test")
		assert.Error(t, err, "Fallo el error")

	})

}

func TestGetCountByPeriod(t *testing.T) {

	t.Run("TestGetCountByPeriod caso 1 camino feliz", func(t *testing.T) {

		//Arrange

		//Action
		val4, _ := GetCountByPeriod(Dataemu, Noche)

		valesperado1 := 2
		//Assert
		assert.Equal(t, valesperado1, val4, "Fallo el test")

	})
	t.Run("TestGetCountByPeriod Error", func(t *testing.T) {

		//Arrange

		//Action
		_, err := GetCountByPeriod(Dataemu, "siesta")

		//Assert
		assert.Equal(t, errors.New("Error time de intput no valido"), err, "Fallo el test")
		assert.Error(t, err, "Fallo el error")

	})

}
