from aocd import data, submit


def parse(lines):
    return lines

def part_a(input):
    return None

def part_b(input):
    return None

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