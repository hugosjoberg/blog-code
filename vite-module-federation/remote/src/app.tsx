import { useState } from 'react'
import './app.css'

function RemoteApp() {
  const [count, setCount] = useState(0)

  return (
    <>
      <h1>This is a remote</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
      </div>
    </>
  )
}

export default RemoteApp
