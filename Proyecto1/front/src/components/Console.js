// Console.js
import React from 'react';

const Console = ({ output }) => {
  return (
    <div className="console-output">
      <pre>{output}</pre>
    </div>
  );
};

export default Console;
