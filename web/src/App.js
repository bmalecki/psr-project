import React from 'react';
import './App.css';
import UploadForm from './UploadForm';
import Documents from './Documents';
import { DocumentsContextProvider } from './context/DocumentsContextProvider';

function App() {
  return (
    <div className="App">
      <h1>My Application</h1>
      <DocumentsContextProvider>
        <UploadForm />
        <Documents />
      </DocumentsContextProvider>
    </div>
  );
}

export default App;
