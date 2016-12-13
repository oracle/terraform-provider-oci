package baremetal

import (
	"errors"
	"reflect"
)

// Namespace is the top level organizational level of the object store
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/objectstorage/20160918/methods/GetNamespace

type Namespace string

// toBeFilled must be a slice of bytes
func (g *Namespace) Unmarshal(b []byte, toBeFilled interface{}) error {
	rv := reflect.ValueOf(toBeFilled)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("Value passed to unmarshal is not a pointer")
	}
	rv.Elem().SetString(string(b))
	return nil
}

// GetNamespace fetches the current user's namespace
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/objectstorage/20160918/methods/GetNamespace
func (c *Client) GetNamespace() (name *Namespace, e error) {
	var opts interface{}
	var required interface{}
	details := &requestDetails{
		ids:      urlParts{resourceNamespaces},
		optional: opts,
		required: required,
	}

	var response *requestResponse
	if response, e = c.objectStorageApi.getRequest(details); e != nil {
		return
	}

	name = new(Namespace)
	e = response.unmarshal(name)
	return
}
