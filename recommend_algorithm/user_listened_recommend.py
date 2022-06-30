# -*- coding =utf-8 -*-
# @time:2022/6/28
import io

# 建立映射关系字典
import os
import random
import pymysql
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
            rid_to_name[line[0]] = line[1].strip()
            name_to_rid[line[1]] = line[0].strip()
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

def get_topN_items(current_item_raw_id, topK):
    """
    得到指定物品的top-N个相似物品
    :param current_item_raw_id: 物品的原始id，必须为字符串类型
    :param topK: 相似度高的前topK首歌曲
    :return: 当前歌曲的相似歌曲id列表
    """

    print("歌曲原始id：")
    print(current_item_raw_id)
    try:
        # 得到矩阵中的歌曲id（内部id），参数为字符串格式
        current_song_inner_id = algo.trainset.to_inner_iid(current_item_raw_id)
        print("歌曲内部id：")
        print(current_song_inner_id)
        # 相似歌曲推荐，得到的是相似歌曲的内部id，得到topK个
        current_song_neighbors = algo.get_neighbors(current_song_inner_id, k=topK)
        # 推荐歌手的内部id如下
        print("推荐歌曲的内部id：")
        print(current_song_neighbors)
        # 从相似歌曲的内部id得到原始id
        current_song_neighbors = (algo.trainset.to_raw_iid(inner_id) for inner_id in current_song_neighbors)
        # 从相似歌曲的原始id得到相似歌曲的名称
        current_song_neighbors = (item_rid_to_name[rid] for rid in current_song_neighbors)
        # 存储topN首歌曲的id
        topN_items_id = []
        print("推荐歌曲如下：")
        for song in current_song_neighbors:
            topN_items_id.append(item_name_to_rid[song])
            print(item_name_to_rid[song], song)

        return topN_items_id
    except:
        topN_items_id=[]
        return topN_items_id


# 获得每首歌曲的相似歌曲推荐
def get_all_topN_songs():
    """
    从文件中读取歌曲信息，然后根据基于用户的协同过滤推荐算法，推荐出每一首歌曲的相似歌曲，并将相似歌曲对应的id存到文件汇总，形成倒排链表
    格式为：当前歌曲id  相似歌曲id1，相似歌曲id2···
    """
    song_info_file_name = './dataset/song_info.txt'
    all_topN_songs_file_name = './resultset/topN_songs_baseline.txt'
    with open(song_info_file_name, 'r', encoding='utf-8') as input_f, open(all_topN_songs_file_name, 'w+', encoding='utf-8') as output_f:
        for line in input_f:
            # print(line)
            # 获得歌曲id
            song_id = line.split('\t')[0]
            print('获得 \"' + item_rid_to_name[song_id] + '\" 的推荐结果：')
            line = song_id + '\t'

            songs_items = get_topN_items(song_id, 10)
            # 将得到的结果拼接成字符串
            print(songs_items)
            if songs_items != []:
                for id in songs_items:
                    line = line + id + ','
            else:
                print("此首歌无相似歌曲")
                line = line+','

            # 写入到文件中
            output_f.write(line.strip(',') + '\n')
            output_f.flush()

        # 关闭文件输入输出流
        input_f.close()
        output_f.close()


def get_recommend_songs(current_user_raw_id, topK):
    '''
    获得一个用户的推荐音乐
    :param current_user_raw_id: 用户id
    :param topK: 推荐的前topK首音乐
    :return:推荐音乐列表
    '''
    user_listened_songs=[]
    recommend_songs=[]
    count = 0
    f = open('./dataset/user_record.txt', "r")
    for line in f.readlines():
        if current_user_raw_id in line:
            count += 1
            print("第 " + str(count) + " 行已找到.")
            print("该行内容: " + line)
            user_listened_songs.append(line.split('\t')[1])

    f.close()
    print(user_listened_songs)

    f = open('./resultset/topN_songs_baseline.txt', "r")
    for current_song_id in user_listened_songs:
        for line in f.readlines():
            if current_song_id == line.split('\t')[0]:
                recommend_songs.extend(line.split('\t')[1].strip('\n').split(','))
    recommend_songs = [x for x in recommend_songs if x != '']
    if len(recommend_songs)>=20:
        recommend_songs = random.sample(recommend_songs, 20)
    else:
        # 打开数据库连接
        conn = pymysql.connect(host='localhost', port=3306, database='music_rs', user='root', password='root',charset='utf8')
        # 获得Cursor对象
        cs = conn.cursor()
        try:
            sql = 'select iid from song'
            cs.execute(sql)
            while 1:
                arr = cs.fetchone()
                if len(recommend_songs) == 20:
                    break

                if arr[0] not in recommend_songs:
                    recommend_songs.append(arr[0])

        except:
            print("sql execute fail")

        cs.close()
        conn.close()

    print(recommend_songs)

    return recommend_songs


def get_all_recommend_songs():
    '''
    根据相似音乐和播放记录生成用户推荐音乐txt
    :return:
    '''
    user_info_file_name = './dataset/user_info.txt'
    all_topN_recommend_music_file_name = './resultset/topN_recommend_music.txt'

    with open(user_info_file_name, 'r', encoding='utf-8') as input_f, open(all_topN_recommend_music_file_name, 'w+',encoding='utf-8') as output_f:
        for line in input_f:

            user_id = line.split('\t')[0]
            line = user_id + '\t'

            for songs_id in get_recommend_songs(user_id, 20):
                line = line + songs_id + ','

            output_f.write(line.strip(',') + '\n')
            output_f.flush()

        input_f.close()
        output_f.close()


if __name__=="__main__":
    '''
    根据音乐相似度和用户听歌记录推荐20首相似音乐
    '''
    # 获取歌曲名称到歌曲id 和 歌曲id到歌曲名的映射
    item_rid_to_name, item_name_to_rid = get_rid2name('./dataset/song_info.txt')
    # 得到用户名称到id 和 用户id到名称的映射    +

    user_rid_to_name, user_name_to_rid = get_rid2name('./dataset/user_info.txt')

    # 得到训练之后的结果
    algo = get_trainset_algo()

    # 根据听歌记录生成与每首歌相似的20首音乐
    # get_all_topN_songs()

    # #结合听歌记录生成每位用户的20首推荐音乐
    get_all_recommend_songs()