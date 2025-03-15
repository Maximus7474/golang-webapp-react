import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

function App() {
  const [message, setMessage] = useState<string | null>(null);
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);

  const handleLogin = async () => {
    const response = await fetch('/api/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        email: 'test@example.com',
        password: 'password123',
      }),
    });

    if (response.ok) {
      setIsAuthenticated(true);
      setMessage('Logged in successfully');
    } else {
      setMessage('Login failed');
    }
  };

  const handleLogout = async () => {
    const response = await fetch('/api/logout', {
      method: 'POST',
      credentials: 'include', // send cookies along with the request
    });

    if (response.ok) {
      setIsAuthenticated(false);
      setMessage('Logged out successfully');
    } else {
      setMessage('Logout failed');
    }
  };

  const handleFetchMessage = async () => {
    const response = await fetch('/api/protected', {
      method: 'GET',
      credentials: 'include', // include credentials (cookies) in request
    });

    if (response.ok) {
      const data = await response.json();
      setMessage(data.message);
    } else {
      setMessage('Failed to fetch protected content');
    }
  };

  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={handleFetchMessage}>
          {message ? 'hello!' : 'hello?'}
        </button>
        {message && <p>{message}</p>}
        {!isAuthenticated ? (
          <button onClick={handleLogin}>Log In</button>
        ) : (
          <button onClick={handleLogout}>Log Out</button>
        )}
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  );
}

export default App;
