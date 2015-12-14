package views

import (
    "net/url"
)

type Form struct {
    Fields map[string]bool // - name required
}

func (form Form) Validate(values url.Values) bool {
    for name, required := range form.Fields {
        if values[name][0] == "" && required { // do separate types?  will have to map string to struct of bool/type then
            return false
        }
    }
    return true
}