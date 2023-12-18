use std::fmt;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

#[derive(PartialEq)]
enum Tile {
    Empty,
    Symbol,
    Number(i32),
    Gear,
}

impl fmt::Display for Tile {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(
            f,
            "{}",
            match self {
                Tile::Empty => " ",
                Tile::Symbol => "&",
                Tile::Number(_) => "#",
                Tile::Gear => "*",
            }
        )
    }
}

fn parse_tile(c: char, n: i32) -> Tile {
    if c.is_digit(10) {
        return Tile::Number(n);
    }
    if c == '.' {
        return Tile::Empty;
    }
    if c == '*' {
        return Tile::Gear;
    }
    return Tile::Symbol;
}

fn get_number_indices(s: &str) -> usize {
    for (i, c) in s.char_indices() {
        if !c.is_digit(10) {
            return i;
        }
    }
    return s.len();
}

fn adjacent_part(map: &Vec<Vec<Tile>>, p: (i32, i32)) -> bool {
    let wee = vec![
        (p.0 + 1, p.1 + 0),
        (p.0 + 0, p.1 + 1),
        (p.0 + 1, p.1 + 1),
        (p.0 - 1, p.1 + 1),
        (p.0 + 1, p.1 - 1),
        (p.0 - 1, p.1 - 0),
        (p.0 - 0, p.1 - 1),
        (p.0 - 1, p.1 - 1),
    ];
    let height = map.len().try_into().unwrap();
    let width = map[0].len().try_into().unwrap();
    for yaoza in wee {
        if yaoza.0 < 0 || yaoza.0 >= width {
            continue;
        }
        if yaoza.1 < 0 || yaoza.1 >= height {
            continue;
        }
        let x = &map[yaoza.1 as usize];
        let y = &x[yaoza.0 as usize];
        match y {
            Tile::Symbol => return true,
            Tile::Gear => return true,
            _ => {}
        }
    }

    return false;
}

fn yeyeyeye(map: &Vec<Tile>, i: usize) -> usize {
    for (j, t) in map[i..].iter().enumerate() {
        match t {
            Tile::Number(_) => {}
            _ => return j,
        }
    }
    return map.len();
}

fn main() {
    if let Ok(lines) = read_lines("./input.txt") {
        let mut map: Vec<Vec<Tile>> = Vec::new();
        for line in lines {
            let mut wee: Vec<Tile> = Vec::new();
            if let Ok(ip) = line {
                let mut aow = -1;
                for (i, c) in ip.char_indices() {
                    if c.is_digit(10) {
                        if aow < 0 {
                            let number_len = get_number_indices(&ip[i..]);
                            aow = ip[i..i + number_len].parse::<i32>().unwrap();
                        }
                    } else if aow >= 0 {
                        aow = -1;
                    }
                    wee.push(parse_tile(c, aow));
                }
            }
            map.push(wee);
        }

        println!("");
        let mut res: Vec<(i32, (usize, usize))> = Vec::new();
        for (j, row) in map.iter().enumerate() {
            let mut calc = false;
            let mut val = false;

            for (i, col) in row.iter().enumerate() {
                match col {
                    Tile::Number(n) => {
                        if calc {
                            if val {
                                print!("g");
                            } else {
                                print!("#");
                            }
                            continue;
                        }
                        let nums = i..i + yeyeyeye(row, i);
                        for t in nums {
                            if adjacent_part(&map, (t as i32, j as i32)) {
                                val = true;
                                calc = true;
                                break;
                            }
                            calc = true
                        }
                        if val {
                            res.push((n.clone(), (i, j)));
                            print!("g");
                        } else {
                            print!("#");
                        }
                    }
                    Tile::Symbol => {
                        calc = false;
                        val = false;
                        print!("~");
                    }
                    Tile::Gear => {
                        calc = false;
                        val = false;
                        print!("*");
                    }
                    _ => {
                        calc = false;
                        val = false;
                        print!(".");
                    }
                }
            }
            println!("");
        }
        println!("");

        let mut res_sum = 0;
        for (r, gg) in res.iter().enumerate() {
            // println!("{:3} {:3}: ({:3},{:3})", r, gg.0, gg.1 .0, gg.1 .1);
            res_sum += gg.0;
        }
        println!("Total: {}", res_sum);
    }
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
