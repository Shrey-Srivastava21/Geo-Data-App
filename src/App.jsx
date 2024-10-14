import React, { useState } from 'react';
import Map from './component/Map'; // Import the Map component
import FileUpload from './component/FileUpload'; // Import the FileUpload component

// Login Component
const Login = ({ onLogin }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleLogin = (e) => {
    e.preventDefault();

    // Simple validation: Check if both fields are non-empty
    if (username.trim() !== '' && password.trim() !== '') {
      onLogin(true); // Successfully logged in
    } else {
      alert('Please enter valid credentials');
    }
  };

  return (
    <div className="d-flex justify-content-center align-items-center vh-100" style={{ backgroundColor: '#f5f5f5' }}>
      <div className="card p-4 shadow-lg" style={{ width: '400px', borderRadius: '10px' }}>
        <h3 className="text-center mb-4">Login</h3>
        <form onSubmit={handleLogin}>
          <div className="form-group mb-3">
            <label htmlFor="username" className="form-label">Username</label>
            <input
              type="text"
              className="form-control"
              id="username"
              placeholder="Enter username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
            />
          </div>
          <div className="form-group mb-3">
            <label htmlFor="password" className="form-label">Password</label>
            <input
              type="password"
              className="form-control"
              id="password"
              placeholder="Enter password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>
          <button type="submit" className="btn btn-primary w-100 mt-3">
            Login
          </button>
        </form>
      </div>
    </div>
  );
};

// Main App Component
const App = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [geojsonData, setGeojsonData] = useState(null); // State for GeoJSON data

  const handleLogin = (status) => {
    console.log('Login status:', status); // Debugging line
    setIsLoggedIn(status); // Update login status
  };

  const handleFileUpload = (fileContent) => {
    setGeojsonData(fileContent); // Update GeoJSON data from file upload
  };

  if (!isLoggedIn) {
    // If not logged in, show the login page
    return <Login onLogin={handleLogin} />;
  }

  return (
    <div>
      <div style={{ textAlign: 'center' }}>
        <p></p>
        <h2>ꜰᴜʟʟ ꜱᴛᴀᴄᴋ ᴡᴇʙ ᴀᴘᴘʟɪᴄᴀᴛɪᴏɴ | ᴍᴀɴᴀɢᴇ ᴀɴᴅ ᴠɪꜱᴜᴀʟɪᴢᴇ ɢᴇᴏꜱᴘᴀᴛɪᴀʟ ᴅᴀᴛᴀ</h2>
        <p></p>
      </div>
      <FileUpload onFileUpload={handleFileUpload} /> {/* File upload component */}
      <Map geojsonData={geojsonData} /> {/* Map component */}
    </div>
  );
};

export default App;
