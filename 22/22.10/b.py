# -*- coding: utf-8 -*
import sys
sys.path.append('../../')
import func

a = 2
p = 1293337
b1 = (p-1)/2
b2 = (p-1)

print(func.aModM(a, b1, p)) #欧拉准则
print(func.aModM(a, b2, p)) #费马小定理
