const express = require('express');
const router = express.Router();
const path = require('path');

const gprc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');


// ===================================================================================


const PROTO_PATH = path.join(__dirname, '../../proto.proto');
const packageDefinition = protoLoader.loadSync(PROTO_PATH);
const CatalogService = gprc.loadPackageDefinition(packageDefinition).proto.CatalogService;

const catalogClient = new CatalogService('localhost:50052', gprc.credentials.createInsecure());


// ===================================================================================


router.get('/ping', (req, res) => {
  catalogClient.ping({}, (err, response) => {
    res.send(response.message);
  });
});

router.get("/", (req, res) => {
  catalogClient.getCatalog({}, (err, response) => {
    if (err) res.json({ error: err });
    else res.json(response);
  });
});


// ===================================================================================


module.exports = router;