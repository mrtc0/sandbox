<?php
  header("Cross-Origin-Opener-Policy: same-origin");
  header("Cross-Origin-Embedder-Policy: require-corp");
?>
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title></title>
</head>
<body>
  <script>
  window.addEventListener("message", function(event) {
    console.log(event)
    event.source.postMessage("test from cross-origin", event.origin);
  })
  </script>
</body>
</html>
