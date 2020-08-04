# simple-log-in-system-jwt<br/>
<p>In This simple JWT authorization system, you can create an in-memory user and get a valid JWT for one hour.</p>
<p>If you have the valid access token then you can get the user information otherwise you will get an error.</p>
<p>There are only three endpoints:</p>
<ul>
  <li>localhost:8080<b>/ping</b></li>
  <li>localhost:8080<b>/users/create</b></li>
  <li>localhost:8080<b>/users/profile</b></li>
</ul>
<h3>/ping</h3>
<p>Send a GET request.</p>
<p>Returns a text/plain: <b>pong</b></p>
<p>This is simply to check if server is runinng or not.</p>
<h3>/users/create</h3>
<p>Send a POST req with a json body. ex) { "name": "Parsa Akbari", "email": "parsaakbari80808080@gmail.com", "password", "12345" }</p>
<p>Returns a application/json and saves the JWT as a cookie:</p>
<p>if succeed: <b>{ "id":<user-id>, "name":<user-name>, "email":<user-email> }</b></p>
<p>if user already exists: <b>{ A json containing the error. }</b></p>
<h3>/users/profile</h3>
<p>Send a GET req.</p>
<p>Returns a application/json:</p>
<p>if jwt is valid: <b>{ "id":<user-id>, "name":<user-name>, "email":<user-email> }</b></p>
<p>else: <b>{ A json containing the error. }</b></p>
  
  
  
