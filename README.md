# informer5-tooling

This tool lets us automate our Informer workflow using the Informer 5 API.
We use it to manage datasets and datasources.

Although the API is exposed as a Swagger/OpenAPI Specification, this project is currently being mostly hand-coded.
However, as the definition becomes more mature, we hope to use [swagger-api/swagger-codegen](github.com/swagger-api/swagger-codegen) to generate much of the code and eliminate the requirement for so much manual work.
This project could then focus on providing a command-line wrapper.

## Get it

```
go get -u github.com/ldfritz/informer5-tooling/cmd/i5
```

## Configure it

```
echo '{"API":"http://example.com/","token":"..."}' >> config.json
```

We still need to implement a stable configuration setup.
We need to set the API base path and the token.
It currently reads from a config JSON file in the current directory.

It is likely that this file will always need to be manually created.
I don't see an obvious way to generate a token using the API without already having a token.

## Use it

```
i5 help
```

The user interface is still very immature and likely to change.
It currently implements a limited set of dataset and datasource manipulations.
