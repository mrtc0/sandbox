<?php

require_once "vendor/autoload.php";

$dirty_html = "<math></p><style><!--</style><img src/onerror=alert(1)>";

$config = HTMLPurifier_HTML5Config::createDefault();
$config->set('HTML.Allowed', 'math');

$purifier = new HTMLPurifier($config);

$clean_html = $purifier->purify($dirty_html);

echo $clean_html;

?>
