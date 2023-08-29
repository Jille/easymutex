[![Go Reference](https://pkg.go.dev/badge/github.com/Jille/easymutex.svg)](https://pkg.go.dev/github.com/Jille/easymutex)

Easymutex is a small wrapper around a sync.Locker that allows you to always defer the unlock, even if you might unlock sooner.

The struct is simply a mutex with a boolean. The boolean knows if we have the lock and skips Lock/Unlock if we already/don't have the lock.

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

or even make use of possibly having the mutex in a defer:

```go
em := easymutex.EasyMutex{Locker: &myMtx}
defer em.Unlock()

defer func() {
	em.Lock()
	// do some cleanup
	em.Unlock()
}()

em.Lock()
if err := something(); err != nil {
	return err
}
em.Unlock()
```
