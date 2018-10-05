<?php

//class SumStringsTest extends TestCase {
//    public function testExamples() {
//        $this->assertEquals('579', sum_strings('123', '456'));
//    }
//}

require_once "sum_strings.php";

function doTest($a, $b, $expect) {
    $r = sum_strings($a, $b);
    assert($r === $expect, "expect $r to equal $expect");
}

//doTest("123", "456", "579");
//doTest("00103", "08567", "8670");

//print_r(PHP_INT_MAX); // 9223372036854775807
doTest("800", "9567", "10367");
doTest("0001", "0002", "3");
doTest("712569312664357328695151392", "8100824045303269669937", "712577413488402631964821329");


//var_dump(sum_strings("712569312664357328695151392", "8100824045303269669937"));
//var_dump(71256.9312664357328695151392 + 0.8100824045303269669937);