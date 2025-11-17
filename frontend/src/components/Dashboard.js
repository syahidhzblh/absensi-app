import React, { useState, useEffect } from 'react';
import axios from 'axios';

const API_URL = 'http://localhost:8080/api';

function Dashboard({ user, onLogout }) {
  const [todayStatus, setTodayStatus] = useState(null);
  const [attendances, setAttendances] = useState([]);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');
  const [loading, setLoading] = useState(false);

  const token = localStorage.getItem('token');

  const axiosConfig = {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  };

  useEffect(() => {
    fetchTodayStatus();
    fetchAttendances();
  }, []);

  const fetchTodayStatus = async () => {
    try {
      const response = await axios.get(`${API_URL}/today-status`, axiosConfig);
      setTodayStatus(response.data);
    } catch (err) {
      console.error('Error fetching status:', err);
    }
  };

  const fetchAttendances = async () => {
    try {
      const response = await axios.get(`${API_URL}/attendances`, axiosConfig);
      setAttendances(response.data || []);
    } catch (err) {
      console.error('Error fetching attendances:', err);
    }
  };

  const handleCheckIn = async () => {
    setError('');
    setSuccess('');
    setLoading(true);

    try {
      await axios.post(`${API_URL}/check-in`, {}, axiosConfig);
      setSuccess('Check-in berhasil!');
      fetchTodayStatus();
      fetchAttendances();
    } catch (err) {
      setError(err.response?.data?.error || 'Check-in gagal');
    } finally {
      setLoading(false);
    }
  };

  const handleCheckOut = async () => {
    setError('');
    setSuccess('');
    setLoading(true);

    try {
      await axios.post(`${API_URL}/check-out`, {}, axiosConfig);
      setSuccess('Check-out berhasil!');
      fetchTodayStatus();
      fetchAttendances();
    } catch (err) {
      setError(err.response?.data?.error || 'Check-out gagal');
    } finally {
      setLoading(false);
    }
  };

  const formatDate = (dateString) => {
    const date = new Date(dateString);
    return date.toLocaleDateString('id-ID', {
      day: 'numeric',
      month: 'long',
      year: 'numeric',
    });
  };

  const formatTime = (dateString) => {
    const date = new Date(dateString);
    return date.toLocaleTimeString('id-ID', {
      hour: '2-digit',
      minute: '2-digit',
    });
  };

  return (
    <div className="dashboard">
      <div className="header">
        <h1>Dashboard Absensi</h1>
        <div className="user-info">
          <span className="user-name">Halo, {user?.name}!</span>
          <button onClick={onLogout} className="btn-logout">
            Logout
          </button>
        </div>
      </div>

      <div className="container">
        <div className="card status-card">
          <h2>Absensi Hari Ini</h2>
          
          {error && <div className="error">{error}</div>}
          {success && <div className="success">{success}</div>}

          {todayStatus && (
            <div className="status-info">
              {todayStatus.checked_in ? (
                <>
                  <p>
                    ✅ Check-in: <span className="time">{formatTime(todayStatus.attendance.check_in)}</span>
                  </p>
                  {todayStatus.attendance.check_out && (
                    <p>
                      🚪 Check-out: <span className="time">{formatTime(todayStatus.attendance.check_out)}</span>
                    </p>
                  )}
                </>
              ) : (
                <p>Anda belum melakukan check-in hari ini</p>
              )}
            </div>
          )}

          <div>
            {!todayStatus?.checked_in && (
              <button
                onClick={handleCheckIn}
                disabled={loading}
                className="btn btn-check btn-check-in"
              >
                {loading ? 'Processing...' : '✓ Check In'}
              </button>
            )}

            {todayStatus?.checked_in && !todayStatus?.checked_out && (
              <button
                onClick={handleCheckOut}
                disabled={loading}
                className="btn btn-check btn-check-out"
              >
                {loading ? 'Processing...' : '🚪 Check Out'}
              </button>
            )}

            {todayStatus?.checked_out && (
              <p style={{ marginTop: '20px', color: '#28a745', fontWeight: 'bold' }}>
                ✅ Anda sudah menyelesaikan absensi hari ini
              </p>
            )}
          </div>
        </div>

        <div className="card">
          <h2>Riwayat Absensi</h2>
          
          {attendances.length > 0 ? (
            <table className="attendance-table">
              <thead>
                <tr>
                  <th>Tanggal</th>
                  <th>Check In</th>
                  <th>Check Out</th>
                  <th>Status</th>
                </tr>
              </thead>
              <tbody>
                {attendances.map((att) => (
                  <tr key={att.id}>
                    <td>{formatDate(att.date)}</td>
                    <td>{formatTime(att.check_in)}</td>
                    <td>{att.check_out ? formatTime(att.check_out) : '-'}</td>
                    <td>
                      <span className={`badge ${att.check_out ? 'badge-success' : 'badge-warning'}`}>
                        {att.check_out ? 'Selesai' : 'Belum Check Out'}
                      </span>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          ) : (
            <div className="empty-state">
              <p>Belum ada riwayat absensi</p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default Dashboard;
