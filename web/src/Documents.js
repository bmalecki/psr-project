import React, { useContext } from 'react';
import { DocumentsContext } from './context/DocumentsContextProvider';
import { url, imagesUrl } from './environment'

import "./Documents.css"

async function deleteImageItem(id) {
  const res = await fetch(`${url}/document/${id}`, {
    method: "DELETE",
  })

  const status = await res.status
  if (status !== 200) {
    console.log("Error");
  }
}

function Documents() {
  const { state, AsyncGetDocuments } = useContext(DocumentsContext);
  const { documents } = state

  const removeItem = async (event, id) => {
    event.target.disabled = true;
    try {
      await deleteImageItem(id)
      await AsyncGetDocuments()
    } catch (e) {
      console.log(e)
    }
  }

  return (
    <div className="Documents">
      <h2>Documents</h2>
      <table>
        <tbody>
          <tr>
            <th></th>
            <th>Id</th>
            <th>Name</th>
            <th>Image status</th>
            <th>Date</th>
            <th>Forbidden words</th>
            <th>Detected forbidden words</th>
          </tr>
          {
            documents != null && documents.map(v =>
              <tr key={v["Id"]}>
                <td>
                  <button onClick={(event) => removeItem(event, v["Id"])}>X</button>
                </td>
                <td>
                  <a href={`${imagesUrl}/${v["Id"]}`} target="_blank" rel="noreferrer noopener">{v["Id"]}</a>
                </td>
                <td>{v["FileName"]}</td>
                <td>{v["ImageStatus"]}</td>
                <td>{v["InsertionDate"]}</td>
                <td>
                  <ul>
                    {v["ForbiddenWords"] && v["ForbiddenWords"].map(words => <li key={words}>{words}</li>)}
                  </ul>
                </td>
                <td>
                  <ul>
                    {v["OccurredForbiddenWords"] && v["OccurredForbiddenWords"].map(words => <li key={words}>{words}</li>)}
                  </ul>
                </td>
              </tr>
            )
          }
        </tbody>
      </table>
    </div>
  );
}

export default Documents;
