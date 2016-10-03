# gobro

gobro is a golang reporting tool provides general statistics on network behaviour based on BRO IDS logs 

# Features for version 1.0.0

* Implement parsers for conn, dns, and ssh logs
* Expose a generic parsing api out of the /parsers package
* Report on a couple of general statistics (defined in doc/stats.md)
* Implement a full unit test suite for /parsers and /db 
* Write integration tests for end to end data flow
* Provide documentation on gobro, and use godoc
* Setup gobro in docker containers

gobro uses semantic versioning 

# Goals 

* gobro will be fast, it should not take longer than ~3s to generate a report 
* gobro will not save state and only report only on 24 hours worth of data
* gobro will report on statistics for a single node running BRO

# Future goals

* gobro will save state
(create a daemon that pumps logs into a datastore, create an API that exposes metrics, and
create a client that speaks to the api and outputs results)
* gobro will report on statistics for an entire network 
(some hardware must be configured in this case)

# High level data flow

gobro concurrently parses each type of log file, buffers x number of rows, and then sends them to 
a data store. gobro then runs a list of queries against the db, obtains the results, and 
outputs the data to stdout

# Project structure 

packages:

/parsers: logic for parsing each type of BRO log (conn,http,etc) 

/db: db specific inserts, queries used to generated reports

/logs: example BRO log files that are used for testing/benchmarking purposes

/stdout: formatting for outputting the results 

/tests: for integration tests 

/config: includes toml files used for setup and testing

/docker: docker-compose files used for testing

/doc: information about the type of statistics that are reported on

# Dependencies

* mysql

# Additional links

https://www.bro.org/
http://gauss.ececs.uc.edu/Courses/c6055/pdf/bro_log_vars.pdf

# Language

### Networking

* lateral connection: a connection under the same network 
* outbound connection: a connection that originates inside a network  
* inbound connection: a connection that originates outside the network 

### Bro logs

* separator: how bro logs are delimited
* fields: column names in bro logs
* entry: a single line in a bro log (the actual data)

# Motivation

* To create a tool that produces useful information about a node's network activity 
* To learn more about bro, networking, data pipelines and golang

# Note

This project is still in the planning stage. Goals, releases, project structure, architecture, etc... are all subject to change


