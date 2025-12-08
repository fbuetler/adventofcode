from collections import defaultdict
from functools import reduce
from itertools import combinations
from math import sqrt

from aocd import data, submit


def parse(lines):
    return [tuple(map(int, line.split(","))) for line in lines]


class UnionFind:
    def __init__(self, size):
        self.parent = list(range(size))

    def find(self, i):
        # root
        if self.parent[i] == i:
            return i

        return self.find(self.parent[i])

    def union(self, i, j):
        irep = self.find(i)
        jrep = self.find(j)
        self.parent[irep] = jrep

    def sets(self):
        groups = defaultdict(set)
        for i in range(len(self.parent)):
            root = self.find(i)
            groups[root].add(i)

        return list(groups.values())


def distance(pair):
    (x1, y1, z1), (x2, y2, z2) = pair
    return sqrt((x1 - x2) ** 2 + (y1 - y2) ** 2 + (z1 - z2) ** 2)


def part_a(input):
    pairs = list(combinations(set(input), 2))
    pairs = sorted(pairs, key=distance)

    uf = UnionFind(len(input))
    for i in range(1000):
        a, b = pairs[i]
        i = input.index(a)
        j = input.index(b)
        uf.union(i, j)

    sets = sorted(list(map(len, uf.sets())), reverse=True)
    res = reduce(lambda a, b: a * b, sets[:3], 1)
    return res


def part_b(input):
    pairs = list(combinations(set(input), 2))
    pairs = sorted(pairs, key=distance)

    uf = UnionFind(len(input))
    for a, b in pairs:
        i = input.index(a)
        j = input.index(b)
        uf.union(i, j)
        if len(uf.sets()) == 1:
            break

    return a[0] * b[0]


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
