import React, { useRef, useContext } from 'react';
import { url } from './environment'
import { DocumentsContext } from './context/DocumentsContextProvider';

async function postImageForm(formData) {
  const res = await fetch(`${url}/document`, {
    method: "POST",
    body: formData
  })

  const status = await res.status
  if (status !== 200) {
    console.log("Error");
  }

  return res.text()
}

async function getImageItem(id) {
  const res = await fetch(`${url}/document/${id}`)

  const status = await res.status
  if (status !== 200) {
    console.log("Error");
  }

  return res.json()
}

function UploadForm() {
  const { state, dispatch } = useContext(DocumentsContext);
  const { uploading } = state;

  const inputFile = useRef(null);
  const forbiddenWords = useRef(null);

  const handleSubmit = async (event) => {
    event.preventDefault();

    let formData = new FormData();
    formData.append("file", inputFile.current.files[0])
    formData.append("forbiddenWords", forbiddenWords.current.value)

    try {
      dispatch({ type: "UPLOADING_START" })
      const id = await postImageForm(formData)
      const item = await getImageItem(id)
      dispatch({ type: "ADD_ITEM", item })
      dispatch({ type: "UPLOADING_END" })
    }
    catch (e) {
      console.log("Error submitting form!");
    }
  }

  return (
    <div className="UploadForm">
      <h2>Upload document</h2>
      <form onSubmit={handleSubmit}>
        <label htmlFor="documentFile">Select a file: </label>
        <input type="file" id="documentFile" ref={inputFile} required></input>
        <br />
        <label htmlFor="forbiddenWords">Forbidden Words: </label>
        <input type="text" id="forbiddenWords" ref={forbiddenWords} required></input>
        <br />
        <input type="submit" value="Submit" disabled={uploading} />
      </form>
    </div>
  );
}

export default UploadForm;
