from heapq import heappop, heappush

from aocd import submit

SIZE = 70 + 1


def parse(lines):
    return [(int(line.split(",")[0]), int(line.split(",")[1])) for line in lines]


def part_a(input):
    return shortest_path(input[:1024])


def part_b(input):
    i = binary_search(input, 0, len(input))
    c = input[i]
    return f"{c[0]},{c[1]}"


def binary_search(corrupted, low, high):
    while low <= high:
        mid = (low + high) // 2

        a = shortest_path(corrupted[:mid])
        b = shortest_path(corrupted[:mid+1])
        if a is not None and b is None:
            return mid
        elif a is not None and b is not None:
            low = mid + 1
        else:
            high = mid - 1

    return None


def shortest_path(corrupted):
    start = (0, 0)

    dist = {}
    pq = []
    dist[start] = 0
    heappush(pq, (0, start))

    while len(pq) != 0:
        d, pos = heappop(pq)
        if dist[pos] < d:
            continue

        for dir in [(-1, 0), (0, 1), (1, 0), (0, -1)]:
            step = (pos[0] + dir[0], pos[1] + dir[1])
            if step[0] < 0 or step[0] >= SIZE or step[1] < 0 or step[1] >= SIZE:
                continue
            if step not in corrupted and (step not in dist or dist[step] > d + 1):
                dist[step] = d + 1
                heappush(pq, (d + 1, step))

    return dist[(SIZE-1, SIZE-1)] if (SIZE-1, SIZE-1) in dist else None


def print_memory(corrupted):
    for i in range(SIZE):
        for j in range(SIZE):
            if (j, i) in corrupted:
                print("#", end="")
            else:
                print(".", end="")
        print()


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
