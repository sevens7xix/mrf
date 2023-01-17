package mrf

import "net/http"

type RoundTripFunc func(req *http.Request) *http.Response

// The Roundtripper will allows us to makes a single HTTP Transaction, and obtain a response for a given request
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

/*
So we're going to create a method that allows us to return a
http client with the transport replaced to avoid making real calls.
*/
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}
