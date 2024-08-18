import WebSocket from "ws";

// USER 1
// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MjQwNzU5MjYsInVzZXIiOiJ1c2VyMSJ9.GUDNaY2kbeOfDwsbyJBLcXqQHSAjUn1adKdSCRTURD0

// USER 0
const socket = new WebSocket("ws://localhost:8080/ws", {
  headers: {
    "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MjQwNzU4NjYsInVzZXIiOiJ1c2VyMCJ9.VcG5nrGxV_-bblCgVqDnjQBLDUeUom_mMXWArSvkpRU" 
  }
});

socket.on("open", () => {
  console.log("Connected to the websocket server");
});

socket.on("message", (data) => {
  // convert the byte array to a string
  const message = JSON.parse(data.toString());
  console.log("Received message: ", message);
});

socket.on("close", () => {
  console.log("Connection closed");
});