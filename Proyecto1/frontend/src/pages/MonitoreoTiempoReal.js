import React, { useEffect, useRef, useMemo } from 'react';
import Chart from 'chart.js/auto';
import './MonitoreoTiempoReal.css';

const MonitoreoTiempoReal = () => {
  const GraficaRAM = useRef(null);
  const GraficaCPU = useRef(null);

  const DatosRAM = useMemo(() => ({
    labels: ['Memoria Libre', 'Memoria Ocupada'],
    datasets: [{
      data: [50, 100],
      backgroundColor: [
        'rgb(54, 162, 235)',
        'rgb(255, 205, 86)'
      ]
    }]
  }), []); 

  const DatosCPU = useMemo(() => ({
    labels: ['CPU Utilizada', 'CPU Libre'],
    datasets: [{
      data: [25, 75], // Datos de ejemplo para la CPU
      backgroundColor: [
        'rgb(255, 99, 132)',
        'rgb(75, 192, 192)'
      ]
    }]
  }), []); 

  useEffect(() => {
    const chartRAM = new Chart(GraficaRAM.current, {
      type: 'pie',
      data: DatosRAM
    });
  
    const chartCPU = new Chart(GraficaCPU.current, {
      type: 'pie',
      data: DatosCPU
    });

    return () => {
      chartRAM.destroy();
      chartCPU.destroy();
    };
  }, [DatosRAM, DatosCPU]); 

  return (
    <div>
      <h1>Monitoreo Tiempo Real</h1>
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

export default MonitoreoTiempoReal;
