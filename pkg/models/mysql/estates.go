package mysql

import (
	"database/sql"

	"github.com/mbichoh/real_estate/pkg/models"
)

type EstateModel struct{
	DB *sql.DB
}

//insert estates in db
func (m *EstateModel) Insert(agentID int, name string, address string, county string, shortDesc string, longDesc string, bedroom int, washroom int, spaceArea int, packing int, price bool)(int, error){
	return 0, nil
}

//fetch single/specific estate using its id
func (m *EstateModel) Get(id int) (*models.Estate, error){
	return nil, nil
}

//fetch all estates limit(10) for pagination later
func (m *EstateModel) Latest() ([]*models.Estate, error) {
	return nil, nil
}
