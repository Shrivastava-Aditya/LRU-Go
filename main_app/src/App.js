import React, { useState } from 'react';

function App() {
  const [getKeyInput, setGetKeyInput] = useState('');
  const [setKeyInput, setSetKeyInput] = useState('');
  const [setValueInput, setSetValueInput] = useState('');
  const [expirationInput, setExpirationInput] = useState('');
  const [data, setData] = useState(null);
  const [error, setError] = useState('');
  const [successMessage, setSuccessMessage] = useState('');

  const fetchData = (key) => {
    if (!key) {
      setError('Please enter a key.');
      return;
    }

    // Fetch data without CORS
    fetch(`http://localhost:8080/get?key=${key}`, { mode: 'no-cors' })
      .then(response => {
        if (!response.ok) {
          throw new Error('');
        }
        return response.json();
      })
      .then(data => {
        setData(data);
        setError('');
      })
      .catch(error => {
        setError(error.message);
        setData(null);
      });
  };

  const setDataPair = () => {
    if (!setKeyInput || !setValueInput || !expirationInput) {
      setError('Please enter both key, value, and expiration.');
      return;
    }

    // Set data with expiration time
    fetch(`http://localhost:8080/set?key=${setKeyInput}&value=${setValueInput}&expiration=${expirationInput}`, { method: 'GET' })
      .then(() => {
        // After setting data, clear input fields
        setSetKeyInput('');
        setSetValueInput('');
        setExpirationInput('');
        setSuccessMessage(`Key '${setKeyInput}' set successfully with expiration ${expirationInput} seconds`);
        fetchData(setKeyInput); // Fetch data for the newly set key
      })
      .catch(error => {
        setError(error.message);
      });
  };

  return (
    <div>
      <h1>Data from Server</h1>
      
      {/* <div>
        <h2>Fetch Data</h2>
        <input type="text" placeholder="Enter Key" value={getKeyInput} onChange={e => setGetKeyInput(e.target.value)} />
        <button onClick={() => fetchData(getKeyInput)}>Fetch</button>
      </div> */}

      <div>
        <h2>Set Data</h2>
        <input type="text" placeholder="Key" value={setKeyInput} onChange={e => setSetKeyInput(e.target.value)} />
        <input type="text" placeholder="Value" value={setValueInput} onChange={e => setSetValueInput(e.target.value)} />
        <input type="text" placeholder="Expiration (seconds)" value={expirationInput} onChange={e => setExpirationInput(e.target.value)} />
        <button onClick={setDataPair}>Set</button>
      </div>

      {successMessage && <p style={{ color: 'green' }}>{successMessage}</p>}
      {/* {error && <p style={{ color: 'red' }}>{error}</p>} */}
      
      {data && (
        <div>
          <p>Key: {data.key}</p>
          <p>Value: {data.value}</p>
        </div>
      )}
    </div>
  );
}

export default App;
