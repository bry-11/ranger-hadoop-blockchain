const express = require('express');
const bodyParser = require('body-parser');
const { exec } = require('child_process');

const app = express();
const port = 8080;

// Middleware para parsear JSON
app.use(bodyParser.json());

// Endpoint para recibir auditorías
app.post('/audit', (req, res) => {
  const { userID, resource, action, result, timestamp } = req.body;

  // Validar los datos de entrada
  if (!userID || !resource || !action || !result || !timestamp) {
    return res.status(400).send('Missing required fields');
  }

  // Construir el comando para ejecutar el script CLI en el contenedor `cli`
  const command = `docker exec cli scripts/register-audit.sh ${userID} ${resource} ${action} ${result} ${timestamp}`;

  // Ejecutar el comando
  exec(command, (error, stdout, stderr) => {
    if (error) {
      console.error(`Error: ${error.message}`);
      return res.status(500).send('Failed to register audit');
    }

    if (stderr) {
      console.error(`Stderr: ${stderr}`);
      return res.status(500).send('Error during audit registration');
    }

    console.log(`Stdout: ${stdout}`);
    res.status(200).send('Audit registered successfully');
  });
});

// Iniciar el servidor
app.listen(port, () => {
  console.log(`Middleware server running at http://localhost:${port}`);
});
