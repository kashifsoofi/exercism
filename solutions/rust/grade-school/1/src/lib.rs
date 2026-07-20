use std::collections::HashMap;

struct Grade {
    grade: u32,
    students: Vec<String>,
}

pub struct School {
    grade_students: HashMap<u32, Vec<String>>,
}

impl School {
    pub fn new() -> School {
        School {
            grade_students: HashMap::new(),
        }
    }

    pub fn add(&mut self, grade: u32, student: &str) {
        let mut students = self.grade_students.entry(grade).or_insert(vec![]);
        students.push(student.to_string());
    }

    pub fn grades(&self) -> Vec<u32> {
        let mut grades = self.grade_students.iter().map(|(grade, _)| *grade).collect::<Vec<u32>>();
        grades.sort();
        grades
    }

    // If grade returned an `Option<&Vec<String>>`,
    // the internal implementation would be forced to keep a `Vec<String>` to lend out.
    // By returning an owned vector instead,
    // the internal implementation is free to use whatever it chooses.
    pub fn grade(&self, grade: u32) -> Option<Vec<String>> {
        let mut students = self.grade_students.get(&grade)?.to_vec();
        students.sort();
        Some(students)
    }
}
