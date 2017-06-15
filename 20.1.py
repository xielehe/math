# quadratic residues modulo 19
# 1, 4, 5, 6, 7, 9, 11, 16, 17
# quadratic nonresidues modulo 19
# 2, 3, 8, 10, 12, 13, 14, 15, 18
import func

def quadraticResidues(num):
    if  not func.is_prime(num):
        print('not prime')
        return
    P = (num - 1) / 2
    residues = map(lambda x: (x*x)%num, range(1,P+1))
    nonresidues = list(set(range(1,num)).difference(set(residues)))
    return residues, nonresidues

if __name__ == '__main__':
    print quadraticResidues(19)