const express = require('express');
const app = express();


// ===================================================================================


app.use(express.json());


// ===================================================================================


// Routes
const authRoutes = require('./routes/authroutes');
const catalogRoutes = require('./routes/catalogroutes');


app.use('/auth', authRoutes);
app.use('/catalog', catalogRoutes);


// ===================================================================================


app.get('/', (req, res) => { res.send('Gateway server'); });


app.listen(3000, () => {
    console.log('Listening on port 3000');
});