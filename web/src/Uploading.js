import React, { useContext } from 'react';
import { DocumentsContext } from './context/DocumentsContextProvider';

import "./Documents.css"

function Uploading() {
  const { state } = useContext(DocumentsContext);
  const { uploading } = state

  return (
    <>
      {uploading ? <h1>Uploading...</h1> : null}
    </>
  );
}

export default Uploading;
