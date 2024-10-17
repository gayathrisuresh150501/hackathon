import React, { useState } from 'react';
import axios from 'axios';

const Register = ({ setCurrentView }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleRegister = async (e) => {
    e.preventDefault();
    setError('');

    try {
      await axios.post('http://localhost:8080/api/register', {
        email,
        password,
      });
      alert('Registration successful!');
      setCurrentView('login'); // Redirect to login after successful registration
    } catch (err) {
      setError('User already exists');
    }
  };

  return (
    <div className="container mt-5">
      <h2>Register</h2>
      <form onSubmit={handleRegister}>
        <div className="form-group">
          <label>Email:</label>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label>Password:</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <button type="submit" className="btn btn-success">Register</button>
        {error && <p className="text-danger">{error}</p>}
      </form>
    </div>
  );
};

export default Register;
