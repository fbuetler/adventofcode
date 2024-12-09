from aocd import submit


def parse(lines):
    fs = list()
    id = 0
    free = 0
    for i, l in enumerate(list(lines[0])):
        if i % 2 == 0:
            # file
            v = id
            id += 1
        else:
            # free
            v = None
            free += int(l)
        for _ in range(int(l)):
            fs.append(v)
    return (fs, id-1)


def part_a(input):
    fs, _ = input
    fs = list(fs)

    first = next_free(fs, 0)
    last = last_occupied(fs, len(fs) - 1)
    while first < last:
        fs[first] = fs[last]
        fs[last] = None
        first = next_free(fs, first)
        last = last_occupied(fs, last)

    return checksum(fs)


def next_free(fs, pos):
    while fs[pos] != None:
        pos += 1
    return pos


def last_occupied(fs, pos):
    while fs[pos] == None:
        pos -= 1
    return pos


def checksum(fs):
    sum = 0
    i = 0
    while i < len(fs):
        if fs[i] != None:
            sum += i * fs[i]
        i += 1
    return sum


def part_b(input):
    fs, max_id = input
    fs = list(fs)

    for id in range(max_id, -1, -1):
        block_start, block_size = block_with_id(fs, id)
        free_start, free_size = free_block_size(fs, 0)
        while free_start < block_start:
            if block_size <= free_size:
                for j in range(block_size):
                    fs[free_start+j] = id
                    fs[block_start+j] = None
                break
            free_start, free_size = free_block_size(fs, free_start+free_size+1)

    return checksum(fs)


def block_with_id(fs, id):
    i = len(fs) - 1
    while i >= 0 and fs[i] != id:
        i -= 1

    j = i
    while j >= 0 and fs[j] == id:
        j -= 1
    return (j+1, i-j)


def free_block_size(fs, pos):
    pos
    while pos < len(fs) and fs[pos] != None:
        pos += 1

    i = pos
    size = 0
    while i < len(fs) and fs[i] == None:
        size += 1
        i += 1
    return (pos, size)


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
