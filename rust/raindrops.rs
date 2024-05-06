pub fn raindrops(n: u32) -> String {
    let drops: Vec<(u32, &str)> = vec![
        (3, "Pling"),
        (5, "Plang"),
        (7, "Plong"),
    ];
    let mut raindrops: String = String::from("");

    for (drop: u32, sound: &str) in drops {
        if n % drop == 0 {
            raindrops.push_str(sound)
        }
    }

    if raindrops == "" {
        return n.to_string();
    }

    return raindrops;
}
