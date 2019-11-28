package v1

import (
	"github.com/walk1ng/k8s-controller-crd/pkg/apis/samplecrd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GroupVersion is the identifier for the API which includes
// the name of the group and the version of the API.
var SchemaGroupVersion = schema.GroupVersion{
	Group:   samplecrd.GroupName,
	Version: samplecrd.Version,
}

func Resource(resource string) schema.GroupResource {
	return SchemaGroupVersion.WithResource(resource).GroupResource()
}

func Kind(kind string) schema.GroupKind {
	return SchemaGroupVersion.WithKind(kind).GroupKind()
}

// addKnownTypes adds ours types to the API scheme by registering
// Bot and BotList
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemaGroupVersion,
		&Bot{},
		&BotList{},
	)

	// register the type in the scheme
	metav1.AddToGroupVersion(scheme, SchemaGroupVersion)
	return nil
}
