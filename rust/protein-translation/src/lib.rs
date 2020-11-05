use std::collections::HashMap;

pub struct CodonsInfo<'a> {
    // This field is here to make the template compile and not to
    // complain about unused type lifetime parameter "'a". Once you start
    // solving the exercise, delete this field and the 'std::marker::PhantomData'
    // import.
    codon_protien_map: HashMap<&'a str, &'a str>,
}

impl<'a> CodonsInfo<'a> {
    pub fn name_for(&self, codon: &str) -> Option<&'a str> {
        if codon.is_empty() {
            return None;
        }

        self.codon_protien_map.get(codon).copied()
    }

    pub fn of_rna(&self, rna: &str) -> Option<Vec<&'a str>> {
        let mut proteins = Vec::new();
        for chunk in rna.chars().collect::<Vec<char>>().chunks(3) {
            if chunk.len() != 3 {
                return None;
            }

            let codon: String = chunk.iter().collect();
            if let Some(protein) = self.name_for(codon.as_str()) {
                if protein.contains("stop") {
                    break;
                }

                proteins.push(protein);
            }
        }
        
        if proteins.is_empty() {
            return None;
        }

        Some(proteins)
    }
}

pub fn parse<'a>(pairs: Vec<(&'a str, &'a str)>) -> CodonsInfo<'a> {
    CodonsInfo {
        codon_protien_map: pairs.into_iter().collect(),
    }
}
