#!/usr/bin/env python2

import os
import shlex
import subprocess
import sys

def run_command(command):
    split_command = shlex.split(command)
    process = subprocess.Popen(split_command, stdout=subprocess.PIPE)
    for line in iter(process.stdout.readline, ''):
        sys.stdout.write(line)

def run_parse_benchmarks():
    run_command("go test -bench=. -benchmem parse.go benchmark_test.go")


if __name__ == '__main__':
    run_parse_benchmarks()
