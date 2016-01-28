package client

import (
  "errors"
  "encoding/json"

  "github.com/parnurzeal/gorequest"
)


type WASP struct {
  Url string
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
