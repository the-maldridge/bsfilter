# Boolean Set Filter

This library provides a performant way to filter sets based on a map
from string keys to map of attribute name to empty struct (set).  The
intent is to provide a performant way to do evaluation of arbitrary
boolean expressions that are expressed as equations whose symbols
appear in the attribute set of the second layer map.

Practically speaking, given the following map:

```go
map[string]map[string]struct{}{
    "foo": {
        "bar": struct{}{},
        "baz": struct{}{},
    },
    "bar: {
        "baz": struct{}{},
    }
}
```

It should be possible to run a filter over the map such as `bar` or
`!bar` and produce a resultant list of keys containing `["foo"]` and
`["bar"]` respectively.
