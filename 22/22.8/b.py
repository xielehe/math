# -*- coding: utf-8 -*
import sys
sys.path.append('../../')
import func

# 22.8(b)，求同余式 x^2 = 5 (mod 541)的一个解.
# 根据(a)中的结论，先求解a^(p-1)/4(mod p) 得 1
# 所以 x = a^(p+3)/8 (mod 541)
#  = 345
a = 5
b = (541 - 1) / 4
m = 541
x = func.aModM(a, b, m)
print (x)

a = 5
b = (541 +3) / 8
m = 541
x = func.aModM(a, b, m)
print (x)
