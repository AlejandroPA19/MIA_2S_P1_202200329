import React, { useState } from 'react';
import FileInput from './components/FileInput';
import Console from './components/Console';
import './App.css'; // Asegúrate de importar el CSS si no lo tienes ya

function App() {
  const [consoleOutput, setConsoleOutput] = useState('Consola de salida...');
  return (
    <div className="App">
      <h1>Mi Proyecto React</h1>
      <div className="Container">
      <FileInput setConsoleOutput={setConsoleOutput} />
      <br />
      <br />
      <Console output={consoleOutput} />
      </div>
    </div>
  );
}

export default App;