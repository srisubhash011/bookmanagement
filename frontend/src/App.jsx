import { useState } from 'react'
import './App.css'

const API_BASE = '/api'

function App() {
  const [bookId, setBookId] = useState('')
  const [searchQuery, setSearchQuery] = useState('')
  const [book, setBook] = useState(null)
  const [searchResults, setSearchResults] = useState([])
  const [newBook, setNewBook] = useState({
    bookid: '',
    bookname: '',
    bookdescription: '',
    bookstatus: 1
  })
  const [message, setMessage] = useState('')
  const [error, setError] = useState('')

  const handleFetchBook = async () => {
    setError('')
    setMessage('')
    if (!bookId) {
      setError('Enter a book ID to fetch')
      return
    }

    try {
      const response = await fetch(`${API_BASE}/book/${bookId}`)
      if (!response.ok) {
        throw new Error(`Book not found (status ${response.status})`)
      }
      const data = await response.json()
      setBook(data)
    } catch (err) {
      setBook(null)
      setError(err.message)
    }
  }

  const handleSearch = async () => {
    setError('')
    setMessage('')
    if (!searchQuery) {
      setError('Enter a search query')
      return
    }

    try {
      const response = await fetch(`${API_BASE}/book/search/${encodeURIComponent(searchQuery)}`)
      if (!response.ok) {
        throw new Error('Search failed')
      }
      const results = await response.json()
      setSearchResults(results)
    } catch (err) {
      setSearchResults([])
      setError(err.message)
    }
  }

  const handleAddBook = async (event) => {
    event.preventDefault()
    setError('')
    setMessage('')

    if (!newBook.bookid || !newBook.bookname || !newBook.bookdescription) {
      setError('All add-book fields are required')
      return
    }

    try {
      const response = await fetch(`${API_BASE}/book/add`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          bookid: Number(newBook.bookid),
          bookname: newBook.bookname,
          bookdescription: newBook.bookdescription,
          bookstatus: Number(newBook.bookstatus)
        })
      })

      if (response.status === 201) {
        setMessage('Book added successfully')
        setNewBook({ bookid: '', bookname: '', bookdescription: '', bookstatus: 1 })
      } else {
        const text = await response.text()
        throw new Error(text || `Add failed (${response.status})`)
      }
    } catch (err) {
      setError(err.message)
    }
  }

  return (
    <div className="app-container">
      <header>
        <h1>Book Management</h1>
        <p>React frontend for the Go backend in <code>cmd/bookapi</code>.</p>
      </header>

      <section className="panel">
        <h2>Get Book by ID</h2>
        <div className="form-row">
          <input
            type="number"
            value={bookId}
            onChange={(e) => setBookId(e.target.value)}
            placeholder="Book ID"
          />
          <button onClick={handleFetchBook}>Fetch Book</button>
        </div>
        {book && (
          <div className="result-card">
            <h3>{book.bookname}</h3>
            <p>{book.bookdescription}</p>
            <small>ID: {book.bookid}</small>
          </div>
        )}
      </section>

      <section className="panel">
        <h2>Search Books</h2>
        <div className="form-row">
          <input
            type="text"
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            placeholder="Search term"
          />
          <button onClick={handleSearch}>Search</button>
        </div>
        {searchResults.length > 0 ? (
          <ul className="search-list">
            {searchResults.map((item, index) => (
              <li key={`${item.bookname}-${index}`}>{item.bookname}</li>
            ))}
          </ul>
        ) : (
          searchQuery && <p className="empty-state">No books found.</p>
        )}
      </section>

      <section className="panel">
        <h2>Add Book</h2>
        <form className="add-form" onSubmit={handleAddBook}>
          <input
            type="number"
            value={newBook.bookid}
            onChange={(e) => setNewBook({ ...newBook, bookid: e.target.value })}
            placeholder="Book ID"
          />
          <input
            type="text"
            value={newBook.bookname}
            onChange={(e) => setNewBook({ ...newBook, bookname: e.target.value })}
            placeholder="Book Name"
          />
          <textarea
            value={newBook.bookdescription}
            onChange={(e) => setNewBook({ ...newBook, bookdescription: e.target.value })}
            placeholder="Book Description"
          />
          <select
            value={newBook.bookstatus}
            onChange={(e) => setNewBook({ ...newBook, bookstatus: Number(e.target.value) })}
          >
            <option value={1}>Active</option>
            <option value={0}>Inactive</option>
          </select>
          <button type="submit">Add Book</button>
        </form>
      </section>

      {(message || error) && (
        <section className="feedback">
          {message && <p className="success">{message}</p>}
          {error && <p className="error">{error}</p>}
        </section>
      )}
    </div>
  )
}

export default App
