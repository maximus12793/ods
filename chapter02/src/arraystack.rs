use chapter01::interface::List;

#[derive(Clone, Debug)]
pub struct Array<T> {
  a: Box<[Option<T>]>,
  n: usize,
}

impl<T> Array<T> {
  pub fn new() -> Self {
    unimplemented!()
  }

  pub fn length(&self) -> usize {
    unimplemented!()
  }
}



impl<T: Clone> List<T> for Array<T> {
  fn size(&self) -> usize {
    self.n
  }

  fn add(&mut self, i: usize, x: T) {
      unimplemented!()
  }

  fn remove(&mut self, i: usize) -> Option<T> {
      unimplemented!()
  }

  fn get(&self, i: usize) -> Option<T> {
    self.a.get(i)?.as_ref().cloned()
  }

  fn set(&mut self, i: usize, x:T) -> Option<T> {
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
        array_stack.remove(4);
        array_stack.remove(4);
        assert_eq!((array_stack.size(), array_stack.length()), (5, 10));
        array_stack.remove(4);
        array_stack.remove(3);
        array_stack.set(2, 'i');
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