import { readFileSync } from "fs";

class Passport {
  byr: string;
  iyr: string;
  eyr: string;
  hgt: string;
  hcl: string;
  ecl: string;
  pid: string;
  cid: string;

  hexFormat = new RegExp(`^[0-9a-f]{6}$`);

  constructor(
    byr: string,
    iyr: string,
    eyr: string,
    hgt: string,
    hcl: string,
    ecl: string,
    pid: string,
    cid: string
  ) {
    this.byr = byr;
    this.iyr = iyr;
    this.eyr = eyr;
    this.hgt = hgt;
    this.hcl = hcl;
    this.ecl = ecl;
    this.pid = pid;
    this.cid = cid;
  }

  isValidPartOne(): boolean {
    return (
      this.byr !== "" &&
      this.iyr !== "" &&
      this.eyr !== "" &&
      this.hgt !== "" &&
      this.hcl !== "" &&
      this.ecl !== "" &&
      this.pid !== ""
    );
  }

  isValidPartTwo(): boolean {
    if (!this.isValidPartOne()) {
      return false;
    }
    let valid = true;

    valid &&=
      this.byr.length === 4 &&
      !isNaN(+this.byr) &&
      1920 <= +this.byr &&
      +this.byr <= 2002;

    valid &&=
      this.byr.length === 4 &&
      !isNaN(+this.iyr) &&
      2010 <= +this.iyr &&
      +this.iyr <= 2020;

    valid &&=
      this.byr.length === 4 &&
      !isNaN(+this.eyr) &&
      2020 <= +this.eyr &&
      +this.eyr <= 2030;

    let hgtValue = this.hgt.substring(0, this.hgt.length - 2);
    let hgtUnit = this.hgt.substring(this.hgt.length - 2);
    valid &&=
      (hgtUnit === "cm" &&
        !isNaN(+hgtValue) &&
        150 <= +hgtValue &&
        +hgtValue <= 193) ||
      (hgtUnit === "in" &&
        !isNaN(+hgtValue) &&
        59 <= +hgtValue &&
        +hgtValue <= 76);

    valid &&=
      this.hcl.substring(0, 1) === "#" &&
      this.hcl.substring(1).match(this.hexFormat) !== null;

    valid &&=
      this.ecl === "amb" ||
      this.ecl === "blu" ||
      this.ecl === "brn" ||
      this.ecl === "gry" ||
      this.ecl === "grn" ||
      this.ecl === "hzl" ||
      this.ecl === "oth";

    valid &&= this.pid.length === 9 && !isNaN(+this.pid);

    return valid;
  }
}

// input is required to have a trailing new line
const input = readFileSync("./input.txt", "utf8").split("\n");

let i = 0;
let passports = new Array<Passport>();
while (i < input.length) {
  let props = new Map<string, string>();
  while (input[i] !== "") {
    const lineParts = input[i].split(" ");
    for (let j = 0; j < lineParts.length; j++) {
      const pair = lineParts[j].split(":");
      if (pair.length !== 2) {
        console.log(`invalid pair: ${pair}`);
        continue;
      }
      props.set(pair[0], pair[1]);
    }
    i++;
  }
  passports.push(
    new Passport(
      props.get("byr") || "",
      props.get("iyr") || "",
      props.get("eyr") || "",
      props.get("hgt") || "",
      props.get("hcl") || "",
      props.get("ecl") || "",
      props.get("pid") || "",
      props.get("cid") || ""
    )
  );
  i++;
}

let validPassportsPartOne = 0;
let validPassportsPartTwo = 0;
for (let i = 0; i < passports.length; i++) {
  if (passports[i].isValidPartOne()) {
    validPassportsPartOne++;
    if (passports[i].isValidPartTwo()) {
      validPassportsPartTwo++;
    }
  }
}

console.log(`part 1: ${validPassportsPartOne}`);
console.log(`part 2: ${validPassportsPartTwo}`);
