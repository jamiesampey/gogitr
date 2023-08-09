package sorting

type Dog struct {
	Name, Breed, Color string
	Age                int
}

type ByAge []Dog

func (s ByAge) Len() int {
	return len(s)
}

func (s ByAge) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByAge) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

type ByName []Dog

func (s ByName) Len() int {
	return len(s)
}

func (s ByName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByName) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}
