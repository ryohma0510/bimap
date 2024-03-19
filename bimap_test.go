package bimap

import (
	"testing"

	"gotest.tools/v3/assert"
)

type testGenderStr string

const (
	testGenderStrMale   testGenderStr = "male"
	testGenderStrFemale testGenderStr = "female"
)

type testGenderInt int

const (
	testGenderIntMale testGenderInt = iota + 1
	testGenderIntFemale
)

func TestNewFromMap(t *testing.T) {

	type args struct {
		forward map[testGenderStr]testGenderInt
	}
	tests := []struct {
		name      string
		args      args
		assertion func(t *testing.T, got *BiMap[testGenderStr, testGenderInt], err error)
	}{
		{
			name: "ok",
			args: args{
				forward: map[testGenderStr]testGenderInt{
					testGenderStrMale:   testGenderIntMale,
					testGenderStrFemale: testGenderIntFemale,
				},
			},
			assertion: func(t *testing.T, got *BiMap[testGenderStr, testGenderInt], err error) {
				assert.NilError(t, err)
				assert.DeepEqual(t, got.forward, map[testGenderStr]testGenderInt{
					testGenderStrMale:   testGenderIntMale,
					testGenderStrFemale: testGenderIntFemale,
				})
				assert.DeepEqual(t, got.inverse, map[testGenderInt]testGenderStr{
					testGenderIntMale:   testGenderStrMale,
					testGenderIntFemale: testGenderStrFemale,
				})
			},
		},
		{
			name: "duplicate values",
			args: args{
				forward: map[testGenderStr]testGenderInt{
					testGenderStrMale:   testGenderIntMale,
					testGenderStrFemale: testGenderIntMale,
				},
			},
			assertion: func(t *testing.T, got *BiMap[testGenderStr, testGenderInt], err error) {
				assert.ErrorContains(t, err, "duplicate values in forward map")
				assert.DeepEqual(t, got, (*BiMap[testGenderStr, testGenderInt])(nil))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFromMap(tt.args.forward)
			tt.assertion(t, got, err)
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name       string
		forwardMap map[testGenderStr]testGenderInt
		k          testGenderStr
		assertion  func(t *testing.T, got testGenderInt, ok bool)
	}{
		{
			name: "ok",
			forwardMap: map[testGenderStr]testGenderInt{
				testGenderStrMale:   testGenderIntMale,
				testGenderStrFemale: testGenderIntFemale,
			},
			k: testGenderStrMale,
			assertion: func(t *testing.T, got testGenderInt, ok bool) {
				assert.Assert(t, ok)
				assert.Equal(t, got, testGenderIntMale)
			},
		},
		{
			name: "not found",
			forwardMap: map[testGenderStr]testGenderInt{
				testGenderStrMale:   testGenderIntMale,
				testGenderStrFemale: testGenderIntFemale,
			},
			k: testGenderStr("not found"),
			assertion: func(t *testing.T, got testGenderInt, ok bool) {
				assert.Assert(t, !ok)
				assert.Equal(t, got, testGenderInt(0))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := NewFromMap(tt.forwardMap)
			assert.NilError(t, err)

			got, ok := b.Get(tt.k)
			tt.assertion(t, got, ok)
		})
	}
}

func TestInverseGet(t *testing.T) {
	tests := []struct {
		name       string
		forwardMap map[testGenderStr]testGenderInt
		v          testGenderInt
		assertion  func(t *testing.T, got testGenderStr, ok bool)
	}{
		{
			name: "ok",
			forwardMap: map[testGenderStr]testGenderInt{
				testGenderStrMale:   testGenderIntMale,
				testGenderStrFemale: testGenderIntFemale,
			},
			v: testGenderIntMale,
			assertion: func(t *testing.T, got testGenderStr, ok bool) {
				assert.Assert(t, ok)
				assert.Equal(t, got, testGenderStrMale)
			},
		},
		{
			name: "not found",
			forwardMap: map[testGenderStr]testGenderInt{
				testGenderStrMale:   testGenderIntMale,
				testGenderStrFemale: testGenderIntFemale,
			},
			v: testGenderInt(0),
			assertion: func(t *testing.T, got testGenderStr, ok bool) {
				assert.Assert(t, !ok)
				assert.Equal(t, got, testGenderStr(""))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := NewFromMap(tt.forwardMap)
			assert.NilError(t, err)

			got, ok := b.InverseGet(tt.v)
			tt.assertion(t, got, ok)
		})
	}
}
