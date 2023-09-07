package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	Madr  = "madrugada (0 → 6)"
	Mania = "mañana (7 → 12)"
	Tarde = "tarde (13 → 19)"
	Noche = "noche (20 → 23)"
)

var (
	ErrtimeNoValid     = errors.New("Error time de intput no valido")
	ErrCountryNotFound = errors.New("No found country...")
)

type Ticket struct {
	Id              string
	Nombre          string
	Email           string
	País_de_destino string
	Hora_del_vuelo  string
	Precio          string
}

func ReadFIleCSV(file string) (datout []Ticket, errout error) {
	fd, err := os.Open(file)
	if err != nil {
		errout = err
		fmt.Println(err)
		return
	}
	//fmt.Println("Archivo abierto exitosamente")
	defer fd.Close()
	//leer el file con csv
	filereader := csv.NewReader(fd)
	info, err2 := filereader.ReadAll()
	if err2 != nil {
		errout = err2
		fmt.Println(err2)
		return datout, err2
	}

	for _, row := range info {
		valor := Ticket{
			Id:              row[0],
			Nombre:          row[1],
			Email:           row[2],
			País_de_destino: row[3],
			Hora_del_vuelo:  row[4],
			Precio:          row[5],
		}

		datout = append(datout, valor)

	}
	errout = nil
	return

}

// ejemplo 1
func GetTotalTickets(info []Ticket, destination string) (cant int, errout error) {
	cant = 0
	errout = nil

	for _, v := range info {
		if v.País_de_destino == destination {
			cant++
		}

	}
	if cant == 0 {
		errout = ErrCountryNotFound
	}
	return

}

// ejemplo 2
func GetMornings(info []Ticket, time string) (cant int, errout error) {
	cant = 0
	errout = nil

	if time == Madr {

		for _, v := range info {

			switch {
			case len(v.Hora_del_vuelo) == 4:

				hh, err := strconv.Atoi(v.Hora_del_vuelo[0:1])
				if err != nil {
					errout = err
					return

				}
				if hh <= 6 && hh >= 0 {
					cant++
				}

			//	fmt.Println("soy 4", v.Hora_del_vuelo[0:1], hh)
			//fmt.Println("soy 4")
			case len(v.Hora_del_vuelo) == 5:
				hh, err := strconv.Atoi(v.Hora_del_vuelo[0:2])
				if err != nil {
					errout = err
					return

				}
				if hh <= 6 && hh >= 0 {
					cant++
				}
			//	fmt.Println("soy 5", v.Hora_del_vuelo[0:2], hh)
			default:
				//	fmt.Println("hora no valida", v.Hora_del_vuelo)

			}

		}
	}

	return
}

func GetMornings2(info []Ticket, time string) (cant int, errout error) {
	cant = 0
	errout = nil

	if time == Mania {

		for _, v := range info {

			switch {
			case len(v.Hora_del_vuelo) == 4:

				hh, err := strconv.Atoi(v.Hora_del_vuelo[0:1])
				if err != nil {
					errout = err
					return

				}
				if hh <= 12 && hh >= 7 {
					cant++
				}

			//	fmt.Println("soy 4", v.Hora_del_vuelo[0:1], hh)
			//fmt.Println("soy 4")
			case len(v.Hora_del_vuelo) == 5:
				hh, err := strconv.Atoi(v.Hora_del_vuelo[0:2])
				if err != nil {
					errout = err
					return

				}
				if hh <= 12 && hh >= 7 {
					cant++
				}
			//	fmt.Println("soy 5", v.Hora_del_vuelo[0:2], hh)
			default:
				//	fmt.Println("hora no valida", v.Hora_del_vuelo)

			}

		}
	}

	return
}
func GetAfernoon(info []Ticket, time string) (cant int, errout error) {
	cant = 0
	errout = nil

	if time == Tarde {

		for _, v := range info {

			switch {
			case len(v.Hora_del_vuelo) == 4:

				hh, err := strconv.Atoi(v.Hora_del_vuelo[0:1])
				if err != nil {
					errout = err
					return

				}
				if hh <= 19 && hh >= 13 {
					cant++
				}

			//	fmt.Println("soy 4", v.Hora_del_vuelo[0:1], hh)
			//fmt.Println("soy 4")
			case len(v.Hora_del_vuelo) == 5:
				hh, err := strconv.Atoi(v.Hora_del_vuelo[0:2])
				if err != nil {
					errout = err
					return

				}
				if hh <= 19 && hh >= 13 {
					cant++
				}
			//	fmt.Println("soy 5", v.Hora_del_vuelo[0:2], hh)
			default:
				//	fmt.Println("hora no valida", v.Hora_del_vuelo)

			}

		}
	}

	return

}
func GetNight(info []Ticket, time string) (cant int, errout error) {
	cant = 0
	errout = nil

	if time == Noche {

		for _, v := range info {

			switch {
			case len(v.Hora_del_vuelo) == 4:

				hh, err := strconv.Atoi(v.Hora_del_vuelo[0:1])
				if err != nil {
					errout = err
					return

				}
				if hh <= 23 && hh >= 20 {
					cant++
				}

			//	fmt.Println("soy 4", v.Hora_del_vuelo[0:1], hh)
			//fmt.Println("soy 4")
			case len(v.Hora_del_vuelo) == 5:
				hh, err := strconv.Atoi(v.Hora_del_vuelo[0:2])
				if err != nil {
					errout = err
					return

				}
				if hh <= 23 && hh >= 20 {
					cant++
				}
			//	fmt.Println("soy 5", v.Hora_del_vuelo[0:2], hh)
			default:
				//	fmt.Println("hora no valida", v.Hora_del_vuelo)

			}

		}
	}

	return

}

func GetCountByPeriod(info []Ticket, time string) (cant int, errout error) {

	switch time {
	case Madr:
		cant, errout = GetMornings(info, Madr)
	case Mania:
		cant, errout = GetMornings2(info, Mania)
	case Tarde:
		cant, errout = GetAfernoon(info, Tarde)
	case Noche:
		cant, errout = GetNight(info, Noche)
	default:
		cant = 0
		errout = ErrtimeNoValid
	}
	return
}

// ejemplo 3
func AverageDestination(info []Ticket, destination string, total int) (out float64, errout error) {
	out = 0.0
	errout = nil
	if total <= 0 {

		//despues hacer mas lindos estos errores
		errout = errors.New("Total invalido no se puede dividir por numeros menores o iguales a cero..")
		return
	}

	cant, err := GetTotalTickets(info, destination)
	if err != nil {
		errout = err
		return
	}

	out = float64(cant) / float64(total)
	return

}
