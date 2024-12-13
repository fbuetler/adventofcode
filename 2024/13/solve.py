import re

import numpy as np
from aocd import submit


def parse(lines):
    button_pattern = re.compile(r"^Button (A|B): X\+(\d+), Y\+(\d+)")
    prize_pattern = re.compile(r"^Prize: X=(\d+), Y=(\d+)")
    machines = list()
    machine = list()
    for line in lines:
        m = re.match(button_pattern, line)
        if m:
            machine.append((int(m.group(2)), int(m.group(3))))
        m = re.match(prize_pattern, line)
        if m:
            machine.append((int(m.group(1)), int(m.group(2))))
            machines.append(machine)
            machine = list()
    return machines


def part_a(input):
    sum = 0
    for m in input:
        (a_x, a_y), (b_x, b_y), (p_x, p_y) = m
        sum += play(a_x, a_y, b_x, b_y, p_x, p_y)
    return sum


def part_b(input):
    sum = 0
    for m in input:
        (a_x, a_y), (b_x, b_y), (p_x, p_y) = m
        sum += play(a_x, a_y, b_x, b_y,
                    10_000_000_000_000+p_x, 10_000_000_000_000+p_y)

    return sum


def play(a_x, a_y, b_x, b_y, p_x, p_y):
    u = [a_x, b_x]
    v = [a_y, b_y]
    w = [p_x, p_y]
    z = np.linalg.solve(np.array([u, v]), np.array(w))
    z = z.round().astype(int)
    if np.allclose(np.dot([u, v], z), w, rtol=0):
        return 3*z[0] + z[1]
    return 0


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
