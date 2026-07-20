#[derive(Debug)]
pub struct CustomSet<T> {
    items: Vec<T>,
}

impl<T: Clone + Ord + PartialEq> PartialEq for CustomSet<T>  {
    fn eq(&self, other: &CustomSet<T>) -> bool {
        self.is_subset(other) && other.is_subset(self)
    }
}

impl<T: Clone + Ord + PartialEq> CustomSet<T> {
    pub fn new(input: &[T]) -> Self {
        CustomSet { items: input.to_vec() }
    }

    pub fn contains(&self, element: &T) -> bool {
        self.items.contains(element)
    }

    pub fn add(&mut self, element: T) {
        if !self.contains(&element) {
            self.items.push(element);
        }
    }

    pub fn is_subset(&self, other: &Self) -> bool {
        self.items.iter().all(|i| other.contains(i))
    }

    pub fn is_empty(&self) -> bool {
        self.items.is_empty()
    }

    pub fn is_disjoint(&self, other: &Self) -> bool {
        self.items.iter().all(|i| !other.contains(i))
    }

    pub fn intersection(&self, other: &Self) -> Self {
        let intersection: Vec<T> = self.items.iter().filter(|&i| other.contains(i)).cloned().collect();
        Self::new(&intersection)
    }

    pub fn difference(&self, other: &Self) -> Self {
        let difference: Vec<T> = self.items.iter().filter(|i| !other.contains(i)).cloned().collect();
        Self::new(&difference)
    }

    pub fn union(&self, other: &Self) -> Self {
        let mut union : Vec<T> = self.items.iter().cloned().collect();
        union.extend(other.items.iter().filter(|i| !self.items.contains(i)).cloned());

        Self::new(&union)
    }
}
