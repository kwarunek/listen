listen
------

A tool server to listen on a given port and print to stdout.

**You should use netcat**, it offers much more than this tool. The listen is a quick'n'dirty tool developed to use on platforms that is easier to scp a staticlinked single binary.

Usage
-----

    Usage of listen:
        listen
    Options:
      -v, --verbose       output verbosely
      --quiet             be quiet, instead
      -a, --addr=0.0.0.0  listen address
      -p, --port=6666     port to listen tos
      -u, --udp           UDP instead of TCP
      -h, --help          Show usage message
      --version           Show version

Todo
----

- toggle printing of a time mark
- toggle printing of the client's ip
- ssl
