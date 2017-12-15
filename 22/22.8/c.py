# -*- coding: utf-8 -*
import sys
sys.path.append('../../')
import func

# 22.8(c)，求同余式 x^2 = 13 (mod 653)的一个解.
# 根据(a)中的结论，先求解a^(p-1)/4(mod p) 得 -1
# 所以 x =2a*4a^(p-5)/8 (mod 653)
#  = 345
p = 653
a = 13
b = (p - 1) / 4
m = p
x = func.aModM(a, b, m)
print (x)

p = 653
a = 13 * 4
b = (p - 5) / 8
m = p
x = func.aModM(a, b, m)
x = (2 * 13 * x) % p
print (x)
