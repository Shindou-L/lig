package lig

import (
	"lig/writer"
)

type builder struct {
	level Level
	hook  WriteHookFunc
	write writer.WriteFunc
}

func (b *builder) WithLevel(level Level) *builder {
	b.level = level
	return b
}

func (b *builder) WithBeforeWriteHook(hook WriteHookFunc) *builder {
	b.hook = hook
	return b
}

func (b *builder) Build() Logger {
	return &abstractLogger{
		level: b.level,
		hook:  b.hook,
		write: b.write,
	}
}

func NewBuilder(writeFunc writer.WriteFunc) *builder {
	return &builder{
		write: writeFunc,
	}
}
