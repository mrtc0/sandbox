use std::io::prelude::*;
use std::net::{Shutdown, TcpStream};
use std::net::{IpAddr, Ipv4Addr, SocketAddr};
use std::time::Duration;

fn main() {
    let socket = SocketAddr::new(IpAddr::V4(Ipv4Addr::new(192, 168, 1, 1)), 12345);
    let five_seconds = Duration::new(5, 0);

    let mut stream = TcpStream::connect_timeout(&socket, five_seconds);
    match stream {
        Ok(x) => {
            println!("Connected");
            match x.shutdown(Shutdown::Both) {
                _ => {}
            }

        }
        Err(e) => {
            println!("Err: {}", e.to_string());
        }
    }
}

