from aocd import submit


def parse(lines):
    grid = list()
    moves = list()
    boxes = set()
    switched = False
    for i, line in enumerate(lines):
        if not switched:
            grid.append(list(line))
            for j, c in enumerate(list(line)):
                if c == "@":
                    robot = (i, j)
                elif c == "O":
                    boxes.add((i, j))
        else:
            moves += list(line)

        if len(line) == 0:
            switched = True

    return grid, boxes, moves, robot


TO_DIR = {
    "^": (-1, 0),
    ">": (0, 1),
    "v": (1, 0),
    "<": (0, -1),
}

WALL = "#"
BOX = "O"
FREE = "."


def part_a(input):
    grid, boxes, moves, (rx, ry) = input
    grid[rx][ry] = FREE
    for m in moves:
        # print(f"Move {m}:")
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
                grid[sx][sy] = BOX
                grid[bx][by] = FREE
                rx, ry = rx+mx, ry+my
        else:
            # walk
            rx, ry = rx+mx, ry+my

    print_grid(grid, rx, ry)

    sum = 0
    for i, l in enumerate(grid):
        for j, c in enumerate(l):
            if c == BOX:
                sum += 100 * i + j

    return sum


def next_space(grid, bx, by, mx, my):
    while grid[bx][by] != WALL:
        if grid[bx][by] == FREE:
            return (bx, by)
        bx, by = bx+mx, by+my
    return None


def print_grid(grid, rx, ry):
    for i, l in enumerate(grid):
        for j, c in enumerate(l):
            if (i, j) == (rx, ry):
                print("@", end="")
            else:
                print(c, end="")
        print()


def part_b(input):
    pass


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
