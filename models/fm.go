package models

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

type FM struct {
	NumFactors    int         `json:"num_factors"`
	NumAttributes int         `json:"num_attributes"`
	K0            bool        `json:"k0"`
	K1            bool        `json:"k1"`
	W0            float64     `json:"w0"`
	W             []float64   `json:"w"` // len(W) == NumAttributes > 0
	V             [][]float64 `json:"v"` // len(V) == NumFactor > 0, len(V[0]) == NumAttributes > 0
}

func (m *FM) Predict(input []float64) (ret float64, err error) {
	inputSize := len(input)
	if inputSize < m.NumAttributes {
		err = fmt.Errorf("invalid input size: %d, expected size is %d", inputSize, m.NumAttributes)
		return
	}
	if m.K0 {
		ret += m.W0
	}
	if m.K1 {
		for i := 0; i < inputSize; i++ {
			ret += m.W[i] * input[i]
		}
	}
	var d, sum, sumSqr float64
	for i := 0; i < m.NumFactors; i++ {
		sum = 0.0
		sumSqr = 0.0
		for j := 0; j < inputSize; j++ {
			d = m.V[i][j] * input[j]
			sum += d
			sumSqr += d * d
		}
		ret += 0.5 * (sum*sum - sumSqr)
	}
	ret = 1.0 / (1.0 + math.Exp(-ret))
	return
}
func (m *FM) LoadModelFromJson(data []byte) error {
	err := json.Unmarshal(data, m)
	if err != nil {
		return err
	}
	if m.NumFactors <= 0 ||
		m.NumAttributes <= 0 ||
		m.NumAttributes != len(m.W) ||
		m.NumAttributes != len(m.V[0]) ||
		m.NumFactors != len(m.V) {
		return fmt.Errorf("invalid model")
	}
	return nil
}
func (m *FM) LoadModelFromJsonFile(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return m.LoadModelFromJson(content)
}
