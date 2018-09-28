

'use strict';

const express = require('express');
const app = express();
const bodyParser = require('body-parser');

const router = express.Router();
var request = require('request');


app.set('view engine', 'html');

app.engine('html', require('ejs').renderFile);

module.exports = router;

app.use(bodyParser.json());



const port = process.env.PORT || 3000;
const server = app.listen(port);

app.use(bodyParser.json());


require('./routes')(router);
app.use('/', router);


app.use(bodyParser.urlencoded({ extended: true }));
console.log(`Application now running on port ${port}`);