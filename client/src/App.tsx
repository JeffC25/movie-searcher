import { BrowserRouter, Routes, Route} from 'react-router-dom'
import Details from './scenes/Details'
import HomePage from './scenes/HomePage'
import Results from './scenes/Results'

function App() {
  return (
    <BrowserRouter>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/home" element={<HomePage />} />
          <Route path="/search" element={<Results />} />
          <Route path="/:category" element={<Results />} />
          <Route path="/details/:id" element={<Details />}/>
        </Routes>
    </BrowserRouter>
  )
}

export default App
