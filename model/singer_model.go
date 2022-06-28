package model

type Singer struct {
	Suid string `gorm:"primary_key;column:suid" json:"suid"`
	SingerName string `gorm:"column:singer_name" json:"sn"`
	SingerUrl string `gorm:"column:singer_url" json:"su"`
}


func (s Singer) TableName() string {
	return "singer"
}

func NewSinger(suid, sn, su string) *Singer {
	return &Singer{
		Suid: suid,
		SingerName: sn,
		SingerUrl: su,
	}
}
