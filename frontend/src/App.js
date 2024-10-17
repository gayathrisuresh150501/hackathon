import React, { useState } from 'react';
import Login from './components/Login';
import Register from './components/Register';
import ForgotPassword from './components/ForgotPassword';
import './App.css';

function App() {
  const [currentView, setCurrentView] = useState('login');

  const renderView = () => {
    console.log(currentView)
    switch (currentView) {
      case 'login':
        return <Login setCurrentView={setCurrentView} />;
      case 'register':
        return <Register setCurrentView={setCurrentView} />;
      case 'forgot':
        return <ForgotPassword setCurrentView={setCurrentView} />;
      default:
        return <Login setCurrentView={setCurrentView} />;
    }

  };

  return (
    <div className="App">
      <div className="container">
        <h1>CIAM Login System</h1>
        {renderView()}
        <div className="text-center mt-3">
          {currentView !== 'login' && (
            <button onClick={() => {setCurrentView('login');}}>Back to Login</button>
          )}
          {currentView === 'login' && (
            <button onClick={() => setCurrentView('register')}>Register</button>
          )}
          {currentView === 'login' && (
            <button onClick={() => setCurrentView('forgot')}>Forgot Password?</button>
          )}
        </div>
      </div>
    </div>
  );
}

export default App;
