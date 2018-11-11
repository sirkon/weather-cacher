package weather

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"math"
)

// IDGenerator генерация идентификатора прогноза по данному местоположению
type IDGenerator interface {
	ID(provID string, lat, lon float64, forecast *Forecast) (string, error)
}

// NewIDGenerator конструктор генератора идентификаторов с данным генератором хешей
func NewIDGenerator(h hash.Hash) IDGenerator {
	return &idGenImpl{
		hash: h,
	}
}

// idGenImpl генератор идентификаторов прогноза в данной точке
type idGenImpl struct {
	hash hash.Hash
}

// ID генерация идентификатора прогноза погоды для данного местоположения
func (ig *idGenImpl) ID(provID string, lat, lon float64, forecast *Forecast) (string, error) {
	ig.hash.Reset()

	io.WriteString(ig.hash, provID)

	var floatBuf [8]byte
	binary.BigEndian.PutUint64(floatBuf[:], math.Float64bits(lat))
	ig.hash.Write(floatBuf[:])
	binary.BigEndian.PutUint64(floatBuf[:], math.Float64bits(lon))
	ig.hash.Write(floatBuf[:])

	encoder := json.NewEncoder(ig.hash)
	if err := encoder.Encode(forecast); err != nil {
		return "", fmt.Errorf("failed to convert forecast when building forecast ID: %s", err)
	}

	return hex.EncodeToString(ig.hash.Sum(nil)), nil
}
