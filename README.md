etime converts between Unix epoch timestamps and human readable forms.

# Installation
    $ go get github.com/quincy/etime

# Usage
    $ etime --help
    usage: etime [-h] {timespec}

    -h, --help   prints this message
    timespec     is a string representation of a time to convert in one of the following formats:
            [Unix epoch time]
            Mon Jan _2 15:04:05 2006
            Mon Jan _2 15:04:05 MST 2006
            Mon Jan 02 15:04:05 -0700 2006
            02 Jan 06 15:04 MST
            02 Jan 06 15:04 -0700
            Monday, 02-Jan-06 15:04:05 MST
            Mon, 02 Jan 2006 15:04:05 MST
            Mon, 02 Jan 2006 15:04:05 -0700
            2006-01-02T15:04:05Z07:00
            2006-01-02T15:04:05.999999999Z07:00
            3:04PM
            Jan _2 15:04:05
            Jan _2 15:04:05.000000000
            Jan _2 15:04:05.000
            Jan _2 15:04:05.000000

