package containers

import (
	"context"
	"fmt"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
)

//type Image struct {
//	context   string
//	imageName string
//}
type container struct {
	imageName string
	context   string
	withsnap  bool
	snapname  string
}

func connectContained() (*containerd.Client, error) {
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		return client, err
	}
	return client, nil
}

// ListImages list existed Images
func ListImages(c string) error {
	// create a new client connected to the default socket path for containerd
	client, err := connectContained()

	if err != nil {
		return err
	}
	defer client.Close()

	// use a context with an "c" namespace
	ctx := namespaces.WithNamespace(context.Background(), c)
	list, err := client.ListImages(ctx)

	for _, i := range list {
		fmt.Printf("%s\n", i.Name())
	}

	if err != nil {
		return err
	}

	return nil
}

// ListContainers list existed containers inside "c" context
func ListContainers(c container) error {

	client, err := connectContained()

	if err != nil {
		return err
	}
	defer client.Close()

	// create a new context with an "example" namespace
	ctx := namespaces.WithNamespace(context.Background(), c.context)

	list, err := client.Containers(ctx)
	fmt.Printf("%v\n", list)
	for _, i := range list {
		fmt.Printf("%T\n", i.Info)
	}

	if err != nil {
		return err
	}
	return nil

}
