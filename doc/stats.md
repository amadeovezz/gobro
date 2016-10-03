### This document contains a list of general statistics that would be useful to report on 

#### Questions for a single node (on 24 hours worth of data)

Lateral statistics 
=================

* All local IP's visited on the network 
	* Displays: IP, port, proto, service

* Logs needed: conn.log, known_hosts.log

Outbound statistics 
==================

* Top 5 domains visisted 
	* Displays: domain, IP, port, proto, service, number of visits

* Top 5 domains with largest payload transfered
	* Displays: domain, IP, port, proto, service, number of bytes transfered

* All failed ssh login attempts
	* Displays: IP, number of attempts 

* Logs needed: dns.log, conn.log, ssh.log

Inbound statistics
==================

* Top 5 domains with largest payload transfered
	* Displays: domain, IP, port, proto, service, number of bytes transfered

* All failed ssh login attempts
	* Displays: IP, number of attempts


* Logs needed: dns.log, conn.log, ssh.log


