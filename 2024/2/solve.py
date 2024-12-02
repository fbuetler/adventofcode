from aocd import submit


def parse(lines):
    reports = list()
    for line in lines:
        levels = [int(l) for l in line.split()]
        reports.append(levels)
    return reports

def part_a(reports):
    save = 0    
    for r in reports:
        if is_safe(r):
            save += 1
    return save

def part_b(reports):
    save = 0    
    for r in reports:
        for ls in sublists(r):
            if is_safe(ls):
                save += 1
                break
    return save

def is_safe(report):
    return is_in_range(report, -3, -1) or is_in_range(report, 1, 3)

def is_in_range(report, left, right):
    for i in range(0, len(report)-1):
        diff = report[i] - report[i+1]
        if diff < left or right < diff:
            return False
    return True

def sublists(ls):
    return [ls[:i] + ls[i+1:] for i in range(len(ls))]
        
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