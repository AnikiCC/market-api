package models

type Deal struct {
	Id    int     `db:"id"`
	Item  Item    `db:"item"`
	User  User    `db:"user"`
	Price float64 `db:"price"`
}

type NewDeal struct {
	Item  Item    `db:"item"`
	User  User    `db:"user"`
	Price float64 `db:"price"`
}
