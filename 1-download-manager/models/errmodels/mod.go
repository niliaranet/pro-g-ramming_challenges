package errmodels

import "errors"

var TooFewArguments error = errors.New("too few arguments")
