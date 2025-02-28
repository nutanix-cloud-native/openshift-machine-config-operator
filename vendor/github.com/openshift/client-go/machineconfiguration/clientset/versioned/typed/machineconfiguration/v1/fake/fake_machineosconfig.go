// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/openshift/api/machineconfiguration/v1"
	machineconfigurationv1 "github.com/openshift/client-go/machineconfiguration/applyconfigurations/machineconfiguration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMachineOSConfigs implements MachineOSConfigInterface
type FakeMachineOSConfigs struct {
	Fake *FakeMachineconfigurationV1
}

var machineosconfigsResource = v1.SchemeGroupVersion.WithResource("machineosconfigs")

var machineosconfigsKind = v1.SchemeGroupVersion.WithKind("MachineOSConfig")

// Get takes name of the machineOSConfig, and returns the corresponding machineOSConfig object, and an error if there is any.
func (c *FakeMachineOSConfigs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.MachineOSConfig, err error) {
	emptyResult := &v1.MachineOSConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(machineosconfigsResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.MachineOSConfig), err
}

// List takes label and field selectors, and returns the list of MachineOSConfigs that match those selectors.
func (c *FakeMachineOSConfigs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.MachineOSConfigList, err error) {
	emptyResult := &v1.MachineOSConfigList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(machineosconfigsResource, machineosconfigsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.MachineOSConfigList{ListMeta: obj.(*v1.MachineOSConfigList).ListMeta}
	for _, item := range obj.(*v1.MachineOSConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested machineOSConfigs.
func (c *FakeMachineOSConfigs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(machineosconfigsResource, opts))
}

// Create takes the representation of a machineOSConfig and creates it.  Returns the server's representation of the machineOSConfig, and an error, if there is any.
func (c *FakeMachineOSConfigs) Create(ctx context.Context, machineOSConfig *v1.MachineOSConfig, opts metav1.CreateOptions) (result *v1.MachineOSConfig, err error) {
	emptyResult := &v1.MachineOSConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(machineosconfigsResource, machineOSConfig, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.MachineOSConfig), err
}

// Update takes the representation of a machineOSConfig and updates it. Returns the server's representation of the machineOSConfig, and an error, if there is any.
func (c *FakeMachineOSConfigs) Update(ctx context.Context, machineOSConfig *v1.MachineOSConfig, opts metav1.UpdateOptions) (result *v1.MachineOSConfig, err error) {
	emptyResult := &v1.MachineOSConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(machineosconfigsResource, machineOSConfig, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.MachineOSConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMachineOSConfigs) UpdateStatus(ctx context.Context, machineOSConfig *v1.MachineOSConfig, opts metav1.UpdateOptions) (result *v1.MachineOSConfig, err error) {
	emptyResult := &v1.MachineOSConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(machineosconfigsResource, "status", machineOSConfig, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.MachineOSConfig), err
}

// Delete takes name of the machineOSConfig and deletes it. Returns an error if one occurs.
func (c *FakeMachineOSConfigs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(machineosconfigsResource, name, opts), &v1.MachineOSConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMachineOSConfigs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(machineosconfigsResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.MachineOSConfigList{})
	return err
}

// Patch applies the patch and returns the patched machineOSConfig.
func (c *FakeMachineOSConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.MachineOSConfig, err error) {
	emptyResult := &v1.MachineOSConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(machineosconfigsResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.MachineOSConfig), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied machineOSConfig.
func (c *FakeMachineOSConfigs) Apply(ctx context.Context, machineOSConfig *machineconfigurationv1.MachineOSConfigApplyConfiguration, opts metav1.ApplyOptions) (result *v1.MachineOSConfig, err error) {
	if machineOSConfig == nil {
		return nil, fmt.Errorf("machineOSConfig provided to Apply must not be nil")
	}
	data, err := json.Marshal(machineOSConfig)
	if err != nil {
		return nil, err
	}
	name := machineOSConfig.Name
	if name == nil {
		return nil, fmt.Errorf("machineOSConfig.Name must be provided to Apply")
	}
	emptyResult := &v1.MachineOSConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(machineosconfigsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.MachineOSConfig), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeMachineOSConfigs) ApplyStatus(ctx context.Context, machineOSConfig *machineconfigurationv1.MachineOSConfigApplyConfiguration, opts metav1.ApplyOptions) (result *v1.MachineOSConfig, err error) {
	if machineOSConfig == nil {
		return nil, fmt.Errorf("machineOSConfig provided to Apply must not be nil")
	}
	data, err := json.Marshal(machineOSConfig)
	if err != nil {
		return nil, err
	}
	name := machineOSConfig.Name
	if name == nil {
		return nil, fmt.Errorf("machineOSConfig.Name must be provided to Apply")
	}
	emptyResult := &v1.MachineOSConfig{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(machineosconfigsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.MachineOSConfig), err
}
