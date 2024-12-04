from aocd import submit


def parse(lines):
    return [list(line) for line in lines]

def part_a(input):
    sum = 0
    n = len(input)
    m = len(input[0])
    for i in range(n):
        for j in range(m):
            for x, y in [(a, b) for a in range(-1,2) for b in range(-1,2)]:
                if x < 0 and i - 3 < 0:
                    continue
                if x > 0 and i + 3 >= n:
                    continue
                if y < 0 and j - 3 < 0:
                    continue
                if y > 0 and j + 3 >= m:
                    continue

                if input[i][j] + input[i+x][j+y] + input[i+2*x][j+2*y] + input[i+3*x][j+3*y] == 'XMAS':
                    sum += 1
    return sum

def part_b(input):
    sum = 0
    n = len(input)
    m = len(input[0])
    for i in range(n):
        for j in range(m):
            if i < 1 or j < 0 or i >= n - 1 or j >= m - 1:
                continue

            diaA = input[i-1][j-1] + input[i][j] + input[i+1][j+1]
            diaB = input[i+1][j-1] + input[i][j] + input[i-1][j+1]

            if (diaA == 'MAS' or diaA == 'SAM') and (diaB == 'MAS' or diaB == 'SAM'):
                sum += 1
    return sum

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