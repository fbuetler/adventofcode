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
    print(order)

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
        # print(a, op, b, c)
        # print(state[a], state[b])
        if op == "AND":
            state[c] = state[a] and state[b]
        elif op == "OR":
            state[c] = state[a] or state[b]
        elif op == "XOR":
            state[c] = state[a] != state[b]
        else:
            raise ValueError("invalid op")
    #     print(state[c])
    # print(state)

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
    adj = to_adj(eqs)
    order = topo_sort(adj)


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
