<?php

require_once "divide_strings.php";

function doTest($a, $b, $expect)
{
    $r = implode(",", divide_strings($a, $b));
    $e = implode(",", $expect);

    assert($r == $e, "expect $r to equal $e");
}

doTest("0", "5", ["0", "0"]);
doTest("1000", "10", ["100", "0"]);
doTest("600001", "100", ["6000", "1"]);