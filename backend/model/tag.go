package model

type Tag struct {
	Base
	Title string `sql:"not null; unique_index"`
}