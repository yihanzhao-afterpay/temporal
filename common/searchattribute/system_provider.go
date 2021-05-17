// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package searchattribute

import (
	enumspb "go.temporal.io/api/enums/v1"
)

type (
	SystemProvider struct{}
)

func NewSystemProvider() *SystemProvider {
	return &SystemProvider{}
}

func (s *SystemProvider) GetSearchAttributes(_ string, _ bool) (NameTypeMap, error) {
	return NameTypeMap{}, nil
}

// FilterCustomOnly returns new search attributes map with only custom search attributes in it.
// TODO: Remove this after 1.10.0 release
func FilterCustomOnly(searchAttributes map[string]enumspb.IndexedValueType) map[string]enumspb.IndexedValueType {
	customSearchAttributes := map[string]enumspb.IndexedValueType{}
	for saName, saType := range searchAttributes {
		if saName == NamespaceID {
			continue
		}
		if _, isSystem := system[saName]; !isSystem {
			customSearchAttributes[saName] = saType
		}
	}
	return customSearchAttributes
}
