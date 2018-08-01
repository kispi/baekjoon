stack = []
ops = []

def parse(stmt):
    return stmt.split(" ")
    
def execute(stmt):
    global stack
    if stmt[0] == 'push':
        stack = [stmt[1]] + stack
    elif stmt[0] == 'front':
        print(stack[len(stack)-1]) if len(stack) > 0 else print(-1)
    elif stmt[0] == 'back':
        print(stack[0]) if len(stack) > 0 else print(-1)
    elif stmt[0] == 'empty':
        print(1) if len(stack) == 0 else print(0)
    elif stmt[0] == 'size':
        print(len(stack))
    elif stmt[0] == 'pop':
        try:
            print(stack.pop())
        except:
            print(-1)

n = int(input())
for i in range(n):
    ops.append(parse(input()))

for op in ops:
    execute(op)