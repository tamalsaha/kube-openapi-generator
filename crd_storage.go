package main

import (
	"fmt"

	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
)

type REST struct {
	gvk  schema.GroupVersionKind
	obj  runtime.Object
	list runtime.Object
}

var _ rest.Getter = &REST{}
var _ rest.Lister = &REST{}
var _ rest.Deleter = &REST{}
var _ rest.GroupVersionKindProvider = &REST{}

func NewREST(gvk schema.GroupVersionKind, obj runtime.Object, list runtime.Object) *REST {
	return &REST{gvk, obj, list}
}

func (r *REST) New() runtime.Object {
	return r.obj
}

func (r *REST) GroupVersionKind(containingGV schema.GroupVersion) schema.GroupVersionKind {
	return r.gvk
}

func (r *REST) Get(ctx apirequest.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	fmt.Println(">>>>--- GET")
	return r.New(), nil
}

func (r *REST) NewList() runtime.Object {
	return r.list
}

func (r *REST) List(ctx apirequest.Context, options *metainternalversion.ListOptions) (runtime.Object, error) {
	fmt.Println(">>>>--- LIST")
	return r.NewList(), nil
}

func (r *REST) Delete(ctx apirequest.Context, name string) (runtime.Object, error) {
	fmt.Println(">>>>--- DELETE")
	return nil, nil
}
