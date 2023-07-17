// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package audit

import (
	"os"
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
			Value:                "",
			IsErrorExpected:      true,
			ExpectedErrorMessage: "format cannot be empty",
		},
		"whitespace": {
			Value:                "     ",
			IsErrorExpected:      true,
			ExpectedErrorMessage: "format cannot be empty",
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

// TestOptions_WithFacility exercises WithFacility Option to ensure it performs as expected.
func TestOptions_WithFacility(t *testing.T) {
	tests := map[string]struct {
		Value         string
		ExpectedValue string
	}{
		"empty": {
			Value:         "",
			ExpectedValue: "",
		},
		"whitespace": {
			Value:         "    ",
			ExpectedValue: "",
		},
		"value": {
			Value:         "juan",
			ExpectedValue: "juan",
		},
		"spacey-value": {
			Value:         "   juan   ",
			ExpectedValue: "juan",
		},
	}

	for name, tc := range tests {
		name := name
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			options := &options{}
			applyOption := WithFacility(tc.Value)
			err := applyOption(options)
			require.NoError(t, err)
			require.Equal(t, tc.ExpectedValue, options.withFacility)
		})
	}
}

// TestOptions_WithTag exercises WithTag Option to ensure it performs as expected.
func TestOptions_WithTag(t *testing.T) {
	tests := map[string]struct {
		Value         string
		ExpectedValue string
	}{
		"empty": {
			Value:         "",
			ExpectedValue: "",
		},
		"whitespace": {
			Value:         "    ",
			ExpectedValue: "",
		},
		"value": {
			Value:         "juan",
			ExpectedValue: "juan",
		},
		"spacey-value": {
			Value:         "   juan   ",
			ExpectedValue: "juan",
		},
	}

	for name, tc := range tests {
		name := name
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			options := &options{}
			applyOption := WithTag(tc.Value)
			err := applyOption(options)
			require.NoError(t, err)
			require.Equal(t, tc.ExpectedValue, options.withTag)
		})
	}
}

// TestOptions_WithSocketType exercises WithSocketType Option to ensure it performs as expected.
func TestOptions_WithSocketType(t *testing.T) {
	tests := map[string]struct {
		Value         string
		ExpectedValue string
	}{
		"empty": {
			Value:         "",
			ExpectedValue: "",
		},
		"whitespace": {
			Value:         "    ",
			ExpectedValue: "",
		},
		"value": {
			Value:         "juan",
			ExpectedValue: "juan",
		},
		"spacey-value": {
			Value:         "   juan   ",
			ExpectedValue: "juan",
		},
	}

	for name, tc := range tests {
		name := name
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			options := &options{}
			applyOption := WithSocketType(tc.Value)
			err := applyOption(options)
			require.NoError(t, err)
			require.Equal(t, tc.ExpectedValue, options.withSocketType)
		})
	}
}

// TestOptions_WithMaxDuration exercises WithMaxDuration Option to ensure it performs as expected.
func TestOptions_WithMaxDuration(t *testing.T) {
	tests := map[string]struct {
		Value                string
		ExpectedValue        time.Duration
		IsErrorExpected      bool
		ExpectedErrorMessage string
	}{
		"empty-gives-default": {
			Value: "",
		},
		"whitespace-give-default": {
			Value: "    ",
		},
		"bad-value": {
			Value:                "juan",
			IsErrorExpected:      true,
			ExpectedErrorMessage: "time: invalid duration \"juan\"",
		},
		"bad-spacey-value": {
			Value:                "   juan   ",
			IsErrorExpected:      true,
			ExpectedErrorMessage: "time: invalid duration \"juan\"",
		},
		"duration-2s": {
			Value:         "2s",
			ExpectedValue: 2 * time.Second,
		},
		"duration-2m": {
			Value:         "2m",
			ExpectedValue: 2 * time.Minute,
		},
	}

	for name, tc := range tests {
		name := name
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			options := &options{}
			applyOption := WithMaxDuration(tc.Value)
			err := applyOption(options)
			switch {
			case tc.IsErrorExpected:
				require.Error(t, err)
				require.EqualError(t, err, tc.ExpectedErrorMessage)
			default:
				require.NoError(t, err)
				require.Equal(t, tc.ExpectedValue, options.withMaxDuration)
			}
		})
	}
}

// TestOptions_WithFileMode exercises WithFileMode Option to ensure it performs as expected.
func TestOptions_WithFileMode(t *testing.T) {
	tests := map[string]struct {
		Value                string
		IsErrorExpected      bool
		ExpectedErrorMessage string
		IsNilExpected        bool
		ExpectedValue        os.FileMode
	}{
		"empty": {
			Value:           "",
			IsErrorExpected: false,
			IsNilExpected:   true,
		},
		"whitespace": {
			Value:           "     ",
			IsErrorExpected: false,
			IsNilExpected:   true,
		},
		"nonsense": {
			Value:                "juan",
			IsErrorExpected:      true,
			ExpectedErrorMessage: "unable to parse file mode: strconv.ParseUint: parsing \"juan\": invalid syntax",
		},
		"zero": {
			Value:           "0000",
			IsErrorExpected: false,
			ExpectedValue:   os.FileMode(0o000),
		},
		"valid": {
			Value:           "0007",
			IsErrorExpected: false,
			ExpectedValue:   os.FileMode(0o007),
		},
	}

	for name, tc := range tests {
		name := name
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			options := &options{}
			applyOption := WithFileMode(tc.Value)
			err := applyOption(options)
			switch {
			case tc.IsErrorExpected:
				require.Error(t, err)
				require.EqualError(t, err, tc.ExpectedErrorMessage)
			default:
				require.NoError(t, err)
				switch {
				case tc.IsNilExpected:
					// Optional Option 'not supplied' (i.e. was whitespace/empty string)
					require.Nil(t, options.withFileMode)
				default:
					// Dereference the pointer, so we can examine the file mode.
					require.Equal(t, tc.ExpectedValue, *options.withFileMode)
				}
			}
		})
	}
}

// TestOptions_Default exercises getDefaultOptions to assert the default values.
func TestOptions_Default(t *testing.T) {
	opts := getDefaultOptions()
	require.NotNil(t, opts)
	require.True(t, time.Now().After(opts.withNow))
	require.False(t, opts.withNow.IsZero())
	require.Equal(t, "AUTH", opts.withFacility)
	require.Equal(t, "vault", opts.withTag)
	require.Equal(t, 2*time.Second, opts.withMaxDuration)
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
		},
		"empty-options": {
			opts:            []Option{},
			IsErrorExpected: false,
			IsNowExpected:   true,
		},
		"with-multiple-valid-id": {
			opts: []Option{
				WithID("qwerty"),
				WithID("juan"),
			},
			IsErrorExpected: false,
			ExpectedID:      "juan",
			IsNowExpected:   true,
		},
		"with-multiple-valid-subtype": {
			opts: []Option{
				WithSubtype("AuditRequest"),
				WithSubtype("AuditResponse"),
			},
			IsErrorExpected: false,
			ExpectedSubtype: ResponseType,
			IsNowExpected:   true,
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
		},
		"with-multiple-valid-then-invalid-now": {
			opts: []Option{
				WithNow(time.Date(2023, time.July, 4, 12, 3, 0, 0, time.Local)),
				WithNow(time.Time{}),
			},
			IsErrorExpected:      true,
			ExpectedErrorMessage: "cannot specify 'now' to be the zero time instant",
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