#! /usr/bin/python
# -*- coding: utf-8 -*-
import func

#数论概论第19张习题--->19.4答案
# if __name__ == "__main__":
#     for x in xrange(1,54):
#         if func.is_prime(x*6+1) and func.is_prime(x*12+1) and func.is_prime(x*18+1):
#             print x , '-->' ,(x*6+1)*(x*12+1)*(x*18+1)

if __name__ == "__main__":
    for k in xrange(1000,10000):
        if func.is_prime(28*1*k+1) and func.is_prime(28*2*k+1) and func.is_prime(28*4*k+1) and func.is_prime(28*14*k+1) and func.is_prime(28*7*k+1):
            print k, '---->',  (28*1*k+1)*(28*2*k+1)*(28*4*k+1)*(28*14*k+1)*(28*7*k+1)