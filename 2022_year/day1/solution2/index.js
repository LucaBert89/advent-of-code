var fs = require("fs");
var calories = fs.readFileSync("../input/input.txt", "UTF-8");

let max = 0;
const caloriesList = calories.split("\n");

caloriesList.reduce((acc, cur) => {
  if (cur === "") {
    if (acc > max) {
      max = acc;
    }
    acc = 0;
  }
  return acc + Number(cur);
}, 0);

console.log(max);
