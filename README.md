# TikTok Hack (athon)

# Contents

- [About](https://github.com/lausan3/hs-hackathon2-tiktok/main/README.md#about)
- [Stack](https://github.com/lausan3/hs-hackathon2-tiktok/main/README.md#stack)
- [Front-end](https://github.com/lausan3/hs-hackathon2-tiktok/main/README.md#demo)
- [Back-end](https://github.com/lausan3/hs-hackathon2-tiktok/main/README.md#api-routes)
- [Demo Video](https://github.com/lausan3/hs-hackathon2-tiktok/main/README.md#front-end)

# About

This is the web app and server built by [Anthony](https://www.linkedin.com/in/ant-lau512/) and [Marouane](https://www.linkedin.com/in/marouane-h/) for Headstarter's second hackathon! 

It features a fully functional web app built in Next.js and web server built in Golang. Features include user authentication, posts, and a real-time notification system using Redis PubSub and WebSockets.


# Stack
- Next.js
- Golang - Gin
- Redis
- Railway
- Docker
- AWS
- Google Cloud
- WebSocket


# Front-end
Anthony was responsible for the front-end design and implementation using Next.js and TypeScript. He interfaced to all of the back-end, as well as the system to keep the WebSocket connection for real-time notifications.

Homepage

![Homepage](https://github.com/user-attachments/assets/bc6e6d7e-922e-4424-82f6-80b182d20445)

Sign-up form

![Sign up form](https://github.com/user-attachments/assets/ad5744c3-20f9-4435-b136-9bcd7e440cea)


# Back-end
The web API was built by Marouane in Golang using the Gin http framework. It includes endpoints for user authentication, a real-time notification system for each post built using WebSockets and Redis PubSub.

### Endpoints
#### Users
- GET /api/users/:username - Get a user

#### Posts
- GET /api/posts - Get all posts
- GET /api/posts/:username - Get all posts for a user
- POST /api/posts - Create a post for a user, body: { title: string, content: string } headers: { Authorization: string }
  
#### Admin
- DELETE /api/admin/users/:username - Delete a post for a user, headers: { admin_key: string }

#### Auth
- POST /api/auth/signup - Create a user, body: { username: string, password: string }
- POST /api/auth/login - Login, body: { username: string, password: string }


# Demo

https://github.com/user-attachments/assets/52c12a75-8203-42d5-a113-7dc07c0ba8fb

