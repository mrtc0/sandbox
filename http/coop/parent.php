<html lang="en">
<head>
  <meta charset="UTF-8">
  <title></title>
</head>
<body>
  <p id="test"></p>
  <iframe id="frame" src="https://shop.local:44301/index.php" style="display:none;"></iframe>

  <script>
window.onload = function() {
  var frame = document.getElementById("frame").contentWindow;
  var sab = new SharedArrayBuffer(1024);
  frame.postMessage(sab, "https://shop.local:44301/index.php");
};

window.addEventListener("message", function(event) {
  document.getElementById("test").textContent = `event.data is ${event.data}`;
}, false);
  </script>
</body>
</html>

