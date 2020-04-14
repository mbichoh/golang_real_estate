package mysql

import (
	"database/sql"

	"github.com/mbichoh/real_estate/pkg/models"
)

type EstateModel struct {
	DB *sql.DB
}

//insert estates in db
func (m *EstateModel) Insert(agentID int, name string, address string, county string, shortDesc string, longDesc string, bedroom int, washroom int, spaceArea int, packing int, price float64) (int, error) {

	stmt := `INSERT INTO estate (agent_id, name, address, county, price, bedroom, washroom, space_area, packing, short_desc, long_desc, created_at)
			  VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, UTC_TIMESTAMP())`

	result, err := m.DB.Exec(stmt, agentID, name, address, county, price, bedroom, washroom, spaceArea, packing, shortDesc, longDesc)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}

//fetch single/specific estate using its id
func (m *EstateModel) Get(id int) (*models.Estate, error) {

	e := &models.Estate{}

	stmt := `SELECT name, address, county, price FROM estate WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	err := row.Scan(&e.Name, &e.Address, &e.County, &e.Price)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return e, nil
}

//fetch all estates limit(10) for pagination later
func (m *EstateModel) Latest() ([]*models.Estate, error) {

	stmt := `SELECT id, name, address, county, bedroom, washroom, price, space_area, short_desc FROM estate`

	rows, err := m.DB.Query(stmt)
	if err != nil{
		return nil, err
	}

	defer rows.Close()

	estates := []*models.Estate{}

	for rows.Next() {
		e := &models.Estate{}

		err = rows.Scan(&e.ID, &e.Name, &e.Address, &e.County, &e.Bedroom, &e.Washroom, &e.Price, &e.SpaceArea, &e.ShortDesc)
		if err != nil{
			return nil, err
		}
		
		estates = append(estates, e)
	}

	if err = rows.Err(); err != nil{
		return nil, err
	}

	return estates, nil
}
