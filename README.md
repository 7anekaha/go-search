# Index Search in GO
The idea of this project is to create a simple index search in GO. The project is based on the idea of an index search, where we have a list of words and their respective indexes. The search is done by looking for the word in the index and returning the indexes of the word.


The data is extracted from ES dump wiki dataset. It is unzipped, processed and stored in a map. The map is then used to search for the indexes of the word.

Next steps:
- Add data processed to a database (MongoDB)
- Add Redis to cache the data and improve search performance (measure the performance before and after adding Redis)
- Add a REST API to search for queries



Idea of index search: https://www.addsearch.com/blog/search-index/ 

## How to run
```bash
go run main.go
```
