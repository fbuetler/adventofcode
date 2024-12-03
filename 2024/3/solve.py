import re

from aocd import submit


def parse(lines):
    instrs = list()
    for line in lines:
        pattern = r"mul\((\d+),(\d+)\)|(do)\(\)|(don't)\(\)"    
        for m in re.findall(pattern, line):
            if m[0] != "":
                instrs.append((int(m[0]), int(m[1])))
            elif m[2] != "":
                instrs.append(True)
            else:
                instrs.append(False)
    return instrs

def part_a(input):
    sum = 0
    for mult in input:
        if type(mult) == bool:
            continue
        sum += mult[0] * mult[1]
    return sum

def part_b(input):
    sum = 0
    enabled = True
    for instr in input:
        if type(instr) == bool:
            enabled = instr
        elif enabled:
            sum += instr[0] * instr[1]
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
        submit(b,  part="b")