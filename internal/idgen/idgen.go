package idgen

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/sirkon/weather-cacher/internal/schema"
	"hash"
	"io"
	"math"
)

// New конструктор генератора идентификаторов с данным генератором хешей
func New(h hash.Hash) *IDGen {
	return &IDGen{
		hash: h,
	}
}

// IDGen генератор идентификаторов прогноза в данной точке
type IDGen struct {
	hash hash.Hash
}

// ID генерация идентификатора прогноза погоды для данного местоположения
func (ig *IDGen) ID(provID string, lat, lon float64, forecast *schema.Forecast) (string, error) {
	ig.hash.Reset()

	io.WriteString(ig.hash, provID)

	var floatBuf [8]byte
	binary.BigEndian.PutUint64(floatBuf[:], math.Float64bits(lat))
	ig.hash.Write(floatBuf[:])
	binary.BigEndian.PutUint64(floatBuf[:], math.Float64bits(lon))
	ig.hash.Write(floatBuf[:])

	marshaler := jsonpb.Marshaler{}
	if err := marshaler.Marshal(ig.hash, forecast); err != nil {
		return "", fmt.Errorf("failed to convert forecast when building forecast ID: %s", err)
	}

	return hex.EncodeToString(ig.hash.Sum(nil)), nil
}
