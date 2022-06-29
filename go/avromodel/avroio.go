package avromodel

import "abeam/libs/model"

type modeler[U model.Event, V any] interface {
	FromModel(u U) (V, error)
	ToModel(v V) (U, error)
}
