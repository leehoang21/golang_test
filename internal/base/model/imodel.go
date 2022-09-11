package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type IModel interface {
	BeforeCreate(prefix string)
	BeforeUpdate()
	BeforeDelete()
	GetID() string
}

type BaseModel struct {
	ID    string `json:"id" bson:"_id"`
	CTime int64  `json:"ctime" bson:"ctime"`
	MTime int64  `json:"mtime" bson:"mtime"`
	DTime int64  `json:"dtime" bson:"dtime"`
}

func (m *BaseModel) BeforeCreate(prefix string) {
	now := time.Now().Unix()
	m.CTime = now
	m.MTime = now
	m.ID = prefix + "_" + uuid.NewString()
	fmt.Println(m)
}

func (m *BaseModel) BeforeUpdate() {
	now := time.Now().Unix()
	m.MTime = now
}

func (m *BaseModel) BeforeDelete() {
	now := time.Now().Unix()
	m.DTime = now
}

func (m *BaseModel) GetID() string {
	return m.ID
}
