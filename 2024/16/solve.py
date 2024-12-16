import sys
from heapq import heappop, heappush

from aocd import submit


def parse(lines):
    n = len(lines)
    m = len(lines[0])
    walls = set()
    for i, line in enumerate(lines):
        for j, l in enumerate(line):
            if l == "S":
                start = (i, j)
            if l == "E":
                stop = (i, j)
            if l == "#":
                walls.add((i, j))
    return n, m, walls, start, stop


# up, right, down, left
ALL_DIRECTIONS = [(-1, 0), (0, 1), (1, 0), (0, -1)]


def part_a(input):
    _, _, walls, start, stop = input
    dist = shortest_paths(walls, [(start, (0, 1))])
    return min([dist[stop, dir] for dir in ALL_DIRECTIONS])


def part_b(input):
    n, m, walls, start, stop = input
    dist_start = shortest_paths(walls, [(start, (0, 1))])
    dist_stop = shortest_paths(walls, [(stop, f) for f in ALL_DIRECTIONS])
    shortest = min([dist_start[stop, dir] for dir in ALL_DIRECTIONS])

    seats = set()
    for i in range(n):
        for j in range(m):
            for a, b in ALL_DIRECTIONS:
                pos_start = ((i, j), (a, b))
                pos_stop = ((i, j), (-a, -b))
                if pos_start in dist_start and pos_stop in dist_stop and dist_start[pos_start] + dist_stop[pos_stop] == shortest:
                    seats.add((i, j))
    return len(seats)


WALKING = 1
TURNING = 1000


def shortest_paths(walls, starts):
    dist = {}
    pq = []
    for pos, facing in starts:
        dist[(pos, facing)] = 0
        heappush(pq, (0, pos, facing))

    while len(pq) != 0:
        d, pos, facing = heappop(pq)
        if dist[(pos, facing)] < d:
            # a faster way is already known
            continue

        # turning
        for turning in ALL_DIRECTIONS:
            if turning == facing:
                # cant rotate in the direction we are facing
                continue
            if (pos, turning) not in dist or dist[(pos, turning)] > d + TURNING:
                dist[(pos, turning)] = d + TURNING
                heappush(pq, (d + TURNING, pos, turning))

        # walking
        step = (pos[0] + facing[0], pos[1] + facing[1])
        if step not in walls and ((step, facing) not in dist or dist[(step, facing)] > d + WALKING):
            dist[(step, facing)] = d + WALKING
            heappush(pq, (d + WALKING, step, facing))

    return dist


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
