import re
import sys

from aocd import submit


def parse(lines):
    maps = list()
    mappings = list()

    seeds_pattern = re.compile(r"^seeds:\s(?P<seeds>[\d\s]+)$")
    range_pattern = re.compile(r"^(?P<to>\d+) (?P<from>\d+) (?P<len>\d+)")
    map_pattern = re.compile(r"^(?P<map_name>[\w\-]+)\smap:$")
    for line in lines:
        m = re.match(seeds_pattern, line)
        if m:
            seeds = list(map(int, m.group("seeds").split()))
        m = re.match(map_pattern, line)
        if m:
            if len(mappings) > 0:
                maps.append(mappings)
                mappings = list()
        m = re.match(range_pattern, line)
        if m:
            ranges = (int(m.group("to")), int(m.group("from")), int(m.group("len")))
            mappings.append(ranges)

    maps.append(mappings)
    return (seeds, maps)


def map_ranges_naive(ranges, maps):
    min_location = sys.maxsize
    for seed_start, seed_end in ranges:
        for s in range(seed_start, seed_end):
            for mappings in maps:
                for m in mappings:
                    (dest, src, len) = m
                    if s in range(src, src + len):
                        s = s - src + dest
                        break
            min_location = min(min_location, s)
    return min_location


def map_ranges(ranges, maps):
    for mappings in maps:
        mapped_ranges = list()
        for m in mappings:
            (dest, src, len) = m
            src_end = src + len

            next_ranges = list()
            while ranges:
                (start, end) = ranges.pop()
                left = (start, min(end, src))
                middle = (max(start, src), min(src_end, end))
                right = (max(src_end, start), end)

                if left[0] < left[1]:
                    next_ranges.append(left)
                if middle[0] < middle[1]:
                    mapped_ranges.append(
                        (middle[0] - src + dest, middle[1] - src + dest)
                    )
                if right[0] < right[1]:
                    next_ranges.append(right)
            ranges = next_ranges

        ranges = mapped_ranges + ranges
    return min(map(lambda l: l[0], ranges))


def part_a(input):
    seeds, maps = input
    ranges = [(seed, seed + 1) for seed in seeds]
    return map_ranges(ranges, maps)


def part_b(input):
    seeds, maps = input
    ranges = [(seed, seed + len) for (seed, len) in zip(seeds[::2], seeds[1::2])]
    return map_ranges(ranges, maps)


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
