# simple-log-in-system-jwt  
In this simple JWT authorization system, you can create an in-memory user and get a valid JWT for one hour.  
If you have the valid access token then you can get the user information otherwise you will get an error.  
There are only three endpoints:  
```
/ping  
/users/create  
/users/profile  
```
## /ping  
Send a GET request.  
Returns a text/plain:  
```
pong
```
This is simply to check if server is runinng or not.  
## /users/create  
Send a POST req with a json body. ex)  
```
{  
  "name": "Parsa Akbari",  
  "email": "parsaakbari80808080@gmail.com",  
  "password": "12345"  
}  
```
Returns an application/json and saves the JWT as a cookie:  
if succeed:  
```
{  
  "id": user-id,  
  "name": user-name;,  
  "email": user-email;  
}  
```
if user already exists:  
```
{  
  json containing the error.  
}  
```
## /users/profile  
Send a GET req.  
Returns an application/json:  
if jwt is valid:  
```
{  
  "id": user-id,  
  "name": user-name,  
  "email": user-email  
}  
```
else: 
```
{  
json containing the error.  
}    
```
  
  
