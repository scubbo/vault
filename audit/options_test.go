// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package audit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// TestOptions_WithFormat exercises WithFormat Option to ensure it performs as expected.
func TestOptions_WithFormat(t *testing.T) {
	tests := map[string]struct {
		Value                string
		IsErrorExpected      bool
		ExpectedErrorMessage string
		ExpectedValue        format
	}{
		"empty": {
			Value:           "",
			IsErrorExpected: false,
			ExpectedValue:   format(""),
		},
		"whitespace": {
			Value:           "     ",
			IsErrorExpected: false,
			ExpectedValue:   format(""),
		},
		"invalid-test": {
			Value:                "test",
			IsErrorExpected:      true,
			ExpectedErrorMessage: "audit.(format).validate: 'test' is not a valid format: invalid parameter",
		},
		"valid-json": {
			Value:           "json",
			IsErrorExpected: false,
			ExpectedValue:   JSONFormat,
		},
		"valid-jsonx": {
			Value:           "jsonx",
			IsErrorExpected: false,
			ExpectedValue:   JSONxFormat,
		},
	}

	for name, tc := range tests {
		name := name
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			options := &options{}
			applyOption := WithFormat(tc.Value)
			err := applyOption(options)
			switch {
			case tc.IsErrorExpected:
				require.Error(t, err)
				require.EqualError(t, err, tc.ExpectedErrorMessage)
			default:
				require.NoError(t, err)
				require.Equal(t, tc.ExpectedValue, options.withFormat)
			}
		})
	}
}

// TestOptions_WithSubtype exercises WithSubtype Option to ensure it performs as expected.
func TestOptions_WithSubtype(t *testing.T) {
	tests := map[string]struct {
		Value                string
		IsErrorExpected      bool
		ExpectedErrorMessage string
		ExpectedValue        subtype
	}{
		"empty": {
			Value:                "",
			IsErrorExpected:      true,
			ExpectedErrorMessage: "subtype cannot be empty",
		},
		"whitespace": {
			Value:                "     ",
			IsErrorExpected:      true,
			ExpectedErrorMessage: "subtype cannot be empty",
		},
		"valid": {
			Value:           "AuditResponse",
			IsErrorExpected: false,
			ExpectedValue:   ResponseType,
		},
	}

	for name, tc := range tests {
		name := name
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			options := &options{}
			applyOption := WithSubtype(tc.Value)
			err := applyOption(options)
			switch {
			case tc.IsErrorExpected:
				require.Error(t, err)
				require.EqualError(t, err, tc.ExpectedErrorMessage)
			default:
				require.NoError(t, err)
				require.Equal(t, tc.ExpectedValue, options.withSubtype)
			}
		})
	}
}

// TestOptions_WithNow exercises WithNow Option to ensure it performs as expected.
func TestOptions_WithNow(t *testing.T) {
	tests := map[string]struct {
		Value                time.Time
		IsErrorExpected      bool
		ExpectedErrorMessage string
		ExpectedValue        time.Time
	}{
		"default-time": {
			Value:                time.Time{},
			IsErrorExpected:      true,
			ExpectedErrorMessage: "cannot specify 'now' to be the zero time instant",
		},
		"valid-time": {
			Value:           time.Date(2023, time.July, 4, 12, 3, 0, 0, time.Local),
			IsErrorExpected: false,
			ExpectedValue:   time.Date(2023, time.July, 4, 12, 3, 0, 0, time.Local),
		},
	}

	for name, tc := range tests {
		name := name
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			options := &options{}
			applyOption := WithNow(tc.Value)
			err := applyOption(options)
			switch {
			case tc.IsErrorExpected:
				require.Error(t, err)
				require.EqualError(t, err, tc.ExpectedErrorMessage)
			default:
				require.NoError(t, err)
				require.Equal(t, tc.ExpectedValue, options.withNow)
			}
		})
	}
}

