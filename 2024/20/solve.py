from heapq import heappop, heappush

from aocd import submit

DIRS = [(-1, 0), (0, 1), (1, 0), (0, -1)]


def parse(lines):
    grid = list()
    for i, line in enumerate(lines):
        l = list()
        for j, c in enumerate(line):
            if c == "S":
                sx, sy = i, j
                l.append(".")
            elif c == "E":
                ex, ey = i, j
                l.append(".")
            else:
                l.append(c)
        grid.append(l)
    return grid, (sx, sy), (ex, ey)


def part_a(input):
    grid, start, end = input
    return cheat(grid, start, end, 2)


def part_b(input):
    grid, start, end = input
    return cheat(grid, start, end, 20)


def cheat(grid, start, end, max_cheat):
    cheats = {}
    dist, path = shortest_path(grid, start, end)

    for u in path:
        for v in path:
            cheat_length = abs(u[0]-v[0]) + abs(u[1]-v[1])
            if cheat_length > max_cheat:
                continue

            d = is_shortcut(dist, u, v, cheat_length)
            if d > 0:
                save_cheat(cheats, d, (u, v))

    return sum([len(v) for k, v in cheats.items() if k >= 100])


def is_shortcut(dist, u, v, w):
    if not u in dist or not v in dist:
        return 0, None
    return dist[v] - (dist[u] + w)


def save_cheat(cheats, saved, start_end):
    start_end = sorted(start_end)
    if saved in cheats:
        if not start_end in cheats[saved]:
            cheats[saved].append(start_end)
    else:
        cheats[saved] = [start_end]


def shortest_path(grid, start, end):
    n = len(grid)
    m = len(grid[0])

    dist = {}
    dist[start] = 0
    pq = []
    heappush(pq, (0, start))
    prev = {}

    while len(pq) != 0:
        d, pos = heappop(pq)
        if dist[pos] < d:
            continue

        for dir in DIRS:
            step = (pos[0] + dir[0], pos[1] + dir[1])
            if step[0] < 0 or step[0] >= n or step[1] < 0 or step[1] >= m:
                continue
            if grid[step[0]][step[1]] != "#" and (step not in dist or dist[step] > d + 1):
                dist[step] = d + 1
                heappush(pq, (d + 1, step))
                prev[step] = pos

    path = []
    n = end
    while n != start:
        path.append(n)
        n = prev[n]
    path.append(start)

    return dist, list(reversed(path))


def print_grid(grid, start, end, path, shortcut):
    GREEN = '\033[92m'
    RED = '\033[91m'
    RESET = '\033[0m'
    BOLD = '\033[1m'
    for i, l in enumerate(grid):
        for j, c in enumerate(l):
            s = ""
            if (i, j) == start:
                s += BOLD
                c = "S"
            if (i, j) == end:
                s += BOLD
                c = "E"
            if (i, j) in shortcut:
                s += RED
            if (i, j) in path and not (i, j) in shortcut:
                s += GREEN

            s += c + RESET
            print(s, end="")
        print("", flush=True)


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
