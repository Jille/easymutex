// Package easymutex is a small wrapper around sync.Locker that allows you to always defer the unlock, even if you might unlock sooner.
//
// The struct is simply a mutex with a boolean. The boolean knows if we have the lock and skips Lock/Unlock if we already/don't have the lock.
package easymutex

import "sync"

// LockLocker wraps a sync.Locker in an EasyLocker and locks it.
func LockLocker(l sync.Locker) *EasyLocker {
	l.Lock()
	return &EasyLocker{L: l, Held: true}
}

// EasyLocker is a wrapper around sync.Locker that keeps track of whether we have the mutex.
type EasyLocker struct {
	L    sync.Locker
	Held bool
}

func (e *EasyLocker) Lock() {
	if !e.Held {
		e.L.Lock()
		e.Held = true
	}
}

func (e *EasyLocker) Unlock() {
	if e.Held {
		e.L.Unlock()
		e.Held = false
	}
}

// LockMutex wraps a *sync.Mutex in an EasyMutex and locks it.
func LockMutex(l *sync.Mutex) *EasyMutex {
	l.Lock()
	return &EasyMutex{L: l, Held: true}
}

// EasyMutex is a wrapper around *sync.Mutex that keeps track of whether we have the mutex.
type EasyMutex struct {
	L    *sync.Mutex
	Held bool
}

func (e *EasyMutex) Lock() {
	if !e.Held {
		e.L.Lock()
		e.Held = true
	}
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

// LockRWMutex wraps a *sync.RWMutex in an EasyRWMutex and locks it exclusively.
func LockRWMutex(l *sync.RWMutex) *EasyRWMutex {
	l.Lock()
	return &EasyRWMutex{L: l, HeldExclusive: true}
}

// RLockRWMutex wraps a *sync.RWMutex in an EasyRWMutex and locks it in shared mode.
func RLockRWMutex(l *sync.RWMutex) *EasyRWMutex {
	l.RLock()
	return &EasyRWMutex{L: l, HeldShared: true}
}

// EasyRWMutex is a wrapper around *sync.RWMutex that keeps track of whether we have the mutex.
type EasyRWMutex struct {
	L             *sync.RWMutex
	HeldExclusive bool
	HeldShared    bool
}

func (e *EasyRWMutex) Lock() {
	if !e.HeldExclusive {
		e.L.Lock()
		e.HeldExclusive = true
	}
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
	if !e.HeldShared {
		e.L.RLock()
		e.HeldShared = true
	}
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
