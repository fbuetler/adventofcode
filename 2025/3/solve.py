from functools import reduce

from aocd import data, submit


def parse(lines):
    return [[int(d) for d in l] for l in lines]


def max_joltage(batteries, required):
    to_remove = len(batteries) - required
    while to_remove > 0:
        for i in range(len(batteries) - 1):
            if batteries[i] < batteries[i + 1]:
                batteries = batteries[:i] + batteries[i + 1 :]
                break
        to_remove -= 1
    batteries = batteries[:required]

    return reduce(lambda a, b: a * 10 + b, batteries)


def part_a(input):
    res = 0
    for l in input:
        res += max_joltage(l, 2)

    return res


def part_b(input):
    res = 0
    for l in input:
        res += max_joltage(l, 12)

    return res


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
