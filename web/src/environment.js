const url = process.env.REACT_APP_URL || 'http://localhost:8000';
const imagesUrl = process.env.REACT_APP_IMAGES_URL || 'http://localhost:8000';

if (process.env.REACT_APP_MOCK === 'true') {
  url = 'http://localhost:8000'
}


export { url, imagesUrl };
