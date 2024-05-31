Exoplanet Microservice
This microservice is designed to manage a catalog of exoplanets, supporting functionalities such as adding, listing, retrieving, updating, deleting exoplanets, and estimating fuel costs for space voyages to these exoplanets.

Features
Add an Exoplanet: Add new exoplanets with properties such as name, description, distance from Earth, radius, mass (for terrestrial planets), and type (Gas Giant or Terrestrial).
List Exoplanets: Retrieve a list of all available exoplanets.
Get Exoplanet by ID: Retrieve information about a specific exoplanet by its unique ID.
Update Exoplanet: Update the details of an existing exoplanet.
Delete Exoplanet: Remove an exoplanet from the catalog.
Fuel Estimation: Retrieve an overall fuel cost estimation for a trip to any particular exoplanet based on crew capacity.
Installation
Prerequisites
Go 1.16+
Docker (for containerization)
Clone the Repository
bash
Copy code
git clone https://github.com/yourusername/exoplanet-microservice.git
cd exoplanet-microservice
Build and Run
Build the application
bash
Copy code
go build -o exoplanet-service
Run the application
bash
Copy code
./exoplanet-service
The server will start on http://localhost:8000.

Docker
Build the Docker image
bash
Copy code
docker build -t exoplanet-service .
Run the Docker container
bash
Copy code
docker run -p 8000:8000 exoplanet-service
API Endpoints
Add an Exoplanet
Endpoint: POST /exoplanets

Request Body:

json
Copy code
{
  "name": "Exoplanet Name",
  "description": "Description of the exoplanet",
  "distance": 150,
  "radius": 2.5,
  "mass": 5.0,
  "type": "Terrestrial"
}
Response:

201 Created with the created exoplanet data.
400 Bad Request if the input data is invalid.
List Exoplanets
Endpoint: GET /exoplanets

Response:

200 OK with a list of exoplanets.
Get Exoplanet by ID
Endpoint: GET /exoplanets/{id}

Response:

200 OK with the exoplanet data.
404 Not Found if the exoplanet does not exist.
Update Exoplanet
Endpoint: PUT /exoplanets/{id}

Request Body:

json
Copy code
{
  "name": "Updated Exoplanet Name",
  "description": "Updated description of the exoplanet",
  "distance": 200,
  "radius": 3.0,
  "mass": 6.0,
  "type": "Terrestrial"
}
Response:

200 OK with the updated exoplanet data.
400 Bad Request if the input data is invalid.
404 Not Found if the exoplanet does not exist.
Delete Exoplanet
Endpoint: DELETE /exoplanets/{id}

Response:

200 OK if the exoplanet is successfully deleted.
404 Not Found if the exoplanet does not exist.
Fuel Estimation
Endpoint: GET /exoplanets/{id}/fuel?crewCapacity={crewCapacity}

Response:

200 OK with the estimated fuel cost.
400 Bad Request if the input data is invalid.
404 Not Found if the exoplanet does not exist.
Running Tests
Unit tests are provided to ensure the reliability of the service.

bash
Copy code
go test -v
Code Structure
main.go: Entry point of the application.
handlers/: Contains the HTTP handlers for different endpoints.
models/: Contains the data models and in-memory store.
validation/: Contains input validation logic.
tests/: Contains the unit tests for the service.
Extensibility
The service is designed with extensibility in mind. New types of exoplanets can be added by:

Defining new gravity calculation logic in the models package.
Extending the validation logic in the validation package.
Updating the handlers to accommodate new exoplanet types.

This project is licensed under the MIT License.
