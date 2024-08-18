# hs-headstarter-hackathon2-tiktok
# Server
### Endpoints
#### Users
- GET /api/users/:username - Get a user
- 
#### Posts
- GET /api/posts - Get all posts
- GET /api/posts/:username - Get all posts for a user
- POST /api/posts/:username - Create a post for a user, body: { title: string, content: string }
  
#### Admin
- DELETE /api/admin/users/:username - Delete a post for a user, headers: { admin_key: string }

#### Auth
- POST /api/auth/signup - Create a user, body: { username: string, password: string }
- POST /api/auth/login - Login, body: { username: string, password: string }