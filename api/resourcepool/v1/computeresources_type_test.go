/*
Copyright 2020 Netflix, Inc.

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
package v1

import (
	"testing"

	"gotest.tools/assert"
)

func TestComputeResource_Add(t *testing.T) {
	result := ComputeResource{CPU: 1, GPU: 2, MemoryMB: 3, DiskMB: 4, NetworkMBPS: 5}.Add(
		ComputeResource{CPU: 6, GPU: 7, MemoryMB: 8, DiskMB: 9, NetworkMBPS: 10})
	expected := ComputeResource{CPU: 7, GPU: 9, MemoryMB: 11, DiskMB: 13, NetworkMBPS: 15}
	assert.Equal(t, result, expected)
}

func TestComputeResource_Sub(t *testing.T) {
	result := ComputeResource{CPU: 10, GPU: 11, MemoryMB: 12, DiskMB: 13, NetworkMBPS: 14}.Sub(
		ComputeResource{CPU: 5, GPU: 4, MemoryMB: 3, DiskMB: 2, NetworkMBPS: 1})
	expected := ComputeResource{CPU: 5, GPU: 7, MemoryMB: 9, DiskMB: 11, NetworkMBPS: 13}
	assert.Equal(t, result, expected)
}

func TestComputeResource_SubWithLimit(t *testing.T) {
	result := ComputeResource{CPU: 1, GPU: 1, MemoryMB: 1, DiskMB: 1, NetworkMBPS: 1}.SubWithLimit(
		ComputeResource{CPU: 2, GPU: 2, MemoryMB: 2, DiskMB: 2, NetworkMBPS: 2}, 0)
	assert.Equal(t, result, Zero)
}

func TestComputeResource_Multiply(t *testing.T) {
	result := ComputeResource{CPU: 1, GPU: 2, MemoryMB: 3, DiskMB: 4, NetworkMBPS: 5}.Multiply(2)
	expected := ComputeResource{CPU: 2, GPU: 4, MemoryMB: 6, DiskMB: 8, NetworkMBPS: 10}
	assert.Equal(t, result, expected)
}

func TestAlignResourceRatios(t *testing.T) {
	reference := ComputeResource{CPU: 8, GPU: 2, MemoryMB: 4096, DiskMB: 8192, NetworkMBPS: 128}

	// CPU dominant resource == reference.CPU
	result := ComputeResource{CPU: 8, GPU: 0, MemoryMB: 100, DiskMB: 200, NetworkMBPS: 100}.AlignResourceRatios(reference)
	expected := ComputeResource{CPU: 8, GPU: 2, MemoryMB: 4096, DiskMB: 8192, NetworkMBPS: 128}
	assert.Equal(t, result, expected)

	// CPU dominant resource < reference.CPU
	result = ComputeResource{CPU: 4, GPU: 0, MemoryMB: 100, DiskMB: 200, NetworkMBPS: 20}.AlignResourceRatios(reference)
	expected = ComputeResource{CPU: 4, GPU: 1, MemoryMB: 2048, DiskMB: 4096, NetworkMBPS: 64}
	assert.Equal(t, result, expected)

	// CPU dominant resource > reference.CPU
	result = ComputeResource{CPU: 12, GPU: 1, MemoryMB: 2, DiskMB: 3, NetworkMBPS: 4}.AlignResourceRatios(reference)
	expected = ComputeResource{CPU: 12, GPU: 3, MemoryMB: 6144, DiskMB: 12288, NetworkMBPS: 192}
	assert.Equal(t, result, expected)

	// reference resource 0, but the target > 0
	target := ComputeResource{CPU: 12, GPU: 1, MemoryMB: 2, DiskMB: 3, NetworkMBPS: 4}
	result = target.AlignResourceRatios(ComputeResource{CPU: 8})
	assert.Equal(t, result, target)
}

func TestComputeResource_CanSplitBy(t *testing.T) {
	dividend := ComputeResource{CPU: 1, GPU: 0, MemoryMB: 2, DiskMB: 0, NetworkMBPS: 3}
	assert.Assert(t, dividend.CanSplitBy(ComputeResource{CPU: 1, GPU: 0, MemoryMB: 1, DiskMB: 0, NetworkMBPS: 1}))
	assert.Assert(t, !dividend.CanSplitBy(ComputeResource{CPU: 0, GPU: 0, MemoryMB: 1, DiskMB: 0, NetworkMBPS: 1}))
}

func TestComputeResource_SplitBy(t *testing.T) {
	dividend := ComputeResource{CPU: 4, GPU: 6, MemoryMB: 8, DiskMB: 10, NetworkMBPS: 12}
	result := dividend.SplitBy(ComputeResource{CPU: 2, GPU: 2, MemoryMB: 2, DiskMB: 2, NetworkMBPS: 2})
	assert.Assert(t, result == 2)
}

func TestComputeResource_SplitByWithCeil(t *testing.T) {
	dividend := ComputeResource{CPU: 5, GPU: 6, MemoryMB: 8, DiskMB: 10, NetworkMBPS: 12}
	result := dividend.SplitByWithCeil(ComputeResource{CPU: 2, GPU: 2, MemoryMB: 2, DiskMB: 2, NetworkMBPS: 2})
	assert.Assert(t, result == 3)
}

func TestComputeResource_StrictlyLessThan(t *testing.T) {
	left := ComputeResource{CPU: 1, GPU: 2, MemoryMB: 3, DiskMB: 4, NetworkMBPS: 5}
	assert.Assert(t, left.StrictlyLessThan(ComputeResource{CPU: 2, GPU: 3, MemoryMB: 4, DiskMB: 5, NetworkMBPS: 6}))
}

func TestComputeResource_LessThan(t *testing.T) {
	left := ComputeResource{CPU: 1, GPU: 2, MemoryMB: 3, DiskMB: 4, NetworkMBPS: 5}
	assert.Assert(t, left.LessThan(ComputeResource{CPU: 2, GPU: 2, MemoryMB: 3, DiskMB: 4, NetworkMBPS: 5}))
}

func TestComputeResource_StrictlyGreaterThan(t *testing.T) {
	left := ComputeResource{CPU: 2, GPU: 3, MemoryMB: 4, DiskMB: 5, NetworkMBPS: 6}
	assert.Assert(t, left.StrictlyGreaterThan(ComputeResource{CPU: 1, GPU: 2, MemoryMB: 3, DiskMB: 4, NetworkMBPS: 5}))
}

func TestComputeResource_GreaterThan(t *testing.T) {
	left := ComputeResource{CPU: 2, GPU: 2, MemoryMB: 3, DiskMB: 4, NetworkMBPS: 5}
	assert.Assert(t, left.GreaterThan(ComputeResource{CPU: 1, GPU: 2, MemoryMB: 3, DiskMB: 4, NetworkMBPS: 5}))
}

func TestComputeResource_IsAboveZero(t *testing.T) {
	assert.Assert(t, !Zero.IsAboveZero())
	assert.Assert(t, ComputeResource{CPU: 0, GPU: 0, MemoryMB: 0, DiskMB: 0, NetworkMBPS: 1}.IsAboveZero())
	assert.Assert(t, !ComputeResource{CPU: 0, GPU: 0, MemoryMB: 0, DiskMB: 0, NetworkMBPS: -1}.IsAboveZero())
}
