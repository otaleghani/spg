# Simple Placeholder Generator

A simple utility used to generate fake data for testing purpose.

## Installation
Add to your imports spg.

``` go
import "github.com/otaleghani/spg"
```
Afterwards run `go get github.com/otaleghani/spg` to get the package.

## Basic Usage
First off you will need to initialize a new spg, which is done by calling `spg.New()`. The `Generator` struct is responsible for creating a rng seed to query the data and defining the language.

Every single call to `spg` will need an `Options` object. This object will define the format and the separator of the output. If not defined the default format and separators will be used.

``` go
package main

import "github.com/otaleghani/spg"

func main(){
    spg := spg.New("en_usa")

    spg.Person().FullName(Options{})
    // FirstName LastName
    
    spg.Person().FullName(Options{Format: "title", Separator: "-"})
    // FIRSTNAME-LASTNAME

    spg.Place().AddressEn(Options{})
    // Number Street
    // City State Zip Code
    // Country

    spg.Product().Thing(Options{})
    // e.g. Smartphone
}
```

Most of the times you will use this package inside of a loop to get more values in one go.

``` go
    spg := spg.New("en_usa")

    for i := 0; i < 5000; i++ {
        fmt.Println(spg.Person().FullName())
    }
```

Keep in mind that this package can generate around 300.000 records in one second for single fields calls (like Person().FirstName) and around 200.000 records for more complex queries (like Product().ProductName)

## Features

- International


## Generations to add

- [x] Names
    - [x] Female First Name
    - [x] Female Full Name
    - [x] Male First Name
    - [x] Male Full Name
    - [x] Full Names
    - [x] First Names
    - [x] Last names
- [ ] Time based
    - [ ] Birthdays
- [ ] Finance records
    - [x] Credit cards
- [ ] Numbers
- [ ] Usernames
- [ ] Passwords
- [ ] Nationality
- [ ] Words
- [ ] Email
- [ ] UUIDs

## To do

- [ ] Localization methods (auto-translation could be good enough)
- [ ] Use GPT4 to create 100 words for different languages

## Data in use

The data in use 

