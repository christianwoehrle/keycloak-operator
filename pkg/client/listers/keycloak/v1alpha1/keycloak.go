// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KeycloakLister helps list Keycloaks.
// All objects returned here must be treated as read-only.
type KeycloakLister interface {
	// List lists all Keycloaks in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Keycloak, err error)
	// Keycloaks returns an object that can list and get Keycloaks.
	Keycloaks(namespace string) KeycloakNamespaceLister
	KeycloakListerExpansion
}

// keycloakLister implements the KeycloakLister interface.
type keycloakLister struct {
	indexer cache.Indexer
}

// NewKeycloakLister returns a new KeycloakLister.
func NewKeycloakLister(indexer cache.Indexer) KeycloakLister {
	return &keycloakLister{indexer: indexer}
}

// List lists all Keycloaks in the indexer.
func (s *keycloakLister) List(selector labels.Selector) (ret []*v1alpha1.Keycloak, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Keycloak))
	})
	return ret, err
}

// Keycloaks returns an object that can list and get Keycloaks.
func (s *keycloakLister) Keycloaks(namespace string) KeycloakNamespaceLister {
	return keycloakNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KeycloakNamespaceLister helps list and get Keycloaks.
// All objects returned here must be treated as read-only.
type KeycloakNamespaceLister interface {
	// List lists all Keycloaks in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Keycloak, err error)
	// Get retrieves the Keycloak from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Keycloak, error)
	KeycloakNamespaceListerExpansion
}

// keycloakNamespaceLister implements the KeycloakNamespaceLister
// interface.
type keycloakNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Keycloaks in the indexer for a given namespace.
func (s keycloakNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Keycloak, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Keycloak))
	})
	return ret, err
}

// Get retrieves the Keycloak from the indexer for a given namespace and name.
func (s keycloakNamespaceLister) Get(name string) (*v1alpha1.Keycloak, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("keycloak"), name)
	}
	return obj.(*v1alpha1.Keycloak), nil
}
