# golang-serverless-blueprint

This is a sample app that was built for the Brest [AWS User Group Meetup of 1st of June 2022](https://www.meetup.com/Brest-AWS-User-Group/events/286113260/?).  
The purpose is to show how to improve the observability of a serverless application by having good logging in place.  
In other words: `How to move from an hello-world to a prod ready serverless application?`.

![image](https://user-images.githubusercontent.com/1489214/171444326-f69d3fdf-9f1b-46ee-a0b8-0be603c5c33a.png)


## Requirements

* AWS CLI already configured with Administrator permission
* [Golang](https://golang.org)
* SAM CLI - [Install the SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)


## Steps for having great logs

1️⃣ Use a logging library: [see commit](https://github.com/jbleduigou/golang-serverless-blueprint/commit/d78aa092cf7e2a33acbb2735f4ed9839cb3c8f68)  
2️⃣ Configure logs as json: [see commit](https://github.com/jbleduigou/golang-serverless-blueprint/commit/21492736572b16f73ed96234f7de27fa5ad5eaed)   
3️⃣ Log aws request id: [see commit](https://github.com/jbleduigou/golang-serverless-blueprint/commit/b3fe5258bffb36f07fa5759dbbe5009ac8da99d7)  
4️⃣ Use structured logging: [see commit](https://github.com/jbleduigou/golang-serverless-blueprint/commit/d75b9518007f951d3bbf9c6bb34e54fcaefb52a9)  
5️⃣ Enable access logs for API GW: [see commit](https://github.com/jbleduigou/golang-serverless-blueprint/commit/6e96212fd48b5c1c1e29ad3dcc93455ec4976f1a)  
6️⃣ Optimize cost by:  
&nbsp;&nbsp;‣ Define retention period for log groups: [see commit](https://github.com/jbleduigou/golang-serverless-blueprint/commit/59782d32058dd53a547316f6f898d50e9ad6e5b9)  
&nbsp;&nbsp;‣ Use specific log level per env: [see commit](https://github.com/jbleduigou/golang-serverless-blueprint/commit/38261090a489fc23c0fb57bdf5dee88f1f7acbf1)  


