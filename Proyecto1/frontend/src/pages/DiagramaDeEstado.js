import React, { useState } from 'react';

const DiagramaDeEstado = () => {
  const [pid, setPid] = useState('');
  const [action, setAction] = useState('');

  const handleInputChange = (event) => {
    setPid(event.target.value);
  };

  const handleButtonClick = (actionType) => {
    setAction(actionType);
    fetch(`/api/estado/${actionType}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ pid: pid })
    })
    .then(response => {
      if (response.ok) {
        console.log(`Solicitud ${actionType} enviada correctamente`);
      } else {
        console.error(`Error al enviar solicitud ${actionType}`);
      }
    })
    .catch(error => {
      console.error(`Error de red: ${error}`);
    });
  };

  return (
    <div>
      <h1>Diagrama de estado</h1>
      <input type="text" placeholder="Ingrese el PID" value={pid} onChange={handleInputChange} />
      <button onClick={() => handleButtonClick('status')}>Status</button>
      <button onClick={() => handleButtonClick('start')}>Start</button>
      <button onClick={() => handleButtonClick('stop')}>Stop</button>
      <button onClick={() => handleButtonClick('resume')}>Resume</button>
      <button onClick={() => handleButtonClick('kill')}>Kill</button>
    </div>
  );
}

export default DiagramaDeEstado;
