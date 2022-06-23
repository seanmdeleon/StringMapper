# StringMapper

To run the code

    cd main/
    go build
    ./main

To add new testcases, update the main function in main.go
I wrote a small test function that compares expected vs actual 'func (s *SkipString) AreStringsEqual(expected string) bool'. This allows for more information in the testing portion of Part 2
If I were to truly test this I would have used the testify golang package to write unittests and assert, but for simplicity and a time limit tests are just written in the main func

Part 1 of the assignment is the function 'func CapitalizeEveryThirdAlphanumericChar(s string) string' in main.go
Part 2 of the assignment is implemented in mapper.go
utils/utils.go holds two utility functions

For ease I kept the SkipString struct and implementation in the mapper package but probably would have created its own package for better organization