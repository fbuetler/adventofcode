import re
from aocd import submit


def parse(lines):
    times = [int(t) for t in re.findall(r"\d+", lines[0])]
    distances = [int(d) for d in re.findall(r"\d+", lines[1])]
    return (times, distances)


def find_hold_time_naive(time, distance):
    wins = list()
    for hold_time in range(time):
        reached = hold_time * (time - hold_time)
        if reached > distance:
            wins.append(hold_time)
    return len(wins)


def find_hold_time(time, distance):
    start = 0
    end = time
    while start <= end:
        hold_time = (start + end) // 2

        reached = hold_time * (time - hold_time)
        if reached > distance:
            end = hold_time - 1
        elif reached < distance:
            start = hold_time + 1
        else:
            return hold_time + 1
    return hold_time


def part_a(input):
    p = 1
    for time, distance in zip(*input):
        h = find_hold_time(time, distance)
        p *= time - 2 * h + 1
    return p


def part_b(input):
    times, distances = input
    time = int("".join([str(t) for t in times]))
    distance = int("".join([str(d) for d in distances]))

    h = find_hold_time(time, distance)
    return time - 2 * h + 1


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
