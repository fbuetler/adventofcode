from aocd import submit

N = 5


def parse(lines):
    locks = list()
    keys = list()

    i = 0
    while i < len(lines):
        form = [0] * N

        for j in range(1, N + 1):
            for k, c in enumerate(lines[i + j]):
                if c == "#":
                    form[k] += 1

        if lines[i] == N * "#":
            locks.append(form)
        else:
            keys.append(form)

        i += N + 3  # omit both indicator rows and empty line

    return keys, locks


def part_a(input):
    keys, locks = input
    sum = 0
    for k in keys:
        for l in locks:
            overlaps = False
            for i in range(N):
                if k[i] + l[i] > 5:
                    overlaps = True
                    break
            if not overlaps:
                sum += 1

    return sum


def part_b(input):
    return "merry christmas"


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
