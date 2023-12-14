const fs = require('fs');

const langs = [
  'js',
  'ts',
  'hs',
  'zig',
  'c',
  'cpp',
  'cs',
  'java',
  'sc',
  'rs',
  'rb',
  'ml',
  'fs',
  'asm',
  'sql',
  'erl',
  'ex',
  'go',
  'py',
  'f90',
  'lisp',
  'kt',
  'SM',
  'jl',
  'bf',
];

function getRandomInt(max) {
  return Math.floor(Math.random() * max);
}

let highestDay = 0;
fs.readdir(__dirname, { withFileTypes: true }, (err, files) => {
  let remainingLangs = [...langs];
  if (err) {
    console.log(err);
  } else {
    let latestDir = undefined;
    files
      .filter((file) => file.isDirectory())
      .forEach((file) => {
        const parts = file.name.split('-');
        if (parts.length === 3) {
          const n = parseInt(parts[1]);
          if (n > highestDay) {
            latestDir = file;
            highestDay = n;
          }
          const l = parts[2];
          remainingLangs = remainingLangs.filter((v) => v !== l);
        }
      });
    if (latestDir !== undefined) {
      const randomLang = remainingLangs[getRandomInt(remainingLangs.length)];
      const prevLang = latestDir.name.split('-')[2];
      const newDirName = [
        latestDir.name.split('-')[0],
        `${highestDay + 1}`,
        randomLang,
      ].join('-');
      fs.cpSync(
        `${latestDir.path}\\${latestDir.name}`,
        `${latestDir.path}\\${newDirName}`,
        {
          recursive: true,
        }
      );
      fs.rename(
        `${latestDir.path}\\${newDirName}\\main.${prevLang}`,
        `${latestDir.path}\\${newDirName}\\main.${randomLang}`,
        () => {}
      );
    }
  }
});
