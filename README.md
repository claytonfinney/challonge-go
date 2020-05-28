# challonge-go

A Golang client for the [Challonge web API](https://api.challonge.com/v1).

### Basic Usage
```go
package main

import (
    "fmt"
    "time""

    "github.com/claytonfinney/challonge-go"
)

func main() {

    // Provide your API creds (use env variables!)
    u := "YOUR_USERNAME"
    k := "YOUR_API_KEY"

    // Initialize the client
    c := challonge.NewChallongeClient(u, k)

    // Give the tournament a starting time
    tme, err := time.Parse(time.RFC3339, "2020-05-13T19:45:07.000Z")
    if err != nil {
        panic(err)    
    }
    trn := challonge.Tournament{challonge.TournamentKey{
        Name:                "challonge go api mock",
        TournamentType:      "single elimination",
        Url:                 "challonge_go_api_mock",
        Description:         "testing a new api client for golang",
        OpenSignup:          false,
        HoldThirdPlaceMatch: true,
        Private:             false,
        StartAt:             tme,
        AcceptAttachments:   true,
    }}

    ct, err := c.CreateTournament(&trn)
    if err != nil {
        panic(err)
    }
    fmt.Println(ct.Name) // challonge go api mock
}

```

### Small Quirks

#### Promoted Structs and New Struct Creation

The structure of Challonge's API responses are best served with promoted struct fields in Go. In the case of any GETs, you don't need to worry about the promoted fields.

```go
t, _ := c.GetTournament("my_test_tournament")
// access the URL with t.Url instead of something like t.Tournament.Url,
//  which would correspond to the json payload's t.tournament.url
fmt.Println(t.Url) 
```

But when you create a new Go struct to send via POST or PUT, you need to use a nested struct that looks like the following:
```go

trn := challonge.Tournament{challonge.TournamentKey{
    Name: "My Test Tournament",
    Url: "my_test_tournament",
}}
```

#### Time

Time has to be formatted according to RFC3339. Challonge responses use a format of `2020-05-26T14:34:55.806-04:00`, so use this too with `time.Parse()`!

#### Naming

Another small "quirk", is that everything is CamelCase for struct fields. `Url`, not `URL`.

Each nested struct is the name of the resource with `Key` appended to the end. For `Tournament`, `TournamentKey`. For `Match`, `MatchKey`.

### Trying a Different Approach to Null Types in Go

I dreaded making this package because the documentation for the Challonge API is a bit lacking, and it seems like any key in a response could be null at any given time. A commonly suggested way to tackle nullable JSON values in Go is to use pointers in the response. In my opinion, requiring the user to first check if the value is nil for each struct field they wish to process seemed ridiculous. Using omitempty seemed fine, but empty string and null are two different things, and Go doesn't really provide a way to differentiate between the two.

Thus, this package implements custom types with their own Unmarshal and Marshal functions implemented. When null JSON values are unmarshalled, they take on the value of its type constant defined in `types.go`. For example, a JSON value that is null, but would otherwise be an integer, would be unmarshalled with the value -16487502. Obviously this leaves rooms for collisions, and this approach may not be suitable for more robust production code. But in the case of the Challonge API, you will effectively never be dealing with negative numbers in the first place. If you, or someone else inputs a score with a value of -16487502, I am no longer responsible for that field marshalling or unmarshalling as `null` :).

This allows for easy null checks:

```go
package main

import (
    "fmt"

    "github.com/claytonfinney/challonge-go"
)

func main() {
    u := "YOUR_USERNAME"
    k := "YOUR_API_KEY"

    c := challonge.NewChallongeClient(u, k)
    t, _ := c.GetTournament("my_test_tournament")
    if challonge.IsNull(t.StartedAt) {
        fmt.Println("A start time has not yet been set for this tournament.")
    }
}
```

### Examples

For the best examples of how to use all the various functions, refer to the example program provided in [examples/](https://github.com/claytonfinney/challonge-go/tree/master/examples), then the unit tests.

### Known Issues With the API
* You cannot add an actual file via the Match Attachments API, but you can still update the URL and description. Challonge throws a 500 when you try via curl or anything else.
