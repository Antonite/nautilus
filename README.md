## Nautilus mini

### Build
go build run/api

### Cmd Args
`-f` [string][optional] Path to input file. Default is `res/ship.csv`<br/>

Example: <br/>
./api -f res/ship.csv

### Solution Nuances
- Designed to generically accept different CSV files with different column count/order
- Easily scalable to compute generic sums of provided data columns (simply add more dataFields)

### Potential TODOs
- Move data to database obviously 
- API should accept more params, such as ship ID
- Expand generic field approach to handle different data types
- Add config file for things like port number