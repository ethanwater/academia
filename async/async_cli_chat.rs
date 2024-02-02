#[allow(unused_imports)]
use tokio::{
    io::{AsyncBufReadExt, AsyncWriteExt, BufReader},
    net::{TcpListener, TcpStream},
    sync::broadcast,
};

const TCP_8080: u16 = 8080;

pub enum ProtocolVersion {
    V4,
}

pub struct IP {
    _version: ProtocolVersion,
    address: String,
    port: u16,
}

impl IP {
    pub async fn fetch8080() -> IP {
        if let Ok(stream) = TcpStream::connect("8.8.8.8:53").await {
            let addr = stream.local_addr().unwrap().ip().to_string();
            IP {
                _version: ProtocolVersion::V4,
                address: addr,
                port: TCP_8080,
            }
        } else {
            panic!("fuck, we need internet");
        }
    }
}

pub async fn server(ip: IP, queue: usize) -> std::io::Result<()> {
    bridge(ip, queue).await?;
    Ok(())
}

async fn bridge(ip: IP, queue: usize) -> std::io::Result<()> {
    let host_address = format!("{}:{}", ip.address, ip.port);
    let listener = TcpListener::bind(host_address).await.unwrap();

    let (sender, _receiver) = tokio::sync::broadcast::channel(queue);
    loop {
        let (mut socket, addr) = listener.accept().await.unwrap();
        let sender = sender.clone();
        let mut receiver = sender.subscribe();

        tokio::spawn(async move {
            let (reader, mut writer) = socket.split();

            let mut reader = BufReader::new(reader);
            let mut line = String::new();

            loop {
                tokio::select! {
                    bytes = reader.read_line(&mut line) => {
                        if bytes.unwrap() == 0 {
                            break;
                        }
                        sender.send((line.clone(), addr)).unwrap();
                        line.clear();
                    }
                    bytes = receiver.recv() => {
                        let (mut msg, addr2) = bytes.unwrap();
                        msg = format!("Yumi: {}", msg);
                        if addr != addr2 {
                            writer.write_all(msg.as_bytes()).await.unwrap();
                        }
                    }
                }
            }
        });
    }
}
