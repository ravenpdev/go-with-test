# Learn go with test

#### Commands

- go test
- go test -bench=.
- go test -bench=. -benchmem
- go test -cover
- go test -race

** Default zero values for all go types **

| Types                                                      | Value |
| ---------------------------------------------------------- | ----- |
| float                                                      | 0.0   |
| int                                                        | 0     |
| bool                                                       | false |
| string                                                     | ""    |
| interfaces, slices, channels, maps, pointers and functions | nil   |
