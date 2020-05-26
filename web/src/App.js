import React from 'react';
import './App.css';
import UploadForm from './UploadForm';
import Documents from './Documents';

function App() {
  return (
    <div className="App">
      <h1>My Application</h1>
      <UploadForm />
      <Documents />
    </div>
  );
}

export default App;
