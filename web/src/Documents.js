import React, { useContext } from 'react';
import { DocumentsContext } from './context/DocumentsContextProvider';

import "./Documents.css"

function Documents() {
  const { documents } = useContext(DocumentsContext);

  return (
    <div className="Documents">
      <h2>Documents</h2>
      <table>
        <tbody>
          <tr>
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
                <td>{v["Id"]}</td>
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
