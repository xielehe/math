# -*- coding: utf-8 -*-
import matplotlib.pyplot as plt

class Line:
   def __init__(self, x, y):
      self.X = x
      self.Y = y

a=6
b=4

# A = 12, N = 7, B = 12
# A-N-B/2 = -1

x = []
y = []
for i in range(b+1):
    for j in range(a+1):
        if not i > j*b/a:
            x.append(j)
            y.append(i)

la = Line([0, a], [0,0])
lb = Line([0, a], [0,b])
lc = Line([a, a], [0,b])

plt.plot(la.X, la.Y)
plt.plot(lb.X, lb.Y)
plt.plot(lc.X, lc.Y)
plt.plot(x, y, 'ro')
plt.ylabel('points')  # 为y轴加注释
plt.show()
