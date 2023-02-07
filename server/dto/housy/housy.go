package housydto

type HouseRequest struct {
	Name      string `json:"name" gorm:"type: varchar(255)"`
	CityID    int    `json:"city_id" form:"city_id"`
	Thumbnail string `json:"thumbnail" form:"thumbnail"`
	Image     string `json:"image" form:"image"`
	Address   string `json:"address" gorm:"type: text"`
	Price     int    `json:"price" gorm:"type: varchar(255)"`
	TypeRent  string `json:"typerent" gorm:"type: varchar(255)"`
	Amenities string `json:"amenities"`
	Bedroom   int    `json:"bedroom"`
	Bathroom  int    `json:"bathroom"`
}
