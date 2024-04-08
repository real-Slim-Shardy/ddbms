cmd/
contains package main files for all instances:
- load-balancer
- server
- client (TODO)

lib/ 
contains libraries for project
- manager
  contains functions for managing resources and work with config
- LB
  contains functions for Load Balancer
- serv
  contains functions for Server
- alarm
  contains functions for innormal situtions
- heartbeat
  contains functions for communication between instances while normal work (KeepAlive)
- scale
  functions for changing configuration of system 
