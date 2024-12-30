import re
from typing import Dict, List

from aocd import submit


def parse(lines):
    input_pattern = re.compile(r"(\w+): (\d)")
    eq_pattern = re.compile(r"(\w+) (AND|OR|XOR) (\w+) -> (\w+)")
    state = dict()
    eqs = dict()
    for line in lines:
        m = re.match(input_pattern, line)
        if m:
            state[m.group(1)] = True if m.group(2) == "1" else False
        m = re.match(eq_pattern, line)
        if m:
            eqs[m.group(4)] = (m.group(1), m.group(2), m.group(3))
    return state, eqs


def part_a(input):
    state, eqs = input
    adj = to_adj(eqs)
    order = topo_sort(adj)
    n = calc(eqs, state, order)
    return n


def to_adj(eqs):
    adj = {k: [] for k in eqs.keys()}
    for c, (a, op, b) in eqs.items():
        if a in eqs.keys():
            adj[a] += [c]

        if b in eqs.keys():
            adj[b] += [c]

    return adj


def topo_sort(adj: Dict):
    stack = []
    visited = dict()

    for v in adj.keys():
        if not v in visited:
            topo_sort_rec(adj, v, visited, stack)

    return list(reversed(stack))


def topo_sort_rec(adj: Dict, v, visited: Dict, stack: List):
    visited[v] = True

    for u in adj[v]:
        if not u in visited:
            topo_sort_rec(adj, u, visited, stack)

    stack.append(v)


def calc(eqs, state, order):
    for c in order:
        a, op, b = eqs[c]
        if op == "AND":
            state[c] = state[a] and state[b]
        elif op == "OR":
            state[c] = state[a] or state[b]
        elif op == "XOR":
            state[c] = state[a] != state[b]
        else:
            raise ValueError("invalid op")

    n = 0
    i = 0
    while f"z{i:02d}" in state:
        v = state[f"z{i:02d}"]
        if v:
            n += 2**i
        i += 1
    return n


def part_b(input):
    state, eqs = input

    def is_input(v):
        return v.startswith("x") or v.startswith("y")

    def is_output(v):
        return v.startswith("z")

    def is_initial(v):
        return v == "x00" or v == "y00"

    def is_final(v):
        return v == "z45"

    swaps = set()
    for c, (a, op, b) in eqs.items():
        if op != "XOR" and is_output(c) and not is_final(c):
            swaps.add(c)

        if op == "AND" and not (is_initial(a) or is_initial(b)):
            for c1, (a1, op1, b1) in eqs.items():
                if (a1 == c or b1 == c) and op1 != "OR":
                    swaps.add(c)

        if (
            op == "XOR"
            and not (is_input(a) or is_output(a))
            and not (is_input(b) or is_output(b))
            and not (is_input(c) or is_output(c))
        ):
            swaps.add(c)

        if op == "XOR":
            for c1, (a1, op1, b1) in eqs.items():
                if (a1 == c or b1 == c) and op1 == "OR":
                    swaps.add(c)

    s = ",".join(sorted(list(swaps)))
    return s


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
