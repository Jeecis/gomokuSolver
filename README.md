# Gomoku solver
My gomoku solver implemented using Minimax algorithm
### Build

To run on ubuntu build the binary file
```
cd game
go build -o ../gomoku
cd ..
```

To run on Windows build the binary file
```
cd game
go build -o ../gomoku.exe
cd ..
```
### Usage

```./gomoku -url "http://37.27.208.205:55555" -id "231RDB342" -gm "api"```
Command Line Flags
* `-gm`: Game mode (required)
    * `api`: Headless mode for server deployment
    * `apiVisualize`: API mode with visualization
    * `pvb`: Player vs Bot (AI starts)
    * `pvh`: Player vs Human (Human starts)
* `-url`: API server URL (required for API modes)
* `-id`: Player ID (required for API modes)

*To run with visualization you must uncomment the code regarding the visualization package in the main.go file (works only on windows)*

