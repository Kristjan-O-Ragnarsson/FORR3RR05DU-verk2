n = int(input())
m = int(input())

def funct(n, m):
    if m == 0:
        return n
    else:
        return funct(m, n % m)

print(funct(n, m))
