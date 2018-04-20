/*
Copyright 2018 Heptio Inc.
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

package util

import (
	"log"
	"regexp"

	"github.com/sirupsen/logrus"
)

// GetFormatter returns a textformatter to customize logs
func GetFormatter() *logrus.TextFormatter {
	return &logrus.TextFormatter{
		FullTimestamp: true,
	}
}

// IsInvalidClusterName returns true if valid cluster name
func IsInvalidClusterName(clustername string) bool {
	matched, err := regexp.MatchString("^[a-zA-Z0-9]+[a-zA-Z0-9-]*[a-zA-Z0-9]+$", clustername)
	if err != nil {
		log.Fatal(err)
	}
	return !matched
}
