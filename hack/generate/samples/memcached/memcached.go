package memcached

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v3/pkg/cli"
	kbutil "sigs.k8s.io/kubebuilder/v3/pkg/plugin/util"
)

type Memcached struct{}

// Generate will generate a Memcached sample Java operator
func (m *Memcached) Generate(cli *cli.CLI, testdir string) error {
	// setup the testdata directory
	log.Print("Setting up the testdata directory for quarkus-memcached-operator")
	dir := filepath.Join(testdir, "quarkus", "quarkus-memcached-operator")
	err := prepareSample(dir)
	if err != nil {
		log.Fatalf("encountered an error preparing the sample directory `%s`: %w", dir, err)
	}

	// Scaffolding
	// -------------------

	log.Print("running the `init` subcommand for quarkus-memcached-operator")
	err = generateInit(cli, dir)
	if err != nil {
		log.Fatalf("encountered an error running the `init` subcommand: %w", err)
	}

	log.Print("running the `create api` subcommand for quarkus-memcached-operator")
	err = generateApi(cli, dir)
	if err != nil {
		log.Fatalf("encountered an error running the `create api` subcommand: %w", err)
	}

	// -------------------

	// Implementation
	// -------------------

	log.Print("implementing reconcile helpers")
	err = implementReconcileHelpers(dir)
	if err != nil {
		return fmt.Errorf("encountered an error implementing reconcile helper functions: %w", err)
	}

	log.Print("implementing reconcile")
	err = implementReconcile(dir)
	if err != nil {
		return fmt.Errorf("encountered an error implementing reconcile function: %w", err)
	}

	log.Print("implementing spec")
	err = implementSpec(dir)
	if err != nil {
		return fmt.Errorf("encountered an error implementing spec: %w", err)
	}

	log.Print("implementing status")
	err = implementStatus(dir)
	if err != nil {
		return fmt.Errorf("encountered an error implementing status: %w", err)
	}

	log.Print("adding reconcile dependencies")
	err = populateReconcileDependencies(dir)
	if err != nil {
		return fmt.Errorf("encountered an error adding reconcile dependencies: %w", err)
	}

	log.Print("adding status dependencies")
	err = populateStatusDependencies(dir)
	if err != nil {
		return fmt.Errorf("encountered an error adding status dependencies: %w", err)
	}

	log.Print("adding pom dependencies")
	err = populatePomDependencies(dir)
	if err != nil {
		return fmt.Errorf("encountered an error adding pom dependencies: %w", err)
	}

	// -------------------

	// Post implementation
	// -------------------
	log.Printf("Running post implementation steps")
	err = postImplementation(dir)
	if err != nil {
		return fmt.Errorf("encountered an error running the post implementation steps: %w", err)
	}
	// -------------------

	return nil
}

// prepareSample will remove and initialize a new directory for generating a sample
func prepareSample(dir string) error {
	log.Printf("Removing directory `%s` if it exists", dir)
	if err := os.RemoveAll(dir); err != nil {
		return fmt.Errorf("encountered an error removing the test directory `%s`: %w", dir, err)
	}

	log.Printf("Recreating directory `%s`", dir)
	if err := os.MkdirAll(dir, 0777); err != nil {
		return fmt.Errorf("encountered an error removing the test directory `%s`: %w", dir, err)
	}

	return nil
}

// generateInit will run the `init` subcommand for scaffolding a memcached operator
func generateInit(cli *cli.CLI, dir string) error {
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

	log.Printf("Setting os.Args to: %v", args)
	// set the args to be used with the cli for scaffolding
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = args

	log.Printf("Changing directory to: %s", dir)
	// change execution directory
	oldDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("encountered an error getting current working directory: %w", err)
	}

	err = os.Chdir(dir)
	if err != nil {
		return fmt.Errorf("encountered an error changing directory to `%s`: %w", dir, err)
	}

	// change execution directory back after function exits
	defer func() {
		err := os.Chdir(oldDir)
		if err != nil {
			log.Printf("encountered an error changing directory back to the previous working directory, this may cause problems: %w", err)
		}
	}()

	err = cli.Run()
	if err != nil {
		return fmt.Errorf("encountered an error when scaffolding using the `init` subcommand: %w", err)
	}

	return nil
}

