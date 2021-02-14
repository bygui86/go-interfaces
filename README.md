
`WORK IN PROGRESS`

# Go testing

Project to understand how to write tests in Golang

## TODO list

- testing types
    - [x] ~~unit-testing~~
    - [x] ~~integration-testing~~
    - [x] ~~integration-testing with testcontainers~~
    - [ ] contract-testing

- testing techniques
    - [x] ~~mocking~~
    - [ ] stubbing

- testing techs
    - [x] ~~basic interface~~
    - [x] ~~database~~
    - [ ] rest
        - [x] ~~server~~
        - [x] ~~client~~
        - [ ] middleware
    - [ ] grpc
        - [ ] server
        - [ ] client
    - [ ] broker
        - [ ] kafka
            - [ ] producer
            - [ ] consumer
        - [ ] kubemq
            - [ ] producer
            - [ ] consumer
    - [x] ~~logs~~
    - [ ] metrics
    - [ ] traces

- libraries
    - [ ] [mockery](github.com/vektra/mockery)

---

## Links

## testing in go

- https://golang.org/pkg/testing/

## code coverage

- https://blog.golang.org/cover
- https://stackoverflow.com/questions/10516662/how-to-measure-code-coverage-in-golang

## frameworks

- https://github.com/stretchr/testify
    - https://github.com/stretchr/testify#mock-package
- https://onsi.github.io/ginkgo/
- https://onsi.github.io/gomega/
- http://labix.org/gocheck

## mocking

- https://github.com/golang/mock
- https://github.com/vektra/mockery
- https://itnext.io/yet-another-tool-to-mock-interfaces-in-go-73de1b02c041
- https://blog.codecentric.de/2019/07/gomock-vs-testify/

## basic-example

- https://medium.com/swlh/using-go-interfaces-for-testable-code-d2e11b02dea

## unit-testing

### database

- https://medium.com/easyread/unit-test-sql-in-golang-5af19075e68e
- https://github.com/DATA-DOG/go-sqlmock
- https://dev.to/pieohpah/mocking-database-in-go-55bo

### http

#### server

- https://golang.org/pkg/net/http/httptest/

#### client

- https://www.thegreatcodeadventure.com/mocking-http-requests-in-golang/

#### middleware

- http://www.inanzzz.com/index.php/post/xgod/testing-a-middleware-within-golang
- https://stackoverflow.com/questions/51201056/testing-golang-middleware-that-modifies-the-request

### grpc

- https://github.com/tokopedia/gripmock

### logging

- https://stackoverflow.com/questions/52734529/testing-zap-logging-for-a-logger-built-from-a-custom-config/52737940
- https://gianarb.it/blog/golang-mockmania-zap-logger
- https://gist.github.com/iwuvjhdva/b0be6405dfd40cdf19467c6100bdf7c8

## integration-testing

- https://medium.com/@victorsteven/understanding-unit-and-integrationtesting-in-golang-ba60becb778d
- https://www.ardanlabs.com/blog/2019/03/integration-testing-in-go-executing-tests-with-docker.html
	- https://www.ardanlabs.com/blog/2019/10/integration-testing-in-go-set-up-and-writing-tests.html
- https://blog.gojekengineering.com/golang-integration-testing-made-easy-a834e754fa4c

## testcontainers

- https://github.com/testcontainers/testcontainers-go
- https://levelup.gitconnected.com/unit-test-sql-in-golang-without-mocking-using-testcontainers-go-postgres-docker-4f61574b1989
    - https://github.com/atkinsonbg/go-gmux-db-testcontainers

## contract-testing

- https://www.youtube.com/watch?v=-6x6XBDf9sQ
- https://medium.com/@kaigo/getting-started-with-contract-testing-with-golang-and-apiary-io-935a46f0e83f
    - https://apiary.io/
	- https://dredd.org/en/latest/

### pact

- https://pkg.go.dev/github.com/pact-foundation/pact-go?tab=doc
- https://github.com/pact-foundation/pact-workshop-go
- https://medium.com/trendyol-tech/writing-pact-contract-tests-with-golang-2c20b5049e0c

### gandalf

- https://github.com/JumboInteractiveLimited/Gandalf
- https://godoc.org/github.com/JumboInteractiveLimited/Gandalf
- https://www.youtube.com/watch?v=PAU-Hi9GEzs

## articles

- https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318
- https://medium.com/@povilasve/go-advanced-tips-tricks-a872503ac859
- https://levelup.gitconnected.com/better-tests-for-golang-apps-681ed2338677

## tutorials

- https://medium.com/better-programming/easy-guide-to-unit-testing-in-golang-4fc1e9d96679
- https://medium.com/better-programming/unit-testing-code-using-the-mongo-go-driver-in-golang-7166d1aa72c0
- https://gianarb.it/blog/unit-testing-kubernetes-client-in-go
