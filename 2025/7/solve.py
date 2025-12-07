from functools import cache

from aocd import data, submit


def parse(lines):
    n = len(lines)
    m = len(lines[0])
    splitters = []
    for i, line in enumerate(lines):
        for j, c in enumerate(line):
            if c == "S":
                start = (i, j)
            elif c == "^":
                splitters.append((i, j))

    return (n, m, start, splitters)


def part_a(input):
    n, m, start, splitters = input

    res = 0
    i = start[0]
    beams = {start[1]}
    while i < n:
        next = set()
        for b in beams:
            if (i + 1, b) in splitters:
                next.add(b - 1)
                next.add(b + 1)
                res += 1
            else:
                next.add(b)

        beams = next
        i += 1

    return res


def part_b(input):
    n, m, start, splitters = input

    @cache
    def send(particle):
        i, j = particle

        if i == n:
            return 0

        if (i + 1, j) in splitters:
            return 1 + send((i + 1, j - 1)) + send((i + 1, j + 1))
        else:
            return send((i + 1, j))

    return 1 + send(start)


if __name__ == "__main__":
    lines = data.split("\n")

    # with open("example.txt", "r") as f:
    #     lines = [line.rstrip() for line in f]

    input = parse(lines)

    a = part_a(input)
    if a:
        print(a)
        submit(a, part="a")

    b = part_b(input)
    if b:
        print(b)
        submit(b, part="b")
