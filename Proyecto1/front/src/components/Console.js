// Console.js
import React from 'react';

const Console = ({ output }) => {
  // Convertir el array de mensajes en un array de elementos <p> separados por salto de línea
  return (
    <div className="console-output">
      {output.map((message, index) => (
        <p key={index}>{message}</p>
      ))}
    </div>
  );
};

export default Console;
