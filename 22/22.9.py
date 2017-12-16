# -*- coding: utf-8 -*
import sys
sys.path.append('../')
import func

# 22.9, 设p是模8余5的素数，求解下x^2=a(mod p)

def coresidual(a, p):
    if p % 8 != 5:
        raise NameError('p != 5 mod(8)')
    b = (p-1)/4
    x = func.aModM(a, b, p)
    if x == 1:
        b = (p+3)/8
        return func.aModM(a, b, p)
    elif x == p-1:
        b = (p - 5) / 8
        return (func.aModM(4*a, b, p) * 2 * a) % p
    else:
        raise NameError('(a/p) != 1 x = ' + str(x))

if __name__ == "__main__":
    print(coresidual(17, 1021)) #494
    print(coresidual(23, 1021)) #858
    print(coresidual(31, 1021)) #不是二次剩余

