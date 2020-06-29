## REST Implementation

I used the Echo framework for the REST implementation. For reference to the framework, please go the the following site: 
[ECHO](https://echo.labstack.com/)

I have also add a Makefile that will clean the existing bin, vendor and Gopkg.lock file and the build should dep ensure with the updated Echo dependency

Run the following command to run the make file in terminal:
``` bash
make
```

If you run into any issue with the following command, you can also run each command individually, for example:
``` bash
make clean
```
To run the http server:
Run ``` bash go run main.go``` from the `rest-service` directory.

## Curl Request:

To get all people in the system:
``` bash
curl -X GET http://localhost:1323/people
```
To get the requested person with an id given:
``` bash
curl -X GET http://localhost:1323/people/df12ce76-767b-4bf0-bccb-816745df9e70
```

To get the requested person base on the first and last name given:
``` bash
curl -X GET 'http://localhost:1323/people?first_name=John&last_name=Doe'
```

To get the requested person base on the phone number given:
``` bash
curl -X  GET 'http://localhost:1323/people?phone_number=+1%20(800)%20555-1313'
```
