import React, { useEffect, useRef, useMemo } from 'react';
import Chart from 'chart.js/auto';
import './MonitoreoHistorico.css';

const MonitoreoHistorico = () => {
  const GraficaRAM = useRef(null);
  const GraficaCPU = useRef(null);

  const datosRAM = useMemo(() => ({
    labels: ['1','2','3','4','5','6','7','8','9','10'], 
    datasets: [{
      label: 'Uso de RAM', 
      data: [50, 65, 60, 40, 52,50,45,67,80,70], // Datos 
      borderColor: 'rgb(75, 192, 192)', 
      backgroundColor: 'rgba(75, 192, 192, 0.2)',
      fill: true // Rellenar el área bajo la línea
    }]
  }), []); 

  const datosCPU = useMemo(() => ({
    labels: ['1','2','3','4','5','6','7','8','9','10'], 
    datasets: [{
      label: 'Uso de CPU', 
      data: [65, 70, 71, 75, 80,75,75,76,79,91], // Datos 
      borderColor: 'rgb(75, 192, 192)', 
      backgroundColor: 'rgba(75, 192, 192, 0.2)',
      fill: true // Rellenar el área bajo la línea
    }]
  }), []); 

  useEffect(() => {
    const charthartRAM = new Chart(GraficaRAM.current, {
      type: 'line', 
      data: datosRAM
    });

    const chartCPU = new Chart(GraficaCPU.current, {
      type: 'line', 
      data: datosCPU
    });

    return () => {
      charthartRAM.destroy();
      chartCPU.destroy();
    };
  }, [datosRAM,datosCPU]); 

  return (
    <div>
      <h1>Monitoreo Histórico</h1>
      <div class="grafica">
        <div class="grafica-item">
          <canvas ref={GraficaRAM}></canvas>
        </div>
        <div class="grafica-item">
          <canvas ref={GraficaCPU}></canvas>
        </div>
      </div>
    </div>
  );
}

export default MonitoreoHistorico;
