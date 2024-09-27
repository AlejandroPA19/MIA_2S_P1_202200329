// FileInput.js
import React, { useState } from 'react';

const FileInput = ({ setConsoleOutput }) => {
  const [text, setText] = useState('');
  const [file, setFile] = useState(null);

  const handleTextChange = (e) => {
    setText(e.target.value);
  };

  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
  };

  const handleFileLoad = () => {
    if (file && file.name.endsWith('.smia')) {
      const reader = new FileReader();
      reader.onload = (e) => {
        setText(e.target.result);
      };
      reader.readAsText(file);
    } else {
      alert('Por favor selecciona un archivo con extensión .smia');
    }
  };

  const handleSubmit = () => {
    fetch('http://localhost:8080/api/analyze', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ text }),
    })
      .then((response) => response.json())
      .then((data) => {
        setConsoleOutput(data.messages);  // Actualizamos la consola con los mensajes devueltos
      })
      .catch((error) => {
        console.error('Error al hacer la petición:', error);
      });
  };

  return (
    <div className="file-input-container">
      <textarea
        value={text}
        onChange={handleTextChange}
        placeholder="Ingresa texto aquí"
        rows="18"
      />
      <br />
      <input className="input-file" type="file" onChange={handleFileChange} />
      <br />
      <button onClick={handleFileLoad}>Cargar Archivo</button>
      <button onClick={handleSubmit}>Ejecutar</button>
    </div>
  );
};

export default FileInput;
