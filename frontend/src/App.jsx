import './App.css'

function App() {
  return (
    <div className="app">
      <header>
        <h1>DevBoard</h1>
        <p>Simple Project & Task Management</p>
      </header>

      <main>
        <section>
          <h2>DevOps Learning</h2>

          <div className="task">
            <span>Learn Git</span>
            <span>Done</span>
          </div>

          <div className="task">
            <span>Learn Docker</span>
            <span>In Progress</span>
          </div>

          <div className="task">
            <span>Build CI/CD Pipeline</span>
            <span>Todo</span>
          </div>
        </section>
      </main>
    </div>
  )
}

export default App