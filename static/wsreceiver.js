
var socket = new WebSocket('ws://localhost:8888');


socket.onopen = function(event) {
   console.log(“Connection established”);
   // Display user friendly messages for the successful establishment of connection
   var.label = document.getElementById(“status”);
   label.innerHTML = ”Connection established”;
}

console.log("Websockets Receiver Loaded");
