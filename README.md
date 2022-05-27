# api-go-gorilla
This API retrieves data from a slice inside the code and displays it in JSON format. This was made as an exercise to practice creating APIs with Mux framework by Gorilla.

## Endpoints
### /getAvenger?name=\<string\>
  Retrieves details of an Avenger when `name` parameter matches exactly.

### /addAvenger
  Takes the body of Request which is formatted in JSON according to the `Avenger` struct and stores it in the slice `db`.

### /deleteAvenger/{name}
  Deletes the Avenger from the slice with their Name the same as {name} parameter in the URL.
