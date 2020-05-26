import React, { useEffect, useState } from 'react';
import { url } from './environment'

import "./Documents.css"

function Documents() {
  const [documents, setDocuments] = useState([]);

  useEffect(() => {
    async function getDocuments() {
      try {
        const res = await fetch(`${url}/document`)

        const status = await res.status
        if (status !== 200) {
          console.log("Error");
        }

        const json = await res.json()
        setDocuments(json)
      }
      catch (e) {
        console.log("Error getting documents!");
      }
    }

    getDocuments()
  }, []);

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
                <td>{v["Name"]}</td>
                <td>{v["ImageStatus"]}</td>
                <td>{v["Timestamp"]}</td>
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