// TestOptions_WithID exercises WithID Option to ensure it performs as expected.
func TestOptions_WithID(t *testing.T) {
	tests := map[string]struct {
		Value                string
		IsErrorExpected      bool
		ExpectedErrorMessage string
		ExpectedValue        string
	}{
		"empty": {
			Value:                "",
			IsErrorExpected:      true,
			ExpectedErrorMessage: "id cannot be empty",
		},
		"whitespace": {
			Value:                "     ",
			IsErrorExpected:      true,
			ExpectedErrorMessage: "id cannot be empty",
		},
		"valid": {
			Value:           "test",
			IsErrorExpected: false,
			ExpectedValue:   "test",
		},
	}

	for name, tc := range tests {
		name := name
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			options := &options{}
			applyOption := WithID(tc.Value)
			err := applyOption(options)
			switch {
			case tc.IsErrorExpected:
				require.Error(t, err)
				require.EqualError(t, err, tc.ExpectedErrorMessage)
			default:
				require.NoError(t, err)
				require.Equal(t, tc.ExpectedValue, options.withID)
			}
		})
	}
}

// TestOptions_WithPrefix exercises WithPrefix Option to ensure it performs as expected.
func TestOptions_WithPrefix(t *testing.T) {
	tests := map[string]struct {
		Value                string
		IsErrorExpected      bool
		ExpectedErrorMessage string
		ExpectedValue        string
	}{
		"empty": {
			Value:           "",
			IsErrorExpected: false,
			ExpectedValue:   "",
		},
		"whitespace": {
			Value:                "     ",
			IsErrorExpected:      false,
			ExpectedErrorMessage: "",
		},
		"valid": {
			Value:           "test",
			IsErrorExpected: false,
			ExpectedValue:   "test",
		},
	}

	for name, tc := range tests {
		name := name
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			options := &options{}
			applyOption := WithPrefix(tc.Value)
			err := applyOption(options)
			switch {
			case tc.IsErrorExpected:
				require.Error(t, err)
				require.EqualError(t, err, tc.ExpectedErrorMessage)
			default:
				require.NoError(t, err)
				require.Equal(t, tc.ExpectedValue, options.withPrefix)
			}
		})
	}
}

// TestOptions_WithRaw exercises WithRaw Option to ensure it performs as expected.
func TestOptions_WithRaw(t *testing.T) {
	tests := map[string]struct {
		Value         bool
		ExpectedValue bool
	}{
		"true": {
			Value:         true,
			ExpectedValue: true,
		},
		"false": {
			Value:         false,
			ExpectedValue: false,
		},
	}

	for name, tc := range tests {
		name := name
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			options := &options{}
			applyOption := WithRaw(tc.Value)
			err := applyOption(options)
			require.NoError(t, err)
			require.Equal(t, tc.ExpectedValue, options.withRaw)
		})
	}
}

// TestOptions_WithElision exercises WithElision Option to ensure it performs as expected.
func TestOptions_WithElision(t *testing.T) {
	tests := map[string]struct {
		Value         bool
		ExpectedValue bool
	}{
		"true": {
			Value:         true,
			ExpectedValue: true,
		},
		"false": {
			Value:         false,
			ExpectedValue: false,
		},
	}

	for name, tc := range tests {
		name := name
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			options := &options{}
			applyOption := WithElision(tc.Value)
			err := applyOption(options)
			require.NoError(t, err)
			require.Equal(t, tc.ExpectedValue, options.withElision)
		})
	}
}

// TestOptions_WithHMACAccessor exercises WithHMACAccessor Option to ensure it performs as expected.
func TestOptions_WithHMACAccessor(t *testing.T) {
	tests := map[string]struct {
		Value         bool
		ExpectedValue bool
	}{
		"true": {
			Value:         true,
			ExpectedValue: true,
		},
		"false": {
			Value:         false,
			ExpectedValue: false,
		},
	}

	for name, tc := range tests {
		name := name
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			options := &options{}
			applyOption := WithHMACAccessor(tc.Value)
			err := applyOption(options)
			require.NoError(t, err)
			require.Equal(t, tc.ExpectedValue, options.withHMACAccessor)
		})
	}
}

// TestOptions_WithOmitTime exercises WithOmitTime Option to ensure it performs as expected.
func TestOptions_WithOmitTime(t *testing.T) {
	tests := map[string]struct {
		Value         bool
		ExpectedValue bool
	}{
		"true": {
			Value:         true,
			ExpectedValue: true,
		},
		"false": {
			Value:         false,
			ExpectedValue: false,
		},
	}

	for name, tc := range tests {
		name := name
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			options := &options{}
			applyOption := WithOmitTime(tc.Value)
			err := applyOption(options)
			require.NoError(t, err)
			require.Equal(t, tc.ExpectedValue, options.withOmitTime)
		})
	}
}

