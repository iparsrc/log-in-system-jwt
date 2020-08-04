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
<code>{ "name": "Parsa Akbari", "email": "parsaakbari80808080@gmail.com", "password": "12345" }</code><br/>
Returns an application/json and saves the JWT as a cookie:<br/>
if succeed:<br/><code>{ "id": &lt;user-id&gt;, "name": &lt;user-name&gt;, "email": &lt;user-email&gt; }</code><br/>
if user already exists: <br/><code>{ json containing the error. }</code></p>
<h3>/users/profile</h3>
<p>Send a GET req.<br/>
Returns an application/json:<br/>
if jwt is valid: <br/><code>{ "id": &lt;user-id&gt;, "name": &lt;user-name&gt;, "email": &lt;user-email&gt; }</code><br/>
else: <br/><code>{ json containing the error. }</code></p>
  
  
  
