#!/usr/bin/env python2

import argparse
import os
import shlex
import subprocess
import sys

gopath = os.environ['GOPATH']

integration_command = """docker run --rm \
                -e "GOPATH=/gopath" \
                -v """ + gopath + """:/gopath \
                -w /gopath/src/github.com/amadeovezz/gobro/db \
                --network=db_default \
                golang:1.7.1 \
                go test -v"""

def run_command(command):
    split_command = shlex.split(command)
    process = subprocess.Popen(split_command, stdout=subprocess.PIPE)
    for line in iter(process.stdout.readline, ''):
        sys.stdout.write(line)

def start_compose_environment():
    print "Spinning up mysql container for testing:\n"
    run_command("docker-compose -f docker-compose.yml up -d")

def stop_compose_environment():
    print "Taking down mysql container:\n"
    run_command("docker-compose -f docker-compose.yml down")

def run_integration_test():
    print "\nRunning database integration tests:\n"
    print "------------------------------------------"
    run_command(integration_command)
    print "------------------------------------------\n"


if __name__ == '__main__':

    parser = argparse.ArgumentParser(description='Test gobro')
    parser.add_argument('--up', action="store_true", help='bring up docker compose environment')
    parser.add_argument('--down', action="store_true", help='take down docker compose environment')
    parser.add_argument('--full', action="store_true", help='run test suite from scratch and clean up after')
    parser.add_argument('--test', action="store_true", help='run test suite assuming compose environment is up')

    args = parser.parse_args()

    if args.up:
        start_compose_environment()
    elif args.down:
        stop_compose_environment()
    elif args.full:
        start_compose_environment()
        run_integration_test()
        stop_compose_environment()
    elif args.test:
        run_integration_test()
