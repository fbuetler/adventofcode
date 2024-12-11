from aocd import submit


def parse(lines):
    heads = list()
    tails = list()
    map = list()
    for i, l in enumerate(lines):
        hs = list()
        for j, h in enumerate(list(l)):
            hs.append(int(h) if h != "." else 100)
            if h == "0":
                heads.append((i, j))
            elif h == "9":
                tails.append((i, j))
        map.append(hs)
    return (map, heads, tails)


def part_a(input):
    map, heads, tails = input
    sum = 0
    for head in heads:
        visited, paths = shortest_path(map, head)
        for (a, b) in tails:
            if visited[a][b]:
                sum += 1
    return sum


def part_b(input):
    map, heads, tails = input
    sum = 0
    for head in heads:
        visited, paths = shortest_path(map, head)
        sum += len(paths)
    return sum


def shortest_path(map, head):
    n = len(map)
    m = len(map[0])

    q = [head]
    paths = [[head]]
    visited = [[False for _ in range(m)] for _ in range(n)]

    while len(q) != 0:
        a, b = q.pop()
        visited[a][b] = True
        forks = []

        for i, j in [(-1, 0), (0, 1), (1, 0), (0, -1)]:
            if a+i < 0 or a+i > n - 1 or b+j < 0 or b+j > m - 1:
                continue
            if map[a+i][b+j] - map[a][b] == 1:
                q.append((a+i, b+j))
                forks.append((a+i, b+j))

        new_paths = list()
        for path in paths:
            if path[-1] == (a, b) and len(forks) > 0:
                for f in forks:
                    new_paths.append(path + [f])
            else:
                new_paths.append(path)
        if len(new_paths) != 0:
            paths = new_paths

    return visited, list(filter(lambda l: len(l) == 10, paths))


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
