A "local" UDP logging server that receives JSON log records and pushes them to a Redis Stream.

Each UDP packet sent to should be json data with the following format with the "fields" field being arbitrary structured data associated with the log record:

```json
{
    "level":"info",
    "timestamp":"2018-07-08T22:44:46.578312921-07:00",
    "message":"A useful log message",
    "fields":{
        "frame":"abcdefg",
        "len":7,
        "session":"96befb055d045bd1f351210747e41c56"
    }
}
```

Test using netcat:

```bash
echo -n "{\"level\":\"info\",\"message\":\"hello world\"}" | nc -4u -w0 localhost 9044
```

The logging server does basic buffering of data and uses batch log updates to reduce load on Redis.

The logging server is configured via Redis and currently supports the following settings:

*
*
