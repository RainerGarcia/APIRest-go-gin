<h1>Tarea de API REST con Goland y GIN</h1>
<p>para la materia de paradigma de programacion I, IUTEPAL, semestre 3, seccion 1301</p>

# API REST Go Gin - Personas

Este proyecto es una API REST construida con Go y el framework Gin para gestionar una lista de personas.

## Rutas disponibles

### Obtener todas las personas
- **GET** `/personas`
- **Descripción:** Devuelve todas las personas.

### Buscar persona por ID
- **GET** `/personas/id/:id`
- **Ejemplo:** `/personas/id/1`
- **Descripción:** Devuelve la persona cuyo ID coincide.

### Buscar persona por nombre exacto
- **GET** `/personas/nombre/:nombre`
- **Ejemplo:** `/personas/nombre/Juan`
- **Descripción:** Devuelve la persona cuyo nombre coincide exactamente.

### Buscar persona por apellido exacto
- **GET** `/personas/apellido/:apellido`
- **Ejemplo:** `/personas/apellido/Pérez`
- **Descripción:** Devuelve la persona cuyo apellido coincide exactamente.

### Buscar persona por edad exacta
- **GET** `/personas/edad/:edad`
- **Ejemplo:** `/personas/edad/30`
- **Descripción:** Devuelve la persona cuya edad coincide exactamente.

### Buscar persona por cédula exacta
- **GET** `/personas/cedula/:cedula`
- **Ejemplo:** `/personas/cedula/123456789`
- **Descripción:** Devuelve la persona cuya cédula coincide exactamente.

### Crear una nueva persona
- **POST** `/personas`
- **Body (JSON):**
  ```json
  {
    "nombre": "Pedro",
    "apellido": "Ramírez",
    "edad": 28,
    "cedula": "1122334455"
  }
  ```
- **Descripción:** Crea una nueva persona.

### Actualizar una persona por ID
- **PUT** `/personas/:id`
- **Body (JSON):** Igual que en POST.
- **Descripción:** Actualiza los datos de la persona con el ID indicado.

### Eliminar una persona por ID
- **DELETE** `/personas/:id`
- **Descripción:** Elimina la persona con el ID indicado.

---

## Notas

- Cambia los valores de ejemplo según tus datos.
- El servidor corre por defecto en `localhost:8080`.
- Para producción, cambia el modo a release.

---
