package main

type Notifier interface {
	Notify(message string)
}

// Simple notifier implementation used later for examples
type PrintNotifier struct{}

func (p PrintNotifier) Notify(message string) {
	// no-op simple notifier
	println(message)
}
