package dao

import (
	"MusicWebsite/model"
	"strconv"
)

type RecordDao struct {
}

// 查询某用户所有记录
func (r *RecordDao) GetAllRecordByUid(uid string) []*model.Record {
	res := make([]*model.Record, 100)
	DBMgr.Where("uid = ?", uid).Find(&res)
	return res
}

// 检查某条记录
func (r *RecordDao) CheckRecord(uid, iid string) (int64, *model.Record) {
	rec := model.NewRecord("","")
	DBMgr.Where("uid = ? and iid = ?",uid,iid).First(rec)
	c,_ := strconv.ParseInt(rec.Count,10,64)
	return c, rec
}

// 添加记录
func (r *RecordDao) AddRecord(uid, iid string, timeStamp string) bool {
	if c, rec := r.CheckRecord(uid, iid); c > 0 {
		c++
		weight := strconv.FormatInt(c, 10)
		rec.Count = ""
		rec.TimeStamp = ""
		res := DBMgr.Model(rec).Where("uid = ? and iid = ?",rec.Uid, rec.Iid).Update(map[string]any{"weight": weight, "timestamp": timeStamp})
		if res.RowsAffected > 0 {
			return true
		} else {
			return false
		}
	} else {
		rec.Count = "1"
		rec.TimeStamp = timeStamp
		DBMgr.Create(rec)
		return true
	}

}