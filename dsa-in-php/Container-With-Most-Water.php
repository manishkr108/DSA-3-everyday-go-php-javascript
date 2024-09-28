<?php
function maxArea($arr)
{
    $left = 0;
    $right = count($arr) - 1;
    $maxarea = 0;

    while ($left < $right) {
        $ht = min($arr[$left], $arr[$right]);

        $maxarea = max($maxarea, $ht * ($right - $left));

        if ($arr[$left] < $arr[$right]) {
            $left++;
        } else {
            $right--;
        }
    }

    return $maxarea;
}

$arr = [1, 8, 6, 2, 5, 4, 8, 3, 7];
echo maxArea($arr);
?>