// generateApi will run the `api` subcommand for scaffolding a memcached operator
func generateApi(cli *cli.CLI, dir string) error {
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

	log.Printf("Setting os.Args to: %v", args)
	// set the args to be used with the cli for scaffolding
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = args

	log.Printf("Changing directory to: %s", dir)
	// change execution directory
	oldDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("encountered an error getting current working directory: %w", err)
	}

	err = os.Chdir(dir)
	if err != nil {
		return fmt.Errorf("encountered an error changing directory to `%s`: %w", dir, err)
	}

	// change execution directory back after function exits
	defer func() {
		err := os.Chdir(oldDir)
		if err != nil {
			fmt.Printf("encountered an error changing directory back to the previous working directory, this may cause problems: %w", err)
		}
	}()

	err = cli.Run()
	if err != nil {
		return fmt.Errorf("encountered an error when scaffolding using the `init` subcommand: %w", err)
	}

	return nil
}

// implementReconcileHelpers inserts the reconciler helper function code to the MemcachedReconciler.java file
func implementReconcileHelpers(dir string) error {
	target := "// TODO Fill in the rest of the reconciler"

	code := `
	private Map<String, String> labelsForMemcached(Memcached m) {
		Map<String, String> labels = new HashMap<>();
		labels.put("app", "memcached");
		labels.put("memcached_cr", m.getMetadata().getName());
		return labels;
	}
	
	private Deployment createMemcachedDeployment(Memcached m) {
		return new DeploymentBuilder()
			.withMetadata(
				new ObjectMetaBuilder()
					.withName(m.getMetadata().getName())
					.withNamespace(m.getMetadata().getNamespace())
					.withOwnerReferences(
						new OwnerReferenceBuilder()
							.withApiVersion("v1")
							.withKind("Memcached")
							.withName(m.getMetadata().getName())
							.withUid(m.getMetadata().getUid())
							.build())
					.build())
			.withSpec(
				new DeploymentSpecBuilder()
					.withReplicas(m.getSpec().getSize())
					.withSelector(
						new LabelSelectorBuilder().withMatchLabels(labelsForMemcached(m)).build())
					.withTemplate(
						new PodTemplateSpecBuilder()
							.withMetadata(
								new ObjectMetaBuilder().withLabels(labelsForMemcached(m)).build())
							.withSpec(
								new PodSpecBuilder()
									.withContainers(
										new ContainerBuilder()
											.withImage("memcached:1.4.36-alpine")
											.withName("memcached")
											.withCommand("memcached", "-m=64", "-o", "modern", "-v")
											.withPorts(
												new ContainerPortBuilder()
													.withContainerPort(11211)
													.withName("memcached")
													.build())
											.build())
									.build())
							.build())
					.build())
			.build();
	}
	`

	return kbutil.InsertCode(filepath.Join(dir, "src", "main", "java", "com", "example", "MemcachedReconciler.java"), target, code)
}

// implementReconcile implements the reconcile function in the MemcachedReconciler.java file
func implementReconcile(dir string) error {
	target := "// TODO: fill in logic"
	code := `
		Deployment deployment = client.apps()
			.deployments()
			.inNamespace(resource.getMetadata().getNamespace())
			.withName(resource.getMetadata().getName())
			.get();

		if (deployment == null) {
			Deployment newDeployment = createMemcachedDeployment(resource);
			client.apps().deployments().create(newDeployment);
			return UpdateControl.noUpdate();
		}

		int currentReplicas = deployment.getSpec().getReplicas();
		int requiredReplicas = resource.getSpec().getSize();

		if (currentReplicas != requiredReplicas) {
			deployment.getSpec().setReplicas(requiredReplicas);
			client.apps().deployments().createOrReplace(deployment);
			return UpdateControl.noUpdate();
		}

		List<Pod> pods = client.pods()
		.inNamespace(resource.getMetadata().getNamespace())
		.withLabels(labelsForMemcached(resource))
		.list()
		.getItems();

		List<String> podNames =
		pods.stream().map(p -> p.getMetadata().getName()).collect(Collectors.toList());


		if (resource.getStatus() == null
			|| !CollectionUtils.isEqualCollection(podNames, resource.getStatus().getNodes())) {
			if (resource.getStatus() == null) resource.setStatus(new MemcachedStatus());
			resource.getStatus().setNodes(podNames);
			return UpdateControl.updateResource(resource);
		}`

	return kbutil.InsertCode(filepath.Join(dir, "src", "main", "java", "com", "example", "MemcachedReconciler.java"), target, code)
}

