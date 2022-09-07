# Ditto - A Simple CLI App for Epicor

## Overview
Ditto allows you to retrieve all cases from an environment and view basic details about them.
You can also search for cases, and download attachments.

## Disclaimer
This is a Hobby side project developed by me, and is not supported by Epicor.
There are likely bugs, and I will not be able to provide support for this. I have only picked up learning Go in the last
few weeks, so the code is likely not the best.


## Installation
Follow the instructions available at Golang to install Go on your system.
The instructions can be found here: https://golang.org/doc/install

Once Go is installed, you can install Ditto by running the following command in the root directory of this repo:
`go install`

This will install Ditto to your `$GOPATH/bin` directory.

Once it is installed you will need to create a .env file. This file should be named `.env` and should be placed in the 
same directory as the Ditto binary. You can view the template.env file to see what is required.

The .env file should contain the following information:
```
BASE_EPICOR_URL=https://YOUR_ENVIRONMENT_URL/ENVIRONMENT_INSTANCE/api/v2/odata/COMPANY/
EPICOR_API_KEY=YOUR_EPICOR_API_KEY
BASIC_AUTH=YOUR_BASIC_EPICOR_AUTH
```

## Usage
Ditto is a CLI app, so you will need to run it from the command line.

To launch Ditto, simply navigate to your `$GOPATH/bin` directory and run the following command:
`DittoV2`

This will launch Ditto and you will be presented with a list of cases and a
help menu along the bottom of the screen.

Alternatively, if you add the `$GOPATH/bin` directory to your PATH, you can run Ditto from anywhere.
