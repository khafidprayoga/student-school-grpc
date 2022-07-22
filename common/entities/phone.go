package entities

type PhoneNumber struct {
	StudentId  uint32 `gorm:"primaryKey"`
	Number     uint32 `gorm:"unique"`
	NumberType uint32 // 0 cellular, 1 home
}

func (p *PhoneNumber) TableName() string {
	return "phone_number"
}
