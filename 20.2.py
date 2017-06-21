import func

def a():
    for x in filter(lambda x: x%2==1, filter(func.is_prime, xrange(1, 20))):
        print func.quadraticResidues(x), '-->', x

def b():
    for x in filter(lambda x: x%2==1, filter(func.is_prime, xrange(1, 20))):
        res = func.quadraticResidues(x)
        print 'A:', sum(res[0]), 'B:', sum(res[1]), 'p-->', x, 'A+B-->', sum(res[0]) + sum(res[1])

if __name__ == '__main__':
    # ([1], [2]) --> 3
    # ([1, 4], [2, 3]) --> 5
    # ([1, 4, 2], [3, 5, 6]) --> 7
    # ([1, 4, 9, 5, 3], [8, 2, 10, 6, 7]) --> 11
    # ([1, 4, 9, 3, 12, 10], [2, 5, 6, 7, 8, 11]) --> 13
    # ([1, 4, 9, 16, 8, 2, 15, 13], [3, 5, 6, 7, 10, 11, 12, 14]) --> 17
    # ([1, 4, 9, 16, 6, 17, 11, 7, 5], [2, 3, 8, 10, 12, 13, 14, 15, 18]) --> 19
    # a()

    # A: 1 B: 2 p--> 3 A+B--> 3
    # A: 5 B: 5 p--> 5 A+B--> 10
    # A: 7 B: 14 p--> 7 A+B--> 21
    # A: 22 B: 33 p--> 11 A+B--> 55
    # A: 39 B: 39 p--> 13 A+B--> 78
    # A: 68 B: 68 p--> 17 A+B--> 136
    # A: 76 B: 95 p--> 19 A+B--> 171
    # p(p-1)/2 = A + B
    b()
