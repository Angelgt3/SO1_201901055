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
        <header className="App-header">
          <div className="contenedor">
            <span> SO1-201901055</span>
            <span className="centro">MODULOS DE KERNEL</span>
          </div>
        </header>
        <nav className="menu">
          <ul>
            <li><a href="/monitoreo-tiempo-real">Monitoreo Tiempo Real</a></li>
            <li><a href="/monitoreo-historico">Monitoreo Historico</a></li>
            <li><a href="/arbol-de-procesos">Arbol de procesos</a></li>
            <li><a href="/diagrama-de-estado">Diagrama de Estado</a></li>
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
