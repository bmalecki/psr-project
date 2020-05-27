import React, { useRef, useContext, useState, createRef, useEffect } from 'react';
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
  const [forbiddenWordsInput, setForbiddenWordsInput] = useState(1)
  const { state, dispatch } = useContext(DocumentsContext);
  const { uploading } = state;

  const inputFile = useRef(null);
  const [forbiddenWords, setForbiddenWords] = React.useState([]);

  useEffect(() => {
    setForbiddenWords(elRefs => (
      Array(forbiddenWordsInput).fill().map((_, i) => elRefs[i] || createRef())
    ));
  }, [forbiddenWordsInput]);


  const handleSubmit = async (event) => {
    event.preventDefault();

    let formData = new FormData();
    formData.append("file", inputFile.current.files[0])

    forbiddenWords.forEach((ref => {
      formData.append("forbiddenWords", ref.current.value)
    }))

    try {
      dispatch({ type: "UPLOADING_START" })
      const id = await postImageForm(formData)
      const item = await getImageItem(id)
      dispatch({ type: "ADD_ITEM", item })
    }
    catch (e) {
      console.log("Error submitting form!");
    } finally {
      dispatch({ type: "UPLOADING_END" })
    }
  }

  const addInput = (event) => {
    event.preventDefault();
    setForbiddenWordsInput(forbiddenWordsInput + 1)
  }

  const removeInput = (event) => {
    event.preventDefault();
    if (forbiddenWordsInput > 1) {
      setForbiddenWordsInput(forbiddenWordsInput - 1)
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

        {[...Array(forbiddenWordsInput)].map((_, i) =>
          <React.Fragment key={i}><br /><input type="text" id="forbiddenWords" ref={forbiddenWords[i]} required></input></React.Fragment>)}

        <button onClick={addInput}>+</button>
        <button onClick={removeInput} disabled={forbiddenWordsInput < 2}>-</button>
        <br />
        <input type="submit" value="Submit" disabled={uploading} />
      </form>
    </div>
  );
}

export default UploadForm;
