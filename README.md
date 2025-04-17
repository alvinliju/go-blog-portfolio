## API Architecture

### Blog Routes
GET /blog/[slug]
POST /blog

### Project Routes
GET /project/[slug]
POST /project

### Admin Login
POST /admin/login -> email, password returns jwt
- A middleware to protect the blog routes and project routes with jwt
