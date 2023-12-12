from functools import cache

from aocd import submit

OPERATIONAL = "."
DAMAGED = "#"
UNKNOWN = "?"


def parse(lines):
    records = list()
    for line in lines:
        parts = line.split(" ")
        springs = parts[0]
        parts = parts[1].split(",")
        groups = tuple(map(int, parts))
        records.append((springs, groups))

    return records


@cache
def reconstruct(springs, groups, processed=0):
    if len(springs) == 0:
        if processed > 0:
            return len(groups) == 1 and groups[0] == processed
        else:
            return len(groups) == 0

    possibilities = 0
    head, tail = springs[0], springs[1:]
    if head == OPERATIONAL:
        if processed > 0:
            if len(groups) > 0 and groups[0] == processed:
                possibilities += reconstruct(tail, groups[1:], 0)
        else:
            possibilities += reconstruct(tail, groups, 0)
    elif head == DAMAGED:
        possibilities += reconstruct(tail, groups, processed + 1)
    elif head == UNKNOWN:
        possibilities += reconstruct(OPERATIONAL + tail, groups, processed)
        possibilities += reconstruct(DAMAGED + tail, groups, processed)

    return possibilities


def part_a(records):
    return sum(reconstruct(springs, groups) for springs, groups in records)


def part_b(records):
    return sum(
        reconstruct("?".join([springs] * 5), groups * 5) for springs, groups in records
    )


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
