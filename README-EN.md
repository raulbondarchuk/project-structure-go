# Complete Project Structure in Golang

![Project Structure](https://github.com/user-attachments/assets/0d39681e-cb80-4f27-8161-4d291fe1611d)

## 📁 Directories

### `/cmd`
The entry point of the application. The directory name should match the executable name. The code in this folder should be minimal, with the `main` function delegating tasks to other parts of the project, especially `/internal` and `/pkg`.

### `/internal`
The core of the application. Here resides all the internal logic that shouldn't be imported by other projects. Go ensures that packages within `/internal` are only accessible inside the current project.

In this folder, we store business logic and database handling. A common architecture organizes `/internal` into three layers:

- **Transport Layer:** Manages user or client interaction (HTTP, gRPC, etc.).
- **Business Layer:** Contains the core business logic of the project.
- **Database Layer:** Handles interaction with databases and persistent storage.

#### Example structure inside `/internal`:

- **`/app`:** The point where all dependencies and logic are connected to run the application.
- **`/config`:** Manages the application's configuration.
- **`/database`:** Handles database interactions.
- **`/models`:** Defines data models and structures.
- **`/services`:** Implements the business logic.
- **`/transport`:** Configures servers, HTTP handlers, etc.

![Internal Structure](https://github.com/user-attachments/assets/b97604da-cba2-43b2-aa71-875956ebfaef)

### `/pkg`
Shared packages that can be reused in other projects. While the code in `/internal` is private, the code in `/pkg` is public and can be imported by other applications.

### `/configs`
Project configuration files, such as YAML files, used during build and execution.

### `/api`
API documentation (OpenAPI/Swagger) and JSON schemas.

### `/build`
Files and scripts related to project building, such as Dockerfile.

### `/deployments`
Manifests and scripts for deployment, such as Kubernetes configurations, Docker Compose files, and Ansible playbooks.

### `/docs`
Project documentation. This folder stores technical specifications and manually written documentation to complement auto-generated docs like Godoc.

### `/README.md`
The README is the first point of contact users have with your project. It should include an overview, installation instructions, and usage examples.

---

## 📂 Additional Common Directories

### `/scripts`
Auxiliary scripts for tasks like building, installation, and code analysis. Keep the main Makefile simple by delegating complex tasks to these scripts.

### `/testdata`
External data used in tests. You can organize this folder as needed.

### `/tools`
Tools used for project management. These tools can import code from `/pkg` or `/internal`.

### `/assets`
Static resources like images, icons, or fonts required by the application.

### `/web`
Directory containing static web resources such as HTML, CSS, JavaScript, or template files for web applications.

### `/migrations`
Database migration files, typically SQL.

---

Follow this structure to maintain a clean, modular, and maintainable Go project!
