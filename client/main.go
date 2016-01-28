package client

import (
  "errors"
  "encoding/json"

  "github.com/parnurzeal/gorequest"
)


type WASP struct {
  Url string
  AuthToken string
}

var Request = gorequest.New()

type Error struct {
  Path string     `json:"path"`
  Message string  `json:"message"`
}

// Get the configuration as a json string for the basepath
func (w *WASP) Get(basepath string) (string, error) {
  return handle(Request.
    Get(w.Url + "/configuration").
    Query("path=" + basepath).
    End())
}

// Get the list of available keys from the basepath
func (w *WASP) List(basepath string) (string, error) {
  return handle(Request.
    Get(w.Url + "/configuration/keys").
    Query("path=" + basepath).
    End())
}

func (w *WASP) Put(path string, config string) (string, error) {
  return handle(Request.
    Post(w.Url + "/configuration").
    Query("path=" + path).
    Query("token=" + w.AuthToken).
    Type("text").
    Send(config).
    End())
}

// Delete a configuration (it deletes recursively)
func (w *WASP) Delete(path string) (string, error) {
  return handle(Request.
    Delete(w.Url + "/configuration").
    Query("path=" + path).
    Query("token=" + w.AuthToken).
    Type("text").
    End())
}

func (w *WASP) SetToken(token string) {
  w.AuthToken = token
}

func handle(response gorequest.Response, body string, errs []error) (string, error) {
  if(response.StatusCode != 200) {
    var errorResponse Error
    json.Unmarshal([]byte(body), &errorResponse)
    errs = append(errs, errors.New(errorResponse.Message))
  }

  e := combineErrors(errs)
  return string(body), e  
}

func combineErrors(errs []error) error {
  if(len(errs) == 1) {
    return errs[0]
  } else if(len(errs) > 1) {
    msg := "Error(s):"
    for _, err := range errs {
      msg += " " + err.Error()
    }
    return errors.New(msg)
  } else {
    return nil
  }
}
