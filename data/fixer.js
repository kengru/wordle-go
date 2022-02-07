"use strict";

const fs = require("fs");
const letters = [
  "a",
  "b",
  "c",
  "d",
  "e",
  "f",
  "g",
  "h",
  "j",
  "k",
  "l",
  "m",
  "n",
  "o",
  "p",
  "q",
  "r",
  "s",
  "t",
  "u",
  "v",
  "w",
  "x",
  "y",
  "z",
];
const total = [];

letters.forEach((l) => {
  let rawdata = fs.readFileSync(l + ".json");
  let student = JSON.parse(rawdata);
  let words = Object.keys(student);
  let filtWords = words
    .filter((w) => w.length === 5 && !w.includes(" "))
    .map((w) => w.toLowerCase());
  total.push(...filtWords);
});

fs.writeFileSync("dict.txt", total.join(","));
