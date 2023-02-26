const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

const path = require('path');

const express = require('express');
const router = express.Router();

// ===================================================================================

const PROTO_PATH = path.join(__dirname, '../../proto/proto.proto');
const packageDefinition = protoLoader.loadSync(PROTO_PATH);
const CartService = grpc.loadPackageDefinition(packageDefinition).proto.CartService;

const cartClient = new CartService('localhost:50052', grpc.credentials.createInsecure());

// ===================================================================================

router.post('/ping', (req, res) => {
  cartClient.ping({}, (err, response) => {
    if (err) res.json({ error: err });
    else res.json(response);
  });
});

// ----------------------------------------------------------------------------------

router.post('/', (req, res) => {
  cartClient.getCart({}, (err, response) => {
    if (err) res.json({ error: err });
    else res.json(response);
  });
});

// ===================================================================================

module.exports = router;