const fs = require('node:fs');
const mappings = {
  one: '1',
  two: '2',
  three: '3',
  four: '4',
  five: '5',
  six: '6',
  seven: '7',
  eight: '8',
  nine: '9',
  0: '0',
  1: '1',
  2: '2',
  3: '3',
  4: '4',
  5: '5',
  6: '6',
  7: '7',
  8: '8',
  9: '9',
};
console.log(JSON.stringify(mappings));
try {
  const data = fs.readFileSync('input.txt', 'utf8');
  process(data);
} catch (err) {
  console.error(err);
}

function process(data) {
  var res = data.split('\n').reduce((acc, curr) => {
    const first = findFirst(curr);
    const second = findLast(curr);
    return acc + parseInt(`${first}${second}`);
  }, 0);
  console.log(res);
}

function findFirst(row) {
  while (row.length > 0) {
    for (var key in mappings) {
      if (row.startsWith(key)) {
        return mappings[key];
      }
    }
    row = row.slice(1);
  }
  console.log(row);
  throw new Error(row);
}

function findLast(row) {
  const rowCopy = row;
  while (row.length > 0) {
    for (var key in mappings) {
      if (row.endsWith(key)) {
        return mappings[key];
      }
    }
    row = row.slice(0, row.length - 1);
  }
  console.log(row);
  throw new Error(rowCopy);
}
