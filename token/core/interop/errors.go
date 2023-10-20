/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package interop

import "github.com/pkg/errors"

var (
	// TokenExistsError is returned when the token already exists
	TokenExistsError = errors.New("token exists")
	// TokenDoesNotExistError is returned when the token does not exist
	TokenDoesNotExistError = errors.New("token does not exists")
)
