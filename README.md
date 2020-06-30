# Rest API - Go


This project was developed to create, update, delete and search for planets.
It was inspired by Star Wars. The database used was MongoDB. In the application environment,
it is necessary to raise the database with docker, with the command: 

````
make dbup
````

Next, to run the application, run the command:

````
make run
````
To run tests of application, execute:

````
go test -v ./router
````
##To test api calls:

GetAll (List all planets)

````
curl -X GET \
  http://localhost:3000/planets \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
````

GetByID (Get the planet with its id)

*ID is required*

````
curl -X GET \
  http://localhost:3000/planet/5efaad69b4726d34ae364263 \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
````

Create (Create a planet)

*Name, Terrain and Climate are required*

````
curl -X POST \
  http://localhost:3000/planet \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
	"name": "Tatooine",
	"terrain": "desert",
	"climate": "arid"
}'
````

Update (Update a planet)

*Id is required*

*Name, Terrain, Climate and movies_count are optional*

````
curl -X PUT \
  http://localhost:3000/planet/5efaaa65b4726d34ae364261 \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
	"ID": "5efaaa65b4726d34ae364261",
	"name": "Tatooine",
	"terrain": "desert",
	"climate": "frio",
    "movies_count": 1
}'
````

Delete (Delete a planet)

*Id is required*

````
curl -X DELETE \
  http://localhost:3000/planet/5efaaa65b4726d34ae364261 \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
````
