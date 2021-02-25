package mocktest_test

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"golang.source-fellows.com/mocktest"
	"golang.source-fellows.com/mocktest/mocks"
)

//go:generate mockgen -source=VehicleService.go -destination=mocks/VehicleRepository.go -package mocks

type dummyRepo struct {
	vehicle    *mocktest.Vehicle
	storeCount int
}

func (dr *dummyRepo) Store(vehicle *mocktest.Vehicle) (*mocktest.Vehicle, error) {
	dr.vehicle = vehicle
	dr.storeCount++
	return vehicle, nil
}

func (dr *dummyRepo) Get(string) (*mocktest.Vehicle, error) {
	return dr.vehicle, nil
}

func TestState(t *testing.T) {

	//given
	vehicleService := mocktest.NewVehicleService(&dummyRepo{})

	dummyVIN := "12345678901234567"
	vehicleToSave := &mocktest.Vehicle{Vin: dummyVIN}

	//when
	vehicle, _ := vehicleService.Create(vehicleToSave)

	//then

	//check errors
	readVehicle, _ := vehicleService.Get(dummyVIN)

	//check errors
	if readVehicle == nil {
		t.Errorf("no vehicle found")
		return
	}

	//compare
	if !reflect.DeepEqual(vehicle, readVehicle) {
		t.Errorf("vehicles differ")
		return
	}
}

func TestBehaviour(t *testing.T) {

	//given
	repo := &dummyRepo{storeCount: 0}
	vehicleService := mocktest.NewVehicleService(repo)

	dummyVIN := "12345678901234567"
	vehicleToSave := &mocktest.Vehicle{Vin: dummyVIN}

	//when
	_, _ = vehicleService.Create(vehicleToSave)

	//then
	if repo.storeCount != 1 {
		t.Error("repo should be called at least once")
		return
	}

}

func TestSave(t *testing.T) {

	//given
	ctrl := gomock.NewController(t)

	repo := mocks.NewMockVehicleRepository(ctrl)
	calcService := mocktest.NewVehicleService(repo)

	vehicle := &mocktest.Vehicle{"12345678901234567"}

	repo.EXPECT().Store(gomock.Any()).Return(vehicle, nil)

	//when
	v, err := calcService.Create(vehicle)

	//then
	if err != nil {
		t.Errorf("could not create because of %v", err)
		return
	}

	if v == nil {
		t.Errorf("no vehicle returned")
		return
	}

}
