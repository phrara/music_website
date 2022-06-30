# -*- coding =utf-8 -*-
# @time:2022/6/28

import os
import random

from surprise import Reader, Dataset, KNNBaseline

def get_rid2name(file_name):
    """
    获得名称到id 和 id到名称的映射，包括歌曲或用户等的映射
    :param:file_name: 需要解析的文件路径名，且文件中每行的第一列为id，第二列为名称，使用制表符隔开
    :return:
        rid_to_name:键为rid，值为name，只有一个键值对
        name_to_rid = {}键为name，值为rid，只有一个键值对
    """
    rid_to_name = {}
    name_to_rid = {}
    with open(file_name, 'r', encoding='utf-8') as f:
        for line in f:
            line = line.split('\t')
            rid_to_name[line[0]] = line[1]
            name_to_rid[line[1]] = line[0]
    return rid_to_name, name_to_rid


def get_trainset_algo():
    """
    KNN算法使用训练集进行训练
    :return:训练结果
    """

    reader = Reader(line_format='user item rating timestamp', sep='\t',  rating_scale=(0, 100),skip_lines=0)
    # 用户听歌记录文件路径
    file_path = "./dataset/user_record.txt"
    # 加载数据集
    data = Dataset.load_from_file(file_path, reader=reader)
    # 将数据集转换成矩阵形式
    trainset = data.build_full_trainset()
    # print(trainset)
    print(trainset.all_items())
    print(trainset.all_users())

    # 基于物品的协同过滤算法，相似度衡量方法：皮尔逊相似度
    # 这是一个用户数量N，矩阵大小为 N*N 的稀疏矩阵，然后get_neighbors得到的是topK个相似用户。如果想要得到相似歌曲，则需要使用基于项目的协同过滤算法，
    # 或者从得到的相似用户中，提取他们的播放记录（这是基于用户的协同过滤算法）
    sim_options = {'name': 'pearson_baseline', 'user_based': True}
    # 选择算法
    algo = KNNBaseline(sim_options=sim_options)
    # algo = KNNBasic(sim_options=sim_options)

    # 训练数据集
    print('开始训练······')
    algo.fit(trainset)
    print('训练结束!!!!')
    return algo

def get_topN_users(current_user_raw_id, topK):

    """
    获得相似音乐好友推荐
    :param current_user_raw_id:  当前用户id
    :param topK:相似度高的前topK个相似音乐好友
    :return:当前用户的相似音乐好友id数组
    """

    print("用户原始id：")
    print(current_user_raw_id)
    try:
        # 得到矩阵中的用户id（内部id），方法是algo.trainset.to_inner_uid(uid)，参数为字符串格式
        current_user_inner_id = algo.trainset.to_inner_uid(current_user_raw_id)
        print("用户内部id：")
        print(current_user_inner_id)

        # 相似音乐好友推荐，得到的是相似音乐好友的内部id，得到topK个
        current_user_neighbors = algo.get_neighbors(current_user_inner_id, k=topK)

        # 推荐相似好友的内部id如下
        print("推荐用户的内部id：")
        print(current_user_neighbors)
        # 从相似音乐好友的内部id转化为原始id
        current_user_neighbors = (algo.trainset.to_raw_uid(inner_id)
                                  for inner_id in current_user_neighbors)
        # 从相似音乐好友的原始id得到名字
        current_user_neighbors = (user_rid_to_name[rid]
                                  for rid in current_user_neighbors)

        # 存储topN个用户的id
        topN_users_id = []
        print("推荐相似好友如下：")
        for user in current_user_neighbors:
            topN_users_id.append(user_name_to_rid[user])
            print(user_name_to_rid[user], user.strip())

        return topN_users_id

    except:
        print("该用户无听歌记录，随机推荐相似好友")
        topN_users_id=[]
        with open("./dataset/user_info.txt",'r', encoding='utf-8') as f:
            for i in range(0,10):
                line = f.readline()
                topN_users_id.append(line.split('\t')[0])

            f.close()
        return topN_users_id

# 获得每个用户的相似音乐好友推荐
def get_all_topN_users():
    """
    从文件中读取用户信息，然后根据基于用户的协同过滤推荐算法，推荐出每一个用户的相似音乐好友，并将音乐好友对应的id存到文件汇总，形成倒排链表
    """
    # 输入文件路径和结果文件路径
    user_info_file_name = './dataset/user_info.txt'
    all_topN_users_file_name = 'resultset/topN_users_baseline.txt'

    with open(user_info_file_name, 'r', encoding='utf-8') as input_f, open(all_topN_users_file_name, 'w+',encoding='utf-8') as output_f:
        for line in input_f:
            # 获得用户id
            user_id = line.split('\t')[0]
            line = user_id + '\t'
            # 将结果拼接成字符串
            for id in get_topN_users(user_id, 10):
                line = line + id + ','
            # 写入到文件中
            output_f.write(line.strip(',') + '\n')
            # 强行把缓冲区中的内容放到磁盘中
            output_f.flush()
        # 关闭文件输入输出流
        input_f.close()
        output_f.close()


if __name__=="__main__":
    '''
    根据听歌记录推荐10位相似用户
    '''
    # 获取歌曲名称到歌曲id 和 歌曲id到歌曲名的映射
    item_rid_to_name, item_name_to_rid = get_rid2name('./dataset/song_info.txt')
    # 得到用户名称到id 和 用户id到名称的映射
    user_rid_to_name, user_name_to_rid = get_rid2name('./dataset/user_info.txt')

    # 得到训练之后的结果
    algo = get_trainset_algo()

    # 获取所有用户的相似好友
    get_all_topN_users()