// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/isurulucky/k8sday/03-advanced-kubernetes/pkg/apis/demo/v1alpha1"
	"github.com/isurulucky/k8sday/03-advanced-kubernetes/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type DemoV1alpha1Interface interface {
	RESTClient() rest.Interface
	HellosGetter
}

// DemoV1alpha1Client is used to interact with features provided by the demo.wso2.com group.
type DemoV1alpha1Client struct {
	restClient rest.Interface
}

func (c *DemoV1alpha1Client) Hellos(namespace string) HelloInterface {
	return newHellos(c, namespace)
}

// NewForConfig creates a new DemoV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*DemoV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &DemoV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new DemoV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *DemoV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new DemoV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *DemoV1alpha1Client {
	return &DemoV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *DemoV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
