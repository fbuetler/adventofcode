import re

from aocd import submit

def parse(lines):
    games = {}
    for line in lines:
        m = re.findall(r"Game (\d+): (.*)", line)
        game = int(m[0][0])
        draws = m[0][1]

        css = list()
        draw = draws.split(";")
        for d in draw:
            cs = {"red": 0, "green": 0, "blue": 0}
            for c in cs.keys():
                m = re.findall(f"(\d+) {c}", d)
                if m:
                    cs[c] = int(m[0])
            css.append(cs)
        games[game] = css
    return games

def part_a(games):
    sum = 0
    for game, draws in games.items():
        works = True
        for draw in draws:
            if draw["red"] > 12:
                works = False
            if draw["green"] > 13:
                works = False
            if draw["blue"] > 14:
                works = False
        if works:
            sum += game

    return sum
            

def part_b(games):
    sum = 0
    for game, draws in games.items():
        red = 0
        green = 0
        blue = 0
        for draw in draws:
            red = max(red, draw["red"])
            green = max(green, draw["green"])
            blue = max(blue, draw["blue"])
        sum += red * green * blue
    return sum

if __name__ == "__main__":
    with open("input.txt", "r") as f:
        lines = [line.rstrip() for line in f]

    a = part_a(parse(lines))
    print(a)
    submit(a, part="a")

    b = part_b(parse(lines))
    print(b)
    submit(b,  part="b")