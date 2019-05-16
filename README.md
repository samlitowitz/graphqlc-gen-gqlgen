# graphqlc-gen-gqlgen

Super preliminary stages, use at your own risk.

This is a code generator designed to work with [graphqlc](https://github.com/samlitowitz/graphqlc).
The intent is to take a gqlgen configuration file and produce identical output allowing a seamless transition from gqlgen to graphqlc.

# Installation
Install [graphqlc](https://github.com/samlitowitz/graphqlc).

`go get -u github.com/samlitowitz/graphqlc-gen-gqlgen/cmd/graphqlc-gen-gqlgen`


# Usage
`graphqlc --gqlgen_out=config=.gqlgen.yml:. schema.graphql`

