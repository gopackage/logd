A "local" UDP logging server that receives JSON log records and pushes them to backend logging services using reliable transports.

The logd server should be run along side each service (similar to the design of nsqd).

In addition to UDP-to-TCP, the logging daemon does basic buffering of data and uses batch log updates for services that support them.
