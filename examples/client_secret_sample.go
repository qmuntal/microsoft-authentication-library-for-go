// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package main

import (
	log "github.com/sirupsen/logrus"

	msalgo "github.com/AzureAD/microsoft-authentication-library-for-go/src/msal"
)

func acquireTokenClientSecret() {
	confidentialClientApp := msalgo.CreateConfidentialClientApplication(
		confidentialConfig.ClientID, confidentialConfig.Authority)
	clientSecretParams := msalgo.CreateAcquireTokenClientCredentialsParameters(
		confidentialConfig.Scopes, confidentialConfig.ClientSecret)
	result, err := confidentialClientApp.AcquireTokenByClientSecret(clientSecretParams)
	if err != nil {
		log.Fatal(err)
	}
	accessToken := result.GetAccessToken()
	log.Info("Access token is: " + accessToken)
}
