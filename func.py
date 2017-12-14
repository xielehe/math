#! /usr/bin/python
# -*- coding: utf-8 -*-
import sympy

#因式分解
def factor(number):
    return sympy.factorint(number, multiple=True)

#判断一个数是否是素数
def is_prime(number):
    if len(factor(number)) == 1:
        return True
    return False

# 判断一个数是否是卡米歇尔数
def Carmichael(number):
    fList = factor(number)
    if len(fList) < 3: return False
    if not len(fList) == len(set(fList)): return False
    for x in fList:
        if not (number-1)%(int(x)-1) == 0: return False
    return True

#计算a^b(mod m)
def aModM(a, b, m):
    binB = list(bin(b)[2:])
    binB.reverse()
    listMod, vaList = [a%m], []
    for x in xrange(1,len(binB)):listMod.append((listMod[x-1] * listMod[x-1]) % m)
    for i,b in enumerate(binB):
        if b == '1': vaList.append(listMod[i])
    return reduce(lambda x, y: (x*y)%m, vaList, 1)

#计算er ci sheng yu 
def quadraticResidues(num):
    if  not is_prime(num):
        print('not prime')
        return
    P = (num - 1) / 2
    residues = map(lambda x: (x*x)%num, range(1,P+1))
    nonresidues = list(set(range(1,num)).difference(set(residues)))
    return residues, nonresidues



#数论概论第19张习题--->19.6答案
if __name__ == "__main__":
    # print aModM(2, 294408, 294409)
    # print (1758249*1758249)%118901521
    # print aModM(45274074, 1, 118901521)
    # print aModM(45274074, 2, 118901521)
    # print aModM(45274074, 4, 118901521)
    # print aModM(45274074, 8, 118901521)
    # print is_prime(285707540662569884530199015485751094149)
    # print is_prime(285 707 540 662 569 884 530 199 015 485 751 094 149)









