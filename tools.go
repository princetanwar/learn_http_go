//go:build tools
// +build tools

// "// +build tools"  this tell go compiler to only include this file in final build if tools tag is provided

package tools

// Import the package as a blank import to include it in the module.
import _ "github.com/unknwon/bra"
