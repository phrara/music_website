package dao

import "MusicWebsite/model"

type SingerDao struct {
}

func (s *SingerDao) GetSingerInfo(suid string) *model.Singer {
	singer := model.NewSinger("","","")
	DBMgr.Where("suid = ?",suid).First(singer)
	return singer
}

func (s *SingerDao) GetSuidByName(singerName string) string {
	singer := model.NewSinger("","","")
	DBMgr.Where("singer_name = ?",singerName).First(singer)
	return singer.Suid
}



