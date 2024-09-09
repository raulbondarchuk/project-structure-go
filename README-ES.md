# Project Structure Go

## Estructura Completa del Proyecto

![Estructura del Proyecto](https://github.com/user-attachments/assets/0d39681e-cb80-4f27-8161-4d291fe1611d)

## 📁 Directorios

### `/cmd`
El punto de entrada de la aplicación. El nombre de la carpeta debe coincidir con el nombre del archivo ejecutable. El código en esta carpeta debe ser mínimo, con la función `main` delegando en otras partes del proyecto, especialmente en `/internal` y `/pkg`.

### `/internal`
El núcleo de la aplicación. Aquí se encuentra toda la lógica interna que no debe ser importada en otros proyectos. Go garantiza que los paquetes en `/internal` solo puedan ser utilizados dentro del proyecto actual.

En esta carpeta almacenamos la lógica de negocio y el manejo de la base de datos. Una arquitectura común organiza `/internal` en tres capas:

- **Capa de Transporte:** Maneja la interacción con el usuario o cliente (HTTP, gRPC, etc.).
- **Capa de Negocio:** Contiene la lógica de negocio principal del proyecto.
- **Capa de Base de Datos:** Encargada de interactuar con las bases de datos y almacenamiento persistente.

#### Ejemplo de la estructura de `/internal`:

- **`/app`:** El punto donde todas las dependencias y lógica se conectan para ejecutar la aplicación.
- **`/config`:** Manejo de la configuración de la aplicación.
- **`/database`:** Interacción con bases de datos.
- **`/models`:** Definición de las estructuras y modelos de datos.
- **`/services`:** Implementación de la lógica de negocio.
- **`/transport`:** Configuración de servidores y controladores HTTP, entre otros.

![Estructura Interna](https://github.com/user-attachments/assets/b97604da-cba2-43b2-aa71-875956ebfaef)

### `/pkg`
Paquetes compartidos que pueden ser reutilizados en otros proyectos. Si el código en `/internal` es privado, el código en `/pkg` es público y puede ser importado por otras aplicaciones.

### `/configs`
Archivos de configuración del proyecto, como los archivos YAML utilizados durante la compilación y ejecución.

### `/api`
Documentación de la API (OpenAPI/Swagger) y esquemas JSON.

### `/build`
Archivos y scripts relacionados con la construcción del proyecto, como Dockerfile.

### `/deployments`
Manifiestos y scripts para el despliegue, como configuraciones para Kubernetes, Docker Compose y Ansible.

### `/docs`
Documentación del proyecto. Aquí puedes almacenar especificaciones técnicas y documentación escrita a mano que complementa la generada automáticamente por herramientas como Godoc.

### `/README.md`
El README es el primer contacto que los usuarios tienen con tu proyecto. Debe incluir una descripción general, instrucciones de instalación y ejemplos de uso.

---

## 📂 Directorios Comunes Adicionales

### `/scripts`
Scripts auxiliares para tareas como compilación, instalación y análisis del código. Mantén tu Makefile sencillo y delega en estos scripts tareas más complejas.

### `/testdata`
Datos externos utilizados en los tests. Puedes organizar esta carpeta como lo necesites.

### `/tools`
Herramientas utilizadas para la gestión del proyecto. Estos pueden importar código desde `/pkg` o `/internal`.

### `/assets`
Recursos estáticos como imágenes, íconos, o fuentes que necesita la aplicación.

### `/web`
Directorio que contiene recursos web estáticos como HTML, CSS, JavaScript o archivos de plantillas si estás desarrollando una aplicación web.

### `/migrations`
Archivos de migraciones de bases de datos, usualmente SQL.

---

¡Sigue este esquema para mantener un proyecto Go limpio, modular y fácil de mantener!
