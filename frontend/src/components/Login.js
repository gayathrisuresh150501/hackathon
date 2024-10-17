import React, { useState } from 'react';
import axios from 'axios';

const Login = ({ setCurrentView }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleLogin = async (e) => {
    try {
      console.log(" DLOGGGGGGGGGG")
      const response = await axios.post('http://localhost:8080/login', {
        email,
        password,
      });
      console.log(response)
      localStorage.setItem('token', response.data.token);
      alert('Login successful!');
    } catch (err) {
      setError('Invalid email or password');
    }
  };

  return (
    <div className="container mt-5">
      <h2>Login</h2>
      <form onSubmit={handleLogin}>
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
        <button type="submit" className="btn btn-primary">Login</button>
        {error && <p className="text-danger">{error}</p>}
      </form>
    </div>
  );
};

export default Login;
