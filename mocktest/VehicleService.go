package mocktest

import "errors"

//Vehicle represents all kind of cars
type Vehicle struct {
	//Vin is a VehicleIdentificationNumber
	Vin string
}

//VehicleRepository manages Vehicles
type VehicleRepository interface {
	Get(vin string) (*Vehicle, error)
	Store(vehicle *Vehicle) (*Vehicle, error)
}

//NewVehicleService creates a new VehicleService with the given VehicleRepository
func NewVehicleService(repo VehicleRepository) VehicleService {
	return VehicleService{repo}
}

//VehicleService manages Vehicles
type VehicleService struct {
	repo VehicleRepository
}

//Create creates a new Vehicle in the data store
func (vs *VehicleService) Create(vehicle *Vehicle) (*Vehicle, error) {

	if vehicle.Vin == "" {
		return vehicle, errors.New("given Vehicle doesn't have a Vin")
	}

	return vs.repo.Store(vehicle)

}

//Get returns the vehicle with the given vin
func (vs *VehicleService) Get(vin string) (*Vehicle, error) {

	return vs.repo.Get(vin)

}
