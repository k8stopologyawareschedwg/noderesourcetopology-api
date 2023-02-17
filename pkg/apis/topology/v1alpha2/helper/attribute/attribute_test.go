/*
Copyright 2023 The Kubernetes Authors.

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

package attribute

import (
	"reflect"
	"testing"

	"github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha2"
)

func TestGet(t *testing.T) {
	testCases := []struct {
		name          string
		attrs         v1alpha2.AttributeList
		attrName      string
		expectedFound bool
		expectedValue string
	}{
		{
			name:     "empty collection and name",
			attrName: "",
		},
		{
			name:     "empty collection",
			attrName: "foobar",
		},
		{
			name: "missing in non-empty collection",
			attrs: v1alpha2.AttributeList{
				{
					Name:  "foo",
					Value: "1",
				},
				{
					Name:  "bar",
					Value: "2",
				},
			},
			attrName: "buz",
		},
		{
			name: "found in collection",
			attrs: v1alpha2.AttributeList{
				{
					Name:  "foo",
					Value: "1",
				},
				{
					Name:  "bar",
					Value: "2",
				},
				{
					Name:  "buz",
					Value: "3",
				},
			},
			attrName:      "bar",
			expectedFound: true,
			expectedValue: "2",
		},
		{
			name: "found in collection with duplicates",
			attrs: v1alpha2.AttributeList{
				{
					Name:  "foo",
					Value: "1",
				},
				{
					Name:  "bar",
					Value: "2",
				},
				{
					Name:  "buz",
					Value: "3",
				},
				{
					Name:  "bar",
					Value: "A2",
				},
			},
			attrName:      "bar",
			expectedFound: true,
			expectedValue: "2",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			val, ok := Get(tt.attrs, tt.attrName)
			if ok != tt.expectedFound {
				t.Errorf("%s presence got=%v expected=%v", tt.attrName, ok, tt.expectedFound)
			}
			expectedAttr := v1alpha2.AttributeInfo{
				Name:  tt.attrName,
				Value: tt.expectedValue,
			}
			if ok && val != expectedAttr {
				t.Errorf("found value for name=%v got=%+#v expected=+#%v", tt.attrName, val, expectedAttr)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	testCases := []struct {
		name            string
		attrs           v1alpha2.AttributeList
		attrInfo        v1alpha2.AttributeInfo
		expectedMissing bool // for special corner cases
	}{
		{
			name:            "empty collection, empty attribute",
			attrInfo:        v1alpha2.AttributeInfo{},
			expectedMissing: true,
		},
		{
			name: "empty collection",
			attrInfo: v1alpha2.AttributeInfo{
				Name:  "foo",
				Value: "42",
			},
		},
		{
			name: "update collection creating element",
			attrs: v1alpha2.AttributeList{
				{
					Name:  "bar",
					Value: "b2",
				},
				{
					Name:  "buz",
					Value: "u3",
				},
			},
			attrInfo: v1alpha2.AttributeInfo{
				Name:  "foo",
				Value: "42",
			},
		},
		{
			name: "update collection changing element in place",
			attrs: v1alpha2.AttributeList{
				{
					Name:  "foo",
					Value: "ff",
				},
				{
					Name:  "bar",
					Value: "b2",
				},
				{
					Name:  "buz",
					Value: "u3",
				},
			},
			attrInfo: v1alpha2.AttributeInfo{
				Name:  "foo",
				Value: "42",
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			origLen := len(tt.attrs)
			origKeys := extractKeys(tt.attrs, "")
			newAttrs := Insert(tt.attrs, tt.attrInfo)
			newLen := len(newAttrs)

			val, ok := Get(newAttrs, tt.attrInfo.Name)
			if ok {
				if val != tt.attrInfo {
					t.Errorf("wrong value in the updated AttributeList for name=%v got=%+#v expected=%+#v", tt.attrInfo.Name, val, tt.attrInfo)
				}
			} else if !tt.expectedMissing {
				t.Errorf("%s missing from updated AttributeList", tt.attrInfo.Name)
			}

			skipKey := ""
			if newLen > origLen {
				// a new item was created, and it must be the provided attribute.
				// Let's check what happened to everything else, expecting no changes.
				skipKey = tt.attrInfo.Name
			}
			newKeys := extractKeys(newAttrs, skipKey)
			if !reflect.DeepEqual(origKeys, newKeys) {
				t.Errorf("update changed the list ordering: origKeys=%v newKeys=%v", origKeys, newKeys)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	testCases := []struct {
		name          string
		existing      v1alpha2.AttributeList
		updated       v1alpha2.AttributeList
		expectedAttrs v1alpha2.AttributeList
	}{
		{
			name:          "empty collections",
			existing:      v1alpha2.AttributeList{},
			updated:       v1alpha2.AttributeList{},
			expectedAttrs: v1alpha2.AttributeList{},
		},
		{
			name: "merge collections, empty into non-empty",
			existing: v1alpha2.AttributeList{
				{
					Name:  "foo",
					Value: "ff",
				},
				{
					Name:  "bar",
					Value: "b2",
				},
				{
					Name:  "buz",
					Value: "u3",
				},
			},
			updated: v1alpha2.AttributeList{},
			expectedAttrs: v1alpha2.AttributeList{
				{
					Name:  "foo",
					Value: "ff",
				},
				{
					Name:  "bar",
					Value: "b2",
				},
				{
					Name:  "buz",
					Value: "u3",
				},
			},
		},
		{
			name:     "merge collections, empty into non-empty",
			existing: v1alpha2.AttributeList{},
			updated: v1alpha2.AttributeList{
				{
					Name:  "foo",
					Value: "ff",
				},
				{
					Name:  "bar",
					Value: "b2",
				},
				{
					Name:  "buz",
					Value: "u3",
				},
			},
			expectedAttrs: v1alpha2.AttributeList{
				{
					Name:  "foo",
					Value: "ff",
				},
				{
					Name:  "bar",
					Value: "b2",
				},
				{
					Name:  "buz",
					Value: "u3",
				},
			},
		},
		{
			name: "merge collections, non-empty into non-empty, no duplicates",
			existing: v1alpha2.AttributeList{
				{
					Name:  "foo",
					Value: "ff",
				},
				{
					Name:  "buz",
					Value: "u3",
				},
			},
			updated: v1alpha2.AttributeList{
				{
					Name:  "bar",
					Value: "b2",
				},
				{
					Name:  "abc",
					Value: "123",
				},
			},
			expectedAttrs: v1alpha2.AttributeList{
				{
					Name:  "foo",
					Value: "ff",
				},
				{
					Name:  "buz",
					Value: "u3",
				},
				{
					Name:  "bar",
					Value: "b2",
				},
				{
					Name:  "abc",
					Value: "123",
				},
			},
		},
		{
			name: "merge collections, non-empty into non-empty, duplicates",
			existing: v1alpha2.AttributeList{
				{
					Name:  "foo",
					Value: "ff",
				},
				{
					Name:  "buz",
					Value: "u3",
				},
			},
			updated: v1alpha2.AttributeList{
				{
					Name:  "buz",
					Value: "b2",
				},
				{
					Name:  "abc",
					Value: "123",
				},
			},
			expectedAttrs: v1alpha2.AttributeList{
				{
					Name:  "foo",
					Value: "ff",
				},
				{
					Name:  "buz",
					Value: "b2",
				},
				{
					Name:  "abc",
					Value: "123",
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			newAttrs := Merge(tt.existing, tt.updated)
			if !reflect.DeepEqual(newAttrs, tt.expectedAttrs) {
				t.Errorf("delete returned unexpected merged list: got=%v expected=%v", newAttrs, tt.expectedAttrs)
			}
		})
	}
}

func TestDeleteAll(t *testing.T) {
	testCases := []struct {
		name          string
		attrName      string
		attrs         v1alpha2.AttributeList
		expectedAttrs v1alpha2.AttributeList
	}{
		{
			name:          "empty collection, empty name",
			attrs:         v1alpha2.AttributeList{},
			expectedAttrs: v1alpha2.AttributeList{},
		},
		{
			name: "collection, missing name",
			attrs: v1alpha2.AttributeList{
				{
					Name:  "foo",
					Value: "ff",
				},
				{
					Name:  "bar",
					Value: "b2",
				},
				{
					Name:  "buz",
					Value: "u3",
				},
			},
			attrName: "missing",
			expectedAttrs: v1alpha2.AttributeList{
				{
					Name:  "foo",
					Value: "ff",
				},
				{
					Name:  "bar",
					Value: "b2",
				},
				{
					Name:  "buz",
					Value: "u3",
				},
			},
		},
		{
			name: "collection, removing attribute",
			attrs: v1alpha2.AttributeList{
				{
					Name:  "foo",
					Value: "ff",
				},
				{
					Name:  "bar",
					Value: "b2",
				},
				{
					Name:  "buz",
					Value: "u3",
				},
			},
			attrName: "bar",
			expectedAttrs: v1alpha2.AttributeList{
				{
					Name:  "foo",
					Value: "ff",
				},
				{
					Name:  "buz",
					Value: "u3",
				},
			},
		},
		{
			name: "collection, removing duplicate attribute",
			attrs: v1alpha2.AttributeList{
				{
					Name:  "foo",
					Value: "ff",
				},
				{
					Name:  "bar",
					Value: "b2",
				},
				{
					Name:  "buz",
					Value: "u3",
				},
				{
					Name:  "buz",
					Value: "pp",
				},
			},
			attrName: "buz",
			expectedAttrs: v1alpha2.AttributeList{
				{
					Name:  "foo",
					Value: "ff",
				},
				{
					Name:  "bar",
					Value: "b2",
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			newAttrs := DeleteAll(tt.attrs, tt.attrName)
			if !reflect.DeepEqual(newAttrs, tt.expectedAttrs) {
				t.Errorf("delete returned unexpected deleted list: got=%v expected=%v", newAttrs, tt.expectedAttrs)
			}
		})
	}
}

func extractKeys(attrs v1alpha2.AttributeList, skipName string) []string {
	keys := make([]string, 0, len(attrs))
	for _, attr := range attrs {
		if skipName != "" && attr.Name == skipName {
			continue
		}
		keys = append(keys, attr.Name)
	}
	return keys
}
