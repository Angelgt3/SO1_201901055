*Universidad de San Carlos de Guatemala*  
*Escuela de Ingeniería en Ciencias y Sistemas, Facultad de Ingenieria*  
*Sistemas Operativos 1, 2023.*
___
# **PROYECTO 2**
___
**Angel Geovany Aragón Pérez**  
**201901055**
___
# Preguntas
## ¿Qué servicio se tardó menos? ¿Por qué?
El servicio gRPC se ejecutó más rápidamente en comparación con el servicio Rust debido a su diseño optimizado para la comunicación eficiente entre servicios distribuidos mediante HTTP/2. 

## ¿En qué casos utilizarías grpc y en qué casos utilizarías rust?
- Se utilizaría gRPC en casos donde se requiera una comunicación eficiente entre servicios distribuidos, alta velocidad y baja latencia.
- Se utilizaria Rust en situaciones donde se necesite un rendimiento extremadamente alto y un control preciso sobre los recursos del sistema.

___
# Manual Técnico 

## Introducción

Este manual técnico detalla la implementación de un sistema distribuido de votaciones para un concurso de bandas de música guatemalteca. El sistema utiliza microservicios desplegados en Kubernetes, con sistemas de mensajería para encolar datos y Grafana como interfaz gráfica de dashboards en tiempo real.

## Objetivos

- Implementar un sistema distribuido con microservicios en Kubernetes.
- Utilizar sistemas de mensajería para encolar distintos servicios.
- Utilizar Grafana como interfaz gráfica de dashboards.

## Tecnologías Utilizadas

- Python: Para desarrollar un generador de tráfico utilizando Locust.
- Golang: Para desarrollar el servidor y cliente productor de gRPC, así como el daemon consumidor.
- Rust: Para desarrollar el servidor y cliente productor para encolar a kafka.
- Node.js: Para desarrollar una API y una webapp con Vue.js.
- Docker: Para la construcción de los contenedores de Redis, MongoDB y los servicios.
- Kubernetes: Para el despliegue de los servicios en un clúster en Google Kubernetes Engine.
- Kafka: Para encolar los datos enviados por los productores.
- Redis: Base de datos para almacenar contadores de votos en tiempo real.
- MongoDB: Base de datos para almacenar logs.
- Grafana: Interfaz gráfica de dashboards para visualizar los contadores en tiempo real.

## Descripción de la Arquitectura

### Locust
Genera tráfico y envía datos a los distintos servidores desplegados en Kubernetes. Desarrollado en Python.

### Kubernetes
Se utiliza un clúster de Kubernetes en Google Cloud. Contiene productores y consumidores, así como el servidor de Kafka.

### Kafka
Servidor dispuesto a recibir peticiones de los productores y encolarlas para el consumidor.

### Productores
- **gRPC:** Desarrollado en Golang.
- **Rust:** Desarrollado en Rust.

### Consumidor
Daemon escrito en Golang, desplegado en un pod con autoescalado.

### Bases de Datos
- **Redis:** Almacena contadores de votos en tiempo real.
- **MongoDB:** Almacena logs.

### Grafana
Interfaz gráfica de dashboards para visualizar los contadores en tiempo real.

### Cloud Run
API en Node.js y webapp con Vue.js desplegadas en Cloud Run para visualizar los registros de logs de MongoDB.

## Descripción de Deployments y Services de Kubernetes
**Deployments**
- rust-deployment: Despliega los contenedores de servidor y cliente para el servicio Rust. Estos contenedores son utilizados para el procesamiento de datos relacionados con las votaciones.
- consumer-deployment: Despliega el consumidor del sistema, que procesa los datos recibidos y los envía a las bases de datos Redis y MongoDB para su almacenamiento.
- kafka-deployment: Despliega un clúster Kafka, que actúa como un sistema de mensajería para el encolado de datos entre los diferentes componentes del sistema distribuido.

**Services**
- service-rust: Expone el servicio Rust, permitiendo que otros servicios dentro del clúster puedan comunicarse con el servidor y el cliente Rust.
- service-grcp: Expone el servicio gRPC, facilitando la comunicación entre los clientes externos y el servidor gRPC dentro del clúster.
- kafka-service: Expone el servicio Kafka, permitiendo que otros servicios dentro del clúster puedan comunicarse con el clúster Kafka para el encolado y procesamiento de datos.
- redis: Expone el servicio Redis, que proporciona una base de datos en memoria para almacenar contadores de votos en tiempo real.
- mongodb: Expone el servicio MongoDB, utilizado para almacenar logs relacionados con las votaciones.

## Conclusiones
En este proyecto se demostró la viabilidad y eficacia de utilizar microservicios desplegados en Kubernetes, junto con sistemas de mensajería como Kafka y bases de datos como Redis y MongoDB, para construir un sistema escalable y robusto

