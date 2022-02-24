package merrors

type CrashError interface {
	error
	Stack() string
	ErrorWithStack() string
}

type crash struct {
	err   error
	stack string
}

func (p *crash) Error() string {
	return p.err.Error()
}

func (p *crash) Stack() string {
	return p.stack
}

func (p *crash) ErrorWithStack() string {
	return p.Error() + "\n" + p.Stack()
}

func NewCrash(err error, stack string) CrashError {
	return &crash{
		err:   err,
		stack: stack,
	}
}
