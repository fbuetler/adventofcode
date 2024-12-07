from aocd import submit


def parse(lines):
    eqs = list()
    for line in lines:
        ps = line.split(": ")
        result = int(ps[0])
        values = list(map(lambda x: int(x), ps[1].split(" ")))
        eqs.append((result, values))
    return eqs


def part_a(input):
    return solve(input, False)


def part_b(input):
    return solve(input, True)


def solve(input, b):
    sum = 0
    for eq in input:
        result, values = eq
        if search(result, values[0], values[1:], b):
            sum += result
    return sum


def search(result, intermediate, values, b):
    if len(values) == 0:
        return False

    value = values[0]
    l = [
        intermediate * value,
        intermediate + value,
    ]
    if b:
        l += [int(str(intermediate) + str(value))]

    for updated in l:
        if (updated == result and len(values) == 1) or search(
            result, updated, values[1:], b
        ):
            return True

    return False


from itertools import product


def search_eval_ftw(result, values, b):
    n = len(values) - 1
    elems = ["*", "+", "||"] if b else ["*", "+"]
    for c in product(elems, repeat=n):
        s = f"{values[0]}"
        for i in range(len(c)):
            if c[i] == "||":
                s = f"{s}{values[i+1]}"
            else:
                s = f"({s}) {c[i]} {values[i+1]}"
        if eval(s) == result:
            return True
    return False


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
