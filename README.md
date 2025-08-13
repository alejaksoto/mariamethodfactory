Analizando el código, el repositorio "mariamethodfactory" es una plataforma web para reservar servicios de belleza (manicure, corte de cabello, coloración) a domicilio, principalmente en Bogotá, Colombia.

# Mariamethodfactory

Plataforma web para reservar servicios de belleza a domicilio (manicure, peluquería, coloración) con atención profesional y personalizada.

## Características principales

- **Frontend (HTML, CSS, JavaScript):**
  - Página informativa y de reservas online.
  - Galería/portafolio de trabajos realizados.
  - Sección de testimonios y contacto.
  - Formulario para agendar citas en línea.

- **Backend (Go):**
  - API REST para recibir y almacenar reservas de servicios (`POST /api/citas`).
  - Notificación automática por WhatsApp al cliente tras reservar.
  - Integración con base de datos para guardar citas.
  - Arquitectura modular (handlers, models, factory para canales de notificación).


## Cómo se desarrollo.

- Se  utiliza el singleton para la conexión con la bd. 
- Se utiliza el metodo factory, para cuando se realice la reserva, se genere la notificación por medio de email o whatsapp. 


## Tecnologías

- **Frontend:** HTML, CSS, JavaScript
- **Backend:** Go (API RESTful)
- **Notificaciones:** WhatsApp (Soporte para Email en arquitectura)
- **Base de datos:** (mariadb)

## Instalación y ejecución

1. Clona el repositorio.
2. Ejecuta el backend en Go (`orfibackend/main.go`) para habilitar la API en `http://localhost:8080`.
3. Abre `index.html` en tu navegador.


---

Este README describe el objetivo, funcionalidad y tecnologías del proyecto según el código actual.
