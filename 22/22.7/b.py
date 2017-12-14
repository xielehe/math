# -*- coding: utf-8 -*
import sys
sys.path.append('../../')
import func

# 22.7(b)，求同余式 x^2 = 7 (mod 787)的一个解.
# 根据(a)中的结论，7^(787+1)/4 是一个解.
# x = 105

a = 7
b = (787 + 1) / 4
m = 787
x = func.aModM(a, b, m)
print (x)
