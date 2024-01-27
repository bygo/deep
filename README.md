# Usage

## Copy

```go
import "github.com/bygo/deep"

var usersCopy = deep.Copy(users)

```

## CopyIgnore

Some structures will be ignored

```go
import "github.com/bygo/deep"

var usersCopy = deep.CopyIgnore(users, User{}, Article{})
```