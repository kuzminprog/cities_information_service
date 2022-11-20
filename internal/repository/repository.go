package repository

import cities "github.com/kuzminprog/cities_information_service"

type CityList interface {
	Create(city cities.CityRequest) (string, error)
	Delete(id int) error
	SetPopulation(id, population int) error
	GetFromRegion(region string) ([]string, error)
	GetFromDistrict(district string) ([]string, error)
	GetFromPopulation(population string) ([]string, error)
	GetFromFoundation(foundation string) ([]string, error)
	findCities(idx_type int, searchText string) []string
	GetFull(id int) (*cities.City, error)
}

type Repository struct {
	CityList
}

func NewRepository(db *DataBase) *Repository {
	return &Repository{
		CityList: NewCityListDB(db),
	}
}