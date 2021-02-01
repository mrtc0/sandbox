<?php

header("Content-Security-Policy: trusted-types default dompurify; require-trusted-types-for 'script'");
?>

<html>
  <head>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/dompurify/2.2.6/purify.min.js" integrity="sha512-rXAHWSMciPq2KsOxTvUeYNBb45apbcEXUVSIexVPOBnKfD/xo99uUe5M2OOsC49hGdUrkRLYsATkQQHMzUo/ew==" crossorigin="anonymous"></script>
  </head>
  <body>
    <script>
    const policy = trustedTypes.createPolicy('default', {
      createHTML: (untsutedValue) => {
        return DOMPurify.sanitize(untsutedValue)
      }
    });
    
    // http://localhost/#'%22%3E%3Csvg/onload=alert(1)%3E
    const rawHTML = decodeURIComponent(location.hash.substring(1));
    document.body.innerHTML = rawHTML;
    // document.body.innerHTML = policy.createHTML(rawHTML);
    </script>
  </body>
</html>
