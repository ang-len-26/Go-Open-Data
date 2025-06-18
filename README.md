# 🌍 Go Open Data API

Este proyecto es una **API RESTful pública** construida con **Go (Golang)**, que proporciona acceso estructurado a información detallada sobre países, ciudades, idiomas, regiones y monedas. Está orientado a ser una fuente educativa, informativa y extensible.

## 📌 ¿Qué hace esta API?

Esta API expone datos públicos sobre:

- 🌐 **Países**: nombre, capital, población, región, subregión, superficie, etc.
- 🏙️ **Ciudades**: población, ubicación (latitud/longitud), país al que pertenecen.
- 🗺️ **Regiones/Subregiones**: agrupación geográfica de países.
- 💬 **Idiomas**: nombre, código ISO, nombre nativo y países donde se hablan.
- 💱 **Monedas**: código, símbolo, países que la utilizan.

Soporta filtros y paginación para facilitar el consumo de datos en aplicaciones frontend, además de una estructura profesional basada en relaciones reales entre entidades.

## 🛠️ Tecnologías utilizadas

- **Go (Golang)** 🦫
- **Gin** (framework HTTP)
- **PostgreSQL** 🐘
- **pgx** (driver para PostgreSQL)
- **godotenv** (para variables de entorno)
- **CORS middleware personalizado**

## 📁 Estructura del proyecto

go-open-data-api/
├── cmd/ # Punto de entrada principal de la app
│ └── main.go # Configuración de servidor y rutas
├── config/ # Lógica de conexión a la base de datos y entorno
├── controllers/ # Controladores por entidad (countries, cities, etc.)
├── database/ # Inicialización y setup opcional de la base
├── models/ # Modelos estructurados para JSON y SQL
├── routes/ # Archivo(s) de rutas centralizados o modulares
├── .env # Variables de entorno locales (PORT, DATABASE_URL, etc.)
├── go.mod / go.sum # Dependencias del proyecto
└── README.md # Este archivo

## 📦 Endpoints implementados

### Countries

- `GET /api/v1/countries` — Lista con filtros y paginación
- `GET /api/v1/countries/{id}` — Detalle de un país
- `GET /api/v1/countries/{id}/cities` — Ciudades del país
- `GET /api/v1/countries/{id}/languages` — Idiomas del país
- `GET /api/v1/countries/{id}/currencies` — Monedas del país

### Cities

- `GET /api/v1/cities` — Lista de ciudades con filtros
- `GET /api/v1/cities/{id}` — Detalle de una ciudad

### Regions / Subregions

- `GET /api/v1/regions`
- `GET /api/v1/regions/{name}/countries`
- `GET /api/v1/subregions`
- `GET /api/v1/subregions/{name}/countries`

### Languages

- `GET /api/v1/languages`
- `GET /api/v1/languages/{iso}`
- `GET /api/v1/languages/{iso}/countries`

### Currencies

- `GET /api/v1/currencies`
- `GET /api/v1/currencies/{code}`
- `GET /api/v1/currencies/{code}/countries`

## 🚀 Cómo ejecutar el proyecto

1. Clona el repositorio:
   ```bash
   git clone https://github.com/ang-len-26/go-open-data-api.git
   cd go-open-data-api
   ```

# Ejecuta la app

go run ./cmd
