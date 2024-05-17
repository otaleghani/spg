# Simple Placeholder Generator

[![Go Report Card](https://goreportcard.com/badge/github.com/otaleghani/spg)](https://goreportcard.com/report/github.com/otaleghani/spg) 
![CI Tests](https://github.com/otaleghani/spg/actions/workflows/tests.yml/badge.svg)

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

## Available Data

### Person

The neutral version (`FirstName()` and `FullName()`) randomly chooses between the male or fema dictionary.

``` go
spg := spg.New("en_usa")

spg.Person().FirstName() 
// e.g. "Jacob"

spg.Person().MaleFirstName() 
// e.g. "Michael"

spg.Person().FemaleFirstName() 
// e.g. "Emily"

spg.Person().LastName() 
// e.g. "Smith"

spg.Person().FullName() 
// e.g. "Hannah Johnson"

spg.Person().MaleFullName() 
// e.g. "Matthew Williams"

spg.Person().FemaleFullName() 
// e.g. "Madison Brown"
```

### Place

The `Locale` version of the places takes data from a pool of names for cities, streets and states related to the chosen language, while the non-locale kind takes a more broad database.
The address function is called `AddressEn()` becaouse some countries use a different approach to writing addresses. 

``` go
spg := spg.New("en_usa")

spg.Place().City() 
// e.g. "St. Petersburg"

spg.Place().CityLocale() 
// e.g. "New York"

spg.Place().Street() 
// e.g. "Bourke Street"

spg.Place().StreetLocale() 
// e.g. "Main Street"

spg.Place().State() 
// e.g. "Tokyo"

spg.Place().StateLocale() 
// e.g. "Alabama"

spg.Place().Country() 
// e.g. "Thailand"

spg.Place().CountryLocale() 
// e.g. "United States"

spg.Place().AddressEn() 
// e.g. 102 La rambla
// Aurora, Yukon 96121
// Zambia

spg.Place().AddressEnLocale() 
// 100 Riverside drive
// Boise, Washington 14139
// United states
```
### Products

N.B.: Complete Product Names and Product descriptions are hilarious.

``` go
spg := spg.New("en_usa")

spg.Product().Thing() 
// e.g. "Smartwatch"

spg.Product().Brand() 
// e.g. "Apple"

spg.Product().Technology() 
// e.g. "Artificial Intelligence"

spg.Product().ProductName() 
// e.g. "QuickWave"

spg.Product().CompleteProductName() 
// e.g. "Heating pad Dolce & gabbana Quickquench with Edge computing"

spg.Product().ProductDescription() 
// e.g. "Optimize your home entertainment with our streaming device, designed with Wi-fi 6. enjoy seamless streaming, high-definition video, and a user-friendly interface, all enabled by this advanced technology. transform your viewing experience with the superior capabilities of Wi-fi 6."
```

### Numbers
``` go
spg := spg.New("en_usa")
max := 5000

spg.StringNumber(max)
// e.g. "512"
// Returns a random positive number converted in a string from 0 to max

spg.StringNumberFixed(5)
// e.g. 92749
// Returns a random combination of numbers. The specified number is the lenght of the returned string..
```

## To do

- [ ] Add italian locale

### Generations to add

#### Products
- [ ] UUIDs

#### Person
- [ ] Email
- [ ] Username
- [ ] Password
- [ ] Birthday
- [ ] Occupations

#### Miscellaneous
- [ ] Industries
- [ ] Hobbies

#### Work
- [ ] 

#### Words
- [ ] Random word
- [ ] Random phrase
- [ ] Random adjective

#### Financial
- [ ] Credit cards
- [ ] Finantial record

#### Animals
- [x] Mammals
- [x] Birds
- [x] Fishes
- [x] Insects
- [x] Pets
- [x] Dogs breeds
- [x] Cats breeds
- [x] Pet names

#### Food
- [ ] 

#### Adjectives
- [ ] Positive
- [ ] Negative
- [ ] Colors
- [ ] Sizes

#### Dates
- [ ] Months
- [ ] Weekdays
- [ ] Time zones

## About the data

The data in use is generated using GPT4.
