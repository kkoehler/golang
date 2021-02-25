package mongodb

import "golang.source-fellows.com/mocktest"

type VehicleRepository struct {
}

func (vr *VehicleRepository) Store(*mocktest.Vehicle) (*mocktest.Vehicle, error) {
	return nil, nil
}

func (vr *VehicleRepository) Get(string) (*mocktest.Vehicle, error) {
	return nil, nil
}
