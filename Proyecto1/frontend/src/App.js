import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css';

import MonitoreoTiempoReal from './pages/MonitoreoTiempoReal';
import MonitoreoHistorico from './pages/MonitoreoHistorico';
import ArbolDeProcesos from './pages/ArbolDeProcesos';
import DiagramaDeEstado from './pages/DiagramaDeEstado';

function App() {
  return (
    <Router>
    <div>
      <header style={{ backgroundColor: '#282c34', minHeight: '10vh', display: 'flex', flexDirection: 'column', justifyContent: 'center', fontSize: 'calc(20px + 2vmin)', color: 'white' }}>
        <div style={{ display: 'flex', justifyContent: 'space-between' }}>
          <span>SO1-201901055</span>
          <span style={{ margin: '0 auto' }}>MODULOS DE KERNEL</span>
        </div>
      </header>
      <nav style={{ backgroundColor: '#f0f0f0', padding: '10px' }}>
        <ul style={{ listStyle: 'none', padding: 0, margin: 0 }}>
          <li style={{ display: 'inline-block', marginRight: '3%' }}><a href="/monitoreo-tiempo-real" style={{ textDecoration: 'none', color: '#333' }}>Monitoreo Tiempo Real</a></li>
          <li style={{ display: 'inline-block', marginRight: '3%' }}><a href="/monitoreo-historico" style={{ textDecoration: 'none', color: '#333' }}>Monitoreo Historico</a></li>
          <li style={{ display: 'inline-block', marginRight: '3%' }}><a href="/arbol-de-procesos" style={{ textDecoration: 'none', color: '#333' }}>Arbol de procesos</a></li>
          <li style={{ display: 'inline-block' }}><a href="/diagrama-de-estado" style={{ textDecoration: 'none', color: '#333' }}>Diagrama de Estado</a></li>
        </ul>
      </nav>
      <Routes>
          <Route path="/monitoreo-tiempo-real" element={<MonitoreoTiempoReal />} />
          <Route path="/monitoreo-historico" element={<MonitoreoHistorico />} />
          <Route path="/arbol-de-procesos" element={<ArbolDeProcesos />} />
          <Route path="/diagrama-de-estado" element={<DiagramaDeEstado />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
