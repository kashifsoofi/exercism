// The code below is a stub. Just enough to satisfy the compiler.
// In order to pass the tests you can add-to or change any of this code.

#[derive(PartialEq, Clone, Copy, Debug)]
pub enum Direction {
    North,
    East,
    South,
    West,
}

#[derive(Clone, Copy, Debug)]
pub struct Robot {
    x: i32,
    y: i32,
    direction: Direction,
}

impl Robot {
    pub fn new(x: i32, y: i32, d: Direction) -> Self {
        Robot { x: x, y: y, direction: d }
    }

    pub fn turn_right(self) -> Self {
        let direction = match self.direction {
            Direction::North => Direction::East,
            Direction::East => Direction::South,
            Direction::South => Direction::West,
            Direction::West => Direction::North,
        };
        Self { x: self.x, y: self.y, direction: direction }
    }

    pub fn turn_left(self) -> Self {
        let direction = match self.direction {
            Direction::North => Direction::West,
            Direction::West => Direction::South,
            Direction::South => Direction::East,
            Direction::East => Direction::North,
        };
        Self { x: self.x, y: self.y, direction: direction }
    }

    pub fn advance(self) -> Self {
        let (x, y) = match self.direction {
            Direction::North => (self.x, self.y + 1),
            Direction::East => (self.x + 1 , self.y),
            Direction::South => (self.x, self.y - 1),
            Direction::West => (self.x - 1, self.y),
        };
        Self { x: x, y: y, direction: self.direction }
    }

    pub fn instructions(self, instructions: &str) -> Self {
        let mut robot = self;
        instructions.chars().for_each(|ch|
            robot = match ch {
                'A' => robot.advance(),
                'L' => robot.turn_left(),
                'R' => robot.turn_right(),
                _ => panic!("Invalid instruction"),
            }
        );

        robot
    }

    pub fn position(&self) -> (i32, i32) {
        (self.x, self.y)
    }

    pub fn direction(&self) -> &Direction {
        &self.direction
    }
}
