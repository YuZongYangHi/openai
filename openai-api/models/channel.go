package models

import "time"

const TableNameChannel = "openai_channel"

type Channel struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (*Channel) TableName() string {
	return TableNameChannel
}

func (c *Channel) Create(title string) (*Channel, error) {
	channel := Channel{
		Title: title,
	}
	return &channel, db.Create(&channel).Error
}

func (c *Channel) List() (result *[]Channel) {
	db.Find(&result)
	return
}

func (c *Channel) Get(id int64) (result *Channel) {
	db.Where("id = ?", id).First(&result)
	return result
}

func (c *Channel) Update(id int64, title string) error {
	return db.Model(&Channel{}).Where("id = ?", id).Update("title", title).Error
}

func (c *Channel) Delete(id int64) error {
	var channel Channel
	return db.Where("id = ?", id).Delete(&channel).Error
}
