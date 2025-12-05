from aocd import data, submit


def parse(lines):
    rs = []
    ings = []
    i = 0
    for i in range(len(lines)):
        line = lines[i]
        if line == "":
            break
        a, b = int(line.split("-")[0]), int(line.split("-")[1])
        rs.append((a, b))

    ings = list(map(int, lines[i + 1 :]))

    return (rs, ings)


def is_fresh(rs, ing):
    for a, b in rs:
        if a <= ing and ing <= b:
            return True

    return False


def part_a(input):
    rs, ings = input
    res = 0
    for ing in ings:
        if is_fresh(rs, ing):
            res += 1

    return res


def part_b(input):
    rs, ings = input
    rs = sorted(rs, key=lambda r: r[0])

    merged = []
    a, b = rs[0]
    for x, y in rs[1:]:
        if x <= b:
            b = max(b, y)
        else:
            merged.append((a, b))
            a, b = x, y

    merged.append((a, b))

    return sum(b - a + 1 for a, b in merged)


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
