const express = require('express');
const path = require('path');

const app = express();
app.use('/static/', express.static(path.resolve(__dirname, '../static')));
app.listen(3003, () => {
    console.log('server startup: 3003');
});