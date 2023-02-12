pub trait Queue<T> {
    fn add(&mut self, x: T);
    fn remove(&mut self) -> Option<T>;
}

pub trait Stack<T> {
    fn push(&mut self, x: T);
    fn pop(&mut self) -> Option<T>;
}

pub trait List<T> {
    fn size(&self) -> usize;
    fn get(&self, i: usize) -> Option<T>;
    fn set(&mut self, i: usize, x: T) -> Option<T>;
    fn add(&mut self, i: usize, x: T);
    fn remove(&mut self, i: usize) -> Option<T>;
}

pub trait USet<T> {
    fn size(&self) -> usize;
    fn add(&mut self, x: T) -> bool;
    fn remove(&mut self, x: T) -> Option<T>;
    fn find(&self, x: T) -> Option<T>;
}

pub trait SSet<T> {
    fn size(&self) -> usize;
    fn add(&mut self, x: T) -> bool;
    fn remove(&mut self, x: T) -> Option<T>;
    fn find(&self, x: T) -> Option<T>;
}
