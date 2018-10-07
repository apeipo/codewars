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
    $ar = str_split(strrev(ltrim($a, "0")));
    $br = str_split(strrev(ltrim($b, "0")));
    $len = max(count($ar), count($br));
    $r = array_fill(0, $len, 0);
    for($i = 0; $i < $len; $i++) {
        $r[$i] += (isset($ar[$i]) ? $ar[$i] : 0) + (isset($br[$i]) ? $br[$i] : 0);
        if($r[$i] > 9) {
            $r[$i+1] = 1;
            $r[$i] = $r[$i] - 10;
        }
    }
    return strrev(ltrim(implode($r, ""), "0"));
}