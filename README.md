### 📝 Task Tracker CLI

A simple command-line tool to manage your tasks efficiently. Supports adding, updating, deleting, marking tasks, and listing them in different states.

---

## 🚀 Features
- Add new tasks
- Update task descriptions
- Delete tasks
- Mark tasks as **done**, **in-progress**, or **to-do**
- List tasks by status

---

## 📦 Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/eralieff/task-tracker.git
   cd task-tracker  
   ```  
2. Install dependencies:
   ```sh
   go mod tidy  
   ```  
3. Build the project (optional):
   ```sh
   go build -o task-tracker  
   ```  

---

## 📌 Usage

Run the program using:
```sh
  go run cmd/main.go <command> [args...]
```
Or, if you built the binary:
```sh
  ./task-tracker <command> [args...]
```

### 🏗 Commands

#### ➕ Add a Task
```sh
  go run cmd/main.go add "Buy groceries"
```

#### ✏️ Update a Task
```sh
  go run cmd/main.go update <task_id> "Updated description"
```

#### ❌ Delete a Task
```sh
  go run cmd/main.go delete <task_id>
```

#### 🚀 Mark Task as In-Progress
```sh
  go run cmd/main.go mark-in-progress <task_id>
```

#### ✅ Mark Task as Done
```sh
  go run cmd/main.go mark-done <task_id>
```

#### 📋 List Tasks
- Show all tasks:
  ```sh
  go run cmd/main.go list  
  ```
- Show only done tasks:
  ```sh
  go run cmd/main.go list done  
  ```
- Show only todo tasks:
  ```sh
  go run cmd/main.go list todo  
  ```
- Show tasks in progress:
  ```sh
  go run cmd/main.go list in-progress  
  ```

---

## 🛠️ Configuration
- Tasks are stored in `tasks.json` by default.
- You can modify `task/storage.go` to change the storage location.

---

## 📜 License
This project is licensed under the MIT License.