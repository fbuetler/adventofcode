from copy import deepcopy
from itertools import combinations
from typing import Dict, List, Set

from aocd import submit


def parse(lines):
    adj = dict()
    for l in lines:
        s = l.split("-")
        c1, c2 = s[0], s[1]
        if c1 in adj:
            adj[c1].append(c2)
        else:
            adj[c1] = [c2]

        if c2 in adj:
            adj[c2].append(c1)
        else:
            adj[c2] = [c1]
    return adj


def part_a(input):
    adj = input
    groups = set()
    find_cliques(adj, adj.keys(), set(), 3, groups)

    sum = 0
    for group in groups:
        for pc in group:
            if pc.startswith("t"):
                sum += 1
                break

    return sum


def part_b(input):
    adj = input
    max_k = max(map(lambda e: len(e[1]), adj.items()))
    clique = biggest_clique(adj, max_k)
    return ",".join(sorted(list(clique)))


def find_cliques(
    adj: Dict[str, List[str]], nodes, clique: Set[str], size: int, groups: frozenset
):
    if len(clique) > size:
        return

    for u in nodes:
        if u in clique:
            continue

        if len(adj[u]) < size - 1:
            continue

        clique.add(u)
        if not is_clique(adj, clique):
            clique.remove(u)
            continue

        if len(clique) == size:
            if not clique in groups:
                groups.add(frozenset(clique))

        find_cliques(adj, nodes, deepcopy(clique), size, groups)
        clique.remove(u)

    return


def biggest_clique(adj, n):
    if n == 0:
        return None

    for u, neighbours in adj.items():
        for subset in combinations(neighbours, n):
            nodes = set(list(subset) + [u])
            if is_clique(adj, nodes):
                return nodes

    return biggest_clique(adj, n - 1)


def is_clique(adj: Dict[str, List[str]], nodes: Set[str]):
    for u in nodes:
        for v in nodes:
            if u == v:
                continue
            if not v in adj[u]:
                return False
    return True


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
