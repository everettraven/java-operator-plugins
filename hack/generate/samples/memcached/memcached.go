package memcached

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v3/pkg/cli"
	kbtestutils "sigs.k8s.io/kubebuilder/v3/test/e2e/utils"
)

type Memcached struct {
	kbtestutils.TestContext
}

// Generate will generate a Memcached sample Java operator
func (m *Memcached) Generate(cli *cli.CLI, testdir string) error {
	// setup the testdata directory
	dir := filepath.Join(testdir, "quarkus-memcached-operator")
	err := prepareSample(dir)
	if err != nil {
		log.Fatalf("encountered an error preparing the sample directory `%s`: %w", dir, err)
	}

	err = generateInit(cli)
	if err != nil {
		log.Fatalf("encountered an error running the `init` subcommand: %w", err)
	}

	err = generateApi(cli)
	if err != nil {
		log.Fatalf("encountered an error running the `create api` subcommand: %w", err)
	}
	return nil
}

// prepareSample will remove and initialize a new directory for generating a sample
func prepareSample(dir string) error {
	if err := os.RemoveAll(dir); err != nil {
		return fmt.Errorf("encountered an error removing the test directory `%s`: %w", dir, err)
	}

	if err := os.MkdirAll(dir, 0777); err != nil {
		return fmt.Errorf("encountered an error removing the test directory `%s`: %w", dir, err)
	}

	return nil
}

// generateInit will run the `init` subcommand for scaffolding a memcached operator
func generateInit(cli *cli.CLI) error {
	args := []string{
		"cli",
		"init",
		"--plugins",
		"quarkus",
		"--domain",
		"example.com",
		"--project-name",
		"memcached-quarkus-operator",
	}

	// set the args to be used with the cli for scaffolding
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = args

	err := cli.Run()
	if err != nil {
		return fmt.Errorf("encountered an error when scaffolding using the `init` subcommand: %w", err)
	}

	return nil
}

// generateApi will run the `api` subcommand for scaffolding a memcached operator
func generateApi(cli *cli.CLI) error {
	args := []string{
		"cli",
		"create",
		"api",
		"--plugins",
		"quarkus",
		"--group",
		"cache",
		"--version",
		"v1",
		"--kind",
		"Memcached",
	}

	// set the args to be used with the cli for scaffolding
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = args

	err := cli.Run()
	if err != nil {
		return fmt.Errorf("encountered an error when scaffolding using the `init` subcommand: %w", err)
	}

	return nil
}
