package table

type Person struct {
	Id         int64  `gorm:"column:id;primary_key"`
	Name       string `gorm:"column:name"`
	Image      string `gorm:"column:image"`
	ThumbImage string `gorm:"column:thumb_image"`
	Lucky      bool
	Level      int64
}

// TableName sets the insert table name for this struct type
func (m *Person) TableName() string {
	return "t_person"
}
