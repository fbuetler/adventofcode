from collections import Counter, deque

from aocd import submit


def parse(lines):
    return [int(l) for l in lines]


def part_a(input):
    sum = 0
    for secret in input:
        for _ in range(2000):
            secret = evolve(secret)
        sum += secret
    return sum


def part_b(input):
    all_seqs = set()
    sellers = list()
    for secret in input:
        bananas = dict()
        q = deque(maxlen=4)
        prev = banana(secret)

        for _ in range(2000):
            secret = evolve(secret)
            b = banana(secret)
            q.appendleft(b - prev)
            prev = b

            if len(q) == 4:
                seq = to_seq(q)
                all_seqs.add(seq)
                if not seq in bananas:
                    bananas[seq] = b

        sellers.append(bananas)

    most_bananas = 0
    for seq in all_seqs:
        b = 0
        for seller in sellers:
            if seq in seller:
                b += seller[seq]
        most_bananas = max(most_bananas, b)

    return most_bananas


def evolve(secret):
    mix = secret * 64
    secret ^= mix
    secret %= 16777216

    mix = secret // 32
    secret ^= mix
    secret %= 16777216

    mix = secret * 2048
    secret ^= mix
    secret %= 16777216

    return secret


def banana(n):
    return int(str(n)[-1])


def to_seq(q):
    return (q[3], q[2], q[1], q[0])


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
