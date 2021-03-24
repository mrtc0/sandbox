<?php

session_set_cookie_params([
  'lifetime' => 0,
  'samesite' => 'lax'
]);
session_start();

echo $_COOKIE['PHPSESSID'];
?>

<form method="POST" action="http://user.shop.local/page.php">
  <input name="text" />
  <button type="submit">post</button>
</form>
