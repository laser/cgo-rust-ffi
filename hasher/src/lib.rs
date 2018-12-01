#![feature(libc)]
#![feature(bufreader_buffer)]

extern crate libc;
extern crate sha2;

use sha2::{Digest, Sha256};
use std::ffi::CStr;
use std::ffi::CString;
use std::fs::OpenOptions;
use std::io::Read;

#[no_mangle]
pub unsafe extern "C" fn checksum_file(file_path: *mut libc::c_char) -> *const libc::c_char {
    let c_str: &CStr = CStr::from_ptr(file_path);
    let str_slice: &str = c_str.to_str().expect("could not convert C string");

    let mut hasher = Sha256::new();

    match OpenOptions::new().read(true).open(str_slice) {
        Ok(mut f) => {
            loop {
                let mut buf = vec![0; 1024 * 1024];
                match f.read(&mut buf) {
                    Ok(n) => {
                        if n == 0 {
                            println!("[rust]   EOF");
                            break;
                        } else {
                            println!("[rust]   read n={} bytes", n);
                            hasher.input(buf);
                        }
                    }
                    Err(err) => panic!(err),
                }
            }

            // take first 32 chars of hex-encoded digest
            let digest: String = format!("{:.32x}", &hasher.result());

            // produce a C string (needs to be freed by caller)
            return CString::new(digest).unwrap().into_raw();
        }
        Err(err) => panic!(err),
    }
}
