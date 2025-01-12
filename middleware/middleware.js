const express = require('express');
const bodyParser = require('body-parser');
const { exec } = require('child_process');

const app = express();
const port = 3000;

// Middleware para parsear JSON
app.use(bodyParser.json());

// Endpoint para recibir auditorías
app.post('/init', (req, res) => {
  // Construir el comando para ejecutar el script CLI en el contenedor `cli`
  const command = `docker exec cli scripts/registerAudit.sh init`;

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

app.post('/audit', (req, res) => {
  const jsonString = JSON.stringify(req.body);
  const command = `docker exec cli scripts/registerAudit.sh register '${jsonString}'`;

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

app.get('/list', (req, res) => {
  const command = `docker exec cli scripts/registerAudit.sh getAll`;

  exec(command, (error, stdout, stderr) => {
    if (error) {
      console.error(`Error: ${error.message}`);
      return res.status(500).send('Failed to list audits');
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
app.listen(port, '0.0.0.0', () => {
  console.log(`Middleware server running at http://localhost:${port}`);
});
