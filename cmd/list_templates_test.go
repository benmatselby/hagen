package cmd_test

import (
	"bytes"
	"testing"

	"github.com/benmatselby/hagen/cmd"
	"github.com/spf13/viper"
)

func TestTemplatesCommand_NoTemplates(t *testing.T) {
	viper.Set("templates", map[string]interface{}{})
	buf := new(bytes.Buffer)
	cmd := cmd.NewListTemplatesCommand()
	cmd.SetOut(buf)
	cmd.SetArgs([]string{})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "No templates found in config file.\n"
	if buf.String() != expected {
		t.Errorf("expected %q, got %q", expected, buf.String())
	}
}

func TestTemplatesCommand_ListNames(t *testing.T) {
	templates := map[string]interface{}{
		"foo": map[string]interface{}{"query": "bar"},
		"baz": map[string]interface{}{"query": "qux"},
	}
	viper.Set("templates", templates)
	buf := new(bytes.Buffer)
	cmd := cmd.NewListTemplatesCommand()
	cmd.SetOut(buf)
	cmd.SetArgs([]string{})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()
	if output != "baz\nfoo\n" && output != "foo\nbaz\n" {
		t.Errorf("unexpected output: %q", output)
	}
}

func TestTemplatesCommand_Verbose(t *testing.T) {
	templates := map[string]interface{}{
		"foo": map[string]interface{}{"query": "bar"},
		"baz": map[string]interface{}{"query": "qux"},
	}
	viper.Set("templates", templates)
	buf := new(bytes.Buffer)
	cmd := cmd.NewListTemplatesCommand()
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"-v"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()
	if output != "baz: qux\nfoo: bar\n" && output != "foo: bar\nbaz: qux\n" {
		t.Errorf("unexpected output: %q", output)
	}
}
