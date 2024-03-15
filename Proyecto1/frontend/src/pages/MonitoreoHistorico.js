import React, { useEffect, useRef, useState } from 'react';
import Chart from 'chart.js/auto';
import './MonitoreoHistorico.css';

const MonitoreoHistorico = () => {
  const GraficaRAM = useRef(null);
  const GraficaCPU = useRef(null);
  const [ramData, setRAMData] = useState([]);
  const [cpuData, setCPUData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const responseRAM = await fetch('/api/historico/ram');
        const responseCPU = await fetch('/api/historico/cpu');

        const dataRAM = await responseRAM.json();
        const dataCPU = await responseCPU.json();

        dataRAM.sort((a, b) => a.id - b.id);
        dataCPU.sort((a, b) => a.id - b.id);

        setRAMData(dataRAM);
        setCPUData(dataCPU);
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    fetchData(); 

    const intervalId = setInterval(fetchData, 2000); 

    return () => clearInterval(intervalId);
  }, []);

  useEffect(() => {
    const datosRAM = {
      labels: ramData.map(entry => entry.id),
      datasets: [{
        label: 'Uso de RAM',
        data: ramData.map(entry => entry.ocupada),
        borderColor: 'rgb(75, 192, 192)',
        backgroundColor: 'rgba(75, 192, 192, 0.2)',
        fill: true
      }]
    };

    const datosCPU = {
      labels: cpuData.map(entry => entry.id),
      datasets: [{
        label: 'Uso de CPU',
        data: cpuData.map(entry => entry.ocupada),
        borderColor: 'rgb(75, 192, 192)',
        backgroundColor: 'rgba(75, 192, 192, 0.2)',
        fill: true
      }]
    };

    const chartRAM = new Chart(GraficaRAM.current, {
      type: 'line',
      data: datosRAM
    });

    const chartCPU = new Chart(GraficaCPU.current, {
      type: 'line',
      data: datosCPU
    });

    return () => {
      chartRAM.destroy();
      chartCPU.destroy();
    };
  }, [ramData, cpuData]);

  return (
    <div>
      <h1>Monitoreo Histórico</h1>
      <div className="grafica">
        <div className="grafica-item">
          <canvas ref={GraficaRAM}></canvas>
        </div>
        <div className="grafica-item">
          <canvas ref={GraficaCPU}></canvas>
        </div>
      </div>
    </div>
  );
}

export default MonitoreoHistorico;
