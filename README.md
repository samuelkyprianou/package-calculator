# package-calculator

This project is a API which accepts a number and determines what quantities of a range of pre-defined pack-sizes are required to fulfil the order.
The Supported packsizes are as follows:
`250, 500, 750, 1000, 2000, 5000`
With the intention to support addition and removal of packsizes in future.
As pack sizes cannot be adjusted, the role of the api is to return the minimum number items, broken down into the minimum number of packs in order to fulfill the order.
The calculation prioritises item numbers over pack numbers, meaning that a greater number of packs which delivers a fewer number of items is favoured over fewer packs which return more items.

Example test cases are as follows:

| Items ordered | Correct number of packs                | Incorrect number of packs |
| ------------- | -------------------------------------- | ------------------------- |
| 1             | 2 x 250                                | 1 x 500                   |
| 250           | 1 x 250                                | 1 x 500                   |
| 251           | 1 x 500                                | 2 x 250                   |
| 501           | 1 x 500, 1 x 250                       | 1 x 1000 OR 3 x 250       |
| 8500          | 1 x 5000, 1 x 2000, 1 x 1000 , 1 x 500 | 1 x 5000, 2 x 2000        |
| 12001         | 2 x 5000, 1 x 2000, 1 x 250            | 3 x 5000                  |


### Deployed Setup

to run this application please visit - https://munearnq2n.us-east-1.awsapprunner.com/

### Local Setup

To test this API locally clone the repository, and run `go get .` in the root of the repository. After dependencies are installed, you can run the application using `go run .`, this will start the application on http://localhost:8080/. To run the test file run `cd calculatepacks` followed by `go test`.
