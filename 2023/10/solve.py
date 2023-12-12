from aocd import submit

VERTICAL = "|"
HORIZONTAL = "-"
UP_RIGHT = "L"
UP_LEFT = "J"
DOWN_RIGHT = "F"
DOWN_LEFT = "7"
GROUND = "."
START = "S"


def parse(lines):
    grid = list()
    adj = dict()
    for i in range(0, len(lines)):
        row = list()
        for j in range(0, len(lines[0])):
            c = lines[i][j]
            row.append(c)
            l = list()
            # up
            if c in [VERTICAL, UP_LEFT, UP_RIGHT, START] and i > 0:
                up = lines[i - 1][j]
                if up in [VERTICAL, DOWN_LEFT, DOWN_RIGHT, START]:
                    l.append((i - 1, j))
            # down
            if c in [VERTICAL, DOWN_LEFT, DOWN_RIGHT, START] and i < len(lines) - 1:
                down = lines[i + 1][j]
                if down in [VERTICAL, UP_LEFT, UP_RIGHT, START]:
                    l.append((i + 1, j))
            # left
            if c in [HORIZONTAL, UP_LEFT, DOWN_LEFT, START] and j > 0:
                left = lines[i][j - 1]
                if left in [HORIZONTAL, UP_RIGHT, DOWN_RIGHT, START]:
                    l.append((i, j - 1))
            # right
            if c in [HORIZONTAL, UP_RIGHT, DOWN_RIGHT, START] and j < len(lines[0]) - 1:
                right = lines[i][j + 1]
                if right in [HORIZONTAL, UP_LEFT, DOWN_LEFT, START]:
                    l.append((i, j + 1))

            if len(l) > 0:
                adj[(i, j)] = l

            if c == START:
                start = (i, j)
        grid.append(row)

    return start, adj, grid


def dfs(adj: dict[str, list], start):
    visited = set()
    stack = [start]

    while len(stack) > 0:
        v = stack.pop()

        if v not in visited:
            visited.add(v)
            for u in adj[v]:
                if u not in visited:
                    stack.append(u)
                elif u == start:
                    n = len(visited) + 1

    return n, visited


def part_a(input):
    start, adj, _ = input
    cycle_len, _ = dfs(adj, start)
    return (cycle_len - 1) // 2


def part_b(input):
    start, adj, grid = input
    _, visited = dfs(adj, start)

    enclosed = 0
    for i, row in enumerate(grid):
        for j, _ in enumerate(row):
            if (i, j) in visited:
                continue

            crossed = 0
            k = 0
            while i + k < len(grid) and j + k < len(grid[0]):
                c = grid[i + k][j + k]
                if (i + k, j + k) in visited and c != UP_RIGHT and c != DOWN_LEFT:
                    crossed += 1
                k += 1

            if crossed % 2 == 1:
                enclosed += 1

    return enclosed


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
