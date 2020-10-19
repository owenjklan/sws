let results_table;
var socket = new WebSocket('/ws');

socket.onopen = function(event) {
   console.log('Connection established');
   // Display user friendly messages for the successful establishment of connection
   var label = document.getElementById('status');
   label.innerHTML = "Connection established";
}

socket.onmessage = function(e){
   	var server_message = e.data;
   	// var message_json = JSON.parse(server_message)
    $('body-div').append('<div class="row">server_message</div>');
   	console.log(server_message);
}

$(document).ready( function () {
	console.log("It's me!");
  $('#map-popover').popover({
   	html: true,
   	trigger: "hover"
	});
});

console.log("Websockets Receiver Loaded");
