use self::Allergen::*;
use std::slice::Iter;

pub struct Allergies {
    score: u32,
}

#[derive(Debug, PartialEq, Copy, Clone)]
pub enum Allergen {
    Eggs = 1 << 0,
    Peanuts = 1 << 1,
    Shellfish = 1 << 2,
    Strawberries = 1 << 3,
    Tomatoes = 1 << 4,
    Chocolate = 1 << 5,
    Pollen = 1 << 6,
    Cats = 1 << 7,
}

impl Allergen {
    pub fn iterator() -> Iter<'static, Allergen> {
        static ALLERGENS: [Allergen; 8] = [Eggs, Peanuts, Shellfish, Strawberries, Tomatoes, Chocolate, Pollen, Cats];
        ALLERGENS.iter()
    }
}

impl Allergies {
    pub fn new(score: u32) -> Self {
        Allergies {
            score,
        }
    }

    pub fn is_allergic_to(&self, allergen: &Allergen) -> bool {
        let value = *allergen as u32;
        self.score & value == value
    }

    pub fn allergies(&self) -> Vec<Allergen> {
        Allergen::iterator().filter(|a| self.is_allergic_to(a)).cloned().collect()
    }
}
