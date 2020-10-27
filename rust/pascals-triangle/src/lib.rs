#[derive(Clone)]
pub struct PascalsTriangle {
    rows: Vec<Vec<u32>>,
}

impl PascalsTriangle {
    pub fn new(row_count: u32) -> Self {
        let mut rows: Vec<Vec<u32>> = Vec::with_capacity(row_count as usize);
        if row_count > 0 {
            rows.push(vec![1u32]);

            for _ in 1..row_count {
                rows.push(Self::generate_row(rows.get(rows.len() - 1).unwrap()));
            }
        }

        PascalsTriangle { rows }
    }

    pub fn rows(&self) -> Vec<Vec<u32>> {
        self.rows.clone()
    }

    fn generate_row(previous: &Vec<u32>) -> Vec<u32> {
        let mut row: Vec<u32> = Vec::with_capacity(previous.len() + 1);
        row.push(1);
        for i in 1..previous.len() {
            let left = previous.get(i - 1).unwrap_or(&0_u32);
            let right = previous.get(i).unwrap_or(&0_u32);
            row.push(left + right);
        }
        row.push(1);
        row
    }
}
