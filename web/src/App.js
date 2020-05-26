import React from 'react';
import './App.css';
import UploadForm from './UploadForm';
import Documents from './Documents';
import { DocumentsContextProvider } from './context/DocumentsContextProvider';
import Uploading from './Uploading';

function App() {
  return (
    <div className="App">
      <h1>My Application</h1>
      <DocumentsContextProvider>
        <UploadForm />
        <Uploading />
        <Documents />
      </DocumentsContextProvider>
    </div>
  );
}

export default App;
