from aocd import submit


def parse(lines):
    grid = [[c == "#" for c in line] for line in lines]
    return grid


def get_galaxies_in_expanded_space(grid, expansion_factor):
    galaxies = list()
    expand_x = 0
    for i in range(len(grid)):
        expand_x += expansion_factor if not any(grid[i]) else 1

        expand_y = 0
        for j in range(len(grid[i])):
            expand_y += expansion_factor if not any([r[j] for r in grid]) else 1

            if grid[i][j]:
                galaxies.append((expand_x, expand_y))

    return galaxies


def sum_of_shorted_paths(galaxies):
    sum = 0
    for i, a in enumerate(galaxies):
        for j, b in enumerate(galaxies):
            if i >= j:
                continue
            d = abs(a[0] - b[0]) + abs(a[1] - b[1])
            sum += d

    return sum


def part_a(grid):
    galaxies = get_galaxies_in_expanded_space(grid, 2)
    return sum_of_shorted_paths(galaxies)


def part_b(grid):
    galaxies = get_galaxies_in_expanded_space(grid, 1_000_000)
    return sum_of_shorted_paths(galaxies)


if __name__ == "__main__":
    with open("input.txt", "r") as f:
        lines = [line.rstrip() for line in f]

    input = parse(lines)

    a = part_a(input)
    if a:
        print(a)
        submit(a, part="a", day=11)

    b = part_b(input)
    if b:
        print(b)
        submit(b, part="b", day=11)
