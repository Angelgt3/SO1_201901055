use tonic::transport::Channel;
use proto_servidor::get_info_client::GetInfoClient;
use proto_servidor::{Data, RequestId};

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Establecer conexión con el servidor gRPC
    let channel = Channel::from_static("http://[::1]:50051")
        .connect()
        .await?;
    let mut client = GetInfoClient::new(channel);

    // Crear solicitud
    let request = tonic::Request::new(RequestId {
        album: "AlbumName".to_string(),
        year: "2022".to_string(),
        artist: "ArtistName".to_string(),
        ranked: "1".to_string(),
    });

    // Enviar solicitud al servidor
    let response = client.return_info(request).await?;

    // Procesar la respuesta del servidor
    let reply = response.into_inner();
    println!("Respuesta del servidor: {:?}", reply);

    Ok(())
}
