# Simple Placeholder Generator

[![Go Report Card](https://goreportcard.com/badge/github.com/otaleghani/spg)](https://goreportcard.com/report/github.com/otaleghani/spg) 
![CI Tests](https://github.com/otaleghani/spg/actions/workflows/tests.yml/badge.svg) 
[![CodeFactor](https://www.codefactor.io/repository/github/otaleghani/spg/badge)](https://www.codefactor.io/repository/github/otaleghani/spg)

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
opt := Options{Format: "camel"}
spg := spg.New("en_usa")

for i := 0; i < 5000; i++ {
    fmt.Println(spg.Person().FullName(opt))
}
```

Keep in mind that this package can generate around 300.000 records in one second for single fields calls (like Person().FirstName) and around 200.000 records for more complex queries (like Product().ProductName)

## Available Data

### Person

The neutral version (`FirstName()` and `FullName()`) randomly chooses between the male or fema dictionary.

``` go
opt := Options{}
spg := spg.New("en_usa")

spg.Person().FirstName(opt) 
// e.g. "Jacob"

spg.Person().MaleFirstName(opt) 
// e.g. "Michael"

spg.Person().FemaleFirstName(opt) 
// e.g. "Emily"

spg.Person().LastName(opt) 
// e.g. "Smith"

spg.Person().FullName(opt) 
// e.g. "Hannah Johnson"

spg.Person().MaleFullName(opt) 
// e.g. "Matthew Williams"

spg.Person().FemaleFullName(opt) 
// e.g. "Madison Brown"
```

### Place

The `Locale` version of the places takes data from a pool of names for cities, streets and states related to the chosen language, while the non-locale kind takes a more broad database.
The address function is called `AddressEn()` becaouse some countries use a different approach to writing addresses. 

``` go
opt := Options{}
spg := spg.New("en_usa")

spg.Place().City(opt) 
// e.g. "St. Petersburg"

spg.Place().CityLocale(opt) 
// e.g. "New York"

spg.Place().Street(opt) 
// e.g. "Bourke Street"

spg.Place().StreetLocale(opt) 
// e.g. "Main Street"

spg.Place().State(opt) 
// e.g. "Tokyo"

spg.Place().StateLocale(opt) 
// e.g. "Alabama"

spg.Place().Country(opt) 
// e.g. "Thailand"

spg.Place().CountryLocale(opt) 
// e.g. "United States"

spg.Place().AddressEn(opt) 
// e.g. 102 La rambla
// Aurora, Yukon 96121
// Zambia

spg.Place().AddressEnLocale(opt) 
// 100 Riverside drive
// Boise, Washington 14139
// United states
```
### Products

N.B.: Complete Product Names and Product descriptions are hilarious.

``` go
opt := Options{}
spg := spg.New("en_usa")

spg.Product().Thing(opt) 
// e.g. "Smartwatch"

spg.Product().Brand(opt) 
// e.g. "Apple"

spg.Product().Technology(opt) 
// e.g. "Artificial Intelligence"

spg.Product().ProductName(opt) 
// e.g. "QuickWave"

spg.Product().CompleteProductName(opt) 
// e.g. "Heating pad Dolce & gabbana Quickquench with Edge computing"

spg.Product().ProductDescription(opt) 
// e.g. "Optimize your home entertainment with our streaming device, designed with Wi-fi 6. enjoy seamless streaming, high-definition video, and a user-friendly interface, all enabled by this advanced technology. transform your viewing experience with the superior capabilities of Wi-fi 6."

spg.Product().UUID(opt) 
// e.g. d303dd94-1a03-0a3a-9c95-35c16d1c2930
```

### Animals

``` go
opt := Options{}
spg := spg.New("en_usa")

spg.Animal().Mammal(opt),
// e.g. Lion

spg.Animal().Fish(opt),
// e.g. Salmon

spg.Animal().Bird(opt),
// e.g. Eagle

spg.Animal().Insect(opt),
// e.g. Ant

spg.Animal().Pet(opt),
// e.g. Dog

spg.Animal().PetName(opt),
// e.g. Bella

spg.Animal().DogBreed(opt),
// e.g. Labrador Retriever

spg.Animal().CatBreed(opt),
// e.g. Persian

spg.Animal().Sex(opt),
// e.g. Male

spg.Animal().MicrochipNumber(opt),
// e.g. 545931400687715
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

spg.ExadecimalNumber()
// e.g. a

spg.ExadecimalNumberFixed(3)
// e.g. ae2
```

### Boolean
``` go
spg := spg.New("en_usa")

spg.Boolean() 
// e.g. true

spg.StringBoolean()
// e.g. true (as a string type)
```

### Internet
``` go
spg := spg.New("en_usa")

spg.Internet().Username(Options{})
// Either a casual Internet().CasualUsername or a Internet().BusinessUsername()

spg.Internet().CasualUsername(Options{})
// e.g. Jhonny87

spg.Internet().BusinessUsername(Options{})
// e.g. jacob.smith

spg.Internet().TLD(Options{})
// Either Internet().CCTLD() or Internet().GTLD()

spg.Internet().TLDLocale(Options{})
// e.g. .com

spg.Internet().CCTLD(Options{})
// e.g. .fr (country code domain)

spg.Internet().GTLD(Options{})
// e.g. .biz (general domains)

spg.Internet().DomainName(Options{})
// e.g. johnson.com

spg.Internet().DomainNameLocale(Options{})
// e.g. cat.us

spg.Internet().EmailProvider(Options{})
// e.g. gmail.com

spg.Internet().EmailProvider(Options{})
// e.g. gmail.com

spg.Internet().Email(Options{})
// e.g. lucy@asics.digital

spg.Internet().CasualEmail(Options{})
// e.g. oscar@gmail.com

spg.Internet().BusinessEmail(Options{})
// e.g. g.weaver@charlie.news
```

## To do

- [ ] Add italian locale

### Generations to add

#### Products
- [x] UUIDs

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
