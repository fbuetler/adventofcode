from itertools import combinations

from aocd import data, submit
from shapely import Polygon, box


def parse(lines):
    return [tuple(map(int, line.split(","))) for line in lines]


def area(pair):
    (x1, y1), (x2, y2) = pair
    return (abs(x1 - x2) + 1) * (abs(y1 - y2) + 1)


def part_a(input):
    pairs = list(combinations(set(input), 2))
    pairs = list(map(area, pairs))
    pairs = sorted(pairs, reverse=True)
    return pairs[0]


def part_b(input):
    input = list(map(lambda a: (a[1], a[0]), input))  # switcheroo

    polygon = Polygon(input)
    pairs = list(combinations(set(input), 2))
    boxes = list(
        map(
            lambda ab: box(
                min(ab[0][0], ab[1][0]),
                min(ab[0][1], ab[1][1]),
                max(ab[0][0], ab[1][0]),
                max(ab[0][1], ab[1][1]),
            ),
            pairs,
        )
    )
    contained = list(filter(polygon.contains, boxes))
    bounds = list(map(lambda b: b.bounds, contained))
    areas = list(map(lambda b: (b[2] - b[0] + 1) * (b[3] - b[1] + 1), bounds))
    areas = sorted(areas, reverse=True)

    return areas[0]


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
