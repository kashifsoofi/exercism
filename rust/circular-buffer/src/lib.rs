use std::collections::VecDeque;

pub struct CircularBuffer<T> {
    buffer: VecDeque<T>,
    capacity: usize,
}

#[derive(Debug, PartialEq)]
pub enum Error {
    EmptyBuffer,
    FullBuffer,
}

impl<T> CircularBuffer<T> {
    pub fn new(capacity: usize) -> Self {
        CircularBuffer { buffer: VecDeque::default(), capacity: capacity }
    }

    pub fn write(&mut self, element: T) -> Result<(), Error> {
        if self.buffer.len() == self.capacity {
            return Err(Error::FullBuffer);
        }

        Ok(self.buffer.push_back(element))
    }

    pub fn read(&mut self) -> Result<T, Error> {
        match self.buffer.pop_front() {
            Some(e) => Ok(e),
            None => Err(Error::EmptyBuffer),
        }
    }

    pub fn clear(&mut self) {
        self.buffer.clear();
    }

    pub fn overwrite(&mut self, element: T) {
        if self.buffer.len() == self.capacity {
            self.buffer.pop_front();
        }

        self.buffer.push_back(element);
    }
}
