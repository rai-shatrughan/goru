#![allow(dead_code)]
struct Person {
    name: String,
    age: u8,
}

pub fn struct_show() {
    let p1 = Person {
        name: String::from("sr"),
        age: 30,
    };
    println!("--");
    println!("person name - {}", p1.name);
    println!("person age - {}", p1.age);
    println!("--");
}

pub enum WebEvent {
    Copy,
    Click,
    Select,
}

pub type MyTypedEvent = WebEvent;

pub fn enum_show(event: WebEvent) {
    println!("--");
    match event {
        WebEvent::Click => println!("Clicked"),
        WebEvent::Copy => println!("Copied"),
        WebEvent::Select => println!("Selected"),
    }
    println!("--");
}

enum Color {
    Red = 10,
    Yellow = 15,
    White = 20,
}

enum Number {
    Zero,
    One,
    Two,
}

pub fn enum_discriminator_show() {
    println!("--");
    println!("Enum Number zero value i32 - {}", Number::Zero as i32);
    println!("Enum Number One value i32 - {}", Number::One as i32);
    println!("--");
}

pub fn enum_discriminator_explicit_show() {
    println!("--");
    println!("Enum Color Red value i32 - {}", Color::Red as i32);
    println!("Enum Color White value i32 - {}", Color::White as i32);
    println!("--");
}
