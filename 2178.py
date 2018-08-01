n = int(input().split(" ")[0])
maze = []

class Guy():
    def __init__(self, path, prev, cur):
        self.prev = prev
        self.cur = cur
        self.path = path

    def check_valid_index(self, maze):
        left = (self.cur[0] - 1, self.cur[1])
        right = (self.cur[0] + 1, self.cur[1])
        up = (self.cur[0], self.cur[1] - 1)
        down = (self.cur[0], self.cur[1] + 1)

        if left >= 0 and right <= len(maze[0]) - 1 and up >= 0 and down <= len(maze):
            return True
        else:
            return False
        
    def move_next(self):
        self.check_valid_index(maze)        

        


def play():
    pass
    
    

for i in range(n):
    maze.append(input())

print(maze)