from collections import Counter
from functools import cache

from aocd import submit


def parse(lines):
    return [int(i) for i in lines[0].split(" ")]


def part_a(input):
    stones = list(input)
    return iterate(stones, 25)


def part_b(input):
    stones = list(input)
    return iterate(stones, 75)


def iterate(stones, max):
    stones = Counter(stones)
    for _ in range(max):
        new_stones = Counter()
        for s, n in stones.items():
            if s == 0:
                new_stones[1] += n
            elif len(str(s)) % 2 == 0:
                h = len(str(s)) // 2
                l = int(str(s)[:h])
                r = int(str(s)[h:])
                new_stones[l] += n
                new_stones[r] += n
            else:
                new_stones[s * 2024] += n
        stones = new_stones

    return stones.total()


def recurse(stones, max):

    @cache
    def deep(s, d):
        if d == 0:
            return 1

        if s == 0:
            return deep(1, d - 1)
        elif len(str(s)) % 2 == 0:
            h = len(str(s)) // 2
            l = int(str(s)[:h])
            r = int(str(s)[h:])
            return deep(l, d - 1) + deep(r, d - 1)
        else:
            return deep(s * 2024, d - 1)

    return sum([deep(s, max) for s in stones])


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
