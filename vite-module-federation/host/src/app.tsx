import { useState, lazy } from 'react'
import './app.css'

const RemoteApp = lazy(() => import('remote/Remote'));

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <h1>This is the host</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
      </div>
      <RemoteApp />
    </>
  )
}

export default App
