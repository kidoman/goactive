# goactive

A simple and easy-to-use command line utility to find out the "current" bandwidth usage on your ACT TV account

It works by simply scraping the http://portal.acttv.in/index.php/myusage URL :)

## Prerequisites

* Go version 1.1 and above
* ACT TV Broadband connection

## Installation

    $ go install github.com/kid0m4n/goactive

## Usage

    $ goactive

## Output

    2013/11/24 03:32:48 Retrieving http://portal.acttv.in/index.php/myusage
    2013/11/24 03:32:49 Looks like you have used 243.41 GB

## Notes

I have tested this from my own home connection (50 Mbps) and it works just fine
