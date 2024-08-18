import WebSocket from "ws";

const socket = new WebSocket("ws://localhost:8080/ws", {
  headers: {
    "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MjQwNDQ3NDcsInVzZXIiOiJ1c2VyMiJ9.UKvig-yiGuYi3uCOX67r6pOh8rlZdTaxDZiMYbWTS9I" 
  }
});

socket.onopen = function() {
    console.log("Connected to WebSocket server");
};
socket.onerror = function(error) {
    console.error("WebSocket Error: ", error);
};
socket.onclose = function() {
    console.log("WebSocket connection closed");
};

socket.onmessage = function(event) {
    console.log("Received: ", event.data);
}