import React from 'react'
import ReactDOM from 'react-dom/client'
import RemoteApp from './app.tsx'
import './index.css'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <RemoteApp />
  </React.StrictMode>,
)
