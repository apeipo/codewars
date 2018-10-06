<?php
function multiply(string $a, string $b): string
{
    $a = ltrim($a, "0");
    $b = ltrim($b, "0");
    $lena = strlen($a);
    $lenb = strlen($b);

    $sums = [];
    $maxpartLen = 0;
    for ($i = $lena - 1; $i >= 0; $i--) {
        $numa = intval($a[$i]);
        $part = "";
        $fix = 0;
        for ($j = $lenb - 1; $j >= 0; $j--) {
            $numb = intval($b[$j]);
            $tmp = $numa * $numb + $fix;
            $fix = floor($tmp / 10);
            $part .= $tmp % 10;
        }

        $fixzero = $lena - $i - 1;
        $part .= $fix > 0 ? $fix : "";
        $part = implode("", array_fill(0, $fixzero, "0")) . $part;
        $sums[] = $part;
        $maxpartLen = (strlen($part) > $maxpartLen) ? strlen($part) : $maxpartLen;
    }

    $fix = 0;
    $r = "";
    for ($i = 0; $i < $maxpartLen; $i++) {
        $tmp = $fix;
        foreach ($sums as $part) {
            if (isset($part[$i])) {
                $tmp += $part[$i];
            }
        }
        $fix = floor($tmp / 10);
        $r .= $tmp % 10;
    }
    $r .= $fix > 0 ? $fix : "";
    $r = empty($r) ? "0" : $r;
    return strrev($r);
}


/*
 * 234
 * 123
 * =======
 * 702
 *468
 234
 *
 * */
function multiply_v2(string $a, string $b): string
{
    $ar = str_split(strrev(ltrim($a, "0")));
    $br = str_split(strrev(ltrim($b, "0")));
    $r = [];
    foreach ($ar as $ai => $av) {
        foreach ($br as $bi => $bv) {
            @$r[$ai + $bi] += $av * $bv;
            if ($r[$ai + $bi] > 9) {
                @$r[$ai + $bi + 1] += floor($r[$ai + $bi] / 10);
                $r[$ai + $bi] = $r[$ai + $bi] % 10;
            }
        }
    }
    return strrev(implode("", $r));
}
