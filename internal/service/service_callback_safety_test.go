// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package service_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestUnitServiceCallbacksDoNotTerminateProviderProcess(t *testing.T) {
	err := filepath.WalkDir(".", func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || filepath.Ext(path) != ".go" || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		fileSet := token.NewFileSet()
		file, err := parser.ParseFile(fileSet, path, nil, 0)
		if err != nil {
			return err
		}
		imports := importAliases(file)
		ast.Inspect(file, func(node ast.Node) bool {
			call, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}
			position := fileSet.Position(call.Pos())
			if selector, ok := call.Fun.(*ast.SelectorExpr); ok {
				identifier, ok := selector.X.(*ast.Ident)
				if ok {
					packagePath, imported := imports[identifier.Name]
					if identifier.Obj == nil && imported && terminatesSharedProvider(packagePath, selector.Sel.Name) {
						t.Errorf("%s:%d calls %s.%s, which terminates the shared provider process", path, position.Line, packagePath, selector.Sel.Name)
					}
				}
			}
			if identifier, ok := call.Fun.(*ast.Ident); ok && identifier.Name == "panic" {
				t.Errorf("%s:%d calls panic, which terminates the shared provider process", path, position.Line)
			}
			return true
		})
		return nil
	})
	if err != nil {
		t.Fatalf("cannot inspect service callbacks: %v", err)
	}
}

func importAliases(file *ast.File) map[string]string {
	aliases := make(map[string]string, len(file.Imports))
	for _, spec := range file.Imports {
		importPath, err := strconv.Unquote(spec.Path.Value)
		if err != nil || importPath == "" || spec.Name != nil && (spec.Name.Name == "_" || spec.Name.Name == ".") {
			continue
		}
		name := path.Base(importPath)
		if spec.Name != nil {
			name = spec.Name.Name
		}
		aliases[name] = importPath
	}
	return aliases
}

func TestUnitTerminatingServiceCallsAreRecognized(t *testing.T) {
	for _, call := range []struct {
		packageName  string
		functionName string
	}{
		{packageName: "log", functionName: "Fatal"},
		{packageName: "log", functionName: "Fatalf"},
		{packageName: "log", functionName: "Fatalln"},
		{packageName: "log", functionName: "Panic"},
		{packageName: "log", functionName: "Panicf"},
		{packageName: "log", functionName: "Panicln"},
		{packageName: "os", functionName: "Exit"},
	} {
		if !terminatesSharedProvider(call.packageName, call.functionName) {
			t.Errorf("%s.%s was not recognized as terminating", call.packageName, call.functionName)
		}
	}
	if terminatesSharedProvider("log", "Printf") {
		t.Error("log.Printf was incorrectly recognized as terminating")
	}
}

func TestUnitTerminatingServiceCallsResolveImportAliases(t *testing.T) {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, "fixture.go", `package fixture
import stdlog "log"
func callback() { stdlog.Panic("stop") }
`, 0)
	if err != nil {
		t.Fatalf("cannot parse fixture: %v", err)
	}
	imports := importAliases(file)
	if got := imports["stdlog"]; got != "log" {
		t.Fatalf("stdlog import resolved to %q, want log", got)
	}

	file, err = parser.ParseFile(fileSet, "fixture.go", `package fixture
type logger struct{}
func (logger) Panic(string) {}
func callback() { log := logger{}; log.Panic("safe") }
`, 0)
	if err != nil {
		t.Fatalf("cannot parse shadowing fixture: %v", err)
	}
	var localSelector *ast.SelectorExpr
	ast.Inspect(file, func(node ast.Node) bool {
		if selector, ok := node.(*ast.SelectorExpr); ok && selector.Sel.Name == "Panic" {
			localSelector = selector
		}
		return true
	})
	identifier, ok := localSelector.X.(*ast.Ident)
	if !ok || identifier.Obj == nil {
		t.Fatal("shadowed log identifier was not resolved as a local object")
	}
}

func terminatesSharedProvider(packageName, functionName string) bool {
	switch packageName {
	case "log":
		switch functionName {
		case "Fatal", "Fatalf", "Fatalln", "Panic", "Panicf", "Panicln":
			return true
		}
	case "os":
		return functionName == "Exit"
	}
	return false
}
