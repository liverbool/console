package console

type Args interface {
	// Get returns the named argument, or else a blank string
	Get(name string) string
	// first returns the first argument, or else a blank string
	first() string
	// Tail returns the rest of the arguments (the last "array" one)
	// or else an empty string slice
	Tail() []string
	// Len returns the length of the wrapped slice
	Len() int
	// Present checks if there are any arguments present
	Present() bool
	// Slice returns a copy of the internal slice
	Slice() []string
}

type args struct {
	values  []string
	command *Command
}

func (a *args) Get(name string) string {
	if a.command == nil {
		return ""
	}

	for i, arg := range a.command.Args {
		if arg.Name != name || arg.Slice {
			continue
		}

		if len(a.values) >= i+1 {
			return a.values[i]
		}

		return arg.Default
	}

	return ""
}

func (a *args) first() string {
	if len(a.values) > 0 {
		return (a.values)[0]
	}
	return ""
}

func (a *args) Tail() []string {
	if a.command != nil {
		for i, arg := range a.command.Args {
			if !arg.Slice {
				continue
			}

			if len(a.values) >= i+1 {
				ret := make([]string, len(a.values)-i)
				copy(ret, a.values[i:])
				return ret
			}

			break
		}
	} else if a.Len() > 1 {
		ret := make([]string, a.Len()-1)
		copy(ret, a.values[1:])
		return ret
	}

	return []string{}
}

func (a *args) Len() int {
	return len(a.values)
}

func (a *args) Present() bool {
	return a.Len() != 0
}

func (a *args) Slice() []string {
	ret := make([]string, len(a.values))
	copy(ret, a.values)
	return ret
}
