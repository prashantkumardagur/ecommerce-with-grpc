const express = require('express');
const router = express.Router();
const path = require('path');

const gprc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');


const { verifyToken } = require('./authroutes');


// ===================================================================================


const PROTO_PATH = path.join(__dirname, '../../proto/proto.proto');
const packageDefinition = protoLoader.loadSync(PROTO_PATH);
const CatalogService = gprc.loadPackageDefinition(packageDefinition).proto.CatalogService;

const catalogClient = new CatalogService('localhost:50053', gprc.credentials.createInsecure());


// ===================================================================================


router.post('/ping', (req, res) => {
  catalogClient.ping({}, (err, response) => {
    res.json(response);
  });
});

// ----------------------------------------------------------------------------------

router.post("/", (req, res) => {
  catalogClient.getCatalog({}, (err, response) => {
    if (err) res.json({ error: err });
    else res.json(response);
  });
});

//----------------------------------------------------------------------------------

router.post("/add", verifyToken, (req, res) => {
  let { product } = req.body;

  catalogClient.addProduct(product, (err, response) => {
    if (err) res.json({ error: err });
    else res.json(response);
  });  
});

//----------------------------------------------------------------------------------

router.post("/delete", verifyToken, (req, res) => { 
  let { id } = req.body;

  catalogClient.deleteProduct({id}, (err, response) => {
    if (err) res.json({ error: err });
    else res.json(response);
  });
});

//----------------------------------------------------------------------------------

router.post("/:id", (req, res) => {
  let { id } = req.params;

  catalogClient.getProduct({id}, (err, response) => {
    if (err) res.json({ error: err });
    else res.json(response);
  });
});



// ===================================================================================


module.exports = router;