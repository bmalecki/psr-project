import React, { useRef } from 'react';

function UploadForm() {

  const input = useRef(null);

  const handleSubmit = (event) => {
    alert('A name was submitted: ' + input.current.value);
    event.preventDefault();
  }

  return (
    <div className="UploadForm">
      <h1>upload form</h1>
      <form onSubmit={handleSubmit}>
        <label>
          Name:
          <input type="text" ref={input} />
        </label>
        <input type="submit" value="Submit" />
      </form>
    </div>
  );
}

export default UploadForm;
