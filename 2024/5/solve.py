from aocd import submit


def parse(lines):
    orders = list()
    updates = list()
    for line in lines:
        if "|" in line:
            os = line.split("|")
            orders.append((int(os[0]), int(os[1])))
        elif "," in line:
            us = line.split(",")
            updates.append([int(u) for u in us])
    return (orders, updates)


def part_a(input):
    orders, updates = input
    oks, noks = classify(orders, updates)
    return middle_sum(oks)


def part_b(input):
    orders, updates = input
    oks, noks = classify(orders, updates)
    for nok in noks:
        changed = True
        while changed:
            changed = False
            for a, b in orders:
                if is_malformed(nok, a, b):
                    swap(nok, a, b)
                    changed = True

    return middle_sum(noks)


def classify(orders, updates):
    oks = list()
    noks = list()
    for update in updates:
        ok = True
        for a, b in orders:
            if is_malformed(update, a, b):
                ok = False
                break
        if ok:
            oks.append(update)
        else:
            noks.append(update)
    return (oks, noks)


def is_malformed(l, a, b):
    return a in l and b in l and l.index(a) > l.index(b)


def swap(l, a, b):
    tmp = l[l.index(a)]
    l[l.index(a)] = l[l.index(b)]
    l[l.index(b)] = tmp


def middle_sum(ls):
    return sum([l[len(l)//2] for l in ls])


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
