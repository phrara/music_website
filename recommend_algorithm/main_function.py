# -*- coding =utf-8 -*-
# @time:2022/6/30

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