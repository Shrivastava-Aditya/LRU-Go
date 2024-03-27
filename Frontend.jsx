import React, { useState } from 'react';
import axios from 'axios';

const App = () => {
  const [key, setKey] = useState('');
  const [value, setValue] = useState('');
  const [expiration, setExpiration] = useState(5); // Default expiration of 5 seconds
  const [message, setMessage] = useState('');

  const handleSet = () => {
    axios.get(`http://localhost:8080/set?key=${key}&value=${value}&expiration=${expiration}`)
      .then(response => {
        setMessage(response.data.message);
      })
      .catch(error => {
        console.error('Error setting key-value pair:', error);
        setMessage('Error setting key-value pair');
      });
  };

  const handleGet = () => {
    axios.get(`http://localhost:8080/get?key=${key}`)
      .then(response => {
        setMessage(`Value for key '${key}' is: ${response.data.value}`);
      })
      .catch(error => {
        console.error('Error getting value for key:', error);
        setMessage(`Error getting value for key '${key}'`);
      });
  };

  return (
    <div>
      <h1>LRU Cache React App</h1>
      <div>
        <label>Key:</label>
        <input type="text" value={key} onChange={e => setKey(e.target.value)} />
      </div>
      <div>
        <label>Value:</label>
        <input type="text" value={value} onChange={e => setValue(e.target.value)} />
      </div>
      <div>
        <label>Expiration (seconds):</label>
        <input type="number" value={expiration} onChange={e => setExpiration(parseInt(e.target.value))} />
      </div>
      <div>
        <button onClick={handleSet}>Set</button>
        <button onClick={handleGet}>Get</button>
      </div>
      {message && <div>{message}</div>}
    </div>
  );
};

export default App;
