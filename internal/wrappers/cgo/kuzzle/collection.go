// Copyright 2015-2017 Kuzzle
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzlesdk.h"
	#include "sdk_wrappers_internal.h"
*/
import "C"
import (
	"encoding/json"
	"unsafe"

	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

// map which stores instances to keep references in case the gc passes
var collectionInstances map[interface{}]bool

// register new instance to the instances map
func registerCollection(instance interface{}) {
	collectionInstances[instance] = true
}

// unregister an instance from the instances map
//export unregisterCollection
func unregisterCollection(d *C.collection) {
	delete(collectionInstances, (*collection.Collection)(d.instance))
}

// Allocates memory
//export kuzzle_new_collection
func kuzzle_new_collection(c *C.collection, k *C.kuzzle) {
	kuz := (*kuzzle.Kuzzle)(k.instance)
	col := collection.NewCollection(kuz)

	if collectionInstances == nil {
		collectionInstances = make(map[interface{}]bool)
	}

	c.instance = unsafe.Pointer(col)
	c.kuzzle = k

	registerCollection(c)
}

//export kuzzle_collection_create
func kuzzle_collection_create(c *C.collection, index *C.char, col *C.char, options *C.query_options) *C.void_result {
	err := (*collection.Collection)(c.instance).Create(C.GoString(index), C.GoString(col), SetQueryOptions(options))
	return goToCVoidResult(err)
}

//export kuzzle_collection_truncate
func kuzzle_collection_truncate(c *C.collection, index *C.char, col *C.char, options *C.query_options) *C.void_result {
	err := (*collection.Collection)(c.instance).Truncate(C.GoString(index), C.GoString(col), SetQueryOptions(options))
	return goToCVoidResult(err)
}

//export kuzzle_collection_exists
func kuzzle_collection_exists(c *C.collection, index *C.char, col *C.char, options *C.query_options) *C.bool_result {
	res, err := (*collection.Collection)(c.instance).Exists(C.GoString(index), C.GoString(col), SetQueryOptions(options))
	return goToCBoolResult(res, err)
}

//export kuzzle_collection_list
func kuzzle_collection_list(c *C.collection, index *C.char, options *C.query_options) *C.string_result {
	res, err := (*collection.Collection)(c.instance).List(C.GoString(index), SetQueryOptions(options))
	var stringResult string
	json.Unmarshal(res, &stringResult)
	return goToCStringResult(&stringResult, err)
}

// Mapping

//export kuzzle_collection_get_mapping
func kuzzle_collection_get_mapping(c *C.collection, index *C.char, col *C.char, options *C.query_options) *C.string_result {
	res, err := (*collection.Collection)(c.instance).GetMapping(C.GoString(index), C.GoString(col), SetQueryOptions(options))
	var stringResult string
	stringResult = string(res)
	return goToCStringResult(&stringResult, err)
}

//export kuzzle_collection_update_mapping
func kuzzle_collection_update_mapping(c *C.collection, index *C.char, col *C.char, body *C.char, options *C.query_options) *C.void_result {
	newBody, _ := json.Marshal(body)
	err := (*collection.Collection)(c.instance).UpdateMapping(C.GoString(index), C.GoString(col), newBody, SetQueryOptions(options))
	return goToCVoidResult(err)
}

// Specifications

//export kuzzle_collection_delete_specifications
func kuzzle_collection_delete_specifications(c *C.collection, index *C.char, col *C.char, options *C.query_options) *C.void_result {
	err := (*collection.Collection)(c.instance).DeleteSpecifications(C.GoString(index), C.GoString(col), SetQueryOptions(options))
	return goToCVoidResult(err)
}

//export kuzzle_collection_get_specifications
func kuzzle_collection_get_specifications(c *C.collection, index *C.char, col *C.char, options *C.query_options) *C.string_result {
	res, err := (*collection.Collection)(c.instance).GetSpecifications(C.GoString(index), C.GoString(col), SetQueryOptions(options))
	var stringResult string
	stringResult = string(res)
	return goToCStringResult(&stringResult, err)
}

//export kuzzle_collection_search_specifications
func kuzzle_collection_search_specifications(c *C.collection, options *C.query_options) *C.search_result {
	res, err := (*collection.Collection)(c.instance).SearchSpecifications(SetQueryOptions(options))
	return goToCSearchResult(res, err)
}

//export kuzzle_collection_update_specifications
func kuzzle_collection_update_specifications(c *C.collection, index *C.char, col *C.char, body *C.char, options *C.query_options) *C.string_result {
	newBody, _ := json.Marshal(body)
	res, err := (*collection.Collection)(c.instance).UpdateSpecifications(C.GoString(index), C.GoString(col), newBody, SetQueryOptions(options))
	var stringResult string
	stringResult = string(res)
	return goToCStringResult(&stringResult, err)
}

//export kuzzle_collection_validate_specifications
func kuzzle_collection_validate_specifications(c *C.collection, body *C.char, options *C.query_options) *C.bool_result {
	newBody, _ := json.Marshal(body)
	res, err := (*collection.Collection)(c.instance).ValidateSpecifications(newBody, SetQueryOptions(options))
	return goToCBoolResult(res, err)
}
