#!/usr/bin/env python2

import argparse
import os
import shlex
import subprocess
import sys

def run_command(command):
    split_command = shlex.split(command)
    process = subprocess.Popen(split_command, stdout=subprocess.PIPE)
    for line in iter(process.stdout.readline, ''):
        sys.stdout.write(line)

def run_tests():
    run_command("go test -v")

def run_benchmarks():
    run_command("go test -bench=. -benchmem parse.go benchmark_test.go")


if __name__ == '__main__':

    parser = argparse.ArgumentParser(description='Test parse package')
    parser.add_argument('--all', action="store_true", help='run unit tests and benchmarks')
    parser.add_argument('--test', action="store_true", help='run unit tests')
    parser.add_argument('--bench', action="store_true", help='run benchmarks')

    args = parser.parse_args()

    if args.all:
        run_tests()
        run_benchmarks()
    elif args.test:
        run_tests()
    elif args.bench:
        run_benchmarks()
