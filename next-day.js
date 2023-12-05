const fs = require("fs");

let highestDay = 0;
fs.readdir(__dirname, { withFileTypes: true }, (err, files) => {
  if (err) {
    console.log(err);
  } else {
    let latestDir = undefined;
    files
      .filter((file) => file.isDirectory())
      .forEach((file) => {
        const parts = file.name.split("-");
        if (parts.length === 2) {
          const n = parseInt(parts[1]);
          if (n > highestDay) {
            latestDir = file;
            highestDay = n;
          }
        }
      });
    if (latestDir !== undefined) {
      fs.cpSync(
        `${latestDir.path}\\${latestDir.name}`,
        `${latestDir.path}\\${latestDir.name.replace(
          highestDay.toString(),
          (highestDay + 1).toString()
        )}`,
        {
          recursive: true,
        }
      );
    }
  }
});
