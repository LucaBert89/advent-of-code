var fs = require("fs");
var calories = fs.readFileSync("../input/input.txt", "UTF-8");

const caloriesList = calories.split("\n");
let array = [];
caloriesList.reduce((acc, cur) => {
  if (cur === "") {
    array.push(acc);
    acc = 0;
  }
  return acc + Number(cur);
}, 0);

const sorting = array.sort((a, b) => (a < b ? 1 : a > b ? -1 : 0));
const result2 = sorting.slice(0, 3).reduce((acc, cur) => {
  return acc + cur;
}, 0);

console.log(result2);
