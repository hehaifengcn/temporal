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

package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryRegistryBuilder_RegisterCategory_DuplicateID(t *testing.T) {
	t.Parallel()

	registry := newCategoryRegistry()
	assert.NoError(t, registry.RegisterCategory(Category{
		id:    100,
		cType: CategoryTypeImmediate,
		name:  "test1",
	}))
	err := registry.RegisterCategory(Category{
		id:    100,
		cType: CategoryTypeImmediate,
		name:  "test2",
	})
	assert.ErrorIs(t, err, ErrCategoryAlreadyRegistered)
	assert.ErrorContains(t, err, "100")
	assert.ErrorContains(t, err, "test1")
	assert.ErrorContains(t, err, "test2")
}

func TestCategoryRegistryBuilder_RegisterCategory_DifferentIDs(t *testing.T) {
	t.Parallel()

	registry := newCategoryRegistry()
	assert.NoError(t, registry.RegisterCategory(Category{
		id:    100,
		cType: CategoryTypeImmediate,
		name:  "test1",
	}))
	assert.NoError(t, registry.RegisterCategory(Category{
		id:    101,
		cType: CategoryTypeImmediate,
		name:  "test1",
	}))
}

func TestCategoryRegistryBuilder_BuildIndex(t *testing.T) {
	t.Parallel()

	registry := newCategoryRegistry()
	assert.NoError(t, registry.RegisterCategory(Category{
		id:    100,
		cType: CategoryTypeImmediate,
		name:  "test1",
	}))
	categoryIndex := registry.BuildCategoryIndex()
	assert.Equal(t, map[int32]Category{
		100: {
			id:    100,
			cType: CategoryTypeImmediate,
			name:  "test1",
		},
	}, categoryIndex.GetCategories())
}
