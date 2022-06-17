package mcli

import "flag"

func init() {
	defaultApp = NewApp()
}

var defaultApp *App

// SetOptions sets optional options for App.
func SetOptions(opts Options) {
	defaultApp.SetOptions(opts)
}

// SetGlobalFlags sets global flags, global flags are available to all commands.
// DisableGlobalFlags may be used to disable global flags for a specific
// command when calling Parse.
func SetGlobalFlags(v interface{}) {
	defaultApp.SetGlobalFlags(v)
}

// Add adds a command.
func Add(name string, f func(), description string) {
	defaultApp.Add(name, f, description)
}

// AddHidden adds a hidden command.
//
// A hidden command won't be showed in help, except that when a special flag
// "--mcli-show-hidden" is provided.
func AddHidden(name string, f func(), description string) {
	defaultApp.AddHidden(name, f, description)
}

// AddGroup adds a group explicitly.
// A group is a common prefix for some commands.
// It's not required to add group before adding sub commands, but user
// can use this function to add a description to a group, which will be
// showed in help.
func AddGroup(name string, description string) {
	defaultApp.AddGroup(name, description)
}

// AddHelp enables the "help" command to print help about any command.
func AddHelp() {
	defaultApp.AddHelp()
}

// Run runs the program, it parses the command line and searches for a
// registered command, it runs the command if a command is found,
// else it will report an error and exit the program.
//
// Optionally you may specify args to parse, by default it parses the
// command line arguments os.Args[1:].
func Run(args ...string) {
	defaultApp.Run(args...)
}

// Parse parses the command line for flags and arguments.
// v should be a pointer to a struct, else it panics.
func Parse(v interface{}, opts ...ParseOpt) (fs *flag.FlagSet, err error) {
	return defaultApp.Parse(v, opts...)
}

// PrintHelp prints usage doc of the current command to stderr.
func PrintHelp() {
	defaultApp.PrintHelp()
}
