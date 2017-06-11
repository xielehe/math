#! /usr/bin/python
# -*- coding: utf-8 -*-
import func
from functools import reduce

def make():
    arr = [17]
    while len(arr) < 5:
        p = reduce(lambda x,y:x*y,arr)
        A = p*p*4 + 1
        arr = arr + map(lambda x: int(x), func.factor(A))
    print arr

if __name__ == "__main__":
    make()