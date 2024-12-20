from typing import Dict, List

from aocd import submit


def parse(lines):
    towels = lines[0].split(", ")
    designs = [l for l in lines[2:]]
    return (towels, designs)


def part_a(input):
    towels, designs = input

    dp = {}
    sum = 0
    for d in designs:
        if count(towels, d, dp) > 0:
            sum += 1

    return sum


def part_b(input):
    towels, designs = input

    dp = {}
    sum = 0
    for d in designs:
        sum += count(towels, d, dp)

    return sum


def count(towels: List[str], design: str, dp):
    if design in dp:
        return dp[design]

    if len(design) == 0:
        return 1

    sum = 0
    for t in towels:
        if design.startswith(t):
            sum += count(towels, design[len(t):], dp)

    dp[design] = sum
    return sum


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
