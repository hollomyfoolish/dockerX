var express = require('express');
var app = express();
 
app.get('/', function (req, res) {
   res.send('Hello World');
})
var server = app.listen(4430, function () {
  var host = server.address().address
  var port = server.address().port
  console.log("server url http://localhost:4430")
})