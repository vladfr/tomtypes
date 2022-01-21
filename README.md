# tomtypes :: Go custom types

Go custom types

### StringInt
StringInt can decode a number from a string, useful when accepting user input in JSON APIs. StringInt implements a custom JSON Unmarshal function that parses ints from strings and floats. 

// Credit for code and idea to https://codesahara.com/blog/golang-cannot-unmarshal-string-into-go-value-of-type-int/

