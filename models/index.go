package models

import (
	"ablog/helpers"
)


type Lists struct{
	Id int `gorm:"primary_key"`
	Title string `gorm:"type:varchar(50)"`
	Article_abstract string `gorm:"type:varchar(200)"`
	Artictl_id int `gorm:"type:int"`
	Type string `gorm:"varchar(50)"`

}

type Navis struct{
	Id int `gorm:"primary_key"`
	Button string `gorm:"varchar(30)"`
	Link string `gorm:"varchar(50)"`
	Css string `gorm:"varchar(50)"`
}

type Articles struct {
	Id int `gorm:"primary_key"`
	Content string `gorm:"varchar(10000)"`
}

func (n *Navis) Get() []*Navis{
	var result []*Navis
	//helpers.Db.Raw("select * from ablog.navi").Scan(&result)
	helpers.Db.Raw("select * from navis").Find(&result)
	return result
}

func (l *Lists) Get() []*Lists{
	var result []*Lists
	helpers.Db.Raw("select * from lists limit 10").Find(&result)
	return result
}

