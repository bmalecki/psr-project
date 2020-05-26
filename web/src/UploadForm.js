import React, { useRef } from 'react';
import { url } from './environment'

function UploadForm() {

  const inputFile = useRef(null);
  const forbiddenWords = useRef(null);

  const handleSubmit = async (event) => {
    event.preventDefault();
    let formData = new FormData();

    const file = inputFile.current.files[0];

    formData.append("file", file)
    formData.append("forbiddenWords", forbiddenWords.current.value)

    try {
      const res = await fetch(`${url}/document`, {
        method: "POST",
        body: formData
      })

      const status = await res.status
      if (status !== 200) {
        console.log("Error");
      }

      const text = await res.text()
      console.log(text)
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
        <input type="submit" value="Submit" />
      </form>
    </div>
  );
}

export default UploadForm;
