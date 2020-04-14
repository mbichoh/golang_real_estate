package mysql

import (
	"database/sql"

	"github.com/mbichoh/real_estate/pkg/models"
)

type EstateModel struct{
	DB *sql.DB
}

//insert estates in db
func (m *EstateModel) Insert(agentID int, name string, address string, county string, shortDesc string, longDesc string, bedroom int, washroom int, spaceArea int, packing int, price float64)(int, error){
	
	stmt := `INSERT INTO estate (agent_id, name, address, county, price, bedroom, washroom, space_area, packing, short_desc, long_desc, created_at)
			  VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, UTC_TIMESTAMP())`;
			  
	result, err := m.DB.Exec(stmt, agentID, name, address, county, price, bedroom, washroom, spaceArea, packing, shortDesc, longDesc)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil{
		return 0, nil
	}

	return int(id), nil
}

//fetch single/specific estate using its id
func (m *EstateModel) Get(id int) (*models.Estate, error){
	return nil, nil
}

//fetch all estates limit(10) for pagination later
func (m *EstateModel) Latest() ([]*models.Estate, error) {
	return nil, nil
}
