# Proyecto de Chat en Golang

Este es un proyecto de chat en tiempo real desarrollado en Golang utilizando el framework [Fiber](https://gofiber.io/).

## Estructura del Proyecto

## Requisitos

- Go 1.22.2 o superior
- [Fiber](https://gofiber.io/)
- [WebSocket](https://github.com/gofiber/websocket/v2)

## Instalación

1. Clona el repositorio:

    ```sh
    git clone https://github.com/tu-usuario/tu-repositorio.git
    cd tu-repositorio
    ```

2. Instala las dependencias:

    ```sh
    go mod tidy
    ```

## Uso

Para iniciar el servidor, ejecuta:

```sh
go run main.go
```
El servidor estará disponible en http://localhost:3000

## Estructura de carpetas
- app/: Contiene la lógica principal de la aplicación.
  - app.go: Configuración y arranque de la aplicación.
  - controllers/: Controladores para manejar la lógica de negocio.
    - broadcast.go: Lógica para la difusión de mensajes.
  - handlers/: Manejadores para las rutas HTTP y WebSocket.
    - main_handler.go: Manejador principal.
    - ws_handler.go: Manejador para WebSocket.
  - models/: Modelos de datos.
    - client.go: Modelo de cliente.
    - message.go: Modelo de mensaje.
- pkg/: Paquetes reutilizables.
  - router/: Configuración de rutas.
    - http_router.go: Rutas HTTP.
    - router.go: Configuración general de rutas.
    - setup.go: Configuración inicial del router.
    - ws_router.go: Rutas WebSocket.
- static/: Archivos estáticos como HTML, CSS, JS.
- views/: Plantillas HTML.
   - index.html: Página principal.
   - layouts/: Layouts para las plantillas.
     - main.html: Layout principal.

****
[Reto de roadmap.sh](https://roadmap.sh/projects/broadcast-server)