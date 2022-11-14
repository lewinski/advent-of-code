<?php

$d = json_decode(file_get_contents('input.txt'), true);
echo "part1: " . part1($d) . "\n";
echo "part2: " . part2($d) . "\n";

function part1($a) {
    if (!is_array($a)) {
        return 0;
    }

    $x = 0;
    if (is_array($a)) {
        foreach ($a as $k => $v) {
            if (is_array($v)) {
                $x += part1($v);
            } elseif (is_numeric($v)) {
                $x += $v;
            }
        }
    }
    return $x;
}

function part2($a) {
    if (!is_array($a)) {
        return 0;
    }

    foreach ($a as $k => $v) {
        if (!is_numeric($k) && $v === 'red') {
            return 0;
        }
    }

    $x = 0;
    foreach ($a as $k => $v) {
        if (is_array($v)) {
            $x += part2($v);
        } elseif (is_numeric($v)) {
            $x += $v;
        }
    }
    return $x;
}