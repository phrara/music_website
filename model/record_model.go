package model

type Record struct{
	Uid string `gorm:"column:uid" json:"uid"`
	Iid string `gorm:"column:iid" json:"iid"`
	Count string `gorm:"column:weight" json:"weight"`
	TimeStamp string `gorm:"column:timestamp" json:"timestamp"`
}


func (r Record) TableName() string {
	return "record"
}


func NewRecord(uid, iid string) *Record {
	return &Record{
		Uid: uid,
		Iid: iid,
		Count: "",
		TimeStamp: "",
	}
}