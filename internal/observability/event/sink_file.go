// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package event

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/hashicorp/eventlogger"
)

// defaultFileMode is the default file permissions (read/write for everyone).
const (
	defaultFileMode = 0o600
	discard         = "discard"
	stdout          = "stdout"
)

// FileSink is a sink node which handles writing events to file.
type FileSink struct {
	file           *os.File
	fileLock       sync.RWMutex
	fileMode       os.FileMode
	path           string
	requiredFormat string
	prefix         string
}

// NewFileSink should be used to create a new FileSink.
// Accepted options: WithFileMode and WithPrefix.
func NewFileSink(path string, format string, opt ...Option) (*FileSink, error) {
	const op = "event.NewFileSink"

	// Parse and check path
	p := strings.TrimSpace(path)
	if p == "" {
		return nil, fmt.Errorf("%s: path is required", op)
	}

	opts, err := getOpts(opt...)
	if err != nil {
		return nil, fmt.Errorf("%s: error applying options: %w", op, err)
	}

	mode := os.FileMode(defaultFileMode)
	// If we got an optional file mode supplied and our path isn't a special keyword
	// then we should use the supplied file mode, or maintain the existing file mode.
	if opts.withFileMode != nil {
		switch {
		case *opts.withFileMode == 0: // Maintain the existing file's mode when set to "0000".
			fileInfo, err := os.Stat(path)
			if err != nil {
				return nil, fmt.Errorf("%s: unable to determine existing file mode: %w", op, err)
			}
			mode = fileInfo.Mode()
		default:
			mode = *opts.withFileMode
		}
	}

	return &FileSink{
		file:           nil,
		fileLock:       sync.RWMutex{},
		fileMode:       mode,
		requiredFormat: format,
		path:           p,
		prefix:         opts.withPrefix,
	}, nil
}

// Process handles writing the event to the file sink.
func (f *FileSink) Process(ctx context.Context, e *eventlogger.Event) (*eventlogger.Event, error) {
	const op = "event.(FileSink).Process"

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	if e == nil {
		return nil, fmt.Errorf("%s: event is nil: %w", op, ErrInvalidParameter)
	}

	// 'discard' path means we just do nothing and pretend we're done.
	if f.path == discard {
		return nil, nil
	}

	formatted, found := e.Format(f.requiredFormat)
	if !found {
		return nil, fmt.Errorf("%s: unable to retrieve event formatted as %q", op, f.requiredFormat)
	}

	err := f.log(formatted)
	if err != nil {
		return nil, fmt.Errorf("%s: error writing file for sink: %w", op, err)
	}

	// return nil for the event to indicate the pipeline is complete.
	return nil, nil
}

// Reopen handles closing and reopening the file.
func (f *FileSink) Reopen() error {
	const op = "event.(FileSink).Reopen"

	switch f.path {
	case stdout, discard:
		return nil
	}

	f.fileLock.Lock()
	defer f.fileLock.Unlock()

	if f.file == nil {
		return f.open()
	}

	err := f.file.Close()
	// Set to nil here so that even if we error out, on the next access open() will be tried.
	f.file = nil
	if err != nil {
		return fmt.Errorf("%s: unable to close file for re-opening on sink: %w", op, err)
	}

	return f.open()
}

// Type describes the type of this node (sink).
func (_ *FileSink) Type() eventlogger.NodeType {
	return eventlogger.NodeTypeSink
}

// open attempts to open a file at the sink's path, with the sink's fileMode permissions
// if one is not already open.
// It doesn't have any locking and relies on calling functions of FileSink to
// handle this (e.g. log and Reopen methods).
func (f *FileSink) open() error {
	const op = "event.(FileSink).open"

	if f.file != nil {
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(f.path), f.fileMode); err != nil {
		return fmt.Errorf("%s: unable to create file %q: %w", op, f.path, err)
	}

	var err error
	f.file, err = os.OpenFile(f.path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, f.fileMode)
	if err != nil {
		return fmt.Errorf("%s: unable to open file for sink: %w", op, err)
	}

	// Change the file mode in case the log file already existed.
	// We special case '/dev/null' since we can't chmod it, and bypass if the mode is zero.
	switch f.path {
	case "/dev/null":
	default:
		if f.fileMode != 0 {
			err = os.Chmod(f.path, f.fileMode)
			if err != nil {
				return fmt.Errorf("%s: unable to change file %q permissions '%v' for sink: %w", op, f.path, f.fileMode, err)
			}
		}
	}

	return nil
}

// log writes the buffer to the file.
// It acquires a lock on the file to do this.
func (f *FileSink) log(data []byte) error {
	const op = "event.(FileSink).log"

	f.fileLock.Lock()
	defer f.fileLock.Unlock()

	reader := bytes.NewReader(data)

	var writer io.Writer
	switch {
	case f.path == stdout:
		writer = os.Stdout
	default:
		if err := f.open(); err != nil {
			return fmt.Errorf("%s: unable to open file for sink: %w", op, err)
		}
		writer = f.file
	}

	// Write prefix before the data if required.
	if f.prefix != "" {
		_, err := writer.Write([]byte(f.prefix))
		if err != nil {
			return fmt.Errorf("%s: unable to write prefix %q for sink: %w", op, f.prefix, err)
		}
	}

	if _, err := reader.WriteTo(writer); err == nil {
		return nil
	} else if f.path == stdout {
		// If writing to stdout there's no real reason to think anything would change on retry.
		return fmt.Errorf("%s: unable write to %q: %w", op, f.path, err)
	}

	// Otherwise, opportunistically try to re-open the FD, once per call (1 retry attempt).
	err := f.file.Close()
	if err != nil {
		return fmt.Errorf("%s: unable to close file for sink: %w", op, err)
	}

	f.file = nil

	if err := f.open(); err != nil {
		return fmt.Errorf("%s: unable to re-open file for sink: %w", op, err)
	}

	_, err = reader.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("%s: unable to seek to start of file for sink: %w", op, err)
	}

	_, err = reader.WriteTo(writer)
	if err != nil {
		return fmt.Errorf("%s: unable to re-write to file for sink: %w", op, err)
	}

	return nil
}
