import re
from aocd import submit


def parse(lines):
    values = list()
    for line in lines:
        v = [int(d) for d in re.findall(r"-?\d+", line)]
        values.append(v)

    return values


def preprocess(seqs):
    i = 0
    while i < len(seqs):
        j = 0
        seq = seqs[i]
        next_seq = list()
        while j < len(seq) - 1:
            next_seq.append(seq[j + 1] - seq[j])
            j += 1

        seqs.append(next_seq)
        i += 1

        if all([e == 0 for e in next_seq]):
            return seqs


def interpolate_forward(seqs):
    for seq in seqs:
        seq.append(0)

    i = len(seqs) - 1
    while i > 0:
        last = seqs[i][-1]
        prev = seqs[i - 1][-2]
        prediction = prev + last
        seqs[i - 1][-1] = prediction
        i -= 1

    return seqs[0][-1]


def interpoloate_backward(seqs):
    for seq in seqs:
        seq.insert(0, 0)

    i = len(seqs) - 1
    while i > 0:
        first = seqs[i][0]
        next = seqs[i - 1][1]
        prediction = next - first
        seqs[i - 1][0] = prediction
        i -= 1

    return seqs[0][0]


def part_a(values):
    return sum([interpolate_forward(preprocess([vs])) for vs in values])


def part_b(values):
    return sum([interpoloate_backward(preprocess([vs])) for vs in values])


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
