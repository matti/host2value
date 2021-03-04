# host2value

Like `dig +short` but works with mDNS (hostname.local) and returns a random IPv4 address (if multiple)

Also `dig +short -t TXT` returns the value in "quotes", host2value does not.

So yet another take on https://unix.stackexchange.com/questions/20784/how-can-i-resolve-a-hostname-to-an-ip-address-in-a-bash-script)

## examples

    $ host2value microsoft.com
    13.77.161.179
    $ host2value microsoft.com
    40.76.4.15

    $ host2value -results 3 microsoft.com
    40.76.4.15
    104.215.148.63
    40.113.200.201

    $ host2value -recordType TXT microsoft.com
    google-site-verification=1TeK8q0OziFl4T1tF-QR65JkzHZ1rcdgNccDFp78iTk

    $ host2value -recordType TXT -results 3 microsoft.com
    8RPDXjBzBS9tu7Pbysu7qCACrwXPoDV8ZtLfthTnC4y9VJFLd84it5sQlEITgSLJ4KOIA8pBZxmyvPujuUvhOg==
    adobe-idp-site-verification=8aa35c528af5d72beb19b1bd3ed9b86d87ea7f24b2ba3c99ffcd00c27e9d809c
    google-site-verification=Zv1IvEEZg4N9wbEXpBSSyAiIjDyyB3S-fzfFClb7D1E

## install

    $ go get github.com/matti/host2value
