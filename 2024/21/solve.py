from collections import defaultdict

from aocd import submit


def parse(lines):
    codes = list()
    for l in lines:
        codes.append(l)
    return codes


def part_a(input):
    sum = 0
    for code in input:
        sum += press(code, 2)
    return sum


def part_b(input):
    sum = 0
    for code in input:
        sum += press(code, 25)
    return sum


def press(code, robots):
    seqs = press_numeric(code)

    for _ in range(robots):
        seqs = press_directional(seqs)

    l = 0
    for seq, n in seqs.items():
        l += len(seq) * n
    c = int(code[:3])
    return l * c


"""
+---+---+---+
| 7 | 8 | 9 |
+---+---+---+
| 4 | 5 | 6 |
+---+---+---+
| 1 | 2 | 3 |
+---+---+---+
    | 0 | A |
    +---+---+
"""

NUMS = {
    "7": (0, 0),
    "8": (0, 1),
    "9": (0, 2),
    "4": (1, 0),
    "5": (1, 1),
    "6": (1, 2),
    "1": (2, 0),
    "2": (2, 1),
    "3": (2, 2),
    "0": (3, 1),
    "A": (3, 2),
}

NUMS_GAP = (3, 0)


def press_numeric(code):
    s = (3, 2)  # A
    seqs = defaultdict(int)
    for digit in code:
        t = NUMS[digit]
        seqs[path(s, t, NUMS_GAP)] += 1
        s = t
    return seqs


"""
    +---+---+
    | ^ | A |
+---+---+---+
| < | v | > |
+---+---+---+
"""

DIRS = {
    "^": (0, 1),
    "A": (0, 2),
    "<": (1, 0),
    "v": (1, 1),
    ">": (1, 2),
}

DIRS_GAP = (0, 0)


def press_directional(seqs):
    s = "A"
    wrapped_seqs = defaultdict(int)
    for seq, n in seqs.items():
        for t in seq:
            wrapped_seqs[path(DIRS[s], DIRS[t], DIRS_GAP)] += n
            s = t
    return wrapped_seqs


def path(s, t, gap):
    sx, sy = s
    tx, ty = t

    dx = sx - tx
    dy = sy - ty
    seq_up = go_vertical(dx)
    seq_lr = go_horizontal(dy)

    if ty > sy and (tx, sy) != gap:
        return seq_up + seq_lr + "A"
    elif (sx, ty) != gap:
        return seq_lr + seq_up + "A"
    else:
        return seq_up + seq_lr + "A"


def go_horizontal(dy):
    if dy == 0:  # hit
        return ""
    elif dy > 0:  # go left
        return dy * "<"
    else:  # go right
        return abs(dy) * ">"


def go_vertical(dx):
    if dx == 0:  # hit
        return ""
    elif dx > 0:  # go up
        return dx * "^"
    else:  # go down
        return abs(dx) * "v"


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
