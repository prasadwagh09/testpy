/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package tlogger initializes the testing logger on import which logs to the
// testing package's T struct.
package tlogger

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"runtime/debug"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"

	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/internal/leakcheck"
)

var logger = tLogger{v: 0}

const callingFrame = 4

type logType int

const (
	logLog logType = iota
	errorLog
	fatalLog
)

type tLogger struct {
	v      int
	t      *testing.T
	errors []*regexp.Regexp
}

func init() {
	vLevel := os.Getenv("GRPC_GO_LOG_VERBOSITY_LEVEL")
	if vl, err := strconv.Atoi(vLevel); err == nil {
		logger.v = vl
	}
	grpclog.SetLoggerV2(&logger)
}

func getStackFrame(stack []byte, frame int) (string, error) {
	s := strings.Split(string(stack), "\n")
	if frame >= (len(s)-1)/2 {
		return "", errors.New("frame request out-of-bounds")
	}
	split := strings.Split(strings.Fields(s[(frame*2)+2][1:])[0], "/")
	return fmt.Sprintf("%v:", split[len(split)-1]), nil
}

func log(t *testing.T, ltype logType, format string, args ...interface{}) {
	s := debug.Stack()
	prefix, err := getStackFrame(s, callingFrame)
	args = append([]interface{}{prefix}, args...)
	if err != nil {
		t.Error(err)
		return
	}
	if format == "" {
		switch ltype {
		case errorLog:
			// fmt.Sprintln is used rather than fmt.Sprint because t.Log uses fmt.Sprintln behavior.
			if logger.expected(fmt.Sprintln(args...)) {
				t.Log(args...)
			} else {
				t.Error(args...)
			}
		case fatalLog:
			panic(fmt.Sprint(args...))
		default:
			t.Log(args...)
		}
	} else {
		format = "%v " + format
		switch ltype {
		case errorLog:
			if logger.expected(fmt.Sprintf(format, args...)) {
				t.Logf(format, args...)
			} else {
				t.Errorf(format, args...)
			}
		case fatalLog:
			panic(fmt.Sprintf(format, args...))
		default:
			t.Logf(format, args...)
		}
	}
}

// Update updates the testing.T that the testing logger logs to. Should be done
// before every test.
func Update(t *testing.T) {
	logger.t = t
	logger.errors = nil
}

// Expect declares an error to be expected. For the next test, all error logs
// matching the expression (using FindString) will not cause the test to fail.
// "For the next test" is includes all the time until the next call to Update().
func Expect(expr string) {
	re, err := regexp.Compile(expr)
	if err != nil {
		logger.t.Error(err)
		return
	}
	logger.errors = append(logger.errors, re)
}

func (g *tLogger) expected(s string) bool {
	for _, re := range g.errors {
		if re.FindStringIndex(s) != nil {
			return true
		}
	}
	return false
}

func (g *tLogger) Info(args ...interface{}) {
	log(g.t, logLog, "", args...)
}

func (g *tLogger) Infoln(args ...interface{}) {
	log(g.t, logLog, "", args...)
}

func (g *tLogger) Infof(format string, args ...interface{}) {
	log(g.t, logLog, format, args...)
}

func (g *tLogger) Warning(args ...interface{}) {
	log(g.t, logLog, "", args...)
}

func (g *tLogger) Warningln(args ...interface{}) {
	log(g.t, logLog, "", args...)
}

func (g *tLogger) Warningf(format string, args ...interface{}) {
	log(g.t, logLog, format, args...)
}

func (g *tLogger) Error(args ...interface{}) {
	log(g.t, errorLog, "", args...)
}

func (g *tLogger) Errorln(args ...interface{}) {
	log(g.t, errorLog, "", args...)
}

func (g *tLogger) Errorf(format string, args ...interface{}) {
	log(g.t, errorLog, format, args...)
}

func (g *tLogger) Fatal(args ...interface{}) {
	log(g.t, fatalLog, "", args...)
}

func (g *tLogger) Fatalln(args ...interface{}) {
	log(g.t, fatalLog, "", args...)
}

func (g *tLogger) Fatalf(format string, args ...interface{}) {
	log(g.t, fatalLog, format, args...)
}

func (g *tLogger) V(l int) bool {
	return l <= g.v
}

var lcFailed uint32

type errorer struct {
	t *testing.T
}

func (e errorer) Errorf(format string, args ...interface{}) {
	atomic.StoreUint32(&lcFailed, 1)
	e.t.Errorf(format, args...)
}

// Tester is an implementation of the x interface parameter to
// grpctest.RunSubTests with default Setup and Teardown behavior. Setup updates
// the tlogger and Teardown performs a leak check. Embed in a struct with tests
// defined to use.
type Tester struct{}

// Setup updates the tlogger.
func (Tester) Setup(t *testing.T) {
	Update(t)
}

// Teardown performs a leak check.
func (Tester) Teardown(t *testing.T) {
	if atomic.LoadUint32(&lcFailed) == 1 {
		return
	}
	leakcheck.Check(errorer{t: t})
	if atomic.LoadUint32(&lcFailed) == 1 {
		t.Log("Leak check disabled for future tests")
	}
}
