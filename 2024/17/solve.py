import re
from copy import deepcopy

from aocd import submit


def parse(lines):
    register_pattern = re.compile(r"Register (A|B|C): (\d+)")
    program_pattern = re.compile(r"Program: ([\d,]+)")
    registers = dict()
    program = list
    for line in lines:
        m = re.match(register_pattern, line)
        if m:
            registers[m.group(1)] = int(m.group(2))
        m = re.match(program_pattern, line)
        if m:
            program = list(map(int, m.group(1).split(",")))
    return registers, program


A = "A"
B = "B"
C = "C"


def part_a(input):
    regs, program = input
    regs = deepcopy(regs)
    print(regs)
    return calculate(regs, program)


def part_b(input):
    regs, program = input


def calculate(regs, program):
    i = 0
    output = list()
    while i < len(program):
        if i+1 >= len(program):
            break

        instr = program[i]
        op = program[i+1]

        if instr == 3:
            if regs[A] != 0:
                i = op
                continue

        elif instr == 2:
            regs[B] = combo(regs, op) % 8

        elif instr == 1:
            regs[B] = regs[B] ^ op
        elif instr == 4:
            regs[B] = regs[B] ^ regs[C]

        elif instr == 0:
            n = regs[A]
            d = 2 ** combo(regs, op)
            regs[A] = n // d
        elif instr == 6:
            n = regs[A]
            d = 2 ** combo(regs, op)
            regs[B] = n // d
        elif instr == 7:
            n = regs[A]
            d = 2 ** combo(regs, op)
            regs[C] = n // d

        elif instr == 5:
            n = combo(regs, op)
            output.append(n % 8)

        else:
            raise ValueError("invalid instr")

        i += 2

    s = ",".join(list(map(str, output)))
    return s


def combo(regs, op):
    if 0 <= op and op <= 3:
        return op
    elif op == 4:
        return regs[A]
    elif op == 5:
        return regs[B]
    elif op == 6:
        return regs[C]
    raise ValueError("invalid combo")


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
