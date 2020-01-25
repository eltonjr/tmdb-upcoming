package genres

// collection holds every genre description cached for later use
var collection map[int]string

// Get returns the genre name given it's id
func Get(id int) string {
	return collection[id]
}
