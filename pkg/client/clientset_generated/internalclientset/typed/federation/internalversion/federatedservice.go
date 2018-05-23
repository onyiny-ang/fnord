/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package internalversion

import (
	federation "github.com/kubernetes-sigs/federation-v2/pkg/apis/federation"
	scheme "github.com/kubernetes-sigs/federation-v2/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FederatedServicesGetter has a method to return a FederatedServiceInterface.
// A group's client should implement this interface.
type FederatedServicesGetter interface {
	FederatedServices(namespace string) FederatedServiceInterface
}

// FederatedServiceInterface has methods to work with FederatedService resources.
type FederatedServiceInterface interface {
	Create(*federation.FederatedService) (*federation.FederatedService, error)
	Update(*federation.FederatedService) (*federation.FederatedService, error)
	UpdateStatus(*federation.FederatedService) (*federation.FederatedService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*federation.FederatedService, error)
	List(opts v1.ListOptions) (*federation.FederatedServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *federation.FederatedService, err error)
	FederatedServiceExpansion
}

// federatedServices implements FederatedServiceInterface
type federatedServices struct {
	client rest.Interface
	ns     string
}

// newFederatedServices returns a FederatedServices
func newFederatedServices(c *FederationClient, namespace string) *federatedServices {
	return &federatedServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedService, and returns the corresponding federatedService object, and an error if there is any.
func (c *federatedServices) Get(name string, options v1.GetOptions) (result *federation.FederatedService, err error) {
	result = &federation.FederatedService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedServices that match those selectors.
func (c *federatedServices) List(opts v1.ListOptions) (result *federation.FederatedServiceList, err error) {
	result = &federation.FederatedServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedServices.
func (c *federatedServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federatedservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a federatedService and creates it.  Returns the server's representation of the federatedService, and an error, if there is any.
func (c *federatedServices) Create(federatedService *federation.FederatedService) (result *federation.FederatedService, err error) {
	result = &federation.FederatedService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federatedservices").
		Body(federatedService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a federatedService and updates it. Returns the server's representation of the federatedService, and an error, if there is any.
func (c *federatedServices) Update(federatedService *federation.FederatedService) (result *federation.FederatedService, err error) {
	result = &federation.FederatedService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedservices").
		Name(federatedService.Name).
		Body(federatedService).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *federatedServices) UpdateStatus(federatedService *federation.FederatedService) (result *federation.FederatedService, err error) {
	result = &federation.FederatedService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedservices").
		Name(federatedService.Name).
		SubResource("status").
		Body(federatedService).
		Do().
		Into(result)
	return
}

// Delete takes name of the federatedService and deletes it. Returns an error if one occurs.
func (c *federatedServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched federatedService.
func (c *federatedServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *federation.FederatedService, err error) {
	result = &federation.FederatedService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federatedservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
