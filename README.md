# First App — Simple Blog System with Go

A simple learning project built with **Go** for practicing backend development. It renders pages server-side using **html/template** and lets a user view a list of blog posts and create new ones.

> This is a personal practice/learning project, not intended for external contributions.

## ✨ Features

- Display a list of blog posts
- Create a new post
- User session handling
- Data persistence in MySQL via GORM
- CLI commands powered by Cobra

## 🛠 Tech Stack

| Library | Purpose |
|---|---|
| [Gin](https://github.com/gin-gonic/gin) | Web framework & HTTP routing |
| [gin-contrib/sessions](https://github.com/gin-contrib/sessions) | User session management |
| [GORM](https://gorm.io/) + [MySQL driver](https://github.com/go-sql-driver/mysql) | ORM and MySQL database access |
| [Cobra](https://github.com/spf13/cobra) | CLI commands |
| [Viper](https://github.com/spf13/viper) | Configuration management |
| `html/template` | Server-side HTML rendering (frontend) |

Go version used: **1.26.1**


> The structure above is a suggestion — adjust it to match your actual project layout.

## ⚙️ Prerequisites

- Go 1.26.1 or later
- A running MySQL server

## 🔧 Configuration

The app is configured via a `config.yaml` file:

```yaml
app:
  name: "Blog"
server:
  host: "localhost"
  port: "8085"
db:
  username: "root"
  password: ""
  host: "127.0.0.1"
  port: "3306"
  name: "blog"
```

Update the `db` section with your own MySQL credentials, and `server` with the host/port you want the app to listen on.

## 🚀 Getting Started

1. Clone the project:
   ```bash
   git clone https://github.com/<username>/first-app.git
   cd first-app
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a MySQL database matching the `db.name` value in `config.yaml` (default: `blog`).

4. Run the project:
   ```bash
   go run main.go help
   ```

5. Open your browser at:
   ```
   http://localhost:8085
   ```
   (or whatever `server.host`/`server.port` you set)

## 📌 Usage

- Visit the home page to see the list of blog posts.
- Use the post form in the app to create a new blog post.
