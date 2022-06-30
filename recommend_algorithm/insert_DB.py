# -*- coding =utf-8 -*-
# @time:2022/6/30
import mysql_util
import pymysql

password = '2001823'

def read_top_users2mysql():
    """
    读取好友推荐结果并写入到数据库的topusers表中
    :return:
    """
    # 打开数据库连接
    conn = pymysql.connect(host='localhost', port=3306, database='music_website', user='root', password=password,
                           charset='utf8')
    # 获得Cursor对象
    cs = conn.cursor()

    # 用户文件路径名
    top_users_file_name = './resultset/topN_users_baseline.txt'
    # sql语句
    sql = "INSERT INTO `topusers`(`uid`, `topusers`) VALUES (%s,%s) ON DUPLICATE KEY UPDATE `topusers`= (%s);"

    # 读取文件存入数据库
    with open(top_users_file_name, 'r', encoding='utf-8') as f:
        for line in f:
            top_user = line.strip().split('\t')
            uid = top_user[0]
            topuers = top_user[1]
            # 定义一个参数元组
            param = (uid, topuers, topuers)
            print(uid + '写入数据库···')
            cs.execute(sql, param)
    # 增、删、改类操作需要提交事务
    conn.commit()
    # 关闭文件
    f.close()
    cs.close()
    # 关闭数据库
    conn.close()


def read_top_songs2mysql():
    """
    读取歌曲推荐结果并写入到数据库的topsongs表中
    :return:
    """
    # 打开数据库连接
    conn = pymysql.connect(host='localhost', port=3306, database='music_website', user='root', password=password,
                           charset='utf8')
    # 获得Cursor对象
    cs = conn.cursor()

    # 用户文件路径名
    top_songs_file_name = './resultset/topN_songs_baseline.txt'
    # sql语句
    sql = "INSERT INTO `topsongs`(`iid`, `topsongs`) VALUES (%s,%s) ON DUPLICATE KEY UPDATE `topsongs`= (%s);"

    # 读取文件存入数据库
    with open(top_songs_file_name, 'r', encoding='utf-8') as f:
        for line in f:
            top_song = line.strip().split('\t')
            iid = top_song[0]
            try:
                topsongs = top_song[1]
            except:
                topsongs = "-1"

            # 定义一个参数元组
            param = (iid, topsongs, topsongs)
            print(iid + '写入数据库···')
            cs.execute(sql, param)
    # 增、删、改类操作需要提交事务
    conn.commit()
    # 关闭文件
    f.close()
    cs.close()
    # 关闭数据库
    conn.close()


def read_recommend_music2mysql():
    """
    读取用户的推荐音乐并写入到数据库的rcmdsongs表中
    :return:
    """
    # 打开数据库连接
    conn = pymysql.connect(host='localhost', port=3306, database='music_website', user='root', password=password,
                           charset='utf8')
    # 获得Cursor对象
    cs = conn.cursor()

    # 用户文件路径名
    top_songs_file_name = './resultset/topN_recommend_music.txt'
    # sql语句
    sql = "INSERT INTO `rcmdsongs`(`uid`, `songs`) VALUES (%s,%s) ON DUPLICATE KEY UPDATE `songs`= (%s);"

    # 读取文件存入数据库
    with open(top_songs_file_name, 'r', encoding='utf-8') as f:
        for line in f:
            top_song = line.strip().split('\t')
            uid = top_song[0]
            songs = top_song[1]
            # 定义一个参数元组
            param = (uid, songs, songs)
            print(uid + '写入数据库···')
            cs.execute(sql, param)
    # 增、删、改类操作需要提交事务
    conn.commit()
    # 关闭文件
    f.close()
    cs.close()
    # 关闭数据库
    conn.close()



if __name__ == "__main__":
    '''
    更新以下三张表的信息
    '''
    # 测试读取topN_users_baseline到数据库中
    print("=====================BEGIN INSERT INTO topusers=====================")
    read_top_users2mysql()
    print("=====================END INSERT INTO topusers=====================")

    # 测试读取topN_songs_baseline到数据库中
    print("=====================BEGIN INSERT INTO topsongs=====================")
    read_top_songs2mysql()
    print("=====================END INSERT INTO topsongs=====================")

    # 测试读取topN_recommend_music到数据库中
    print("=====================BEGIN INSERT INTO rcmdsongs=====================")
    read_recommend_music2mysql()
    print("=====================END INSERT INTO rcmdsongs=====================")