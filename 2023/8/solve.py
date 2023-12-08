import math
import re
from functools import reduce

from aocd import submit


def parse(lines):
    directions = [c for c in lines[0]]
    network = {}
    for line in lines[2:]:
        m = re.findall(r"(\w+) = \((\w+), (\w+)\)", line)[0]
        network[m[0]] = (m[1], m[2])
    return (directions, network)


def get_steps(start_position, end_position_suffix, directions, network):
    steps = 0
    p = start_position
    while not p.endswith(end_position_suffix):
        if directions[steps % len(directions)] == "L":
            p = network[p][0]
        else:
            p = network[p][1]
        steps += 1

    return steps


def part_a(input):
    directions, network = input
    return get_steps("AAA", "ZZZ", directions, network)


def part_b(input):
    directions, network = input
    start_positions = [p for p in network.keys() if p.endswith("A")]
    lcms = [get_steps(p, "Z", directions, network) for p in start_positions]
    lcm = reduce(lambda x, y: math.lcm(x, y), lcms)
    return lcm


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
