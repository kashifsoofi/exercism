use std::iter::FromIterator;

struct Node<T> {
    data: T,
    next: Option<Box<Node<T>>>,
}

impl<T> Node<T> {
    fn new(data: T, next: Option<Box<Node<T>>>) -> Self {
        Node { data: data, next: next }
    }
}

pub struct SimpleLinkedList<T> {
    head: Option<Box<Node<T>>>,
    len: usize,
}

impl<T> SimpleLinkedList<T> {
    pub fn new() -> Self {
        SimpleLinkedList { head: None, len: 0 }
    }

    // You may be wondering why it's necessary to have is_empty()
    // when it can easily be determined from len().
    // It's good custom to have both because len() can be expensive for some types,
    // whereas is_empty() is almost always cheap.
    // (Also ask yourself whether len() is expensive for SimpleLinkedList)
    pub fn is_empty(&self) -> bool {
        self.len == 0
    }

    pub fn len(&self) -> usize {
        self.len
    }

    pub fn push(&mut self, element: T) {
        let node = Box::new(Node::new(element, self.head.take()));

        self.len += 1;
        self.head = Some(node);
    }

    pub fn pop(&mut self) -> Option<T> {
        let node = self.head.take();

        match node {
            Some(mut x) => {
                self.head = x.next.take();
                self.len -= 1;

                Some(x.data)
            },
            None => None
        }
    }

    pub fn peek(&self) -> Option<&T> {
        match self.head {
            None => None,
            _ => Some(&self.head.as_ref().unwrap().data)
        }
    }

    pub fn rev(self) -> SimpleLinkedList<T> {
        let mut revlist = Self::new();

        let mut current_node = self.head;
        while let Some(node) = current_node {
            revlist.push(node.data);
            current_node = node.next;
        }

        revlist
    }
}

impl<T> FromIterator<T> for SimpleLinkedList<T> {
    fn from_iter<I: IntoIterator<Item = T>>(iter: I) -> Self {
        let mut list = Self::new();

        for e in iter {
            list.push(e);
        }

        list
    }
}

// In general, it would be preferable to implement IntoIterator for SimpleLinkedList<T>
// instead of implementing an explicit conversion to a vector. This is because, together,
// FromIterator and IntoIterator enable conversion between arbitrary collections.
// Given that implementation, converting to a vector is trivial:
//
// let vec: Vec<_> = simple_linked_list.into_iter().collect();
//
// The reason this exercise's API includes an explicit conversion to Vec<T> instead
// of IntoIterator is that implementing that interface is fairly complicated, and
// demands more of the student than we expect at this point in the track.

impl<T> Into<Vec<T>> for SimpleLinkedList<T> {
    fn into(self) -> Vec<T> {
        let mut v: Vec<T> = Vec::new();

        let mut current_node = self.head;
        while let Some(node) = current_node {
            v.push(node.data);
            current_node = node.next;
        }

        v.reverse();
        v
    }
}
