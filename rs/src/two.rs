use std::usize;



static ROCK: i32 = 1;
static PAPER: i32 = 2;
static SCISSORS: i32 = 3;
static LOST: i32 = 0;
static DRAW: i32 = 3;
static WON: i32 = 6;

fn score(opponent: i32, player: i32) -> i32 {
    let moves: Vec<i32> = vec![
        ROCK,
        PAPER,
        SCISSORS,
    ];

    let outcomes: Vec<i32> = vec![
        DRAW,
        WON,
        LOST,
    ];

    let a: usize = (opponent - 1) as usize;
    let b: usize = (player - 1) as usize;
    let play = moves[a] + moves[b];
    let index = ((play + player) % 3) as usize;
    return outcomes[index];
}

fn letter_to_shit_one(letter: char) -> i32 {
    match letter {
        'A' | 'X' => ROCK,
        'B' | 'Y' => PAPER,
        'C' | 'Z' => SCISSORS,
        _ => panic!("fuck yo shit"),
    }
}

pub fn day_two() {
    let input = include_str!("../two.txt");
    let result = input.lines().map(|play| {
        let opp_letter = play.chars().next().unwrap();
        let you_letter = play.chars().rev().next().unwrap();
        let opp = letter_to_shit_one(opp_letter);
        let you = letter_to_shit_one(you_letter);
        let bitch = score(opp, you);
        let total = you + bitch;
        println!("Play {opp_letter} ({opp}) {you_letter} ({you}) {bitch} total {total}");
        return total;
    });
    println!("Day Two Part One: {}", result.sum::<i32>());
}
