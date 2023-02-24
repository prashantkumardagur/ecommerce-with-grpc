const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

const path = require('path');

const express = require('express');
const router = express.Router();


// ===================================================================================


const PROTO_PATH = path.join(__dirname, '../../proto.proto');
const packageDefinition = protoLoader.loadSync(PROTO_PATH);
const AuthService = grpc.loadPackageDefinition(packageDefinition).proto.AuthService;

const authClient = new AuthService('localhost:50051', grpc.credentials.createInsecure());


// ===================================================================================


router.get('/', (req, res) => { res.send('Auth route'); });

router.get('/ping', (req, res) => {
    authClient.ping({}, (err, response) => {
      res.send(response.message);
    });
});

router.get('/profile', (req, res) => {
  authClient.getProfile({id: "1"}, (err, response) => {
    if(err) res.json({ error: err });
    else res.json(response);
  });
});

router.get('/register', (req, res) => {
  let user = {
    id: "2",
    name: "John Doe",
    email: "johhny@gmail.com",
    password: "12345678",
    phone: "1234567890"
  };

  authClient.register(user, (err, response) => {
    if(err) res.json({ error: err });
    else res.json(response);
  });
});


// ===================================================================================

module.exports = router;
