package model

import (
	"github.com/mkideal/cli"
)

type InputFormatT struct {
	InputFormat string `cli:"if,input-format" usage:"define the stricture format {j,json,t,toml,y,yml,yaml}" dft:"extension"`
}

type OutputFormatT struct {
	OutputFormat string `cli:"of,output-format" usage:"format of resulting file {j,json,t,toml,y,yml,yaml}" dft:"json"`
}

type PrettyOutputT struct {
	PrettyOutput bool `cli:"po,pretty-output" usage:"make a pretty output (only supported by JSON format)" dft:"false"`
}

// OutT define the arguments for output
type OutT struct {
	OutputFormatT
	FileOutput string `cli:"fo,file-output" usage:"write to file" dft:""`
	StdOutput  bool   `cli:"so,stdout,standard-streams-output" usage:"write to standard streams" dft:"true"`
	PrettyOutputT
}

// MultiInputT define the arguments for multiple inputs
type MultiInputT struct {
	InputFormatT
	FileInput []string `cli:"fi,file-input" usage:"read from files"`
	StdInput  []string `cli:"si,stdin,standard-streams-input" usage:"read from standard streams"`
}

// SingleInputT define the arguments for single input
type SingleInputT struct {
	InputFormatT
	FileInput string `cli:"fi,file-input" usage:"read from file"`
	StdInput  string `cli:"si,stdin,standard-streams-input" usage:"read from standard stream"`
}

// SingleQueryT define the arguments for single query
type SingleQueryT struct {
	Query string `cli:"q,query" usage:"query to define a position"`
}

// SingleItemT define the arguments for single item
type SingleItemT struct {
	Item string `cli:"i,item" usage:"item to insert"`
}

// ConcatenateT define the arguments for the concatenate command
type ConcatenateT struct {
	cli.Helper
	OutputFormatT
	PrettyOutputT
	MultiInputT
}

// CreateT define the arguments for the create command
type CreateT struct {
	cli.Helper
	OutputFormatT
	FileOutput []string `cli:"fo,file-output" usage:"write to files" dft:""`
}

// DeleteT define the arguments for the delete command
type DeleteT struct {
	cli.Helper
	SingleInputT
	SingleQueryT
	OutT
}

// DifferenceT define the arguments for the difference command
type DifferenceT struct {
	cli.Helper
	MultiInputT
	OutputFormatT
}

// ValidateT define the arguments for the equal command
type EqualT struct {
	cli.Helper
	MultiInputT
}

// SelectT define the arguments for the get and delete command
type SelectT struct {
	cli.Helper
	SingleInputT
	SingleQueryT
}

// MergeT define the arguments for the merge command
type MergeT struct {
	cli.Helper
	MultiInputT
	AppendArray bool `cli:"a,append,append-array" usage:"arrays can be replaced or appended" dft:"false"`
	OutT
}

// PutT define the arguments for the put command
type PutT struct {
	cli.Helper
	SingleInputT
	SingleQueryT
	SingleItemT
	OutT
}

// RooT define the arguments for the root command
type RootT struct {
	cli.Helper
	Version bool `cli:"v,version" usage:"display version"`
	List    bool `cli:"l,list" usage:"list all commands"`
}

// ValidateT define the arguments for the validate command
type ValidateT struct {
	cli.Helper
	SingleInputT
}
