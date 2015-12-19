package views

import (
    "net/url"
    "fmt"
    "strconv"
)

var _ = fmt.Println

type Form struct {
    Fields map[string]Field // - name required
    SuccessUrl string
    Validate func(values url.Values) []string
}

type Field struct {
    Required bool
    Type string
}

func (form Form) HandleValidate(values url.Values) []string {
    var errors []string
    for name, field := range form.Fields {
        switch field.Type { // specific validation
            case "number": 
                _, err := strconv.Atoi(values[name][0])
                if err != nil {
                    errors = append(errors, fmt.Sprintf("Field \"%v\" must be a number.", name))
                }
        }
        if (len(values[name]) == 0 || values[name][0] == "") && field.Required {
            errors = append(errors, fmt.Sprintf("Field \"%v\" is required.", name))
        }
    }
    if form.Validate != nil {
        return form.Validate(values)
    }
    return  errors
}