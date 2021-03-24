<?php

require_once "vendor/autoload.php";

$smarty = new Smarty();
$smarty->template_dir = "./smarty/template/";
$smarty->compile_dir = "./smarty/template_c/";

$smarty->assign("name", "{ 1+1 }");

$smarty->display("hello.tpl");
?>
