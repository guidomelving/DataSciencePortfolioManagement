# Data Science Portfolio Management 🚀

A CLI (*Command Line Interface*) application based on **Go (Golang)** to manage Data Science portfolio projects. This application was built to fulfill the **Algorithm and Programming Final Project** for the Bachelor of Data Science Program, Faculty of Informatics, Telkom University.

## 👥 Contributors

- **Guidomelvin**
- **Ryan Azrya Altriyananda**

## 📌 Project Description

This application is designed for Data Science beginners to record, manage, and track their completed projects. Users can store comprehensive project details, including project names, categories, technologies used, difficulty levels, and creation dates.

## ✨ Key Features

### 📂 Project Management (CRUD)
- Add new portfolio projects
- Edit existing project information
- Delete projects from the portfolio
- Duplicate project name validation
- Input validation for project data

### 📊 Portfolio Information
Each project contains:
- Project Name
- Category
  - Machine Learning
  - Data Visualization
  - Data Analysis
  - Natural Language Processing (NLP)
  - Computer Vision
- Technologies Used
- Project Description
- Difficulty Level (1–5)
- Creation Date

### 🔍 Searching Features
- Search projects by name using **Sequential Search**
- Supports partial keyword matching
- Search projects by category using **Recursive Binary Search**

### 📈 Sorting Features
- Sort projects by difficulty level using **Selection Sort**
- Sort projects by creation date using **Insertion Sort**
- Supports ascending and descending order

### 📉 Portfolio Dashboard
Displays:
- Total number of projects
- Project statistics by category
- Most and least populated categories
- List of acquired skills/technologies
- Easiest and hardest projects (Min-Max Analysis)

## 🧠 Implemented Algorithm Concepts

This application implements core Algorithm and Programming concepts:

### Data Structures
- Array of Struct
- Fixed-size data storage using `NMAX`

### Programming Concepts
- Modular Programming
- Functions and Procedures
- Pass-by-Value
- Pass-by-Reference (Pointers)
- Input Validation

### Searching Algorithms
- **Sequential Search**
  - Search projects by name
  - Supports partial text matching

- **Recursive Binary Search**
  - Search projects by category
  - Data is sorted before searching

### Sorting Algorithms
- **Selection Sort**
  - Sort projects by difficulty level

- **Insertion Sort**
  - Sort projects by creation date

### Additional Algorithms
- **Recursion**
  - Display all projects recursively
  - Recursive Binary Search implementation

- **Min-Max**
  - Find the easiest project
  - Find the hardest project

## 🛠 Technologies Used

- Go (Golang)
- Standard Library Packages:
  - `fmt`
  - `bufio`
  - `os`
  - `strconv`
  - `strings`

## 📋 How to Run

Make sure Go is installed on your machine.

```bash
go run main.go
```

Or compile the program first:

```bash
go build main.go
./main
```

## 📚 Learning Objectives

This project demonstrates the implementation of:

- Arrays and Structures
- Searching Algorithms
- Sorting Algorithms
- Recursion
- Data Validation
- Data Processing
- Portfolio Analytics
- Console-Based Application Development

## 🎓 Course Information

**Final Project**  
Algorithm and Programming
Faculty of Informatics  
Telkom University

---

⭐ Developed as a learning project for beginner Data Science portfolio management.
