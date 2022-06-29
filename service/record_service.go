package service

import (
	"MusicWebsite/dao"
	"MusicWebsite/model"
	"MusicWebsite/tool"
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


func (rs *RecordService) GetRecord(rec *model.Record) tool.Res {
	i, r := rs.rd.CheckRecord(rec.Uid, rec.Iid)
	if i > 0 {
		return tool.GetGoodResult(*r)
	} else {
		return tool.GetGoodResult(*rec)
	}
}


func (rs *RecordService) AddRecord(rec *model.Record) tool.Res {
	t := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)	
	if b := rs.rd.AddRecord(rec.Uid, rec.Iid, t); b {
		rec.TimeStamp = t
		if b2 := rs.sd.UpdatePlayCnt(rec.Iid); b2 {
			return tool.GetGoodResult(*rec)
		} else {
			return tool.GetGoodResult(*rec)
		}
	} else {
		return tool.GetGoodResult(*rec)
	}

}
