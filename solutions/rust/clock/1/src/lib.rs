use std::fmt;

#[derive(PartialEq, Debug)]
pub struct Clock {
    time_in_minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let max_minutes = 24 * 60;
        let total_minutes = hours * 60 + minutes;
        Clock {
            time_in_minutes: (max_minutes + total_minutes % max_minutes) % max_minutes,
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Clock::new(self.hours(), self.minutes() + minutes)
    }

    fn hours(&self) -> i32 {
        self.time_in_minutes / 60
    }

    fn minutes(&self) -> i32 {
        self.time_in_minutes % 60
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{:02}:{:02}", self.hours(), self.minutes())
    }
}
