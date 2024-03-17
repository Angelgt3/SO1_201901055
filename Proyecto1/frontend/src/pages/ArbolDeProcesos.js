import React, { useState, useEffect } from 'react';
import Chart from 'chart.js/auto';

const ArbolDeProcesos = () => {
  const [arbolData, setArbolData] = useState(null);

  useEffect(() => {
    const obtenerDatosArbol = async () => {
      try {
        const response = await fetch('/api/arbol');
        if (!response.ok) {
          throw new Error('Error al obtener los datos del árbol de procesos');
        }
        const data = await response.json();
        setArbolData(data);
      } catch (error) {
        console.error(error);
      }
    };

    obtenerDatosArbol();
  }, []);

  useEffect(() => {
    if (arbolData) {
      const nombres = arbolData.map(proceso => proceso.Name);
      const ids = arbolData.map(proceso => proceso.PID);

      const ctx = document.getElementById('arbolChart');
      new Chart(ctx, {
        type: 'bar',
        data: {
          labels: nombres,
          datasets: [{
            label: 'PID',
            data: ids,
            backgroundColor: 'rgba(54, 162, 235, 0.2)',
            borderColor: 'rgba(54, 162, 235, 1)',
            borderWidth: 1
          }]
        },
        options: {
          scales: {
            y: {
              beginAtZero: true
            }
          }
        }
      });
    }
  }, [arbolData]);

  return (
    <div>
      <h1>Arbol de Procesos</h1>
      <canvas id="arbolChart" width="400" height="200"></canvas>
    </div>
  );
}

export default ArbolDeProcesos;
