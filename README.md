# BackendAPI-Instagram
This is an API for the submission of appointy

## How to run the project
- put the files in a dirctory
- open comand prompt or teminal in that directry
- run ```go get .```
- it will download the dependencies and export all the required functions of Handler, DatabaseConnector and Encrypt
- to run the server run ```go run ./main.go``
- the server should be up and running

## Checking endpoints
GET all users
- to fetech all the users you can use ```curl localhost:8080/users``` or ```curl localhost:8080/users/```

GET user by ID
- to fetch user by a specific object id use ```curl localhost:8080/users/<id>```

Create User
- to create a POST request to the server you can use ```curl -X POST -d "{\"_id\": \"6161a7d45f16fd708f0bb812\", \"email\":\"TestEmail@gmail.com\",\"name\":\"InsertignUSer\",\"password\":\"myPassword\"}" http://localhost:8080/users```

To get all post by a user
- to fetch all post by a user use ```curl localhost:8080/posts/users/<user-id>```

To get all posts by id
- to fetch all posts use ```curl localhost:8080/posts/<post-id>```

To create a post 
- to create a post use ```curl -X POST -d "{\"_id\": \"6160a930cc5e3bf1bf21d5bd\", \"email\":\"test2@gmail.com\", \"caption\": \"MyCoolCap\", \"url\": \"image.comasdas\", \time\": \"2021-09-07T00:00:00+05:30\"}" http://localhost:8080/posts```

