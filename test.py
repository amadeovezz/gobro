#!/usr/bin/env python2

import commands
import time
import argparse

# This command does the following:

# Runs the golang:1.7.1 image
# It clones your $PWD to /go/src/github.com/amadeovezz/gobro in the container
# It changes directory in the container to the packaage you are testing

unit_command = """ docker run --rm \
              -v "$PWD":/go/src/github.com/amadeovezz/gobro \
              -w /go/src/github.com/amadeovezz/gobro/db/ \
              --link mysql:mysql \
              golang:1.7.1 \
              go test -v """

def unit():
    print "Spinning up docker dependencies for testing:\n"
    print commands.getoutput("docker-compose -f db/docker-compose.yml up -d")
    print "------------------------------------------"

    print "running tests inside docker container:\n"
    print commands.getoutput(unit_command)
    print "-------------------------------------------"

    print "Bringing down docker dependencies:\n"
    print commands.getoutput("docker-compose -f db/docker-compose.yml down")

def bring_up_containers():
    print "Spinning up docker dependencies for testing:\n"
    print commands.getoutput("docker-compose -f db/docker-compose.yml up -d")

def bring_down_containers():
    print "Bringing down docker dependencies:\n"
    print commands.getoutput("docker-compose -f db/docker-compose.yml down")

def run_test():
    print "running tests inside docker container:\n"
    print commands.getoutput(unit_command)
    print "-------------------------------------------"


if __name__ == '__main__':

    parser = argparse.ArgumentParser(description='Test DB package')
    parser.add_argument('--unit',action="store_true", help='run unit tests')
    parser.add_argument('--test',action="store_true", help='run unit tests')
    parser.add_argument('--up',action="store_true", help='bring up containers')
    parser.add_argument('--down', action="store_true", help='bring down containers')

    args = parser.parse_args()

    if args.unit:
        unit()

    elif args.up:
        bring_up_containers()

    elif args.down:
        bring_down_containers()

    elif args.test:
        run_test()
