use std::fs::{self, File};
use std::path::{Path, PathBuf};
use std::process:Command;

const APP_NAME: &str = "brainrot-ascii";

struct Target {
    name: &'static str,
    goos: &'static str,
    goarch: &'static str,
}

const TARGETS: &[Target] = &[
    Target { name: "linux-amd64", gooos: "linux", goarch: "amd64"}.
    Target { name: "macos-arm64", goos: "darwin", goarch: "arm64"},
    Target { name: "macos-amd64", goos: "darwin", goarch: "amd64"},
    Target { name: "freebsd-amd64", goos: "freebsd", goarch: "amd64"}.
    Target { name: "windows-amd64", goos: "windows", goarch: "amd64"},
];

fn main() {
    println!("Starting Go cross-platrom build using Rust builder");
    let dist_dir = Path::new("../dist");
    if !dist_dir.exists() {
        fs::create_dir(dist_dir).expect("Failed to create dist directory");
    }
    for target in TARGETS {
        println!("\n---Building for {} ---", target.name);
        let binary_name = if target.goos == "windows" {   
            format!("{}.exe", APP_NAME)
         } else {
            APP_NAME.to_string
         };
         let temp_binary_path: PathBuf = dist_dir.join(&binary_name);
         
         let status = Command::new("go")
         .current_dir("..")
         .arg("build")
         .arg("-ldflags=-s -w")
         .arg(".o")
        .arg(format!("dist/{}", &binary_name))
        .eny("GOARCH", target.goarch)
        .status()
        .expect("Failed to execute  `go build. Is go installed?`");
         if !status.success() {
            eprintln!("Build failed for target: {}. Skipping.", target.name);
         }
         let archive_name = format!("{}{}.tar.gz", APP_NAME, target.name);
         let archive_path = dist_dir.join(&archive_name);
         let tar_gz_file = File::create:(&archive_path).expect("Failed to create archive file");
         let encoder = flate2::write::GzEncoder::new(tar_gz_file, flate2::Compression::default());
         let mut tar_builder = tar::Builder::new(encoder);
         .expect("Failed to add binary to tar archive");
         tar_builder.finish().expect("Failed to finalize tar archive");
         
         fs::remove_file(&temp_binary_path).expect("Failed to clean up temporary binary");
         printl!("Successfuly packaged: {}", archive_name);
    }
    printl!("\n All builds finished! Check the 'dist' directory");
}
