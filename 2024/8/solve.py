from itertools import combinations

from aocd import submit


def parse(lines):
    map = list()
    freqs = dict()
    for i, l in enumerate(lines):
        map.append(list(l))
        for j, c in enumerate(list(l)):
            if c != ".":
                if c in freqs:
                    freqs[c] = freqs[c] + [(i, j)]
                else:
                    freqs[c] = [(i, j)]
    n = len(map)
    m = len(map[0])
    return (freqs, n, m)


def part_a(input):
    freqs, n, m = input
    locs = set()
    for pos in freqs.values():
        for a, b in combinations(pos, 2):
            d = sub(b, a)

            a1 = sub(a, d)
            if in_bounds(n, m, a1):
                locs.add(a1)

            a2 = add(b, d)
            if in_bounds(n, m, a2):
                locs.add(a2)
    return len(locs)


def part_b(input):
    freqs, n, m = input
    locs = set()
    for pos in freqs.values():
        for a, b in combinations(pos, 2):
            d = sub(b, a)

            i = 0
            while True:
                a1 = sub(a, d, factor=i)
                if not in_bounds(n, m, a1):
                    break
                locs.add(a1)
                i += 1

            i = 0
            while True:
                a2 = add(b, d, factor=i)
                if not in_bounds(n, m, a2):
                    break
                locs.add(a2)
                i += 1
    return len(locs)


def add(a, b, factor=1):
    return (a[0] + factor * b[0], a[1] + factor * b[1])


def sub(a, b, factor=1):
    return (a[0] - factor * b[0], a[1] - factor * b[1])


def in_bounds(n, m, pos):
    return 0 <= pos[0] and pos[0] < n and 0 <= pos[1] and pos[1] < m


def print_map(map, locs):
    for i, l in enumerate(map):
        for j, c in enumerate(l):
            if (i, j) in locs:
                print("#", end="")
            elif c != ".":
                print(c, end="")
            else:
                print(".", end="")
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
