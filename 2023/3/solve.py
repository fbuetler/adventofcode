from functools import reduce
from aocd import submit


def parse_grid(lines):
    grid = list()
    for line in lines:
        row = list()
        for c in line:
            if c.isdigit():
                row.append(int(c))
            else:
                row.append(c)
        grid.append(row)

    return grid


def parse_number(grid, i, j):
    digits = list()
    # walk left
    k = -1
    while j + k >= 0 and isinstance(grid[i][j + k], int):
        digits.insert(0, f"{grid[i][j+k]}")
        k -= 1
    # walk right
    k = 0
    while j + k < len(grid[i]) and isinstance(grid[i][j + k], int):
        digits.append(f"{grid[i][j+k]}")
        k += 1
    return int("".join(digits)), k


def get_part_numbers(grid, i, j):
    numbers = list()
    for u in range(-1, 2):
        v = -1
        while v < 2:
            if i + u < 0 or i + u >= len(grid):
                v += 1
            elif j + v < 0 or j + v >= len(grid[0]):
                v += 1
            elif isinstance(grid[i + u][j + v], int):
                n, v_step = parse_number(grid, i + u, j + v)
                numbers.append(n)
                v += v_step
            else:
                v += 1
    return numbers


def parse(lines):
    grid = parse_grid(lines)

    part_numbers = list()
    for i, row in enumerate(grid):
        for j, c in enumerate(row):
            if isinstance(c, str) and c != ".":
                part_numbers.append((c, get_part_numbers(grid, i, j)))

    return part_numbers


def part_a(part_numbers):
    s = 0
    for _, numbers in part_numbers:
        s += sum(numbers)
    return s


def part_b(part_numbers):
    s = 0
    for symbol, numbers in part_numbers:
        if symbol == "*" and len(numbers) == 2:
            s += reduce((lambda x, y: x * y), numbers)
    return s


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
