# CRUD Operations Using GORM
A simple microservice written in Golang, that connects to a DB and performs CRUD operations.

## Requirements
- Golang
- Docker

## API Endpoints
- `/user` -POST request to add a new Row in the DB
- `/user/<id>` -GET request for the specified id
- `/user` -GET request to display all rows in the DB
- `/user/<id>` -PATCH request to update the row specified by id
- `/user/<id>` -DELETE request to delete the row specified by the id 

## Getting Started 
1. Open a terminal or command prompt and navigate to the project directory.

2. Run the following command to start the container and execute the program in the scripts directory:

  	 - chmod +x run.sh
   followed by:
	 - ./run.sh
3. Postman Collection for the POST request in docs folder. 
4. Press CNTRL + C in the terminal to exit. 