// implementSpec populates the MemcachedSpec.java file with the `size` field
func implementSpec(dir string) error {
	target := "// Add Spec information here"
	code := `
	// Size is the size of the memcached deployment
    private Integer size;

    public Integer getSize() {
        return size;
    }

    public void setSize(Integer size) {
        this.size = size;
    }
	`

	return kbutil.InsertCode(filepath.Join(dir, "src", "main", "java", "com", "example", "MemcachedSpec.java"), target, code)
}

// implementStatus populates the MemcachedStatus.java file with the proper status fields
func implementStatus(dir string) error {
	target := "// Add Status information here"
	code := `
	// Nodes are the names of the memcached pods
    private List<String> nodes;

    public List<String> getNodes() {
        if (nodes == null) {
            nodes = new ArrayList<>();
        }
        return nodes;
    }

    public void setNodes(List<String> nodes) {
        this.nodes = nodes;
    }
	`
	return kbutil.InsertCode(filepath.Join(dir, "src", "main", "java", "com", "example", "MemcachedStatus.java"), target, code)
}

func populateReconcileDependencies(dir string) error {
	target := `import io.javaoperatorsdk.operator.api.reconciler.UpdateControl;`
	code := `

import io.fabric8.kubernetes.api.model.ContainerBuilder;
import io.fabric8.kubernetes.api.model.ContainerPortBuilder;
import io.fabric8.kubernetes.api.model.LabelSelectorBuilder;
import io.fabric8.kubernetes.api.model.ObjectMetaBuilder;
import io.fabric8.kubernetes.api.model.OwnerReferenceBuilder;
import io.fabric8.kubernetes.api.model.Pod;
import io.fabric8.kubernetes.api.model.PodSpecBuilder;
import io.fabric8.kubernetes.api.model.PodTemplateSpecBuilder;
import io.fabric8.kubernetes.api.model.apps.Deployment;
import io.fabric8.kubernetes.api.model.apps.DeploymentBuilder;
import io.fabric8.kubernetes.api.model.apps.DeploymentSpecBuilder;
import org.apache.commons.collections.CollectionUtils;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;
`
	return kbutil.InsertCode(filepath.Join(dir, "src", "main", "java", "com", "example", "MemcachedReconciler.java"), target, code)

}

func populateStatusDependencies(dir string) error {
	target := `package com.example;`
	code := `

import java.util.ArrayList;
import java.util.List;
`
	return kbutil.InsertCode(filepath.Join(dir, "src", "main", "java", "com", "example", "MemcachedStatus.java"), target, code)
}

func populatePomDependencies(dir string) error {
	target := `    <dependency>
      <groupId>io.quarkus</groupId>
      <artifactId>quarkus-micrometer-registry-prometheus</artifactId>
      <version>${quarkus.version}</version>
    </dependency>`
	code := `
    <dependency>
      <groupId>commons-collections</groupId>
      <artifactId>commons-collections</artifactId>
      <version>3.2.2</version>
    </dependency>
`
	return kbutil.InsertCode(filepath.Join(dir, "pom.xml"), target, code)
}

func postImplementation(dir string) error {
	log.Printf("Changing directory to: %s", dir)
	// change execution directory
	oldDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("encountered an error getting current working directory: %w", err)
	}

	err = os.Chdir(dir)
	if err != nil {
		return fmt.Errorf("encountered an error changing directory to `%s`: %w", dir, err)
	}

	// change execution directory back after function exits
	defer func() {
		err := os.Chdir(oldDir)
		if err != nil {
			fmt.Printf("encountered an error changing directory back to the previous working directory, this may cause problems: %w", err)
		}
	}()

	log.Printf("Running `mvn clean install`")
	out, err := exec.Command("mvn", "clean", "install", "-X").CombinedOutput()
	if err != nil {
		return fmt.Errorf("encountered an error running `mvn clean install`: %w | `mvn clean install` output: %s", err, out)
	}

	log.Printf("Running `make bundle`")
	out, err = exec.Command("make", "bundle").CombinedOutput()
	if err != nil {
		return fmt.Errorf("encountered an error running `make bundle`: %w | `make bundle` output: %s", err, out)
	}

	return nil
}
