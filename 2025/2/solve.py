from aocd import data, submit


def parse(lines):
    return list(
        map(lambda r: (int(r.split("-")[0]), int(r.split("-")[1])), lines[0].split(","))
    )


def is_invalid_2(n):
    s = str(n)
    l = len(s)
    if l % 2 == 1:
        return False

    m = l // 2
    a = s[:m]
    b = s[m:]

    return a == b


def is_invalid_n(n):
    s = str(n)
    l = len(s)

    for i in range(1, (l // 2) + 1):
        if l % i != 0:
            continue

        m = True
        a = s[:i]
        for j in range((l // i) - 1):
            b = s[i + j * i : i + (j + 1) * i]
            if a != b:
                m = False
                break

        if m:
            return True

    return False


def find_invalids(a, b, is_invalid):
    invalids = []
    for n in range(a, b + 1):
        if is_invalid(n):
            invalids.append(n)
    return invalids


def part_a(input):
    s = 0
    for a, b in input:
        invalids = find_invalids(a, b, is_invalid_2)
        # print(len(invalids), invalids)
        s += sum(invalids)
    return s


def part_b(input):
    s = 0
    for a, b in input:
        invalids = find_invalids(a, b, is_invalid_n)
        s += sum(invalids)
    return s


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
