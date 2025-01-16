# Running the application
In order to run the application, you should use this command to build it:

`go build -o main main.go`

And after the binary appeared, just launch it by typing:

`./main`

The application will fetch data from Asana REST API regarding projects and users of a given workspace.
The data will be saved in json format in the corresponding `projects.json` and `users.json` files in the directory where the 
application was run.

The application can be configured using the config.json file. There, the developer can specify the 
access token and workspace GID for Asana, as well as the API base URL. Additionally, the app can be configured
to have custom pull intervals `extractionRateString`: by setting this value to a duration in seconds, the app 
will fetch data until this duration expires, will wait for this duration, and will start pulling data again.