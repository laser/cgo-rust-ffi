extern crate cbindgen;

const VERSION: &'static str = env!("CARGO_PKG_VERSION");

fn main() {
    let crate_dir = std::env::var("CARGO_MANIFEST_DIR").unwrap();

    let cfg = cbindgen::Config::from_root_or_default(std::path::Path::new(&crate_dir));

    let c = cbindgen::Builder::new()
        .with_config(cfg)
        .with_crate(crate_dir)
        .with_header(format!("/* libhasher header version {} */", VERSION))
        .with_language(cbindgen::Language::C)
        .generate();

    match c {
        Ok(res) => {
            res.write_to_file("libhasher.h");
        }
        Err(err) => {
            eprintln!("unable to generate bindings: {:?}", err);
            std::process::exit(1);
        }
    }
}
