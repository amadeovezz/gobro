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

integration_command = """ docker run --rm \
              -v "$PWD":/go/src/github.com/amadeovezz/gobro \
              -w /go/src/github.com/amadeovezz/gobro/tests/ \
              --link mysql:mysql \
              golang:1.7.1 \
              go test -v """


def db_unit():
    print "Spinning up docker dependencies for testing:\n"
    print commands.getoutput("docker-compose -f db/docker-compose.yml up -d")
    print "------------------------------------------"

    print "running tests inside docker container:\n"
    print commands.getoutput(unit_command)
    print "-------------------------------------------"

    print "Bringing down docker dependencies:\n"
    print commands.getoutput("docker-compose -f db/docker-compose.yml down")

def integration():
    print "Spinning up docker dependencies for testing:\n"
    print commands.getoutput("docker-compose -f tests/docker-compose.yml up -d")
    print "------------------------------------------"

    print "running tests inside docker container:\n"
    print commands.getoutput(integration_command)
    print "-------------------------------------------"

    print "Bringing down docker dependencies:\n"
    print commands.getoutput("docker-compose -f tests/docker-compose.yml down")


def run_db_unit_test():
    print "running tests inside docker container:\n"
    print commands.getoutput(unit_command)
    print "-------------------------------------------"

def run_int_test():
    print "running tests inside docker container:\n"
    print commands.getoutput(integration_command)
    print "-------------------------------------------"


def bring_up_db_containers():
    print "Spinning up db docker dependencies for testing (not daemonized):\n"
    print commands.getoutput("docker-compose -f db/docker-compose.yml up")

def bring_down_db_containers():
    print "Bringing down db docker dependencies:\n"
    print commands.getoutput("docker-compose -f db/docker-compose.yml down")


def bring_up_int_containers():
    print "Spinning up docker dependencies for testing (not daemonized):\n"
    print commands.getoutput("docker-compose -f tests/docker-compose.yml up")

def bring_down_int_containers():
    print "Spinning up docker dependencies for testing (not daemonized):\n"
    print commands.getoutput("docker-compose -f tests/docker-compose.yml down")


if __name__ == '__main__':

    parser = argparse.ArgumentParser(description='Test gobro')
    parser.add_argument('--unit',action="store_true", help='run db unit tests')
    parser.add_argument('--int', action="store_true", help='run integration tests')
    parser.add_argument('--testunit',action="store_true", help='run db unit tests with assumption contains are up')
    parser.add_argument('--testint',action="store_true", help='run integration tests with assumption contains are up')
    parser.add_argument('--updb',action="store_true", help='bring up db containers')
    parser.add_argument('--downdb', action="store_true", help='bring down db containers')
    parser.add_argument('--upint',action="store_true", help='bring up integration containers')
    parser.add_argument('--downint', action="store_true", help='bring down integration containers')


    args = parser.parse_args()

    if args.int:
        integration()

    elif args.unit:
        db_unit()

    elif args.testunit:
        run_db_unit_test()

    elif args.testint:
        run_int_test()

    elif args.updb:
        bring_up_db_containers()

    elif args.downdb:
        bring_down_db_containers()

    elif args.upint:
        bring_up_int_containers()

    elif args.downint:
        bring_down_int_containers()

