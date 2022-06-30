# -*- coding =utf-8 -*-
# @time:2022/6/30
'''
Author:张靖宇
Function：
    1.调用该py文件后，会先从数据库中读取信息，更新dataset文件夹中的数据源
    2.根据KNN算法生成每位用户的10位相似好友，存储在./resultset/topN_users_baseline.txt
    3.根据KNN算法生成每位用户的20首推荐音乐，存储在./resultset/topN_recommend_music.txt
    4.将./resultset文件加下的三个文件写回数据库
'''
import os
print("\n========================Update Data Source Begin========================\n")
os.system('python ./update_data_source.py')
print("\n========================Update Data Source End========================\n")

print("\n========================Recommend Same Users Begin========================\n")
os.system('python ./other_user_recommend.py')
print("\n========================Recommend Same Users End========================\n")

# 运行时间约20min
print("\n========================Recommend Music Begin========================\n")
os.system('python ./user_listened_recommend.py')
print("\n========================Recommend Music End========================\n")

print("\n========================Insert Into Mysql Begin========================\n")
os.system('python ./insert_DB.py')
print("\n========================Insert Into Mysql End========================\n")