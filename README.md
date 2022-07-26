Basic webserver written in golang which operates on GET and POST request
POST takes data in form of json, with fields "Name" and "Email", and stores it to a file called test.txt
GET returns the contents of the file, or an error if file does not exist.
