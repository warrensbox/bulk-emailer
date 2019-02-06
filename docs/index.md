# Bulk Emailer

The `bulk-emailer` command line tool lets you send messages to multiple recipients simultaneously. You need to provide is a cvs file with recipient emails and email content file. 

<hr>

## Installation

`bulk-emailer` is available for MacOS and Linux based operating systems.

### Homebrew

Installation for MacOS is the easiest with Homebrew. [If you do not have homebrew installed, click here](https://brew.sh/){:target="_blank"}. 


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
<img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/bulk-emailer/bulk-emailer-demo.gif" alt="drawing" style="width: 480px;"/>

1.  Type `bulk-emailer` on the command line with the following parameters: 
2.  Parameter `--from` (string) for sender's email.
3.  Parameter `--subject`(string) for the subject of the email.
4.  Parameter `--message` (file) path to the file of your email content.
5.  Parameter `--contacts` (file) path to the cvs file of your recipients.
6.  Hit **Enter** to send messages to recipient simutaneouly.

## Emails sent  
<img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/bulk-emailer/bulk-emailer-ouput.jpeg" alt="drawing" style="width: 200px;"/>

##Prerequisite 
Again, you must have AWS SES set up to use this tool.

<hr>

## Issues

Please open  *issues* here: [New Issue](https://github.com/warrensbox/bulk-emailer/issues){:target="_blank"}

<hr>

See how to *upgrade*, *uninstall*, *troubleshoot* here:
[Additional Info](additional)