import { readFileSync } from "fs";

class Food {
  ingredients = new Array<string>();
  allergens = new Array<string>();

  constructor(rawFood: string) {
    const parts = rawFood.split("(");
    parts[0]
      .trim()
      .split(" ")
      .forEach((el) => this.ingredients.push(el));
    parts[1]
      .replace(")", "")
      .replace("contains", "")
      .trim()
      .split(",")
      .forEach((el) => this.allergens.push(el.trim()));
  }
}

function countAppearances(map: Map<string, string[]>): Map<string, number> {
  let appearances = new Map<string, number>();
  for (let key of Array.from(map.keys())) {
    const values = map.get(key);
    values.forEach((value) =>
      appearances.set(value, (appearances.get(value) || 0) + 1)
    );
  }
  return appearances;
}

const input = readFileSync("./input.txt", "utf8").split("\n");
let allergeneMayBeIn = new Map<string, string[]>();
let foods = new Array<Food>();
input.forEach((rawFood) => {
  const food = new Food(rawFood);
  foods.push(food);
  food.allergens.forEach((allergen) => {
    if (allergeneMayBeIn.has(allergen)) {
      const ingrds = allergeneMayBeIn.get(allergen);
      allergeneMayBeIn.set(
        allergen,
        ingrds.filter((ingrd) => food.ingredients.includes(ingrd))
      );
    } else {
      allergeneMayBeIn.set(allergen, food.ingredients);
    }
  });
});

let ingredientsWithAllergene = new Set<string>();
for (let allergene of Array.from(allergeneMayBeIn.keys())) {
  allergeneMayBeIn
    .get(allergene)
    .forEach((ingrd) => ingredientsWithAllergene.add(ingrd));
}
let countApps = 0;
foods.forEach((food) =>
  food.ingredients.forEach(
    (ingrd) => (countApps += !ingredientsWithAllergene.has(ingrd) ? 1 : 0)
  )
);
console.log(`Part 1: ${countApps}`);

let ingredientContainsAllergene = new Array<[string, string]>();
while (allergeneMayBeIn.size !== 0) {
  const appearances = countAppearances(allergeneMayBeIn);
  let ingredient: string;
  for (ingredient of Array.from(appearances.keys())) {
    if (appearances.get(ingredient) === 1) {
      break;
    }
  }
  for (let allergene of Array.from(allergeneMayBeIn.keys())) {
    let ingrds = allergeneMayBeIn.get(allergene);
    if (!ingrds.includes(ingredient)) {
      continue;
    }
    ingredientContainsAllergene.push([ingredient, allergene]);
    allergeneMayBeIn.delete(allergene);
  }
}

console.log(
  `Part 2: ${ingredientContainsAllergene
    .sort(([a, b], [c, d]) => (b < d ? -1 : 1))
    .map(([a, b]) => a)
    .join(",")}`
);
