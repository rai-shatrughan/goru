use std::fmt::Display;
use std::ops::Add;

pub fn loop_example(init_value: i64, max_value: i64) {
    println!("====");
    let mut i: i64 = init_value;
    println!("i init : {}", i);
    loop {
        i = i + 1;
        println!("i in loop : {}", i);
        if i >= max_value {
            break;
        }
    }
    println!("i after: {}", i);
    println!("====");
}

pub fn while_example(init_value: i64, max_value: i64) {
    println!("----");
    let mut i = init_value;
    println!("i init : {}", i);
    while i < max_value {
        i = i + 1;
        println!("i in loop : {}", i);
    }
    println!("i after: {}", i);
    println!("----");
}

pub fn for_example(init_value: i64, max_value: i64) {
    println!("++");
    println!("i init : {}", init_value);
    for i in init_value..max_value + 1 {
        println!("i in loop : {}", i);
    }
    println!("i after: {}", max_value);
    println!("++");
}

pub fn vec_example<T: Display>(array: &Vec<T>) {
    println!("--");
    for item in array {
        println!("Vec Item : {}", item)
    }
    println!("--");
}

pub fn transform_vec_example<T: Add<Output = T> + Copy>(array: &mut Vec<T>, increment_by: T) {
    for i in 0..array.len() {
        array[i] = array[i] + increment_by;
    }
}

pub fn show() {
    let mut my_vec = vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
    let mut my_float_vec = vec![1.1, 2.1, 3.1, 4.1, 5.1, 6.1, 7.1, 8.1, 9.1, 10.1];
    loop_example(0, 10);
    while_example(0, 10);
    for_example(0, 10);

    transform_vec_example(&mut my_vec, 5);
    transform_vec_example(&mut my_float_vec, 5.0);

    vec_example(&my_vec);
    vec_example(&my_float_vec);
}
