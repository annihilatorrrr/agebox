package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/alecthomas/kingpin/v2"
	"github.com/sirupsen/logrus"

	"github.com/slok/agebox/cmd/agebox/commands"
	"github.com/slok/agebox/internal/log"
	loglogrus "github.com/slok/agebox/internal/log/logrus"
)

// Version is the application version.
var Version = "dev"

// Run runs the main application.
func Run(ctx context.Context, args []string, stdin io.Reader, stdout, stderr io.Writer) error {
	app := kingpin.New("agebox", "Age based repo file encrypt helper.")
	app.Version(Version)
	app.DefaultEnvars()
	config := commands.NewRootConfig(app)

	// Setup commands (registers flags).
	initCmd := commands.NewInitCommand(app)
	encryptCmd := commands.NewEncryptCommand(app)
	decryptCmd := commands.NewDecryptCommand(app)
	reencryptCmd := commands.NewReencryptCommand(app)
	untrackCmd := commands.NewUntrackCommand(app)
	catCmd := commands.NewCatCommand(app)
	validateCmd := commands.NewValidateCommand(app)

	cmds := map[string]commands.Command{
		initCmd.Name():      initCmd,
		encryptCmd.Name():   encryptCmd,
		decryptCmd.Name():   decryptCmd,
		reencryptCmd.Name(): reencryptCmd,
		untrackCmd.Name():   untrackCmd,
		catCmd.Name():       catCmd,
		validateCmd.Name():  validateCmd,
	}

	// Parse commandline.
	cmdName, err := app.Parse(args[1:])
	if err != nil {
		return fmt.Errorf("invalid command configuration: %w", err)
	}

	// Set up global dependencies.
	config.Stdin = stdin
	config.Stdout = stdout
	config.Stderr = stderr
	config.Logger = getLogger(*config)

	// Execute command.
	err = cmds[cmdName].Run(ctx, *config)
	if err != nil {
		return fmt.Errorf("%q command failed: %w", cmdName, err)
	}

	return nil
}

// getLogger returns the application logger.
func getLogger(config commands.RootConfig) log.Logger {
	if config.NoLog {
		return log.Noop
	}

	// If not logger disabled use logrus logger.
	logrusLog := logrus.New()
	logrusLog.Out = config.Stderr // By default logger goes to stderr (so it can split stdout prints).
	logrusLogEntry := logrus.NewEntry(logrusLog)

	if config.Debug {
		logrusLogEntry.Logger.SetLevel(logrus.DebugLevel)
	}

	// Log format.
	switch config.LoggerType {
	case commands.LoggerTypeDefault:
		logrusLogEntry.Logger.SetFormatter(&logrus.TextFormatter{
			ForceColors:   !config.NoColor,
			DisableColors: config.NoColor,
		})
	case commands.LoggerTypeJSON:
		logrusLogEntry.Logger.SetFormatter(&logrus.JSONFormatter{})
	}

	logger := loglogrus.NewLogrus(logrusLogEntry).WithValues(log.Kv{
		"version": Version,
	})

	logger.Debugf("Debug level is enabled") // Will log only when debug enabled.

	return logger
}

func main() {
	ctx := context.Background()
	err := Run(ctx, os.Args, os.Stdin, os.Stdout, os.Stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
