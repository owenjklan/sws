var hostname = window.location.hostname;
var host = window.location.host;
var wsAddr1 = 'ws://' + host + '/ws';
var wsAddr2 = 'ws://' + host + '/notifications';

var webSocket1 = new WebSocket(wsAddr1);
var webSocket2 = new WebSocket(wsAddr2);

webSocket1.onopen = function(event) {
   console.log('Connection established');
   // Display user friendly messages for the successful establishment of connection
   var label = document.getElementById('status');
   label.innerHTML = "Connection established";
}

webSocket1.onmessage = function(e){
   	var server_message = e.data;
   	// var message_json = JSON.parse(server_message)
    $('#body-div').append('<div class="row">' + server_message + '</div>');
   	console.log(server_message);
}

webSocket2.onmessage = function(e){
    var server_message = e.data;
    // var message_json = JSON.parse(server_message)
    $('#status').html('<div class="row">' + server_message + '</div>');
    console.log(server_message);
}

$(document).ready( function () {
	console.log("It's me!");
});

console.log("WebSockets Receiver Loaded");
