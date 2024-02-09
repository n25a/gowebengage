<div align="center">
    <img src="webengage.png" align="center" width="600" height="315"/>
    <div align="center"> 
        <h1> 
            <b>
                <span style="color:white;"> 
                    Webengage Golang SDK
                </span>
            </b>
        </h1>
    </div>
</div>


## Introduction
This SDK allows you to integrate WebEngage with your Golang application. 
Using this SDK, you can track user events, send transactional events, and more. 
You can also use this SDK to trigger campaigns and track user activity in your application.  
The SDK is built using the WebEngage REST APIs. You can check the source document [here](https://docs.webengage.com/docs).

## Installation
To install the WebEngage Golang SDK, run the following command:
```bash
go get github.com/n25a/gowebengage
```

## Usage
To use the WebEngage Golang SDK, you need to import the package and initialize the client.

```go
package main

import (
    "github.com/n25a/gowebengage"
)

func main() {
    // Initialize the client
    client := gowebengage.NewClient("YOUR_WEBENGAGE_ADDRESS")
}
```

for more information about all functions and how to use them, please check `Webengage` interface in `client.go` file.  