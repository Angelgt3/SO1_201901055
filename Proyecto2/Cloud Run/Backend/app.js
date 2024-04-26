const express = require('express');
const mongoose = require('mongoose');
const bodyParser = require('body-parser'); 
const cors = require('cors');
require('dotenv').config();


const app = express();


app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(cors());

// Conexión a MongoDB
mongoose.connect(process.env.MONGODB_URI, {
  useNewUrlParser: true,
  useUnifiedTopology: true
})
.then(() => console.log('MongoDB conectado'))
.catch(err => console.log(err));


// Definición de un modelo de logs
const LogSchema = new mongoose.Schema({
  date: { type: Date, default: Date.now }, 
  data: String
});
const Log = mongoose.model('Log', LogSchema);

// Ruta para guardar un nuevo log
app.post('/logs', async (req, res) => {
  try {
    const { data } = req.body; 
    const log = new Log({ data }); 
    await log.save(); 
    res.status(201).json({ message: 'Log creado exitosamente' });
  } catch (err) {
    console.error(err);
    res.status(500).json({ message: 'Error' });
  }
});


app.get('/logs', async (req, res) => {
  try {
    const logs = await Log.find().sort({ date: -1 }).limit(20); 
    res.json(logs);
  } catch (err) {
    console.error(err);
    res.status(500).json({ message: 'Error' });
  }
});

const PORT = process.env.PORT || 8080;
app.listen(PORT, () => {
  console.log(`Servidor corriendo en ${PORT}`);
});
