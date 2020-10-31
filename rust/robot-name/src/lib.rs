use lazy_static::lazy_static;
use rand::{thread_rng, Rng};
use std::collections::HashSet;
use std::sync::Mutex;

const MAX_REGISTRY_SIZE: usize = 26 * 26 * 10 * 10 * 10;

pub struct Registry {
    used_names: HashSet<String>,
}

impl Registry {
    pub fn new() -> Self {
        Self {
            used_names: HashSet::new(),
        }
    }

    fn get_unique_name(&mut self) -> String {
        if self.used_names.len() >= MAX_REGISTRY_SIZE {
            panic!("All names used");
        }

        let mut random_name = Registry::generate_random_name();
        while self.used_names.contains(&random_name) {
            random_name = Registry::generate_random_name();
        }
        self.used_names.insert(random_name.clone());
        random_name
    }

    fn remove_name(&mut self, name: String) {
        self.used_names.remove(&name);
    }

    fn generate_random_name() -> String {
        let mut rng = thread_rng();
        format!(
            "{}{}{:03}",
            rng.gen_range(b'A', b'Z' + 1) as char,
            rng.gen_range(b'A', b'Z' + 1) as char,
            rng.gen_range(0, 1000)
        )
    }
}

lazy_static! {
    static ref REGISTRY: Mutex<Registry> = Mutex::new(Registry::new());
}

pub struct Robot {
    name: String,
}

impl Robot {
    pub fn new() -> Self {
        Self {
            name: REGISTRY.lock().unwrap().get_unique_name(),
        }
    }

    pub fn name(&self) -> &str {
        self.name.as_str()
    }

    pub fn reset_name(&mut self) {
        let new_name = REGISTRY.lock().unwrap().get_unique_name();
        REGISTRY.lock().unwrap().remove_name(self.name.clone());
        self.name = new_name;
    }
}
