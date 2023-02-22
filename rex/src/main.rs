use example::custom::{MyTypedEvent, WebEvent};

mod example;
fn main() {
    example::loops::show();
    example::primitive::show();
    example::slice::show();

    example::custom::struct_show();
    let event = WebEvent::Click;
    example::custom::enum_show(event);

    example::custom::enum_show(MyTypedEvent::Select);
    example::custom::enum_discriminator_show();
    example::custom::enum_discriminator_explicit_show();
}
