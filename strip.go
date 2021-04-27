package main

import (
	"bytes"
	"github.com/z7zmey/php-parser/pkg/visitor/printer"
)

func Strip(code []byte) ([]byte, error) {
	buf := &bytes.Buffer{}

	code = bytes.TrimSpace(code)

	node, err := Parse(code)
	if err != nil {
		return nil, err
	}

	stripPrinter := printer.NewPrinter(buf).IgnoreFreeFloating()
	node.Accept(stripPrinter)

	return buf.Bytes(), nil
}