import re

from aocd import submit


def sum_digits(lines):
    sum = 0
    for line in lines:
        matches = re.findall(r'\d', line)
        sum += int(matches[0]) * 10 + int(matches[-1])

    return sum

def part_a(lines):
    return sum_digits(lines)

def part_b(lines):
    ps = list()
    for line in lines:
        i = 0
        p = ""
        while i < len(line):
            found = False
            for (number, value) in [("one", "1"), ("two", "2"), ("three", "3"), ("four", "4"), ("five", "5"), ("six", "6"), ("seven", "7"), ("eight", "8"), ("nine", "9")]:
                if line[i:i+len(number)] == number:
                    p += value
                    found = True
            if not found:
                p += line[i]
            i+= 1
        ps.append(p)
    
    return sum_digits(ps)


if __name__ == "__main__":
    with open("input.txt", "r") as f:
        lines = [line.rstrip() for line in f]

    a = part_a(lines)
    print(a)
    submit(a, part="a")

    b = part_b(lines)
    print(b)
    submit(b,  part="b")
