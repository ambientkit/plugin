package docgen

import (
	"bytes"
	"context"
	"fmt"
	"go/ast"
	"go/doc"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/ambientkit/ambient"
	"golang.org/x/tools/imports"
)

// Generate will generate the README.md file for the plugin.
func Generate(t *testing.T, p interface{}, outputFile string) {
	ctx := context.Background()

	if len(outputFile) == 0 {
		outputFile = "README.md"
	}

	// Create doc store.
	docInfo := new(Doc)
	docInfo.PackagePath = packagePath(p)

	// Get the working directory.
	dir, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	// Parse the current package.
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		t.Error(err)
	}

	// Loop through packages to build up files for documentation.
	files := make([]*ast.File, 0)
	for _, v := range pkgs {
		for fileName, f := range v.Files {
			files = append(files, f)

			// Take the first example and then stop.
			if strings.HasSuffix(fileName, "_test.go") {
				efs := token.NewFileSet()
				ef, err := parser.ParseFile(efs, fileName, nil, parser.ParseComments)
				if err != nil {
					t.Error(err)
				}
				examples := doc.Examples(ef)
				for _, v := range examples {
					if len(docInfo.Example) == 0 {
						docInfo.Example = formatFile(t, efs, v.Play)
						break
					}
				}
			}
		}
	}

	// Parse files as documentation.
	p2, err := doc.NewFromFiles(fset, files, docInfo.PackagePath)
	if err != nil {
		t.Error(err)
	}

	docInfo.PackageDescription = StripNewline(p2.Doc)

	// Determine if the plugin satisfies the core plugin interface.
	pluginCore, isPluginCore := p.(ambient.PluginCore)
	if isPluginCore {
		docInfo.PackageVersion = pluginCore.PluginVersion(ctx)
		docInfo.PackageName = pluginCore.PluginName(ctx)

		if docInfo.PackageName != p2.Name {
			t.Error(fmt.Errorf("package names do not match"))
		}
	}

	// Determine if the plugin is runnable.
	plugin, isPlugin := p.(ambient.Plugin)
	if isPlugin {
		os.Setenv("AMB_LOGLEVEL", "FATAL")
		app := LighweightAppSetup(ctx, "docgen", plugin, true)

		docInfo.GrantRequests = plugin.GrantRequests(ctx)
		docInfo.Settings = plugin.Settings(ctx)
		docInfo.FuncMap = plugin.FuncMap(ctx)
		assets, embedfs := plugin.Assets(ctx)
		docInfo.Assets = assets
		docInfo.EmbeddedFS = embedfs
		docInfo.Routes = app.Mux.Routes(ctx)
	}

	// Determine if the plugin contains middleware.
	mw, ok := p.(ambient.MiddlewarePlugin)
	if ok {
		docInfo.Middleware = mw.Middleware(ctx)
	}

	// Determine if the plugin satisfies the logger interface.
	_, ok = p.(ambient.LoggingPlugin)
	if ok {
		docInfo.Logger = true
	}

	// Determine if the plugin satisfies the storage system interface.
	_, ok = p.(ambient.StoragePlugin)
	if ok {
		docInfo.StorageSystem = true
	}

	// Determine if the plugin satisfies the router interface.
	_, ok = p.(ambient.RouterPlugin)
	if ok {
		docInfo.Router = true
	}

	// Determine if the plugin satisfies the template engine interface.
	_, ok = p.(ambient.TemplateEnginePlugin)
	if ok {
		docInfo.TemplateEngine = true
	}

	// Determine if the plugin satisfies the session manager interface.
	_, ok = p.(ambient.SessionManagerPlugin)
	if ok {
		docInfo.SessionManager = true
	}

	docInfo.Output(outputFile)
}

func packagePath(v interface{}) string {
	// Source: https://stackoverflow.com/a/60846213
	if v == nil {
		return ""
	}

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		return val.Elem().Type().PkgPath()
	}
	return val.Type().PkgPath()
}

func formatFile(t *testing.T, fset *token.FileSet, n *ast.File) string {
	if n == nil {
		return ""
	}
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, n); err != nil {
		t.Fatal(err)
	}

	// Order the imports.
	b, err := imports.Process("", buf.Bytes(), &imports.Options{
		Fragment:  true,
		AllErrors: true,
		Comments:  true,
		TabIndent: true,
		TabWidth:  8,
	})
	if err != nil {
		t.Fatal(err)
	}

	return string(b)
}
