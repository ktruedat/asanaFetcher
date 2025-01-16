# Running the application
In order to run the application, you should use this command to build it:

`go build -o main main.go`

And after the binary appeared, just launch it by typing:

`./main`

The application will fetch data from Asana REST API regarding projects and users of a given workspace.
The data will be saved in json format in the corresponding `projects.json` and `users.json` files in the directory where the 
application was run.