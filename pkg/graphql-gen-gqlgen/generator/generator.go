package generator

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/99designs/gqlgen/api"

	gqcmd "github.com/99designs/gqlgen/cmd"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/stubgen"
	"github.com/samlitowitz/graphqlc/pkg/graphqlc/compiler"
)

type Generator struct {
	*bytes.Buffer
	Request  *compiler.CodeGeneratorRequest  // The input
	Response *compiler.CodeGeneratorResponse // The output

	Param map[string]string // Command-line parameters

	gqlgenConfig      *config.Config // gqlgen configuration
	serverFileName    string         // file name for gqlgen server output
	stubFileName      string         // file name for gqlgen stub output
	overwriteResolver bool
	overwriteServer   bool
}

func New() *Generator {
	g := new(Generator)
	g.Buffer = new(bytes.Buffer)
	g.Request = new(compiler.CodeGeneratorRequest)
	g.Response = new(compiler.CodeGeneratorResponse)
	return g
}

// Error reports a problem, including an error, and exits the program.
func (g *Generator) Error(err error, msgs ...string) {
	s := strings.Join(msgs, " ") + ":" + err.Error()
	log.Print("graphqlc-gen-test: error:", s)
	os.Exit(1)
}

// Fail reports a problem and exits the program.
func (g *Generator) Fail(msgs ...string) {
	s := strings.Join(msgs, " ")
	log.Print("graphqlc-gen-test: error:", s)
	os.Exit(1)
}

func (g *Generator) CommandLineArguments(parameter string) {
	g.Param = make(map[string]string)
	for _, p := range strings.Split(parameter, ",") {
		if i := strings.Index(p, "="); i < 0 {
			g.Param[p] = ""
		} else {
			g.Param[p[0:i]] = p[i+1:]
		}
	}

	for k, v := range g.Param {
		switch k {
		case "config":
			config, err := config.LoadConfig(v)
			if err != nil {
				g.Error(err)
			}
			g.gqlgenConfig = config
		case "overwriteResolver":
			switch v {
			case "true":
				g.overwriteResolver = true
			case "false":
				g.overwriteResolver = false
			default:
				g.Fail(fmt.Sprintf(`unknown %q, want "true" or "false"`, v))
			}
		case "overwriteServer":
			switch v {
			case "true":
				g.overwriteServer = true
			case "false":
				g.overwriteServer = false
			default:
				g.Fail(fmt.Sprintf(`unknown %q, want "true" or "false"`, v))
			}
		case "server":
			if v != "" {
				g.serverFileName = v
			}
		case "stub":
			if v != "" {
				g.stubFileName = v
			}
		}
	}
	// Use default config if no config provided
	if g.gqlgenConfig == nil {
		g.gqlgenConfig = config.DefaultConfig()
	}
	// Add default resolver if not provided
	if g.gqlgenConfig.Resolver == (config.PackageConfig{}) {
		g.gqlgenConfig.Resolver = config.PackageConfig{
			Filename: "resolver.go",
			Type:     "Resolver",
		}
	}
	// Use default server file if none provided
	if g.serverFileName == "" {
		g.serverFileName = "server/server.go"
	}
	// We've only parsed files to generate, let's not touch anything else
	g.gqlgenConfig.SchemaFilename = g.Request.FileToGenerate
}

func (g *Generator) GenerateAllFiles() {
	oldStderr := os.Stderr
	oldStdout := os.Stdout

	rStderr, wStderr, err := os.Pipe()
	if err != nil {
		g.Error(err)
	}
	rStdout, wStdout, err := os.Pipe()
	if err != nil {
		g.Error(err)
	}

	os.Stderr = wStderr
	os.Stdout = wStdout

	stderrC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, rStderr)
		stderrC <- buf.String()
	}()
	stdoutC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, rStdout)
		stdoutC <- buf.String()
	}()

	if g.overwriteResolver {
		_ = os.Remove(g.gqlgenConfig.Resolver.Filename)
	}

	if g.overwriteServer {
		_ = os.Remove(g.serverFileName)
	}
	_, err = os.Stat(g.serverFileName)
	if os.IsNotExist(err) {
		gqcmd.GenerateGraphServer(g.gqlgenConfig, g.serverFileName)
	} else {
		var options []api.Option
		if g.stubFileName != "" {
			options = append(options, api.AddPlugin(stubgen.New(g.stubFileName, "Stub")))
		}

		err = api.Generate(g.gqlgenConfig, options...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}

	wStderr.Close()
	wStdout.Close()

	os.Stderr = oldStderr
	os.Stdout = oldStdout

	g.Response.File = append(g.Response.File, &compiler.CodeGeneratorResponse_File{
		Name:    "gqlgen_output.log",
		Content: <-stderrC + <-stdoutC,
	})
}
