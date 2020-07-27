// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package msalbase

type IAccessToken interface {
	GetSecret() string
	GetExpiresOn() string
	GetScopes() string
}
