import { readFileSync } from "fs";

const input = readFileSync("./input.txt", "utf8").split("\n");

const earliestDepart = +input[0];
const buses = input[1].split(",");

let departed = false;
let timestamp = earliestDepart;
let busID = 0;
while (!departed) {
  for (let i = 0; i < buses.length; i++) {
    if (buses[i] !== "x" && timestamp % +buses[i] === 0) {
      departed = true;
      busID = +buses[i];
      break;
    }
  }
  if (!departed) {
    timestamp++;
  }
}

console.log(
  `Part 1 - Waited time: (${timestamp} - ${earliestDepart}) * ${busID} = ${
    (timestamp - earliestDepart) * busID
  }`
);

function mod(n: number, m: number): number {
  return ((n % m) + m) % m;
}

function egcd(a: number, b: number): [number, number, number] {
  if (b === 0) {
    return [a, 1, 0];
  }
  let [d, s, t] = egcd(b, mod(a, b));
  return [d, t, s - Math.floor(a / b) * t];
}

function crt(remainders: number[], modulos: number[]): number {
  let M = modulos.reduce((prod, val) => prod * +val, 1);
  let sum = 0;
  for (let i = 0; i < modulos.length; i++) {
    const Mi = M / modulos[i];
    const [gcd, r, s] = egcd(modulos[i], Mi);
    sum += remainders[i] * s * Mi;
    sum = mod(sum, M);
  }
  return sum;
}

let modulos = new Array<number>();
let remainders = new Array<number>();
for (let i = 0; i < buses.length; i++) {
  if (!isNaN(+buses[i])) {
    const m = +buses[i];
    modulos.push(m);
    remainders.push(mod(-i, m));
  }
}
console.log(`Part 2 - First time: ${crt(remainders, modulos)}`);
