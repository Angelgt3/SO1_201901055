<template>
  <body class="log-viewer">
    <div>
      <h2 style="color: white;">Registros de Logs</h2>
      <textarea v-model="logText" rows="15" cols="50" readonly :style="{ backgroundColor: '#222', color: '#fff', marginBottom: '10px', width: '100%', resize: 'none' }"></textarea>
      <button class="update-button" @click="fetchLogs">Actualizar</button>
    </div>
  </body>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      logText: ''
    };
  },
  created() {
    this.fetchLogs();
  },
  methods: {
    async fetchLogs() {
      try {
        const response = await axios.get('/logs');
        this.logText = response.data.map(log => log.data).join('\n');
      } catch (error) {
        console.error('Error al obtener los registros de logs:', error);
      }
    }
  }
};
</script>

<style>
body {
  background-color: #555;
}

.log-viewer {
  padding: 20px;
}

.update-button {
  margin-top: 20px; /* Aumenta el espacio entre el textarea y el botón */
  display: block; /* Asegura que el botón ocupe toda la anchura */
  width: 100%;
  background-color: #888; /* Color de fondo gris para el botón */
  color: #fff; /* Color de texto blanco para el botón */
  border: none; /* Elimina el borde del botón */
  padding: 10px; /* Ajusta el relleno del botón */
  cursor: pointer; /* Cambia el cursor al pasar sobre el botón */
  border-radius: 5px; /* Agrega bordes redondeados al botón */
}
</style>
