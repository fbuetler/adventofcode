from aocd import submit


def parse(lines):
    ls = list()
    rs = list()
    for line in lines:
        parts = line.split()
        ls.append(int(parts[0]))
        rs.append(int(parts[1]))
    return (ls, rs)


def part_a(input):
    ls, rs = input
    ls = sorted(ls)
    rs = sorted(rs)
    diff = 0
    for i in range(len(ls)):
        diff += abs(ls[i] - rs[i])
    return diff


def part_b(input):
    ls, rs = input
    sum = 0
    for l in ls:
        for r in rs:
            if l == r:
                sum += l
    return sum


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
