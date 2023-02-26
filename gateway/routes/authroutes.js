const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

const path = require('path');

const express = require('express');
const router = express.Router();


// ===================================================================================


const PROTO_PATH = path.join(__dirname, '../../proto/proto.proto');
const packageDefinition = protoLoader.loadSync(PROTO_PATH);
const AuthService = grpc.loadPackageDefinition(packageDefinition).proto.AuthService;

const authClient = new AuthService('localhost:50051', grpc.credentials.createInsecure());


// ===================================================================================


router.get('/', (req, res) => { res.send('Auth route'); });

//----------------------------------------------------------------------------------

router.post('/ping', (req, res) => {
    authClient.ping({}, (err, response) => {
      res.json(response);
    });
});

//----------------------------------------------------------------------------------

router.post('/login', (req, res) => {
  let { email, password } = req.body;

  authClient.login({email, password}, (err, response) => {
    if(err) res.json({ error: err });
    else res.json(response);
  });
});

//----------------------------------------------------------------------------------

router.post('/profile', (req, res) => {
  authClient.getProfile({id: "1"}, (err, response) => {
    if(err) res.json({ error: err });
    else res.json(response);
  });
});

//----------------------------------------------------------------------------------

router.post('/register', (req, res) => {
  let { user } = req.body;

  authClient.register(user, (err, response) => {
    if(err) res.json({ error: err });
    else res.json(response);
  });
});

//----------------------------------------------------------------------------------

router.post('/verify', (req, res) => {
  let { token } = req.body;

  authClient.verify({id: token}, (err, response) => {
    if(err) res.json({ error: err });
    else res.json(response);
  });
});


// ===================================================================================

const verifyToken = (req, res, next) => {
  try {
    let token = req.headers.authorization.split(" ")[1];
    authClient.verify({id: token}, (err, response) => {
      if(response.success){
        req.user = response.message;
        next();
      } else {
        return res.json({ success: false, message: "Unauthorized" });
      }
    });
  } catch (error) {
    return res.json({ success: false, message: "Unauthorized" });
  }
}

module.exports = router;
module.exports.verifyToken = verifyToken;
