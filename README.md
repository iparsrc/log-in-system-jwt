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
<p>Send a get request.</p>
<p>Returns a text/plain: <b>pong</b></p>
<p>This is simply to check if server is runinng or not.</p>
<h3>/users/create</h3>
<p>Returns a application/json:</p>
<p>if succeed: <b>{ "name":<user-name>, "email":<user-email>, "id":<user-id> }</b></p>

