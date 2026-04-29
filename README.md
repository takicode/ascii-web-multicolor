# ASCII Art Multicolor 
This project is a Go-powered web application that generates ASCII art from input text and allows you to apply multiple colors to specific words within the output. It turns plain text into glowing, expressive ASCII art then paint it.

# Features
-- Convert text into ASCII art using different banners
-- Apply multiple colors to selected words
-- Supports multi-line input
-- Fast and lightweight (pure Go, no heavy frameworks)
-- Dockerized for easy deployment
-- Styled UI with animated ASCII background
-- How It Works (Simple Breakdown)

## The system has two main brains:

1. ASCII Generator
Reads banner files from /banner
Converts each character into ASCII patterns
Builds output line-by-line
2. Color Engine
Maps words → colors (like: "hello,world" → "red,blue")
Breaks input into segments
Applies HTML <span style="color:..."> around matching parts
Reconstructs the final colored ASCII art

So instead of coloring the whole text, it surgically paints only the words you choose.


## Getting Started
Run Locally
go run main.go

Open your browser:

http://localhost:8080
🐳 Run with Docker

## Build the image:

docker build -t ascii-multicolor .

## Run the container:

docker run -p 8080:8080 ascii-multicolor

Then visit:

http://localhost:8080


## Usage
-Enter your text
-Choose a banner:
-standard or shadow or thinkertoy
-click generate

## Error Handling
404 → Custom ASCII error page
400 → Invalid input
405 → Method not allowed
500 → Internal server error

## UI Highlights
Animated ASCII background ✨
Glowing neon theme 🌌
Responsive layout
Modal instructions panel

## Tech Stack
Go (Golang) – backend logic
HTML/CSS – frontend
Docker – containerization

## Notes
ASCII rendering depends on banner files

Each character is mapped using ASCII index math:
asciiValue := (char - 32) * 9 + 1
Each character is drawn using 8 rows

## Author
Taki Pelumi Emmanuel

### Future Improvements
1.Support hex colors (#ff0000)
2.Gradient coloring
3.Download ASCII as file
4.More font styles
5.API endpoint for programmatic use