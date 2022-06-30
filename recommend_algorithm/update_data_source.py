# -*- coding =utf-8 -*-
# @time:2022/6/29
import pymysql

import other_user_recommend
import user_listened_recommend

def update_user_listened_record():
    '''
    更新 ./dataset/user_record.txt
    :return:
    '''
    try:
        sql = 'select * from record'
        cs.execute(sql)
        all_record = cs.fetchall()
        print(all_record)
    except:
        print("select record sql execute fail")

    with open("./dataset/user_record.txt","w+", encoding='utf-8') as f:
        for i in range(len(all_record)):
            line = all_record[i][0] + "\t" + all_record[i][1] + "\t" + all_record[i][2] + "\t" + all_record[i][3] + "\n"
            f.write(line)
            f.flush()
        f.close()


def update_songs_info():
    '''
    更新 ./dataset/song_info.txt
    :return:
    '''
    try:
        sql = 'select iid,song_name from song'
        cs.execute(sql)
        all_song = cs.fetchall()
        print(all_song)
    except:
        print("select song sql execute fail")

    with open("./dataset/song_info.txt", "w+", encoding='utf-8') as f:
        for i in range(len(all_song)):
            line = all_song[i][0] + "\t" + all_song[i][1] + "\n"
            f.write(line)
            f.flush()
        f.close()


def update_users_info():
    '''
    更新 ./dataset/user_info.txt
    :return:
    '''
    try:
        sql = 'select uid,name from user'
        cs.execute(sql)
        all_users = cs.fetchall()
        print(all_users)
    except:
        print("select user sql execute fail")

    with open("./dataset/user_info.txt", "w+", encoding='utf-8') as f:
        for i in range(len(all_users)):
            line = all_users[i][0] + "\t" + all_users[i][1] + "\n"
            f.write(line)
            f.flush()
        f.close()


if __name__ == "__main__":
    '''
    更新用户听歌记录、音乐库和用户列表，实时进行音乐推荐和相似用户推荐
    '''
    conn = pymysql.connect(host='localhost', port=3306, database='music_rs', user='root', password='root', charset='utf8')
    cs = conn.cursor()

    update_user_listened_record()

    update_songs_info()
    
    update_users_info()
