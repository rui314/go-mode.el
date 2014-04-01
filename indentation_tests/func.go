package main

func (T)
func (p.T)
func (x T)
func (x, y T, z p.T)
func (interface{}, map[T]T, []T, *T)

func (T) T
func (T) *T
func (T) (T)
func (T) (*T)
func (T) (T, T)
func (T) (n T, err T)
func (T) (m, n T)

func fn(T)
func fn(T, T, T)
func fn(x T)
func fn(x, y T, z T)

func (T) method(T) T
func (x T) method(T) (T, T)
func (x T) method(T) (n T, err T)

func (x func(x T), y func(T, T) T) func() T
