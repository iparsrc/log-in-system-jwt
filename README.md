# simple-log-in-system-jwt<br/>
<p>In This simple JWT authorization system, you can create an in-memory user and get a valid JWT for one hour.<br/>
If you have the valid access token then you can get the user information otherwise you will get an error.</p>
<p>There are only three endpoints:</p>
<ul>
  <li>localhost:8080<b>/ping</b></li>
  <li>localhost:8080<b>/users/create</b></li>
  <li>localhost:8080<b>/users/profile</b></li>
</ul>
<h3>/ping</h3>
<p>Send a GET request.<br/>
Returns a text/plain: <b>pong</b><br/>
This is simply to check if server is runinng or not.</p>
<h3>/users/create</h3>
<p>Send a POST req with a json body. ex)<br/>
<b>{<br/> "name": "Parsa Akbari",<br/> "email": "parsaakbari80808080@gmail.com",<br/> "password": "12345"<br/> }</b><br/>
Returns an application/json and saves the JWT as a cookie:<br/>
if succeed:<br/> <b>{<br/> "id": &lt;user-id&gt;,<br/> "name": &lt;user-name&gt;,<br/> "email": &lt;user-email&gt;<br/> }</b><br/>
if user already exists: <b>{ json containing the error. }</b></p>
<h3>/users/profile</h3>
<p>Send a GET req.<br/>
Returns an application/json:<br/>
if jwt is valid:<br/> <b>{<br/> "id": &lt;user-id&gt;,<br/> "name": &lt;user-name&gt;,<br/> "email": &lt;user-email&gt;<br/> }</b><br/>
else: <b>{ json containing the error. }</b></p>
  
  
  
