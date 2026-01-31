package fail

import (
	"container/list"
	"sync"
)

// Mapper converts errors in any direction: generic->fail, fail->generic, fail->fail, etc.
type Mapper interface {
	Name() string
	Priority() int

	// Map attempts to map the input error to some other error (generic or fail)
	// Returns ok = true if the mapper handled the error, false otherwise
	Map(error) (error, bool)

	// MapToFail : Optional convenience for mapping from generic error to fail.Error type
	MapToFail(err error) (*Error, bool)

	// MapFromFail : Optional convenience for mapping from fail.Error to another error type
	MapFromFail(*Error) (error, bool)
}

// RegisterMapper adds a generic error mapper
func RegisterMapper(mapper Mapper) {
	global.RegisterMapper(mapper)
}

func (r *Registry) RegisterMapper(mapper Mapper) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Insert in priority order (higher first)
	r.genericMappers.Add(mapper)
}

// MapperList keeps mappers sorted by priority using container/list
type MapperList struct {
	mu      sync.RWMutex
	mappers *list.List // *list.Element.Value will be Mapper
}

// NewMapperList creates a new MapperList. If includeDefault is true, adds the default mapper with priority -1
func NewMapperList() *MapperList {
	ml := &MapperList{
		mappers: list.New(),
	}

	return ml
}

// Add inserts a mapper into the list by descending priority
func (ml *MapperList) Add(m Mapper) {
	ml.mu.Lock()
	defer ml.mu.Unlock()

	priority := m.Priority()
	for e := ml.mappers.Front(); e != nil; e = e.Next() {
		existing := e.Value.(Mapper)
		if priority > existing.Priority() {
			ml.mappers.InsertBefore(m, e)
			return
		}
	}
	// If we didn't insert yet, add at the end
	ml.mappers.PushBack(m)
}

// All returns all mappers as a slice
func (ml *MapperList) All() []Mapper {
	ml.mu.RLock()
	defer ml.mu.RUnlock()

	out := make([]Mapper, 0, ml.mappers.Len())
	for e := ml.mappers.Front(); e != nil; e = e.Next() {
		out = append(out, e.Value.(Mapper))
	}
	return out
}

// MapError tries to map an error using mappers in priority order
func (ml *MapperList) MapError(err error) (error, bool) {
	ml.mu.RLock()
	defer ml.mu.RUnlock()

	for e := ml.mappers.Front(); e != nil; e = e.Next() {
		if mapped, ok := e.Value.(Mapper).Map(err); ok {
			return mapped, true
		}
	}
	return nil, false
}

// MapToFail maps to *fail.Error
func (ml *MapperList) MapToFail(err error) (*Error, bool) {
	ml.mu.RLock()
	defer ml.mu.RUnlock()

	for e := ml.mappers.Front(); e != nil; e = e.Next() {
		if fe, ok := e.Value.(Mapper).MapToFail(err); ok {
			return fe, true
		}
	}
	return nil, false
}

// MapFromFail maps from *fail.Error
func (ml *MapperList) MapFromFail(err *Error) (error, bool) {
	ml.mu.RLock()
	defer ml.mu.RUnlock()

	for e := ml.mappers.Front(); e != nil; e = e.Next() {
		if mapped, ok := e.Value.(Mapper).MapFromFail(err); ok {
			return mapped, true
		}
	}
	return nil, false
}
