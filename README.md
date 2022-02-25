# prototag
**Go style tag parser for protocol buffers**

## Install

for cli  
```go install github.com/wirekang/prototag/cmd/prototag@latest```

for go modules   
```go get github.com/wirekang/prototag```

## Example
*example.proto*
```protobuf
syntax = "proto3";
package example;

message Person {
  string name = 1; // insert key:"value" in backticks like go `json:"foo" xml:"bar"`
  int32 id = 2;  // comments not in backticks will be ignored
  string email = 3; // use can escape double quotes `key:"va\"u\"es"`

  enum PhoneType {
    MOBILE = 0; // "not backtick"
    HOME = 1; // `key:"value" key2:"value for key2"`
    WORK = 2; // `wrong:syntax:will:be:ignored`
  }
}

message AddressBook {
  repeated Person people = 1;
}

```

### CLI
```shell
$ prototag
prototag [flags] [file or stdin]
  -d    debug mode
  -j    json output
  -o string
        output to file
  -p    pretty output
```
```shell
$ prototag -j example.proto # same with `prototag -j < example.proto`
{"messages":[{"name":"Person","fields":[{"name":"name","number":1,"tags":[{"key":"json","value":"foo"},{"key":"xml","value":"bar"}]},{"name":"id","number":2,"tags":[]},{"name":"email","number":3,"tags":[{"key":"key","value":"va\"u\"es"}]}]},{"name":"AddressBook","fields":[{"name":"people","number":1,"tags":[]}]}],"enums":[{"name":"PhoneType","fields":[{"name":"MOBILE","number":0,"tags":[]},{"name":"HOME","number":1,"tags":[{"key":"key","value":"value"},{"key":"key2","value":"value for key2"}]},{"name":"WORK","number":2,"tags":[]}]}]}
```
```shell
$ prototag -j -p example.proto
{
    "messages": [
        {
            "name": "Person",
            "fields": [
                {
                    "name": "name",
                    "number": 1,
                    "tags": [
                        {
                            "key": "json",
                            "value": "foo"
                        },
                        {
                            "key": "xml",
                            "value": "bar"
                        }
                    ]
                },
                {
                    "name": "id",
                    "number": 2,
                    "tags": []
                },
                {
                    "name": "email",
                    "number": 3,
                    "tags": [
                        {
                            "key": "key",
                            "value": "va\"u\"es"
                        }
                    ]
                }
            ]
        },
        {
            "name": "AddressBook",
            "fields": [
                {
                    "name": "people",
                    "number": 1,
                    "tags": []
                }
            ]
        }
    ],
    "enums": [
        {
            "name": "PhoneType",
            "fields": [
                {
                    "name": "MOBILE",
                    "number": 0,
                    "tags": []
                },
                {
                    "name": "HOME",
                    "number": 1,
                    "tags": [
                        {
                            "key": "key",
                            "value": "value"
                        },
                        {
                            "key": "key2",
                            "value": "value for key2"
                        }
                    ]
                },
                {
                    "name": "WORK",
                    "number": 2,
                    "tags": []
                }
            ]
        }
    ]
}
```

### Go

```go
package main

import (
	"fmt"

	"github.com/wirekang/prototag/pkg/prototag"
)

func main() {
	m, err := prototag.ParseFile("e.proto")
	if err != nil {
		return
	}

	fmt.Println(m.Messages[0].Name) // Person
	fmt.Println(m.Messages[0].Fields[0].Tags[1].Value) // bar
	
	// for Message(name) Field(name) Tag(name) instead of Message[n]
	m.Cache() 
	fmt.Println(m.Message("Person").Field("email").Tag("key").Value) // va"u"es
}


```