package service

import (
	"MusicWebsite/dao"
	"MusicWebsite/model"
	"strconv"
	"time"
)

type RecordService struct {
	rd *dao.RecordDao
	sd *dao.SongDao
}

func NewRecordService() *RecordService {
	rs := new(RecordService)
	rs.rd = new(dao.RecordDao)
	rs.sd = new(dao.SongDao)
	return rs
}


func (rs *RecordService) GetRecord(rec *model.Record) *model.Record {
	i, r := rs.rd.CheckRecord(rec.Uid, rec.Iid)
	if i > 0 {
		return r
	} else {
		return rec
	}
}


func (rs *RecordService) AddRecord(rec *model.Record) *model.Record{
	t := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)	
	if b := rs.rd.AddRecord(rec.Uid, rec.Iid, t); b {
		rec.TimeStamp = t
		if b2 := rs.sd.UpdatePlayCnt(rec.Iid); b2 {
			return rec
		} else {
			return rec
		}
	} else {
		return rec
	}

}
