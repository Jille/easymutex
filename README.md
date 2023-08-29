[![Go Reference](https://pkg.go.dev/badge/github.com/Jille/easymutex.svg)](https://pkg.go.dev/github.com/Jille/easymutex)

Easymutex is a small wrapper around a sync.Locker that allows you to always defer the unlock, even if you might unlock sooner.

The struct is simply a mutex with a boolean. The boolean knows if we have the lock and skips Unlock if we didn't have it.

```go
var myMtx sync.Mutex

em := easymutex.EasyMutex{Locker: &myMtx}
em.Lock()
defer em.Unlock()
// do stuff
if err != nil {
	return err
}

em.Unlock()

// do more slow stuff

// The defered Unlock won't do anything because it sees we no longer have the lock.
```
