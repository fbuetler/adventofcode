from functools import cmp_to_key, partial

from aocd import submit


def parse(lines):
    cbs = list()
    for line in lines:
        parts = line.split(" ")
        cards = list()
        for card in parts[0]:
            cards.append(card)
        bid = parts[1]
        cbs.append((cards, int(bid)))
    return cbs


CARDS = ["A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"]
CARDS_JOKER = ["A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"]


def get_card_types(cards, joker, hand):
    jokers = len(list(filter(lambda k: k == joker, hand)))

    # count cards per value
    card_counts = {}
    for card in cards:
        n = 0
        for c in hand:
            if card == c:
                n += 1
        card_counts[card] = n

    # kinds
    kinds = list()
    for i in range(2, 6):
        kind = False
        for k in card_counts:
            if k == joker:
                continue
            for j in range(jokers + 1):
                if not kind and card_counts[k] + j == i:
                    kind = True
        kinds.append(kind)

    # full house/two pairs
    full_house = False
    two_pairs = False
    for k1 in cards:
        for k2 in cards:
            if k1 == k2 or k1 == joker or k2 == joker:
                continue
            if not full_house and (
                (card_counts[k1] == 3 and card_counts[k2] == 2)
                or (card_counts[k1] == 2 and card_counts[k2] == 2 and jokers == 1)
                or (card_counts[k1] == 1 and card_counts[k2] == 2 and jokers == 2)
                or (card_counts[k1] == 0 and card_counts[k2] == 2 and jokers == 3)
            ):
                full_house = True
            if not two_pairs and (
                (card_counts[k1] == 2 and card_counts[k2] == 2)
                or (card_counts[k1] == 1 and card_counts[k2] == 2 and jokers == 1)
                or (card_counts[k1] == 0 and card_counts[k2] == 2 and jokers == 2)
            ):
                two_pairs = True

    # highest card
    highest_card = [c for c in cards if c in hand][0]

    return [kinds[3], kinds[2], full_house, kinds[1], two_pairs, kinds[0], highest_card]


def compare_cards(cards, joker, h1, h2):
    c1 = h1[0]
    c2 = h2[0]

    t1 = get_card_types(cards, joker, c1)
    t2 = get_card_types(cards, joker, c2)

    for i in range(len(t1)):
        if t1[i] and t2[i]:
            break  # tie
        if t1[i] and not t2[i]:
            return 1
        if not t1[i] and t2[i]:
            return -1

    for i in range(len(c1)):
        for c in cards:
            if c1[i] == c and c2[i] != c:
                return 1
            if c1[i] != c and c2[i] == c:
                return -1

    return 0


def part_a(cbs):
    cbs = sorted(cbs, key=cmp_to_key(partial(compare_cards, CARDS, "")))
    total = 0
    for i, (_, bids) in enumerate(cbs):
        total += (i + 1) * bids
    return total


def part_b(cbs):
    cbs = sorted(cbs, key=cmp_to_key(partial(compare_cards, CARDS_JOKER, "J")))
    total = 0
    for i, (_, bids) in enumerate(cbs):
        total += (i + 1) * bids
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
