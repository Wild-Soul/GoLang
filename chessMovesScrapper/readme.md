### HOW TO RUN

##### STEP 1
Run fetchData.go in ./fetchData
Purpose -> This will scrap the data from https://www.chessgames.com/chessecohelp.html and store it in csv format.
command -> go run fetchData/fetchData.go


##### STEP 2
Run loadData.go in ./loadData
Purpose -> Since this data is a static data we can just store it once in DB, it's a migration script to fulfill this purpose
Command -> go run loadData/loadData.go

##### STEP 3
Start the server.
Command -> go run main.go

###### NOTE : All the above mentioned commands need to be run from root directory.

