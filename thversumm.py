st = input()

def thversumm(st, i):
    if i == 0:
        return int(st[i])
    else:
        return int(st[i]) + thversumm(st, i -1)

print(thversumm(st, len(st) -1))
