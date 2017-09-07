import func
import math

def iss(m, n):
    erci , nerci = func.quadraticResidues(n)
    return m in erci

if __name__ == "__main__":
    print iss(85, 101)
    print iss(29, 541)
    print iss(101, 1987)
    print iss(31706, 43789)
    print iss(13, 31957)
    for x in xrange(1,31957):
        if (x * x) % 31957 == 13:
            print x
    print (14241*14241-3*14241-1)/31957