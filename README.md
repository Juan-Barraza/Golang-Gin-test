# Holiday API Service (Golang)

Servicio REST desarrollado en Go utilizando el framework **Gin Gonic** y **Arquitectura Hexagonal**. Consume la información base de feriados desde la API pública de `boostr.cl` al momento de inicializar el contexto del servicio y almacena una caché local para evitar reconsultar el servicio externo en llamadas subsecuentes.

## Arquitectura

El proyecto sigue los principios de Arquitectura Hexagonal (Puertos y Adaptadores):

```text
├── cmd/
│   └── api/                # Punto de entrada de la aplicación
├── internal/
│   ├── domain/             # Entidades, errores e interfaces de repositorios (Core Domain)
│   ├── application/        # Casos de uso, DTOs y puertos
│   ├── infrastructure/     # Adaptadores HTTP (Gin, handlers, router), cliente HTTP y almacenamiento en archivo
│   └── config/             # Carga de variables de entorno
├── docs/                   # Especificación OpenAPI 3.0.3
├── Dockerfile              # Construcción multi-stage en Docker
└── docker-compose.yml      # Despliegue en contenedor
```

---

## Requisitos de Ejecución

- **Docker** y **Docker Compose**
- **Go 1.26** (para desarrollo local sin Docker)

---

## Instrucciones de Ejecución

### 1. Variables de Entorno

Crear el archivo `.env` en la raíz del proyecto basándose en las variables requeridas:

```env
PORT=8080
DATA_FILE_PATH=data/holidays.json
BOOSTR_API_URL=https://api.boostr.cl/holidays.json
```

### 2. Ejecutar con Docker Compose

```bash
docker-compose up --build -d
```

### 3. Detener el Contenedor

```bash
docker-compose down
```

---

## Documentación OpenAPI 3.x.x

La especificación OpenAPI 3.0.3 se encuentra en el archivo:
- `docs/openapi.yaml`

---

## Ejemplos de Uso (cURL)

### 1. Obtener todos los feriados en formato JSON

```bash
curl.exe -i -H "Accept: application/json" "http://localhost:8080/holidays"
```

### 2. Obtener todos los feriados en formato XML

```bash
curl.exe -i -H "Accept: application/xml" "http://localhost:8080/holidays"
```

### 3. Filtrar por Tipo (Civil / Religioso)

```bash
curl.exe -i "http://localhost:8080/holidays?type=Civil"
```

### 4. Filtrar por Rango de Fechas

```bash
curl.exe -i "http://localhost:8080/holidays?startDate=2026-05-01&endDate=2026-10-31"
```

### 5. Filtrar por Tipo y Rango de Fechas

```bash
curl.exe -i "http://localhost:8080/holidays?type=Civil&startDate=2026-09-01&endDate=2026-12-31"
```

### 6. Prueba de Error - Tipo Inválido (HTTP 400)

```bash
curl.exe -i "http://localhost:8080/holidays?type=Invalido"
```
