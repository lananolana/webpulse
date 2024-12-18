import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import { Provider } from 'react-redux';
import reportWebVitals from './reportWebVitals';
import App from './App';
import { compose } from 'redux';
import { BrowserRouter } from 'react-router-dom';
import { store } from './services/store';

declare global {
  interface Window {
    __REDUX_DEVTOOLS_EXTENSION_COMPOSE__?: typeof compose;
  }
}

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLFormElement,
);

root.render(
  <React.StrictMode>
    <Provider store={store}>
      <BrowserRouter>
        <App />
      </BrowserRouter>
    </Provider>
  </React.StrictMode>,
);

reportWebVitals();
