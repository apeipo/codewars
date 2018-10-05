<?php

function sum_strings($a, $b)
{
    $a = ltrim($a, "0");
    $b = ltrim($b, "0");
    $a = strrev($a);
    $b = strrev($b);
    if (strlen($a) > strlen($b)) {
        $tmp = $a;
        $a = $b;
        $b = $tmp;
    }

    $lena = strlen($a);
    $lenb = strlen($b);
    $fix = 0;
    $res = "";
    for ($i = 0; $i < $lenb; $i++) {
        $tmp = ($i < $lena) ? (intval($a[$i]) + intval($b[$i]) + $fix) : (intval($b[$i]) + $fix);
        $fix = $tmp > 9 ? 1 : 0;
        $res .= $tmp > 9 ? ($tmp - 10) : $tmp;
    }
    $res = $fix > 0 ? $res . $fix : $res;

    return strrev($res);
}

function sum_strings_v2($a, $b)
{
    return $a + $b;
}