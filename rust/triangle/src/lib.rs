pub struct Triangle<T> {
    a: T,
    b: T,
    c: T,
}

impl<T> Triangle<T>
where T : Copy + ::std::cmp::PartialOrd + ::std::ops::Add<Output = T> + num_traits::Zero {
    pub fn build(sides: [T; 3]) -> Option<Self> {
        let [a, b, c] = sides;

        let is_not_inequal = a + b < c || a + c < b || b + c < a;
        if a.is_zero() || b.is_zero() || c.is_zero() || is_not_inequal {
            None
        }
        else {
            Some(Triangle{
                a,
                b,
                c,
            })
        }
    }

    pub fn is_equilateral(&self) -> bool {
        self.a == self.b && self.b == self.c
    }

    pub fn is_scalene(&self) -> bool {
        self.a != self.b && self.a != self.c && self.b != self.c
    }

    pub fn is_isosceles(&self) -> bool {
        self.a == self.b || self.b == self.c || self.a == self.c
    }
}
