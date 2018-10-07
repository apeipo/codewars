<?php

function divide_strings($a, $b)
{
    $a = ltrim($a, "0");
    $b = ltrim($b, "0");
    if ($a < $b) {
        return ["0", "$b"];
    }

    $lenb = strlen($b); //$lenb <= strlen(a)

    $q = ""; $r = "";
    while (true) {
        $asub = substr($a, 0, $lenb);
        $fix = "";
        for ($i = 0; $i < 9; $i++) {
            if (multiply($b, (string)$i) <= $b && multiply($b, (string)($i + 1)) > $b) {
                $q .= $i;
                $fix = "";
                break;
            }
        }
    }

    return [];
}

function sub_string(string $a, string $b): string
{
    $a = ltrim($a, "0");
    $b = ltrim($b, "0");
    $flag = false;
    if ($a < $b) {
        $flag = true;
        $tmp = $a;
        $a = $b;
        $b = $tmp;
    }
}