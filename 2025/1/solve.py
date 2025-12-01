from aocd import data, submit


def parse(lines):
    return [(line[:1], int(line[1:])) for line in lines]

def part_a(input):
    count = 0 
    c = 50
    for d, n in input:
        if d == "L":
            c -= n 
        else:
            c += n
        c %= 100

        if c == 0:
            count += 1

    return count

def part_b(input):
    count = 0
    c = 50

    for d, n in input:
        if d == "L":
            s = -1
        else:
            s = 1

        for _ in range(n):
            c = (c + s) % 100
            if c == 0:
                count += 1

    return count

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
        submit(b,  part="b")