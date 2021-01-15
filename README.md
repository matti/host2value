# host2ipv4

Like `dig +short` but works with mDNS (hostname.local) and returns a random IPv4 address (if multiple)

So yet another take on https://unix.stackexchange.com/questions/20784/how-can-i-resolve-a-hostname-to-an-ip-address-in-a-bash-script)

## examples

    $ host2ipv4 microsoft.com
    13.77.161.179
    $ host2ipv4 microsoft.com
    40.76.4.15

## install

    $ go get github.com/matti/host2ipv4
