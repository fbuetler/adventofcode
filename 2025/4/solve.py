from aocd import data, submit


def parse(lines):
    return [[p for p in line] for line in lines]


def is_accessible(grid, i, j):
    c = 0
    for m in range(-1, 2, 1):
        for n in range(-1, 2, 1):
            if m == 0 and n == 0:
                continue
            elif i + m < 0 or i + m >= len(grid):
                continue
            elif j + n < 0 or j + n >= len(grid[0]):
                continue
            elif grid[i + m][j + n] == "@":
                c += 1

    return c < 4


def part_a(input):
    res = 0
    for i, line in enumerate(input):
        for j, p in enumerate(line):
            if p != "@":
                continue

            if is_accessible(input, i, j):
                res += 1

    return res


def part_b(input):
    res = 0
    removed = True
    while removed:
        removed = False

        next = [[p for p in line] for line in input]
        for i in range(len(input)):
            for j in range(len(input[0])):
                if input[i][j] != "@":
                    continue

                if is_accessible(input, i, j):
                    next[i][j] = "."
                    res += 1
                    removed = True
        input = next

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
