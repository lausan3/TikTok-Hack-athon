# hs-headstarter-hackathon2-tiktok
# Server
### Endpoints
#### Posts
- GET /api/posts/:username - Get all posts for a user
- POST /api/posts/:username - Create a post for a user, body: { title: string, content: string }

#### Users
- POST /api/users - Create a user, body: { username: string, password: string }
- DELETE /api/users/:username - Delete a post for a user
- GET /api/users/:username - Get a user