// TestOptions_Default exercises getDefaultOptions to assert the default values.
func TestOptions_Default(t *testing.T) {
	opts := getDefaultOptions()
	require.NotNil(t, opts)
	require.True(t, time.Now().After(opts.withNow))
	require.False(t, opts.withNow.IsZero())
}

// TestOptions_Opts exercises GetOpts with various Option values.
func TestOptions_Opts(t *testing.T) {
	tests := map[string]struct {
		opts                 []Option
		IsErrorExpected      bool
		ExpectedErrorMessage string
		ExpectedID           string
		ExpectedSubtype      subtype
		ExpectedFormat       format
		IsNowExpected        bool
		ExpectedNow          time.Time
	}{
		"nil-options": {
			opts:            nil,
			IsErrorExpected: false,
			IsNowExpected:   true,
			ExpectedFormat:  JSONFormat,
		},
		"empty-options": {
			opts:            []Option{},
			IsErrorExpected: false,
			IsNowExpected:   true,
			ExpectedFormat:  JSONFormat,
		},
		"with-multiple-valid-id": {
			opts: []Option{
				WithID("qwerty"),
				WithID("juan"),
			},
			IsErrorExpected: false,
			ExpectedID:      "juan",
			IsNowExpected:   true,
			ExpectedFormat:  JSONFormat,
		},
		"with-multiple-valid-subtype": {
			opts: []Option{
				WithSubtype("AuditRequest"),
				WithSubtype("AuditResponse"),
			},
			IsErrorExpected: false,
			ExpectedSubtype: ResponseType,
			IsNowExpected:   true,
			ExpectedFormat:  JSONFormat,
		},
		"with-multiple-valid-format": {
			opts: []Option{
				WithFormat("json"),
				WithFormat("jsonx"),
			},
			IsErrorExpected: false,
			ExpectedFormat:  JSONxFormat,
			IsNowExpected:   true,
		},
		"with-multiple-valid-now": {
			opts: []Option{
				WithNow(time.Date(2023, time.July, 4, 12, 3, 0, 0, time.Local)),
				WithNow(time.Date(2023, time.July, 4, 13, 3, 0, 0, time.Local)),
			},
			IsErrorExpected: false,
			ExpectedNow:     time.Date(2023, time.July, 4, 13, 3, 0, 0, time.Local),
			IsNowExpected:   false,
			ExpectedFormat:  JSONFormat,
		},
		"with-multiple-valid-then-invalid-now": {
			opts: []Option{
				WithNow(time.Date(2023, time.July, 4, 12, 3, 0, 0, time.Local)),
				WithNow(time.Time{}),
			},
			IsErrorExpected:      true,
			ExpectedErrorMessage: "cannot specify 'now' to be the zero time instant",
			ExpectedFormat:       JSONFormat,
		},
		"with-multiple-valid-options": {
			opts: []Option{
				WithID("qwerty"),
				WithSubtype("AuditRequest"),
				WithFormat("json"),
				WithNow(time.Date(2023, time.July, 4, 12, 3, 0, 0, time.Local)),
			},
			IsErrorExpected: false,
			ExpectedID:      "qwerty",
			ExpectedSubtype: RequestType,
			ExpectedFormat:  JSONFormat,
			ExpectedNow:     time.Date(2023, time.July, 4, 12, 3, 0, 0, time.Local),
		},
	}

	for name, tc := range tests {
		name := name
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			opts, err := getOpts(tc.opts...)

			switch {
			case tc.IsErrorExpected:
				require.Error(t, err)
				require.EqualError(t, err, tc.ExpectedErrorMessage)
			default:
				require.NotNil(t, opts)
				require.NoError(t, err)
				require.Equal(t, tc.ExpectedID, opts.withID)
				require.Equal(t, tc.ExpectedSubtype, opts.withSubtype)
				require.Equal(t, tc.ExpectedFormat, opts.withFormat)
				switch {
				case tc.IsNowExpected:
					require.True(t, time.Now().After(opts.withNow))
					require.False(t, opts.withNow.IsZero())
				default:
					require.Equal(t, tc.ExpectedNow, opts.withNow)
				}

			}
		})
	}
}
