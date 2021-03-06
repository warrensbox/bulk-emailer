[![Build Status](https://travis-ci.org/warrensbox/bulk-emailer.svg?branch=master)](https://travis-ci.org/warrensbox/bulk-emailer)
[![Go Report Card](https://goreportcard.com/badge/github.com/warrensbox/bulk-emailer)](https://goreportcard.com/report/github.com/warrensbox/bulk-emailer)
[![CircleCI](https://circleci.com/gh/warrensbox/bulk-emailer/tree/release.svg?style=shield&circle-token=9ae193b5ccf01cf421d909a55d76600efa1556f5)](https://circleci.com/gh/warrensbox/bulk-emailer/tree/release)


# Bulk Emailer

<img style="text-allign:center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/bulk-emailer/smallerlogo.png" alt="drawing" width="90" height="130"/>

The `bulk-emailer` command line tool lets you send messages to multiple recipients simultaneously. You need to provide is a cvs file with recipient emails and email content file. 

See installation guide here: [bulk-emailer installation](https://warrensbox.github.io/bulk-emailer)

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

## Prerequisite 
You must have AWS SES set up to use this tool.

See how to set up AWS SES: [Set Up AWS SES](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/quick-start.html)

<hr>

## How to use:
### Pass in the required parameters 
<img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/bulk-emailer/bulk-emailer-demo.gif" alt="drawing" width="400" />

1.  Type `bulk-emailer` on the command line with the following parameters: 
2.  Parameter `--from` (string) for sender's email.
3.  Parameter `--subject`(string) for the subject of the email.
4.  Parameter `--message` (file) path to the file of your email content.
5.  Parameter `--contacts` (file) path to the cvs file of your recipients.
6.  Hit **Enter** to send messages to recipient simutaneouly.

## Emails sent  
<img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/bulk-emailer/bulk-emailer-ouput.jpeg" alt="drawing" width="300" height="533"/>

<hr>

## Issues

Please open  *issues* here: [New Issue](https://github.com/warrensbox/bulk-emailer/issues)

