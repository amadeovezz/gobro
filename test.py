#!/usr/bin/env python2

import argparse
import os
import shlex
import subprocess
import sys
import db.test as db
import parse.test as parse

# This command does the following:
# Runs the golang:1.7.1 image
# It clones your $PWD to /go/src/github.com/amadeovezz/gobro in the container
# It changes directory in the container to the package you are testing

gopath = os.environ['GOPATH']

integration_command = """docker run --rm \
              -e "GOPATH=/gopath" \
              -v """ + gopath + """:/gopath \
              -v """ + gopath + """/src/github.com/amadeovezz/gobro/sample_logs:/go/sample_logs \
              -w /gopath/src/github.com/amadeovezz/gobro/tests \
              --network=tests_default \
              golang:1.7.1 \
              go test -v"""

integration_bench_command = """docker run --rm \
              -e "GOPATH=/gopath" \
              -v """ + gopath + """:/gopath \
              -w /gopath/src/github.com/amadeovezz/gobro/tests \
              --network=tests_default \
              golang:1.7.1 \
              go test -bench=. -benchmem integration_test.go"""

def run_command(command):
    split_command = shlex.split(command)
    process = subprocess.Popen(split_command, stdout=subprocess.PIPE)
    for line in iter(process.stdout.readline, ''):
        sys.stdout.write(line)

def start_compose_environment():
    print "Spinning up integration test containers:\n"
    run_command("docker-compose -f " + gopath + "/src/github.com/amadeovezz/gobro/tests/docker-compose.yml up -d")

def stop_compose_environment():
    print "Taking down integration test containers:\n"
    run_command("docker-compose -f " + gopath + "/src/github.com/amadeovezz/gobro/tests/docker-compose.yml down")

def run_integration_test():
    print "\nRunning integration test suite:\n"
    print "------------------------------------------"
    run_command(integration_command)
    print "------------------------------------------"

def run_integration_benchmarks():
    print "\nRunning integration benchmarks:\n"
    print "------------------------------------------"
    run_command(integration_bench_command)
    print "------------------------------------------"

def run_package_tests():
    # db package tests
    db.start_compose_environment()
    db.run_integration_test()
    db.stop_compose_environment()

    print "\n------------------------------------------\n"

    # parse package tests
    os.chdir(gopath + "/src/github.com/amadeovezz/gobro/parse")
    parse.run_tests()

def run_package_benchmarks():
    # parse benchmarks
    os.chdir(gopath + "/src/github.com/amadeovezz/gobro/parse")
    parse.run_benchmarks()


if __name__ == '__main__':

    parser = argparse.ArgumentParser(description='Test gobro')
    parser.add_argument('--up', action="store_true", help='bring up docker compose environment')
    parser.add_argument('--down', action="store_true", help='take down docker compose environment')
    parser.add_argument('--full', action="store_true", help='run package and integration tests and benchmarks from scratch and clean up after')
    parser.add_argument('--package', action="store_true", help='run individual package tests')
    parser.add_argument('--package-bench', action="store_true", help='run individual package benchmarks')
    parser.add_argument('--int', action="store_true", help='run integration tests assuming environment is up')
    parser.add_argument('--bench', action="store_true", help='run integration benchmarks assuming environment is up')

    args = parser.parse_args()

    if args.up:
        start_compose_environment()
    elif args.down:
        stop_compose_environment()
    elif args.full:
        run_package_tests()
        print "\n------------------------------------------\n"
        run_package_benchmarks()
        print "\n------------------------------------------\n"
        start_compose_environment()
        run_integration_test()
        run_integration_benchmarks()
        stop_compose_environment()
    elif args.package:
        run_package_tests()
    elif args.package_bench:
        run_package_benchmarks()
    elif args.int:
        run_integration_test()
    elif args.bench:
        run_integration_benchmarks()
