import re

from aocd import submit
from z3 import BitVec, Solver, sat


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
    return registers["A"], registers["B"], registers["C"], program


def part_a(input):
    a, b, c, program = input
    output = calculate(a, b, c, program)
    return ",".join(list(map(str, output)))


def calculate(a, b, c, program):
    i = 0
    output = list()
    while i < len(program):
        if i+1 >= len(program):
            break

        instr = program[i]
        op = program[i+1]

        # print(f"{instr} {op} - {a} {b} {c}")

        if instr == 3:  # jnz
            if a != 0:
                i = op
                continue
        elif instr == 2:  # bst
            b = combo(a, b, c, op) % 8
        elif instr == 1:  # bxl
            b ^= op
        elif instr == 4:  # bxc
            b ^= c
        elif instr == 0:  # adv
            d = 2 ** combo(a, b, c, op)
            a //= d
        elif instr == 6:  # bdv
            d = 2 ** combo(a, b, c, op)
            b = a // d
        elif instr == 7:  # cdv
            d = 2 ** combo(a, b, c, op)
            c = a // d
        elif instr == 5:  # out
            n = combo(a, b, c, op)
            output.append(n % 8)
        else:
            raise ValueError("invalid instr")

        i += 2

    return output


def combo(a, b, c, op):
    if 0 <= op and op <= 3:
        return op
    elif op == 4:
        return a
    elif op == 5:
        return b
    elif op == 6:
        return c
    raise ValueError("invalid combo")


"""
the program ends with 0, 3 i.e. jumping back to the start of the program, until a == 0
the following (instr, op) pairs are executed repeatedly
2 4 # bst a
1 7 # bxl 7
7 5 # cdv b
0 3 # adv 3
1 7 # bxl 7
4 1 # bxc 1
5 5 # out b
3 0 # jnz 0
translating into logic formulas give the loop below
"""


def part_b(input):
    a, b, c, program = input

    a = BitVec("A", 64)
    constraints = list()
    for p in program:
        b = a % 8
        b ^= 7
        c = a >> b
        a >>= 3
        b ^= 7
        b ^= c
        constraints.append(b % 8 == p)

    constraints.append(a == 0)

    s = Solver()
    s.add(constraints)
    assert s.check() == sat
    m = s.model()
    for v in m:
        return m[v].as_long()


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
