use actix_web::{web, App, HttpServer, HttpResponse, Responder};
use serde::{Deserialize, Serialize};
use reqwest::Client;

#[derive(Debug, Deserialize, Serialize)]
struct Data {
    album: String,
    year: String,
    artist: String,
    ranked: String,
}

async fn insert_data(data: web::Json<Data>) -> impl Responder {
    println!("Data: {:?}", data);

    let url = "http://localhost:3001"; 

    let json_data = serde_json::to_string(&data.into_inner()).unwrap();

    match Client::new().post(url)
        .header("Content-Type", "application/json")
        .body(json_data)
        .send()
        .await {
        Ok(response) => {
            println!("Response: {:?}", response);
            HttpResponse::Ok().body("Successful POST request")
        },
        Err(err) => {
            println!("Error sending request: {:?}", err);
            HttpResponse::InternalServerError().body("Error sending POST request")
        }
    }
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .service(web::resource("/insert").route(web::post().to(insert_data)))
    })
    .bind("127.0.0.1:3000")?
    .run()
    .await
}
