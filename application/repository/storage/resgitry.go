package storage

import "challenge/domain/model"

type RegistryStorage interface {
	CreateRegistryStorage(registry model.Registry) error
}
