/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package log

import (
	"mosn.io/pkg/log"
)

// errorLogger is a default implementation of ErrorLogger
// we use ErrorLogger to write common log message.
type errorLogger struct {
	*log.SimpleErrorLog
}

func CreateDefaultErrorLogger(output string, level log.Level) (log.ErrorLogger, error) {
	lg, err := log.GetOrCreateLogger(output, nil)
	if err != nil {
		return nil, err
	}
	return &errorLogger{
		SimpleErrorLog: &log.SimpleErrorLog{
			Logger:    lg,
			Formatter: log.DefaultFormatter,
			Level:     level,
		},
	}, nil
}

// default logger error level format:
// {time} [{level}] [{error code}] {content}
// default error code is normal
const defaultErrorCode = "normal"

func (l *errorLogger) Errorf(format string, args ...interface{}) {
	if l.Disable() {
		return
	}
	if l.Level >= log.ERROR {
		s := l.SimpleErrorLog.Formatter(log.ErrorPre, defaultErrorCode, format)
		l.Logger.Printf(s, args...)
	}
}
