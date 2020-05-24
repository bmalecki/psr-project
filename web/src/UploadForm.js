import React, { useRef } from 'react';
import { url } from './environment'

function UploadForm() {

  const inputFile = useRef(null);

  const handleSubmit = (event) => {
    event.preventDefault();
    let formData = new FormData();

    const file = inputFile.current.files[0];

    formData.append("file", file)
    formData.append("words", "aaaaa")


    fetch(`${url}/document`, {
      mode: 'no-cors',
      method: "POST",
      body: formData
    }).then(function (res) {
      if (res.ok) {
        console.log("Perfect! ");
      } else {
        console.log("Oops! ");
      }
    }, function (e) {
      console.log("Error submitting form!");
    });
  }

  return (
    <div className="UploadForm">
      <h1>Upload document</h1>
      <form onSubmit={handleSubmit}>
        <label htmlFor="documentFile">Select a file:</label>
        <input type="file" id="documentFile" ref={inputFile} required></input>
        <br />
        <input type="submit" value="Submit" />
      </form>
    </div>
  );
}

export default UploadForm;
