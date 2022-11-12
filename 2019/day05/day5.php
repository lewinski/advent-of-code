<?php

require_once __DIR__.'/../lib/intcode.php';

$input = file_get_contents(__DIR__.'/input.txt');

$computer = new \Intcode\Computer($input);
$computer->input(1);
$computer->run();
$output = $computer->output(0);
$output = end($output);
echo "day5 part1: $output\n";

$computer = new \Intcode\Computer($input);
$computer->input(5);
$computer->run();
$output = $computer->output(0);
$output = end($output);
echo "day5 part2: $output\n";
