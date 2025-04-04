package main

type authenticationInfo struct {
	username string
	password string
}

// create the method below
func (authZ authenticationInfo) getBasicAuth() string {
	basicAuth := "Authorization: Basic " + authZ.username + ":" + authZ.password
	return basicAuth
}
