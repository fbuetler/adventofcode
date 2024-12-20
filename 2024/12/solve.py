from copy import deepcopy

from aocd import submit

DIRS = [(-1, 0), (0, 1), (1, 0), (0, -1)]


def parse(lines):
    map = [list(l) for l in lines]
    # group by plant type
    plants = dict()
    for i, l in enumerate(map):
        for j, plant_type in enumerate(l):
            if plant_type in plants:
                plants[plant_type] += [(i, j)]
            else:
                plants[plant_type] = [(i, j)]
    return (map, plants)


def part_a(input):
    map, plants = input
    sum = 0
    plants = deepcopy(plants)  # we gonna mangle the locations
    for plant_type, plant_locs in plants.items():
        while len(plant_locs) != 0:
            region = connected_plants_bfs(map, plant_type, plant_locs[0])
            area = len(region)
            perimeter = fences(map, plant_type, region)
            sum += area * perimeter
            for r in region:
                plant_locs.remove(r)
    return sum


def part_b(input):
    map, plants = input
    sum = 0
    plants = deepcopy(plants)  # we gonna mangle the locations
    for plant_type, plant_locs in plants.items():
        while len(plant_locs) != 0:
            region = connected_plants_bfs(map, plant_type, plant_locs[0])
            area = len(region)
            sides = corners(map, plant_type, region)  # corners == sides
            sum += area * sides
            for r in region:
                plant_locs.remove(r)
    return sum


def connected_plants_bfs(map, plant_type, s):
    visited = {s}
    q = [s]

    while len(q) != 0:
        i, j = q.pop(0)
        for (u, v) in DIRS:
            a = i + u
            b = j + v
            if a < 0 or a >= len(map) or b < 0 or b >= len(map[0]):
                continue
            if (a, b) not in visited and map[a][b] == plant_type:
                visited.add((a, b))
                q.append((a, b))

    return list(visited)


def fences(map, plant_type, region):
    f = 0
    for i, j in region:
        for (u, v) in DIRS:
            a = i + u
            b = j + v
            if a < 0 or a >= len(map) or b < 0 or b >= len(map[0]):
                f += 1
            elif map[a][b] != plant_type:
                f += 1
    return f


def corners(map, plant_type, region: list):
    c = 0
    n = len(DIRS)

    def is_plant(x, y):
        if x < 0 or x >= len(map) or y < 0 or y >= len(map[0]):
            return False
        return map[x][y] == plant_type

    for px, py in sorted(region):
        for i in range(n):
            x1, y1 = DIRS[i]
            x2, y2 = DIRS[(i+1) % n]
            if not is_plant(px+x1, py+y1) and not is_plant(px+x2, py+y2):
                # outside corners
                c += 1

            dx, dy = x1 + x2, y1 + y2
            if is_plant(px+x1, py+y1) and is_plant(px+x2, py+y2) and not is_plant(px+dx, py+dy):
                # inside corners
                c += 1

    return c


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
