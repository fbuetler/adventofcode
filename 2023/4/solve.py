import re

from aocd import submit


def parse(lines):
    cards = dict()
    for line in lines:
        parts = line.split(":")
        card = int(re.findall(r"\d+", parts[0])[0])
        parts = parts[1].split("|")
        winners = re.findall(r"\d+", parts[0])
        numbers = re.findall(r"\d+", parts[1])
        cards[card] = {"winners": winners, "numbers": numbers, "copies": 1}
    return cards


def part_a(cards):
    points = 0
    for card in cards.values():
        w = set(card["winners"]).intersection(set(card["numbers"]))
        if len(w) > 0:
            points += 2 ** (len(w) - 1)

    return points


def part_b(cards):
    total = len(cards)
    for card_id, card in cards.items():
        for _ in range(card["copies"]):
            w = set(card["winners"]).intersection(set(card["numbers"]))
            for j in range(len(w)):
                cards[card_id + j + 1]["copies"] += 1
                total += 1

    return total


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
