package clark

import "fmt"

type Command interface {
	fmt.Stringer

	Labels() []Label
}

func NewCommand(labels ...Label) Command {
	if len(labels) == 0 {
		panic("a command needs at least one label")
	}

	return &command{labels: labels}
}

///////////////////////////////////////////////////////////////////////////
// Default Implementation
///////////////////////////////////////////////////////////////////////////

type command struct {
	labels []Label
}

func (command *command) String() string {
	return string(command.labels[0])
}

func (command *command) Labels() []Label {
	return command.labels
}
