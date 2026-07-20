#[derive(Debug, PartialEq)]
pub struct Dna {
    dna: String
}

#[derive(Debug, PartialEq)]
pub struct Rna {
    rna: String
}

impl Dna {
    pub fn new(dna: &str) -> Result<Dna, usize> {
        for (i, c) in dna.chars().enumerate() {
            match c {
                'A' | 'C' | 'G' | 'T' => (),
                _ => return Err(i),
            }
        }
        Ok(Dna { dna: dna.to_string()})
    }

    pub fn into_rna(self) -> Rna {
        let mut rna = String::new();
        for c in self.dna.chars() {
            match c {
                'A' => rna.push('U'),
                'C' => rna.push('G'),
                'G' => rna.push('C'),
                'T' => rna.push('A'),
                _ => (),
            }
        }
        Rna { rna: rna }
    }
}

impl Rna {
    pub fn new(rna: &str) -> Result<Rna, usize> {
        for (i, c) in rna.chars().enumerate() {
            match c {
                'A' | 'C' | 'G' | 'U' => (),
                _ => return Err(i),
            }
        }
        Ok(Rna { rna: rna.to_string()})
    }
}
