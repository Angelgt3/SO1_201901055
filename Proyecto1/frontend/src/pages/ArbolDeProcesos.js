import React, { useState, useEffect } from 'react';
import { Tree } from 'react-d3-tree';

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

  return (
    <div style={{ width: '100%', height: '500px' }}>
      {arbolData && (
        <Tree
          data={construirArbol(arbolData)}
          orientation="vertical"
          translate={{ x: 100, y: 50 }}
          separation={{ siblings: 2, nonSiblings: 2 }}
        />
      )}
    </div>
  );
}

const construirArbol = (data) => {
  const root = {
    name: data.Name,
    attributes: {
      PID: data.PID
    },
    children: []
  };
  construirHijos(root, data.Hijos);
  return root;
};

const construirHijos = (padre, hijos) => {
  if (!hijos || hijos.length === 0) return;
  hijos.forEach(hijo => {
    const nodoHijo = {
      name: hijo.Name,
      attributes: {
        PID: hijo.PID
      },
      children: []
    };
    padre.children.push(nodoHijo);
    construirHijos(nodoHijo, hijo.Hijos);
  });
};

export default ArbolDeProcesos;
