import React, { useState, useEffect } from 'react';
import { DataSet } from 'vis-data';
import { Network } from 'vis-network';

const DiagramaDeEstado = () => {
const [pid, setPid] = useState('');
const [estados, setEstados] = useState([]);
const [transiciones, setTransiciones] = useState([]);
const [network, setNetwork] = useState(null);

const handleInputChange = (event) => {
  setPid(event.target.value);
};

const handleButtonClick = (actionType) => {
  fetch(`/api/estado/${actionType}?pid=${pid}`, {
    method: 'GET'  
  })
  .then(response => {
    if (response.ok) {
      console.log(`Solicitud ${actionType} enviada correctamente`);
      obtenerDatosEstado();
    } else {
      console.error(`Error al enviar solicitud ${actionType}`);
    }
  })
  .catch(error => {
    console.error(`Error de red: ${error}`);
  });
};

const obtenerDatosEstado = () => {
  fetch('/api/estado/status')
    .then(response => {
      if (!response.ok) {
        throw new Error('Error al obtener datos de estado ');
      }
      return response.json();
    })
    .then(data => {
      if (!data || !data.states || !Array.isArray(data.states)) {
        throw new Error('Los datos recibidos del servidor no tienen la estructura esperada.');
      }

      const estados = data.states.map((estado, index) => ({ id: index + 1, label: estado }));
      const transiciones = data.states.slice(0, -1).map((_, index) => ({ from: index + 1, to: index + 2 }));

      setEstados(estados);
      setTransiciones(transiciones);
    })
    .catch(error => {
      console.error(error);
    });
};

useEffect(() => {
  obtenerDatosEstado();
}, []); 

useEffect(() => {
  if (estados.length > 0 && transiciones.length > 0) {
    const container = document.getElementById('network');
    const data = {
      nodes: new DataSet(estados),
      edges: new DataSet(transiciones)
    };
    const options = {};

    if (network === null) {
      const nuevaRed = new Network(container, data, options);
      setNetwork(nuevaRed);
    } else {
      network.setData(data);
    }
  }
}, [estados, transiciones, network]);

return (
  <div>
    <h1>Diagrama de estado</h1>
    <input type="text" placeholder="Ingrese el PID" value={pid} onChange={handleInputChange} />
    <button onClick={() => handleButtonClick('status')}>Status</button>
    <button onClick={() => handleButtonClick('start')}>Start</button>
    <button onClick={() => handleButtonClick('stop')}>Stop</button>
    <button onClick={() => handleButtonClick('resume')}>Resume</button>
    <button onClick={() => handleButtonClick('kill')}>Kill</button>
    <div id="network" style={{ width: '800px', height: '600px' }}></div>
  </div>
);
}

export default DiagramaDeEstado;