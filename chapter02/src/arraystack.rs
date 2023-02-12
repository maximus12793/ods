use chapter01::interface::List;

#[derive(Clone, Debug)]
pub struct Array<T> {
    a: Box<[Option<T>]>,
    n: usize,
}

impl<T: Clone> Array<T> {
    pub fn new() -> Self {
        Self::with_length(1)
    }

    fn allocate_on_heap(size: usize) -> Box<[Option<T>]> {
        vec![None; size].into_boxed_slice()
    }

    pub fn with_length(capacity: usize) -> Self {
        Self {
            a: Self::allocate_on_heap(capacity),
            n: 0,
        }
    }

    pub fn length(&self) -> usize {
        self.a.len()
    }

    fn resize(&mut self) {
        let mut b = Self::allocate_on_heap(std::cmp::max(self.n * 2, 1));
        // Note: simply swapping the variables does not copy the elements from
        // the old array to the new array. This is because swapping only
        // transfers ownership of the memory allocation, not the data stored in
        // the arrays. The data stored in the arrays remains unchanged after the
        // swap.
        std::mem::swap(&mut self.a, &mut b);
        // The assignment here will acctually transfer the data.
        for i in 0..self.n {
            self.a[i] = b[i].take();
        }
    }
}

impl<T: Clone> List<T> for Array<T> {
    fn size(&self) -> usize {
        self.n
    }

    fn add(&mut self, i: usize, x: T) {
        let n = self.n;
        if n + 1 >= self.length() {
            self.resize();
        }
        if i >= n {
            self.a[n] = Some(x);
        } else {
            self.a[i..n].rotate_right(1);
            let end = self.a[i].replace(x);
            self.a[n] = end;
        }
        self.n += 1;
    }

    fn remove(&mut self, i: usize) -> Option<T> {
        let x = self.a.get_mut(i)?.take();
        if i < self.n {
            self.a[i..self.n].rotate_left(1);
            self.n -= 1;
            if self.length() >= 3 * self.n {
                self.resize();
            }
        }
        x
    }

    fn get(&self, i: usize) -> Option<T> {
        self.a.get(i)?.as_ref().cloned()
    }

    fn set(&mut self, i: usize, x: T) -> Option<T> {
        self.a.get_mut(i)?.replace(x)
    }
}

#[cfg(test)]
mod test {
    use super::Array;
    use chapter01::interface::List;

    #[test]
    fn test_arraystack() {
        let mut array_stack: Array<char> = Array::new();
        assert_eq!(array_stack.size(), 0);
        for (i, elem) in "bred".chars().enumerate() {
            array_stack.add(i, elem);
        }
        array_stack.add(2, 'e');
        array_stack.add(5, 'r');
        assert_eq!((array_stack.size(), array_stack.length()), (6, 10));
        for (i, elem) in "breedr".chars().enumerate() {
            assert_eq!(array_stack.get(i), Some(elem));
        }
        array_stack.add(5, 'e');
        assert_eq!((array_stack.size(), array_stack.length()), (7, 10));
        array_stack.remove(4);
        array_stack.remove(4);
        assert_eq!((array_stack.size(), array_stack.length()), (5, 10));
        array_stack.remove(4);
        array_stack.set(2, 'i');
        // Consistent with book.
        println!("\nArrayStack = {:?}\n", array_stack);
        array_stack.remove(3);
        assert_eq!((array_stack.size(), array_stack.length()), (3, 6));
        for (i, elem) in "bri".chars().enumerate() {
            assert_eq!(array_stack.get(i), Some(elem));
        }
        assert_eq!(array_stack.get(4), None);
        println!("\nArrayStack = {:?}\n", array_stack);
        let mut array_stack: Array<i32> = Array::new();
        let num = 10;
        for i in 0..num {
            array_stack.add(array_stack.size(), i);
        }
        while array_stack.remove(0).is_some() {}
        println!("\nArrayStack = {:?}\n", array_stack);
    }
}
