n = input()
if len(n) == 3:
    print('A')
elif len(n) == 1:
    print('F')
elif n[0] == '9':
    print('A')
elif n[0] == '8':
    print('B')
elif n[0] == '7':
    print('C')
elif n[0] == '6':
    print('D')
else:
    print('F')
