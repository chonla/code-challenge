package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

type dataSet struct {
	amount          float64
	paid            float64
	expectedChanges [][]float64
}

func main() {
	data := []dataSet{
		dataSet{
			amount: 17.75,
			paid:   1000,
			expectedChanges: [][]float64{
				[]float64{500, 1},
				[]float64{100, 4},
				[]float64{50, 1},
				[]float64{20, 1},
				[]float64{10, 1},
				[]float64{2, 1},
				[]float64{0.25, 1},
			},
		},
		dataSet{
			amount: 17.75,
			paid:   4000,
			expectedChanges: [][]float64{
				[]float64{1000, 3},
				[]float64{500, 1},
				[]float64{100, 4},
				[]float64{50, 1},
				[]float64{20, 1},
				[]float64{10, 1},
				[]float64{2, 1},
				[]float64{0.25, 1},
			},
		},
	}

	for _, v := range data {
		res, _ := getChanges(v.amount, v.paid)
		if reflect.DeepEqual(res, v.expectedChanges) {
			fmt.Println("Cool!")
		} else {
			fmt.Printf("expecting result to be [%v] but [%v]\n", v.expectedChanges, res)
		}
	}
}

func getChanges(amt, cash float64) ([][]float64, error) {
	a := int(amt * 100.0)
	c := int(cash * 100.0)

	changes := c - a
	if changes < 0 {
		return nil, errors.New("not enough cash")
	}

	avail := []int{100000, 50000, 10000, 5000, 2000, 1000, 500, 200, 100, 50, 25}
	result := [][]float64{}

	for cursor := 0; changes > 0 && cursor < len(avail); cursor++ {
		if changes >= avail[cursor] {
			v := avail[cursor]
			va := math.Floor(float64(changes / v))
			changes = changes % v
			result = append(result, []float64{float64(v) / 100.0, va})
		}
	}

	return result, nil
}
