from copy import deepcopy

from aocd import submit


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
            sides = all_sides(map, plant_type, region)
            sum += area * sides
            for r in region:
                plant_locs.remove(r)
    return sum


def connected_plants_bfs(map, plant_type, s):
    visited = [[False for _ in range(len(map[0]))] for _ in range(len(map))]
    q = [s]
    visited[s[0]][s[1]] = True

    while len(q) != 0:
        i, j = q.pop(0)
        for (u, v) in [(0, -1), (1, 0), (0, 1), (-1, 0)]:
            a = i + u
            b = j + v
            if a < 0 or a >= len(map) or b < 0 or b >= len(map[0]):
                continue
            if not visited[a][b] and map[a][b] == plant_type:
                visited[a][b] = True
                q.append((a, b))

    return [(i, j) for j in range(len(visited[0])) for i in range(len(visited)) if visited[i][j]]


def fences(map, plant_type, region):
    f = 0
    for i, j in region:
        for (u, v) in [(0, -1), (1, 0), (0, 1), (-1, 0)]:
            a = i + u
            b = j + v
            if a < 0 or a >= len(map) or b < 0 or b >= len(map[0]):
                f += 1
            elif map[a][b] != plant_type:
                f += 1
    return f


def all_sides(m, plant_type, region: list):
    return sides_horizontal(m, plant_type, region) + sides_vertical(m, plant_type, region)


def sides_horizontal(m, plant_type, region):
    min_x = min(list(map(lambda x: x[0], region)))
    max_x = max(list(map(lambda x: x[0], region)))

    # scan horizontal sides
    sides_line = 0
    for i in range(min_x, max_x+1, 1):
        # go line by line
        line = list(filter(lambda x: x[0] == i, region))

        # find borders
        up = list()
        down = list()
        for u, v in line:
            if u == 0 or m[u-1][v] != plant_type:
                # border left/right
                up.append(v)
            if i == len(m) - 1 or m[u+1][v] != plant_type:
                # neighbouring another plant type up/down
                down.append(v)

        # sum up sides facing up/down
        sides_line += sides_facing(up) + sides_facing(down)

    return sides_line


def sides_vertical(m, plant_type, region):
    min_y = min(list(map(lambda x: x[1], region)))
    max_y = max(list(map(lambda x: x[1], region)))

    # scan vertical sides
    sides_col = 0
    for i in range(min_y, max_y+1, 1):
        # go column by column
        col = list(filter(lambda x: x[1] == i, region))

        # find borders
        left = list()
        right = list()
        for u, v in col:
            if v == 0 or m[u][v-1] != plant_type:
                # border left/right
                left.append(u)
            if i == len(m[0]) - 1 or m[u][v+1] != plant_type:
                # neighbouring another plant type left/right
                right.append(u)

        # sum up sides facing left/right
        sides_col += sides_facing(left) + sides_facing(right)

    return sides_col


def sides_facing(l):
    s = 0
    for j in range(len(l)):
        if j == 0:
            # fence-post
            s += 1
        elif l[j] - l[j-1] != 1:
            # neighbouring
            s += 1
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
        submit(b,  part="b")
