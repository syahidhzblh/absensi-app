import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';

const API_URL = 'http://localhost:8080/api';

function Login({ onLogin }) {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setLoading(true);

    try {
      const response = await axios.post(`${API_URL}/login`, {
        email,
        password,
      });

      onLogin(response.data.token, response.data.user);
    } catch (err) {
      setError(err.response?.data?.error || 'Login gagal. Silakan coba lagi.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="auth-container">
      <div className="auth-card">
        <h2>Login Absensi</h2>
        
        {error && <div className="error">{error}</div>}
        
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label>Email</label>
            <input
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
              placeholder="email@example.com"
            />
          </div>

          <div className="form-group">
            <label>Password</label>
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
              placeholder="Password"
            />
          </div>

          <button type="submit" className="btn" disabled={loading}>
            {loading ? 'Loading...' : 'Login'}
          </button>
        </form>

        <div className="text-center">
          Belum punya akun? <Link to="/register" className="link">Daftar di sini</Link>
        </div>
      </div>
    </div>
  );
}

export default Login;
