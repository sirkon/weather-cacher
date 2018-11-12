/*
rawclient предоставляет базовую абстракцию и реализации т.н. "rawclient"-ов, т.е. объектов инкапсулирующих транспортный
уровень
*/

package rawclient

import (
	"context"
	"io"
)

// RawClient абстрация закрывающая подробности работы на транспортном уровне
type RawClient interface {
	Get(ctx context.Context, lat, log float64) (io.ReadCloser, error)
}
