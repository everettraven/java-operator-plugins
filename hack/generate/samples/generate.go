package main

import (
	"log"

	"github.com/operator-framework/java-operator-plugins/hack/generate/samples/memcached"
	quarkusv1 "github.com/operator-framework/java-operator-plugins/pkg/quarkus/v1alpha"
	"sigs.k8s.io/kubebuilder/v3/pkg/cli"
	cfgv3 "sigs.k8s.io/kubebuilder/v3/pkg/config/v3"
)

func main() {

	testdir := "testdata"

	// cli for generation
	c, err := cli.New(
		cli.WithCommandName("java-sample-cli"),
		cli.WithVersion("v0.0.0"),
		cli.WithPlugins(
			&quarkusv1.Plugin{},
		),
		cli.WithDefaultPlugins(cfgv3.Version, quarkusv1.Plugin{}),
		cli.WithDefaultProjectVersion(cfgv3.Version),
	)
	if err != nil {
		log.Fatal("encountered an error creating a cli for scaffolding: ", err)
	}

	memcached := &memcached.Memcached{}
	err = memcached.Generate(c, testdir)

	if err != nil {
		log.Fatal("encountered an error generating the memcached sample:", err)
	}
}
