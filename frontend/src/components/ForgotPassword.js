import React, { useState } from 'react';
import axios from 'axios';

const ForgotPassword = ({ setCurrentView }) => {
  const [email, setEmail] = useState('');
  const [message, setMessage] = useState('');
  const [error, setError] = useState('');

  const handleForgotPassword = async (e) => {
    e.preventDefault();
    setMessage('');
    setError('');

    try {
      await axios.post('http://localhost:8080/api/forgot-password', { email });
      setMessage('Password reset email sent!');
    } catch (err) {
      setError('Failed to send reset email');
    }
  };

  return (
    <div className="container mt-5">
      <h2>Forgot Password</h2>
      <form onSubmit={handleForgotPassword}>
        <div className="form-group">
          <label>Email:</label>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <button type="submit" className="btn btn-primary">Send Reset Email</button>
        {message && <p className="text-success">{message}</p>}
        {error && <p className="text-danger">{error}</p>}
      </form>
    </div>
  );
};

export default ForgotPassword;
