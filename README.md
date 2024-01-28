# Usage

## Copy

```go
import "github.com/bygo/deep"

var usersCopy = deep.Copy(users)

```

## CopyIgnore

Some structs will not be deconstructed and their fields copied, but the struct will be copied.

```go
import "github.com/bygo/deep"

var usersCopy = deep.CopyIgnore(users, User{}, Article{})
```