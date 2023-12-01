from aocd import submit

def part_a(lines):
    pass

def part_b(lines):
    pass

if __name__ == "__main__":
    with open("input.txt", "r") as f:
        lines = [line.rstrip() for line in f]

    a = part_a(lines)
    print(a)
    submit(a, part="a")

    b = part_b(lines)
    print(b)
    submit(b,  part="b")