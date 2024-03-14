import React, { useEffect, useRef, useState } from 'react';
import Chart from 'chart.js/auto';
import './MonitoreoTiempoReal.css';

const MonitoreoTiempoReal = () => {
  const GraficaRAM = useRef(null);
  const GraficaCPU = useRef(null);
  const [ramData, setRAMData] = useState({ usedPct: 50, freePct: 50 });
  const [cpuData, setCPUData] = useState({ usedPct: 50, freePct: 50 });

  useEffect(() => {
    const fetchRAMData = async () => {
      try {
        const response = await fetch('/api/ram');
        const data = await response.json();

        setRAMData({
          usedPct: data.used_ram_pct,
          freePct: data.free_ram_pct
        });
      } catch (error) {
        console.error('Error fetching RAM data:', error);
      }
    };

    const fetchCPUData = async () => {
      try {
        const response = await fetch('/api/cpu');
        const data = await response.json();

        setCPUData({
          usedPct: data.used_cpu_pct,
          freePct: data.free_cpu_pct
        });
      } catch (error) {
        console.error('Error fetching CPU data:', error);
      }
    };

    const fetchDataPeriodically = () => {
      fetchRAMData();
      fetchCPUData();
    };

    const intervalId = setInterval(fetchDataPeriodically, 5000);

    return () => clearInterval(intervalId);
  }, []);

  useEffect(() => {
    const datosRAM = {
      labels: ['Memoria Ocupada', 'Memoria Libre'],
      datasets: [{
        data: [ramData.usedPct, ramData.freePct],
        backgroundColor: [
          'rgb(255, 99, 132)',
          'rgb(75, 192, 192)'
        ]
      }]
    };

    const datosCPU = {
      labels: ['CPU Utilizada', 'CPU Libre'],
      datasets: [{
        data: [cpuData.usedPct, cpuData.freePct],
        backgroundColor: [
          'rgb(255, 99, 132)',
          'rgb(75, 192, 192)'
        ]
      }]
    };

    const chartRAM = new Chart(GraficaRAM.current, {
      type: 'pie',
      data: datosRAM
    });

    const chartCPU = new Chart(GraficaCPU.current, {
      type: 'pie',
      data: datosCPU
    });

    return () => {
      chartRAM.destroy();
      chartCPU.destroy();
    };
  }, [ramData, cpuData]);

  return (
    <div>
      <h1>Monitoreo Tiempo Real</h1>
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

export default MonitoreoTiempoReal;

const styles = `
  .grafica {
    position: absolute;
    left: 5%;
    width: 50%;
    height: 50%;
    margin: 0 auto;
    display: flex;
  }

  .grafica-item {
    flex: 1;
    margin: 0 3%;
  }
`;

const styleElement = document.createElement('style');
styleElement.type = 'text/css';
styleElement.appendChild(document.createTextNode(styles));
document.head.appendChild(styleElement);
