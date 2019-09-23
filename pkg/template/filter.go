package template

import "k8s.io/apimachinery/pkg/runtime"

var (
	// RetainNamespaces a func to retain only namespaces
	RetainNamespaces FilterFunc = func(obj runtime.Object) bool {
		gvk := obj.GetObjectKind().GroupVersionKind()
		return gvk.Kind == "Namespace"
	}

	// RetainAllButNamespaces a func to retain all but namespaces
	RetainAllButNamespaces FilterFunc = func(obj runtime.Object) bool {
		gvk := obj.GetObjectKind().GroupVersionKind()
		return gvk.Kind != "Namespace"
	}
)

// FilterFunc a function to retain an object or not
type FilterFunc func(runtime.Object) bool

// Filter filters the given objs to return only those matching the given filters (if any)
func Filter(objs []runtime.Object, filters ...FilterFunc) []runtime.Object {
	result := make([]runtime.Object, 0, len(objs))
loop:
	for _, obj := range objs {
		for _, filter := range filters {
			if !filter(obj) {
				continue loop
			}

		}
		result = append(result, obj)
	}
	return result
}
