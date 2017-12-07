# informer5-tooling

Some tooling we are writing against the Informer 5 API.

This tool lets us automate our Informer workflow.
We can create and manage datasets.

We will slowly add more functionality to manage the data (and reports).

## Get it

```
go get -u github.com/dghubble/sling
go get -u github.com/ldfritz/informer5-tooling/cmd/i5
```

## Use it

We still need to implement a stable configuration setup.
We need to set the API base path and the token.
It currently reads from a config JSON file in the current directory.

It is likely that this file will always need to be manually created.
I don't see an obvious way to generate a token using the API without already having a token.

The user interface is still very immature and likely to change.
It currently implements a limited set of dataset manipulations.

```
Usage: i5 RESOURCE COMMAND [ARGUMENTS]

Resources:
  dataset

dataset Commands:
  append   ID FILENAME           Add data from FILE to dataset.
  create   NAME [DESCRIPTION]    Create a new dataset.
  delete   ID                    Delete a dataset.
  download ID                    Download dataset.
  info     ID                    Display dataset information.
  list                           List datasets.
  upload   ID FILENAME           Upload data from FILE to dataset.
```
