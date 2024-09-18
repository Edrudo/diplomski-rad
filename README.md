Server that knows how to communicate with master thesis HTTP3 client.
Before starting the server you need to set enviroment variables and 
create two directories `./images` and `./json`. This can be done using the `env.sh` script.  

The server is expecting a JSON file with the structure that can be found in
`./internal/domain/models/geoshot.json`. It will save the json and image file in 
appropriate directories when all data parts arrive.

