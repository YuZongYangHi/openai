package models

import "time"

const TableNameMessage = "openai_message"

var MessageType = map[int]string{
	0: "robot",
	1: "user",
}

type Message struct {
	ID          int64     `json:"id"`
	ChannelId   int64     `json:"channelId" gorm:"column:channel_id"`
	Content     string    `json:"content" gorm:"size(256)"`
	DialogType  int       `json:"dialogType" gorm:"column:dialog_type"`
	ContentType int       `json:"contentType" gorm:"column:content_type"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:created_at"`
}

func (m *Message) IsValid() bool {
	if len(m.Content) == 0 {
		return false
	}
	return true
}

func (*Message) TableName() string {
	return TableNameMessage
}

func (m *Message) Add(message *Message) (*Message, error) {
	return message, db.Create(message).Error
}

func (m *Message) List(channelId int64) (result *[]Message) {
	db.Where("channel_id = ?", channelId).Order("created_at").Find(&result)
	return
}

func (m *Message) Delete(channelId int64) error {
	var message Message
	return db.Where("channel_id = ?", channelId).Delete(&message).Error
}
