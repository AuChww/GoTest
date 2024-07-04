package main

// import (
//     "gorm.io/gorm"
// )

type Book struct {
    ID          int64
    Name        string
    Author      string
  	Publisher      string
    Description string
}

var books = []Book{
	{ID: 1, Name: "To Kill a Mockingbird", Author: "Harper Lee", Publisher: "J.B. Lippincott & Co.", Description: "A novel about the serious issues of rape and racial inequality."},
	{ID: 2, Name: "1984", Author: "George Orwell", Publisher: "Secker & Warburg", Description: "A dystopian social science fiction novel and cautionary tale about the dangers of totalitarianism."},
	{ID: 3, Name: "Moby-Dick", Author: "Herman Melville", Publisher: "Harper & Brothers", Description: "A sailor's narrative of the obsessive quest of Ahab for revenge on Moby Dick."},
	{ID: 4, Name: "Pride and Prejudice", Author: "Jane Austen", Publisher: "T. Egerton", Description: "A romantic novel that charts the emotional development of the protagonist Elizabeth Bennet."},
	{ID: 5, Name: "The Great Gatsby", Author: "F. Scott Fitzgerald", Publisher: "Charles Scribner's Sons", Description: "A novel that critiques the disillusionment and moral decay of society during the 1920s."},
	{ID: 6, Name: "War and Peace", Author: "Leo Tolstoy", Publisher: "The Russian Messenger", Description: "A historical novel that intertwines the lives of private and public individuals during the time of the Napoleonic wars."},
	{ID: 7, Name: "The Catcher in the Rye", Author: "J.D. Salinger", Publisher: "Little, Brown and Company", Description: "A story about adolescent angst and alienation."},
	{ID: 8, Name: "Crime and Punishment", Author: "Fyodor Dostoevsky", Publisher: "The Russian Messenger", Description: "A novel about the mental anguish and moral dilemmas of an impoverished ex-student."},
	{ID: 9, Name: "The Hobbit", Author: "J.R.R. Tolkien", Publisher: "George Allen & Unwin", Description: "A fantasy novel that introduces the world of Middle-earth and the adventures of Bilbo Baggins."},
	{ID: 10, Name: "Brave New World", Author: "Aldous Huxley", Publisher: "Chatto & Windus", Description: "A dystopian novel that explores the implications of a world where technology and conditioning control society."},
}