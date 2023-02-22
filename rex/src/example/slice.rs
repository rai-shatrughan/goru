pub fn show() {
    let stock_list: [&str; 2] = ["suzene", "t_motor"];
    let stock_price: [f64; 5] = [111.11; 5];

    borrow_print_slice(&stock_list);

    println!("stock list {:?}", stock_list);
    println!("stock price {:?}", stock_price);
    println!("array length - {}", stock_price.len());
}

fn borrow_print_slice(stocks: &[&str]) {
    for stock in stocks {
        println!("stock - {}", stock);
    }
}
