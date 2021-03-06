package profile

import (
	"io/ioutil"
	"path/filepath"
)

// NamedRequest is a representation of a request that can be loaded from a profile.
type NamedRequest struct {
	AllowInsecure     bool
	Body              string
	FileToUpload      string
	Headers           map[string][]string
	Method            string
	Name              string
	PostProcessScript string
	Source            string // File where this request was loaded from
	URL               string
	Values            map[string][]string
}

// GetAllowInsecure returns if this named request allow insecure HTTP connections
func (req NamedRequest) GetAllowInsecure() bool {
	return req.AllowInsecure
}

// GetBody returns the body for this NamedRequest
func (req NamedRequest) GetBody() (string, error) {
	if req.Body != "" {
		return req.Body, nil
	}

	// If there's a file to upload from profile, load it
	if req.FileToUpload != "" {
		path := req.FileToUpload

		// If not absolute, it's relative to the profiles dir
		if !filepath.IsAbs(path) {
			profileDir, profileDirError := GetProfilesDir()
			if profileDirError != nil {
				return "", profileDirError
			}
			path = filepath.Join(profileDir, req.FileToUpload)
		}

		data, loadErr := ioutil.ReadFile(path)
		if loadErr != nil {
			return "", loadErr
		}

		return string(data), nil
	}

	return "", nil
}

// GetHeaders returns the headers for this NamedRequest
func (req NamedRequest) GetHeaders() map[string][]string {
	return req.Headers
}

// GetMethod return the HTTP method for this NamedRequest
func (req NamedRequest) GetMethod() string {
	return req.Method
}

// GetValues returns the values for this NamedRequest
func (req NamedRequest) GetValues() map[string][]string {
	return req.Values
}
