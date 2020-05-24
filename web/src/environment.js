let url = process.env.REACT_APP_URL || 'http://aaaaa:8000';

if (process.env.REACT_APP_MOCK === 'true') {
  url = 'http://localhost:8000'
}


export { url };
