// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePerson = "people"

// Person mapped from table <people>
type Person struct {
	ID          int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name        string         `gorm:"column:name" json:"name"`
	Age         int32          `gorm:"column:age" json:"age"`
	Flag        bool           `gorm:"column:flag" json:"flag"`
	Commit      string         `gorm:"column:commit" json:"commit"`
	First       bool           `gorm:"column:First" json:"First"`
	FlagAnother int32          `gorm:"column:flag_another" json:"flag_another"`
	Bit         []byte         `gorm:"column:bit" json:"bit"`
	Small       int32          `gorm:"column:small" json:"small"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Score       float64        `gorm:"column:score" json:"score"`
	Type        int32          `gorm:"column:type" json:"type"`
	Birth       time.Time      `gorm:"column:birth" json:"birth"`
}

// TableName Person's table name
func (*Person) TableName() string {
	return TableNamePerson
}
