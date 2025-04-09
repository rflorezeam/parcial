# 🧪 Microservicio de Libros (Parcial - Sistemas operativos)

Parcial de Sistemas operativos: Crear microservicio desarrollado en Golang con MongoDB, contenerizado con Docker, probado con Thunder Client y cubierto con pruebas unitarias usando Testify.

---

## 🚀 Crud

CRUD completo sobre la entidad `Libro`, que incluye:

- [x] Crear libro (POST `/libros`)
- [x] Obtener libros (GET `/libros`)
- [x] Actualizar libro por ID (PUT `/libros/{id}`)
- [x] Eliminar libro por ID (DELETE `/libros/{id}`)

---

## 📦 Tecnologías utilizadas

- Golang
- MongoDB
- Docker y Docker-Compose
- Thunder Client (Extension de VSCode)
- Testify para pruebas unitarias
- Godotenv para manejar entornos

---

## 🛠️ Instrucciones para ejecutar

### 1. Clonar el repositorio y cargar variables

```bash
git clone https://github.com/rflorezeam/parcial.git
cd parcial
```

Crear un archivo `.env` con:

```env
ENV=prod
MONGO_URI=mongodb://mongo:27017
MONGO_URI_TEST=mongodb://localhost:27017
MONGO_DB=librosdb
PORT=8080
```

### 2. Ejecutar el microservicio con Docker

```bash
# O usar el script automatizado:
./automatizar.ps1     # Windows PowerShell
```

La app estará disponible en: [http://localhost:8080](http://localhost:8080)

---

## 🐳 Imagen en Docker Hub

La imagen del microservicio fue construida y publicada en Docker Hub. Puede ser ejecutada en cualquier entorno compatible con Docker:

🔗 **https://hub.docker.com/r/rflorezeam/parcial**

### Para ejecutarla directamente:

```bash
docker pull rflorezeam/parcial
docker run -p 8080:8080 --env-file .env rflorezeam/parcial
```

> Asegúrate de tener un archivo `.env` válido al momento de correr el contenedor.

---

## 🧪 Pruebas

### Pruebas unitarias

```bash
go test -cover ./...
```

- Cubre servicios: crear, obtener, actualizar y eliminar tareas

📊 Reporte de cobertura generado:
Archivo: coverage.out  
```bash
go test -coverprofile=coverage.out ./...
```

---

### Thunder Client (colección de pruebas)

- Incluye todos los endpoints
- Validaciones de estado, respuestas y estructura
- Archivo de colección exportado: `thunder-collection_parcial2.json`

---

## 🗄️ Backup de base de datos

Se genera backup dentro del contenedor Mongo con:

```powershell
docker exec -t mongo mongodump --archive="/data/db/backup-$(Get-Date -Format yyyy-MM-dd).gz" --gzip --db=tareasdb
docker cp mongo:/data/db/backup-2025-04-04.gz ./backup-2025-04-04.gz
```



## 🎯 Entregables

- Código fuente con estructura modular
- Dockerfile y docker-compose.yml funcionales
- Archivo `.env` y scripts
- Colección Thunder Client exportada
- Pruebas unitarias y cobertura
- Backup generado desde MongoDB
- BONUS: Script de automatización (`automatizar.ps1`)
- 🐳 Imagen publicada en Docker Hub: [rflorezeam/parcial](https://hub.docker.com/r/rflorezeam/parcial)

---

Desarrollado por **Ricardo Florez** | Sistemas operativos
