/*
Package htmllog, a simple logger that produces an append-only, reloading HTML file

Features

	usable by multiple go-routines
	protection against printing recursive data structures in message
	configureable log event function
	configureable Html style
	configureable limit on message

Example

	h, err := htmllog.New("myapp.log")
	h.Infof("just now is %v", time.Now())
*/
package htmllog
