import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import Login from './components/Login';
import Register from './components/Register';
import Dashboard from './components/Dashboard';

function App() {
  const [token, setToken] = useState(localStorage.getItem('token'));
  const [user, setUser] = useState(JSON.parse(localStorage.getItem('user') || 'null'));

  const handleLogin = (authToken, userData) => {
    localStorage.setItem('token', authToken);
    localStorage.setItem('user', JSON.stringify(userData));
    setToken(authToken);
    setUser(userData);
  };

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    setToken(null);
    setUser(null);
  };

  return (
    <Router>
      <Routes>
        <Route
          path="/login"
          element={!token ? <Login onLogin={handleLogin} /> : <Navigate to="/dashboard" />}
        />
        <Route
          path="/register"
          element={!token ? <Register onLogin={handleLogin} /> : <Navigate to="/dashboard" />}
        />
        <Route
          path="/dashboard"
          element={token ? <Dashboard user={user} onLogout={handleLogout} /> : <Navigate to="/login" />}
        />
        <Route path="/" element={<Navigate to={token ? "/dashboard" : "/login"} />} />
      </Routes>
    </Router>
  );
}

export default App;
