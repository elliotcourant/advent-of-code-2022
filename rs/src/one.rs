use std::fs;

pub fn read_data() {
    let input = match fs::read_to_string("src/one.txt") {
        Ok(data) => data,
        Err(e) => panic!("{}", e),
    };

    let elves = input.split("\n\n");
    let elves_and_their_fuckin_calories = elves.map(|elf| elf.lines()
            .map(|shit| shit.parse::<i32>().unwrap())
            .sum::<i32>());
    println!("Part One: {}", elves_and_their_fuckin_calories.clone().max().unwrap());

    let mut shit: Vec<i32> = elves_and_their_fuckin_calories.clone().collect::<Vec<i32>>();
    shit.sort();
    shit.reverse();

    println!("Part Two: {:?}", (&shit[0..3]).iter().sum::<i32>());
}

