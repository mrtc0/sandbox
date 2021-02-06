<?php
function output_chunk($chunk)
{
    echo sprintf("%x\r\n", strlen($chunk));
    echo $chunk . "\r\n";
}

// header("Content-type: application/octet-stream");
header("Content-Type: application/json");
header("Transfer-encoding: chunked");
ob_flush();
flush();

$i = 0;
while ( !connection_aborted() ) {
    $id        = $i;
    $message   = "hoge".$i;
    $createdAt = date("Y-m-d H:i:s");

    $json = json_encode( 
        array(
            "id" => $id, 
            "message" => $message, 
            "createdAt" => $createdAt
        ) 
    ); 

    output_chunk(
        $json . "\n"
    );
    ob_flush();
    flush();

    sleep(1);

    $i++;
}
echo "0\r\n\r\n";
