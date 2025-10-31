# Shortest Path API

API en Go para encontrar el depósito más cercano a una ubicación de accidente usando Dijkstra.

## Contenido
- Descripción
- Requisitos
- Ejecutar localmente
- Docker
- Ejemplo de petición
- Estructura del proyecto

## Descripción
Servicio HTTP que recibe un grafo, una ubicación de accidente y una lista de depósitos, y devuelve el depósito más cercano, la ruta y la distancia.

## Requisitos
- Go 1.20+ instalado
- Docker (opcional para ejecutar en contenedor)

## Ejecutar localmente
1. Descargar dependencias:
   ```
   go mod download
   ```
2. Ejecutar:
   ```
   go run ./cmd/api
   ```
   Por defecto el binario sirve en el puerto 8080 (puede cambiarse vía variable de entorno).

Para compilar el binario:
```
go build -o bin/shortest-path-api ./cmd/api
```

## Docker
Build:
```
docker build -t shortest-path-api:latest .
```
Run:
```
docker run --rm -p 8080:8080 shortest-path-api:latest
```

## API
POST /shortest-path
- Content-Type: application/json
- Body (ejemplo):
```json
{
  "Graph": {
    "A": {"B": 7, "C": 9, "F": 14},
    "B": {"A": 7, "C": 10, "D": 15},
    "C": {"A": 9, "B": 10, "D": 11, "F": 2},
    "D": {"B": 15, "C": 11, "E": 6},
    "E": {"D": 6, "F": 9},
    "F": {"A": 14, "C": 2, "E": 9}
  },
  "AccidentLocation": "E",
  "Depots": ["A", "B", "C"]
}
```

- Respuesta (ejemplo):
```json
{
  "FromDepot": "C",
  "To": "E",
  "Path": ["C","F","E"],
  "Distance": 11
}
```

Errores comunes:
- Asegurarse que el JSON del grafo corresponde a la estructura esperada por `domain.Graph`.
- Si no existe ruta entre depósito y accidente se devuelve un error indicando que no se encontró camino.

## Estructura del proyecto (resumen)
- cmd/api - punto de entrada
- internal/application - lógica de negocio (servicio)
- internal/domain - tipos y errores
- internal/infrastructure - handlers y algoritmos (Dijkstra)
- Dockerfile, go.mod, .gitignore, README.md

## Notas
- Si el nombre del módulo en `go.mod` cambia, actualizar los import paths para que coincidan.
- El endpoint espera grafos con claves de nodo como strings y pesos numéricos; ajustar según definición en `internal/domain`.