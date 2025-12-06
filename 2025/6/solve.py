from functools import reduce

from aocd import data, submit


def parse(lines):
    return lines


def compute(problems, ops):
    res = 0
    for i, p in enumerate(problems):
        if ops[i] == "+":
            res += reduce(lambda a, b: a + b, p, 0)
        else:
            res += reduce(lambda a, b: a * b, p, 1)

    return res


def part_a(input):
    problems = [[] for _ in range(len(input[0].split()))]
    for line in input[:-1]:
        for i, p in enumerate(line.split()):
            problems[i].append(int(p))
    ops = input[-1].split()

    return compute(problems, ops)


def part_b(input):
    ops = input[-1].split()
    l = max(map(len, input))
    problems = [[] for _ in range(len(ops))]
    j = 0
    for i in range(l - 1, -1, -1):
        s = ""
        for line in input[:-1]:
            s += line[i].strip()

        if len(s) == 0:
            j += 1
        else:
            problems[j].append(int(s))

    return compute(problems, list(reversed(ops)))


if __name__ == "__main__":
    lines = data.split("\n")

    # with open("example.txt", "r") as f:
    #     lines = [line for line in f]

    input = parse(lines)

    a = part_a(input)
    if a:
        print(a)
        submit(a, part="a")

    b = part_b(input)
    if b:
        print(b)
        submit(b, part="b")
