s an extendable golang toolkit to work with BRO IDS data.

# Features for version 0.1.0

* Expose a generic parsing api 
* Expose a generic api to populate a db using golangs mysql package
* Implement a full unit test suite for /parsers and /db 
* Implement integration tests to test end to end data flow
* Provide documentation on gobro, and use godoc

gobro uses semantic versioning 

# Project structure 

packages:

/parsers: logic for parsing BRO logs 

/db: db inserts and general utility functions 

/tests: for integration tests 

/config: includes toml file and schema for launching gobro

/doc: information about the type of statistics that are reported on

# Dependencies

* toml config file located in /config
* mysql schema (optional) also located in /config

# Additional links

https://www.bro.org/
http://gauss.ececs.uc.edu/Courses/c6055/pdf/bro_log_vars.pdf

# Language

### Bro logs

* separator: how bro logs are delimited
* fields: column names in bro logs
* entry: a single line in a bro log (the actual data)

# Examples

Please see the dev branch README.md for examples

# Note

This project is still in the planning stage. Goals, releases, project structure, architecture, etc... are all subject to change


