/*
The loader packages abstracts the top-level xslate package from the job of
loading the bytecode from a key value.
*/
package loader

import (
  "errors"

  "github.com/lestrrat/go-xslate/vm"
)

const (
  DUMP_BYTECODE = 1 << iota
  DUMP_AST
)

type Flags struct { flags int32 }

func NewFlags() *Flags {
  return &Flags { 0 }
}

func (f *Flags) DumpAST (b bool) {
  if b {
    f.flags |= DUMP_AST
  } else {
    f.flags &= ^DUMP_AST
  }
}

func (f *Flags) DumpByteCode (b bool) {
  if b {
    f.flags |= DUMP_BYTECODE
  } else {
    f.flags &= ^DUMP_BYTECODE
  }
}

func (f *Flags) ShouldDumpAST() bool {
  return f.flags & DUMP_AST == DUMP_AST
}

func (f Flags) ShouldDumpByteCode() bool {
  return f.flags & DUMP_BYTECODE == 1
}

type DebugDumper interface {
  DumpAST(bool)
  DumpByteCode(bool)
  ShouldDumpAST() bool
  ShouldDumpByteCode() bool
}

// Teh ByteCodeLoader loads the ByteCode given a key
type ByteCodeLoader interface {
  DebugDumper
  LoadString(string) (*vm.ByteCode, error)
  Load(string) (*vm.ByteCode, error)
}

// The template loader loads the template string given a key
type TemplateLoader interface {
  Load(string) ([]byte, error)
}

var ErrTemplateNotFound = errors.New("Specified template was not found")

