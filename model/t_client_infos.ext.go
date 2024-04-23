package model

type TClientInfoService interface {
	Insert(client TClientInfo) bool
}

func (c *TClientInfo) Insert() bool {
	if err := DB.Create(&c).Error; err != nil {
		return false
	}
	return true
}
