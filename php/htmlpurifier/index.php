<?php

require_once "vendor/autoload.php";

$dirty_html = "<math></p><style><!--</style><img src/onerror=alert(1)>";

$config = HTMLPurifier_Config::createDefault();
$config->set('HTML.Allowed', 'svg');

$purifier = new HTMLPurifier($config);
$clean_html = $purifier->purify($dirty_html);

echo $clean_html;

?>
