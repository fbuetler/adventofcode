import copy

from aocd import submit

FREE = "."
OBSTACLE = "#"

OFFSETS = [
    (-1, 0),  # up
    (0, 1),  # right
    (1, 0),  # down
    (0, -1),  # left
]


def parse(lines):
    map = [list(l) for l in lines]
    pos = [(i, l.index("^")) for i, l in enumerate(lines) if "^" in l][0]
    map[pos[0]][pos[1]] = FREE  # mark start as free
    return (map, pos)


def part_a(input):
    map, pos = input
    return len(run(map, pos))


def part_b(input):
    map, pos = input
    sum = 0

    visited = run(map, pos)
    for i, j in visited:
        if (i, j) == pos:
            continue
        mc = copy.deepcopy(map)
        mc[i][j] = OBSTACLE
        if is_looping(mc, pos):
            sum += 1
    return sum


def run(map, pos):
    n = len(map)
    m = len(map[0])
    facing = 0
    visited = list()
    while True:
        if not pos in visited:
            visited.append(pos)

        if leaving(pos, n, m, facing):
            break
        if is_blocked(map, pos, facing):
            facing = turn_right(facing)
        else:
            pos = step(pos, facing)

    return visited


def is_looping(map, pos):
    n = len(map)
    m = len(map[0])
    facing = 0
    visited = set()
    while True:
        if leaving(pos, n, m, facing):
            return False
        if is_blocked(map, pos, facing):
            facing = turn_right(facing)
            if pos in visited:
                return True
            else:
                visited.add(pos)
        else:
            pos = step(pos, facing)


def leaving(pos, n, m, facing):
    offset = OFFSETS[facing]
    return (
        pos[0] + offset[0] < 0
        or pos[0] + offset[0] >= n
        or pos[1] + offset[1] < 0
        or pos[1] + offset[1] >= m
    )


def is_blocked(map, pos, facing):
    offset = OFFSETS[facing]
    return map[pos[0] + offset[0]][pos[1] + offset[1]] == OBSTACLE


def turn_right(facing):
    return (facing + 1) % 4


def step(pos, facing):
    offset = OFFSETS[facing]
    return (pos[0] + offset[0], pos[1] + offset[1])


def print_map(map, pos, facing):
    for i, line in enumerate(map):
        for j, col in enumerate(line):
            if (i, j) == pos:
                print("*", end="")
            else:
                print(col, end="")
        print()
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
        submit(b, part="b")
