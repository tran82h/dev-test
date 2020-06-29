## REST Implementation

I used the Echo framework for the REST implementation. For reference to the framework, please go the the following site: 
[ECHO](https://echo.labstack.com/)

I have also add a Makefile that will clean the existing bin, vendor and Gopkg.lock file and the build should dep ensure with the updated Echo dependency

Run the following command to run the make file in terminal:
`make`

If you run into any issue with the following command, you can also run each command individually, for example:
`make clean`

To run the http server:
Run `go run main.go` from the `rest-service` directory.
# dev-test
