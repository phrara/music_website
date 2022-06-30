package service

import (
	"MusicWebsite/dao"
	"MusicWebsite/model"
	"MusicWebsite/tool"
	"strconv"
	"strings"
)

type RCMDService struct {
	rd *dao.RCMDDao
	ud *dao.UserDao
	sd *dao.SongDao
	srd *dao.SingerDao
	rcd *dao.RecordDao
}

func NewRCMDService() *RCMDService {
	rs := new(RCMDService)
	rs.ud = new(dao.UserDao)
	rs.rd = new(dao.RCMDDao)
	rs.sd = new(dao.SongDao)
	rs.srd = new(dao.SingerDao)
	rs.rcd = new(dao.RecordDao)
	return rs
}

// 获取推荐用户
func (rs *RCMDService) GetRcmdUsers(user *model.UserInfo) tool.Res {
	ur := rs.rd.GetRcmdUsers(user.Uid)
	ulist := make([]model.User,0)
	favor := rs.ud.GetUserInfo(user.Uid).Des
	records := rs.rcd.GetAllRecordByUid(user.Uid)
	recNum := 0
	for _, v := range records {
		rec, _ := strconv.ParseInt(v.Count, 10, 64)
		recNum += int(rec)
	}
	if recNum < 50 {
		return tool.GetGoodResult(rs.rd.GetSimilarUsers(favor))
	}

	for _, v := range ur.GetTopUsers() {
		u := rs.ud.GetUserInfo(v)
		u.Password = "*********"
		ulist = append(ulist, *u)
	}
	return tool.GetGoodResult(ulist)
}

// 获取推荐歌曲
func (rs *RCMDService) GetRcmdSongs(user *model.UserInfo) tool.Res {
	slist := make([]model.SongInfo,0)
	records := rs.rcd.GetAllRecordByUid(user.Uid)
	recNum := 0
	for _, v := range records {
		rec, _ := strconv.ParseInt(v.Count, 10, 64)
		recNum += int(rec)
	}
	if recNum < 50 {
		songs := rs.sd.GetSongList(100)
		favor := rs.ud.GetUserInfo(user.Uid).Des
		for _, v := range songs {
			iid, _ := strconv.ParseInt(v.Iid, 10, 64)
			kind := strconv.FormatInt(iid % 10, 10)
			if strings.Contains(favor, kind) {
				si := model.NewSongInfo(&v, rs.srd.GetSingerInfo(v.SingerId).SingerName)
				slist = append(slist, *si)
			}
		}
		return tool.GetGoodResult(slist)
	}
	sr := rs.rd.GetRcmdSongs(user.Uid)
	for _, v := range sr.GetSongs() {
		song := rs.sd.GetSongInfo(v)
		si := model.NewSongInfo(song, rs.srd.GetSingerInfo(song.SingerId).SingerName)
		slist = append(slist, *si)
	}
	return tool.GetGoodResult(slist)
}

// 获取相似歌曲
func (rs *RCMDService) GetSimilarSongs(iid string) tool.Res {
	ss := rs.rd.GetSimilarSongs(iid)
	slist := make([]model.SongInfo,0)
	for _, v := range ss.GetTopSongs() {
		song := rs.sd.GetSongInfo(v)
		si := model.NewSongInfo(song, rs.srd.GetSingerInfo(song.SingerId).SingerName)
		slist = append(slist, *si)
	}
	return tool.GetGoodResult(slist)
}