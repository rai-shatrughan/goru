pub fn show() {
    let is_active: bool = true;

    let runners: i32 = 100;
    let programs = 20i32;

    let suzene_price: f64 = 1120.34;
    let tata_motor_price = 1110.23f64;

    let one_million = 1_00_000u32;

    let row_one = ("ritz", "trips", 20, 2023);
    println!("--");
    println!("is active : {is_active}");
    println!("runners : {runners}");
    println!("programs : {programs}");
    println!("suzene_price : {suzene_price}");
    println!("tata_motor_price : {tata_motor_price}");

    println!("sum of runners and programs : {}", runners + programs);
    println!("one million = {}", one_million);

    println!("tuple pretty : {:#?}", row_one);
    println!("tuple : {:?}", row_one);
    println!("--");
}
