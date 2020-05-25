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
      <h1>Documents</h1>
      <table>
        <tbody>
          <tr>
            <th>Id</th>
            <th>Name</th>
            <th>Image status</th>
            <th>Forbidden words</th>
            <th>Detected forbidden words</th>
          </tr>
          {
            documents.map(v =>
              <tr>
                <td>{v["Id"]}</td>
                <td>{v["Name"]}</td>
                <td>{v["ImageStatus"]}</td>
                <td>
                  <ul>
                    {v["ForbiddenWords"].map(words => <li>{words}</li>)}
                  </ul>
                </td>
                <td>
                  <ul>
                  {v["OccurredForbiddenWords"] && v["OccurredForbiddenWords"].map(words => <li>{words}</li>)}

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
