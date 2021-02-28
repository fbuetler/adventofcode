from sympy.ntheory.modular import crt

with open("input.txt") as f:
    input = f.read().split("\n")

modulos = []
remainders = []
buses = input[1].split(",")
for i in range(len(buses)):
    if buses[i].isnumeric():
        b = int(buses[i])
        modulos.append(b)
        remainders.append(-i % b)

print(crt(modulos, remainders, check=False)[0])