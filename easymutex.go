// Package easymutex is a small wrapper around sync.Locker that allows you to always defer the unlock, even if you might unlock sooner.
//
// The struct is simply a mutex with a boolean. The boolean knows if we have the lock and skips Unlock if we didn't have it.
package easymutex

import "sync"

// EasyLocker is a wrapper around sync.Locker that keeps track of whether we have the mutex.
type EasyLocker struct {
	L    sync.Locker
	Held bool
}

func (e *EasyLocker) Lock() {
	e.L.Lock()
	e.Held = true
}

func (e *EasyLocker) Unlock() {
	if e.Held {
		e.L.Unlock()
		e.Held = false
	}
}

// EasyMutex is a wrapper around *sync.Mutex that keeps track of whether we have the mutex.
type EasyMutex struct {
	L    *sync.Mutex
	Held bool
}

func (e *EasyMutex) Lock() {
	e.L.Lock()
	e.Held = true
}

func (e *EasyMutex) TryLock() bool {
	if e.L.TryLock() {
		e.Held = true
		return true
	}
	return false
}

func (e *EasyMutex) Unlock() {
	if e.Held {
		e.L.Unlock()
		e.Held = false
	}
}

// EasyRWMutex is a wrapper around *sync.RWMutex that keeps track of whether we have the mutex.
type EasyRWMutex struct {
	L             *sync.RWMutex
	HeldExclusive bool
	HeldShared    bool
}

func (e *EasyRWMutex) Lock() {
	e.L.Lock()
	e.HeldExclusive = true
}

func (e *EasyRWMutex) TryLock() bool {
	if e.L.TryLock() {
		e.HeldExclusive = true
		return true
	}
	return false
}

func (e *EasyRWMutex) Unlock() {
	if e.HeldExclusive {
		e.L.Unlock()
		e.HeldExclusive = false
	}
}

func (e *EasyRWMutex) RLock() {
	e.L.RLock()
	e.HeldShared = true
}

func (e *EasyRWMutex) TryRLock() bool {
	if e.L.TryRLock() {
		e.HeldShared = true
		return true
	}
	return false
}

func (e *EasyRWMutex) RUnlock() {
	if e.HeldShared {
		e.L.RUnlock()
		e.HeldShared = false
	}
}

// EasyUnlock calls Unlock() or RUnlock() or neither based on whether have the lock exclusive/shared or not at all.
func (e *EasyRWMutex) EasyUnlock() {
	if e.HeldExclusive {
		e.L.Unlock()
		e.HeldExclusive = false
	} else if e.HeldShared {
		e.L.RUnlock()
		e.HeldShared = false
	}
}
