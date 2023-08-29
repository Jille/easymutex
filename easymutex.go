// Package easymutex is a small wrapper around sync.Locker that allows you to always defer the unlock, even if you might unlock sooner.
//
// The struct is simply a mutex with a boolean. The boolean knows if we have the lock and skips Unlock if we didn't have it.
package easymutex

import "sync"

// EasyLocker is a wrapper around sync.Locker that keeps track of whether we have the mutex.
type EasyLocker struct {
	Locker sync.Locker
	Held   bool
}

func (e *EasyLocker) Lock() {
	e.Locker.Lock()
	e.Held = true
}

func (e *EasyLocker) Unlock() {
	if e.Held {
		e.Locker.Unlock()
		e.Held = false
	}
}

// EasyMutex is a wrapper around *sync.Mutex that keeps track of whether we have the mutex.
type EasyMutex struct {
	Locker *sync.Mutex
	Held   bool
}

func (e *EasyMutex) Lock() {
	e.Locker.Lock()
	e.Held = true
}

func (e *EasyMutex) TryLock() bool {
	if e.Locker.TryLock() {
		e.Held = true
		return true
	}
	return false
}

func (e *EasyMutex) Unlock() {
	if e.Held {
		e.Locker.Unlock()
		e.Held = false
	}
}

// EasyRWMutex is a wrapper around *sync.RWMutex that keeps track of whether we have the mutex.
type EasyRWMutex struct {
	Locker        *sync.RWMutex
	HeldExclusive bool
	HeldShared    bool
}

func (e *EasyRWMutex) Lock() {
	e.Locker.Lock()
	e.HeldExclusive = true
}

func (e *EasyRWMutex) TryLock() bool {
	if e.Locker.TryLock() {
		e.HeldExclusive = true
		return true
	}
	return false
}

func (e *EasyRWMutex) Unlock() {
	if e.HeldExclusive {
		e.Locker.Unlock()
		e.HeldExclusive = false
	}
}

func (e *EasyRWMutex) RLock() {
	e.Locker.RLock()
	e.HeldShared = true
}

func (e *EasyRWMutex) TryRLock() bool {
	if e.Locker.TryRLock() {
		e.HeldShared = true
		return true
	}
	return false
}

func (e *EasyRWMutex) RUnlock() {
	if e.HeldShared {
		e.Locker.RUnlock()
		e.HeldShared = false
	}
}

// EasyUnlock calls Unlock() or RUnlock() or neither based on whether have the lock exclusive/shared or not at all.
func (e *EasyRWMutex) EasyUnlock() {
	if e.HeldExclusive {
		e.Locker.Unlock()
		e.HeldExclusive = false
	} else if e.HeldShared {
		e.Locker.RUnlock()
		e.HeldShared = false
	}
}
