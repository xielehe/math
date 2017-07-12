import func

def a(n):
    return map(lambda x: (x*x*x)%n, range(1,n))

if __name__ == '__main__':

    print(a(5))
    print(a(7))
    print(a(11))
    print(a(13))



 # 1= 1^3(mod 5)
 # 3= 2^3(mod 5)
 # 2= 3^3(mod 5)
 # 4= 4^3(mod 5)

 # 1= 1^3(mod 7)
 # 1= 2^3(mod 7)
 # -1= 3^3(mod 7)
 # 1= 4^3(mod 7)
 # -1= 5^3(mod 7)
 # -1= 6^3(mod 7)

 # 1= 1^3(mod 11)
 # 8= 2^3(mod 11)
 # 5= 3^3(mod 11)
 # 9= 4^3(mod 11)
 # 4= 5^3(mod 11)
 # 7= 6^3(mod 11)
 # 2= 7^3(mod 11)
 # 6= 8^3(mod 11)
 # 3= 9^3(mod 11)
 # 10= 10^3(mod 11)
 
 # 1= 1^3(mod 13)
 # 8= 2^3(mod 13)
 # 1= 3^3(mod 13)
 # -1= 4^3(mod 13)
 # 8= 5^3(mod 13)
 # 8= 6^3(mod 13)
 # 5= 7^3(mod 13)
 # 5= 8^3(mod 13)
 # 1= 9^3(mod 13)
 # -1= 10^3(mod 13)
 # 5= 11^3(mod 13)
 # -1= 12^3(mod 13)
