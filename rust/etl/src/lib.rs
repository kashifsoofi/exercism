use std::collections::BTreeMap;

pub fn transform(h: &BTreeMap<i32, Vec<char>>) -> BTreeMap<char, i32> {
    let mut result = BTreeMap::new();
    for (score, letters) in h {
        for l in letters {
            result.insert(l.to_ascii_lowercase(), *score);
        }
    }
    result
}
