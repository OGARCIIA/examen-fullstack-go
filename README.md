# Examen Fullstack Go + Next.js

Este proyecto es la entrega del examen técnico para la posición Fullstack Senior.

## Tecnologías utilizadas

- Backend: Go (Golang), Gin, GORM, MySQL
- Frontend: React (Next.js) + Material UI
- CI/CD: GitHub Actions
- Contenedores: Docker + Docker Compose

## Estructura del proyecto

```
examen-fullstack-go/
├── backend/                 # Backend en Go (Clean Architecture)
├── frontend/                # Frontend en Next.js + Material UI
├── .github/workflows/ci.yml # CI/CD pipeline con GitHub Actions
├── docker-compose.yml       # Orquestación de servicios
└── README.md
```

## Funcionalidades implementadas

### Backend

✅ API REST para productos (CRUD completo)  
✅ API REST para órdenes (con control de stock y transacción atómica)  
✅ Autenticación con JWT temporal  
✅ Logs con logrus  
✅ Manejo de concurrencia en creación de órdenes  
✅ Expuesto en contenedor Docker

### Frontend

✅ Listado de productos con stock  
✅ Formulario para crear órdenes (selección de producto + cantidad)  
✅ Uso de JWT para autenticación en las peticiones  
✅ Estilo responsive con Material UI  
✅ Expuesto en contenedor Docker

### CI/CD

✅ Pipeline de CI configurado en GitHub Actions:

- Build del backend
- Build del frontend
- Tests unitarios del backend

## Cómo ejecutar el proyecto localmente

### Requisitos

- Docker
- Docker Compose

### Instrucciones

1️⃣ Clonar el repositorio:

```bash
git clone https://github.com/OGARCIIA/examen-fullstack-go.git
cd examen-fullstack-go
```

2️⃣ Levantar los contenedores:

```bash
docker-compose up --build -d
```

3️⃣ Acceder a la aplicación:

- Frontend → [http://localhost:3000](http://localhost:3000)
- Backend → [http://localhost:8080](http://localhost:8080)

## CI/CD Pipeline

El pipeline se ejecuta automáticamente al realizar `git push` sobre la rama `main`:

- Verifica compilación del backend
- Verifica compilación del frontend
- Corre tests del backend

## Notas

- Los productos deben ser creados inicialmente vía API (Postman) para que el frontend los consuma.
- El JWT temporal es generado desde `/generate-token` y se usa en las peticiones protegidas.

## Autor

Omar García  
GitHub: [@OGARCIIA](https://github.com/OGARCIIA)
