const express = require('express');
const path = require('path');
const fs = require('fs');


for(let p in process.env){
    console.log(`${p}: ${process.env[p]}`)
}

// console.log('secret: ' + fs.readFileSync('/run/secrets/db_password.txt', "utf8").trim());

const app = express();
app.use('/static/', express.static(path.resolve(__dirname, '../static')));
app.get('/api/header', (req, res) => {
    res.status(200).send(req.headers)
  });
app.listen(3003, () => {
    console.log('server startup: 3003');
});