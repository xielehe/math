#! /usr/bin/python
# -*- coding: utf-8 -*-
import func
def base(m, a):
    for x in xrange(1,m):
        if func.aModM(x, 2, m) == a:
            return (str(x) + '----> yes')
    return 'none'

def a(): print base(5987, 5987-1)
def b(): print base(6781, 6781-1)
def c(): print base(337, 84)
def d(): print base(3011, 81)

if __name__ == "__main__":
    # a()
    # b()
    # d()
    # print 41*41-41*64+943
    # c()
    print (67*67+14*67-35)/337











