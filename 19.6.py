#! /usr/bin/python
# -*- coding: utf-8 -*-
import func

#数论概论第19张习题--->19.6答案
if __name__ == "__main__":
    for x in xrange(1000001,2000000,2):
        if func.Carmichael(x):
            print x, '-->',len(func.factor(x))
            break