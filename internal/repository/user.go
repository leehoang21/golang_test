package repository

import (
	"base/internal/base/mgo"
)

type User interface {
	mgo.BaseTable
}
