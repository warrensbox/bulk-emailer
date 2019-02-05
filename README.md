[![Build Status](https://travis-ci.org/warrensbox/bulk-emailer.svg?branch=master)](https://travis-ci.org/warrensbox/bulk-emailer)
[![Go Report Card](https://goreportcard.com/badge/github.com/warrensbox/bulk-emailer)](https://goreportcard.com/report/github.com/warrensbox/bulk-emailer)
[![CircleCI](https://circleci.com/gh/warrensbox/bulk-emailer/tree/release.svg?style=shield&circle-token=9ae193b5ccf01cf421d909a55d76600efa1556f5)](https://circleci.com/gh/warrensbox/bulk-emailer/tree/release)


# Bulk Emailer

<img style="text-allign:center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/bulk-emailer/smallerlogo.png" alt="drawing" width="110" height="140"/>

The `bulk-emailer` command line tool lets you send messages to multiple recipients simultaneously. You need to provide is a cvs file with recipient emails and email content file. 

<hr>

## Installation

`bulk-emailer` is available for MacOS and Linux based operating systems.

### Homebrew

Installation for MacOS is the easiest with Homebrew. [If you do not have homebrew installed, click here](https://brew.sh/). 


```ruby
brew install warrensbox/tap/bulk-emailer
```

### Linux

Installation for Linux operation systems.

```sh
curl -L https://raw.githubusercontent.com/warrensbox/bulk-emailer/release/install.sh | bash
```

### Install from source

Alternatively, you can install the binary from the source [here](https://github.com/warrensbox/bulk-emailer/releases) 

<hr>

## How to use:
### Pass in the required parameters 
<img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/github-bulk-emailer/bulk-emailerdemo1.gif" alt="drawing" style="width: 480px;"/>

1.  Type `bulk-emailer` on the command line with the following parameters: 
2.  Parameter `--from` (string) for sender's email.
3.  Parameter `--subject`(string) for the subject of the email.
4.  Parameter `--message` (file) path to the file of your email content.
5.  Parameter `--contacts` (file) path to the cvs file of your recipients.
6.  Hit **Enter** to send messages to recipient simutaneouly.

Again, you must have AWS SES set up to use this tool.

<hr>

## Issues

Please open  *issues* here: [New Issue](https://github.com/warrensbox/bulk-emailer/issues)

<hr>

See how to *upgrade*, *uninstall*, *troubleshoot* here:
[Additional Info](additional)