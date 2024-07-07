import React, { useEffect, useState } from 'react';

const BookList = () => {
    const [books, setBooks] = useState([]);

    useEffect(() => {
        fetch('http://localhost:8080')
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok ' + response.statusText);
                }
                return response.json();
            })
            .then(data => {
                console.log('Data:', data);
                setBooks(data);
            })
            .catch(error => console.error('Error fetching data:', error));
    }, []);

    return (
        <div>
            <h1>Book List</h1>
            <ul>
                {books.map(book => (
                    <li key={book.ID}>
                        <h2>{book.Name}</h2>
                        <p><strong>Author:</strong> {book.Author}</p>
                        <p><strong>Publisher:</strong> {book.Publisher}</p>
                        <p><strong>Description:</strong> {book.Description}</p>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default BookList;
