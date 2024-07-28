package models

type Item struct {
	Id      int     `db:"id"`
	Name    string  `db:"name"`
	Price   float64 `db:"price"`
	OwnerId int     `db:"owner_id"`
}

type NewItem struct {
	Name    string  `db:"name"`
	Price   float64 `db:"price"`
	OwnerId int     `db:"owner_id"`
}
