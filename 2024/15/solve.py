from copy import deepcopy

from aocd import submit


def parse(lines):
    grid = list()
    moves = list()
    switched = False
    for i, line in enumerate(lines):
        if len(line) == 0:
            switched = True
            continue
        elif not switched:
            grid.append(list(line))
            for j, c in enumerate(list(line)):
                if c == "@":
                    rx, ry = i, j
        else:
            moves += list(line)

    grid[rx][ry] = FREE
    return grid, moves, (rx, ry)


def blow_up(grid, rx, ry):
    new_grid = list()
    for i, l in enumerate(grid):
        g = list()
        for j, c in enumerate(l):
            if c == "O":
                g += ["[", "]"]
            else:
                g += [c, c]
        new_grid.append(g)
    return new_grid, rx, 2*ry


TO_DIR = {
    "^": (-1, 0),
    ">": (0, 1),
    "v": (1, 0),
    "<": (0, -1),
}

FREE = "."
WALL = "#"
BOX = "O"
BOX_LEFT = "["
BOX_RIGHT = "]"


def part_a(input):
    grid, moves, (rx, ry) = input
    grid = deepcopy(grid)
    for m in moves:
        mx, my = TO_DIR[m]

        if grid[rx+mx][ry+my] == WALL:
            # wall
            continue
        elif grid[rx+mx][ry+my] == BOX:
            # box
            bx, by = rx+mx, ry+my
            space = next_space(grid, bx, by, mx, my)
            if space is not None:
                sx, sy = space
                move_boxes(grid, bx, by, sx, sy)
                rx, ry = rx+mx, ry+my
        else:
            # walk
            rx, ry = rx+mx, ry+my

    return gps(grid)


def next_space(grid, bx, by, mx, my):
    while grid[bx][by] != WALL:
        if grid[bx][by] == FREE:
            return (bx, by)
        bx, by = bx+mx, by+my
    return None


def move_boxes(grid, bx, by, sx, sy):
    grid[sx][sy] = BOX
    grid[bx][by] = FREE


def part_b(input):
    grid, moves, (rx, ry) = input
    grid, rx, ry = blow_up(grid, rx, ry)

    for m in moves:
        mx, my = TO_DIR[m]
        if try_push(grid, rx, ry, mx, my):
            push(grid, rx, ry, mx, my)
            rx, ry = rx+mx, ry+my

    return gps(grid)


def try_push(grid, sx, sy, mx, my):
    nx = sx + mx
    ny = sy + my

    if grid[nx][ny] == FREE:
        return True
    elif grid[nx][ny] == WALL:
        return False
    elif my == 0:  # up/down
        if grid[nx][ny] == BOX_LEFT:
            return try_push(grid, nx, ny, mx, my) and try_push(grid, nx, ny+1, mx, my)
        elif grid[nx][ny] == BOX_RIGHT:
            return try_push(grid, nx, ny-1, mx, my) and try_push(grid, nx, ny, mx, my)
    elif my == -1:  # left
        if grid[nx][ny] == BOX_RIGHT:
            return try_push(grid, nx, ny-1, mx, my)
    elif my == 1:  # right
        if grid[nx][ny] == BOX_LEFT:
            return try_push(grid, nx, ny+1, mx, my)


def push(grid, sx, sy, mx, my):
    nx = sx + mx
    ny = sy + my

    if grid[nx][ny] == WALL:
        return
    elif grid[nx][ny] == FREE:
        grid[nx][ny] = grid[sx][sy]
        grid[sx][sy] = FREE
    elif my == 0:  # up/down
        if grid[nx][ny] == BOX_LEFT:
            push(grid, nx, ny, mx, my)
            push(grid, nx, ny+1, mx, my)

            grid[nx][ny] = grid[sx][sy]
            grid[sx][sy] = FREE
        elif grid[nx][ny] == BOX_RIGHT:
            push(grid, nx, ny, mx, my)
            push(grid, nx, ny-1, mx, my)

            grid[nx][ny] = grid[sx][sy]
            grid[sx][sy] = FREE
    elif my == -1:  # left
        if grid[nx][ny] == BOX_RIGHT:
            push(grid, nx, ny-1, mx, my)

            grid[nx][ny-1] = grid[nx][ny]
            grid[nx][ny] = grid[sx][sy]
            grid[sx][sy] = FREE
    elif my == 1:  # right
        if grid[nx][ny] == BOX_LEFT:
            push(grid, nx, ny+1, mx, my)

            grid[nx][ny+1] = grid[nx][ny]
            grid[nx][ny] = grid[sx][sy]
            grid[sx][sy] = FREE


def gps(grid):
    sum = 0
    for i, l in enumerate(grid):
        for j, c in enumerate(l):
            if c == BOX or c == BOX_LEFT:
                sum += 100 * i + j

    return sum


def print_grid(grid, rx, ry):
    for i, l in enumerate(grid):
        for j, c in enumerate(l):
            if (i, j) == (rx, ry):
                print("@", end="")
            else:
                print(c, end="")
        print()


if __name__ == "__main__":
    with open("input.txt", "r") as f:
        lines = [line.rstrip() for line in f]

    input = parse(lines)

    a = part_a(input)
    if a:
        print(a)
        submit(a, part="a")

    b = part_b(input)
    if b:
        print(b)
        submit(b,  part="b")
