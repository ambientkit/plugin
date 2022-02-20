package docgen

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/ambientkit/ambient"
	"github.com/ambientkit/plugin/router/routerecorder"
	"github.com/oxtoacart/bpool"
)

//go:embed *.tmpl
var assets embed.FS

// Doc represents the variables passed to a doc template.
type Doc struct {
	PackageName        string
	PackagePath        string
	PackageDescription string
	PackageVersion     string
	GrantRequests      []ambient.GrantRequest
	Settings           []ambient.Setting
	FuncMap            func(r *http.Request) template.FuncMap
	Assets             []ambient.Asset
	EmbeddedFS         *embed.FS
	Routes             []routerecorder.Route
	Middleware         []func(next http.Handler) http.Handler
	Example            string

	Logger         bool
	StorageSystem  bool
	Router         bool
	TemplateEngine bool
	SessionManager bool
}

// bufpool is used to write out HTML after it's been executed and before it's
// written to the ResponseWriter to catch any partially written templates.
var bufpool *bpool.BufferPool = bpool.NewBufferPool(64)

func (d *Doc) funcMap() template.FuncMap {
	fm := make(template.FuncMap)
	fm["GrantRequests"] = func() string {
		out := ""

		if len(d.GrantRequests) == 0 {
			out += "The plugin does not request any grants."
		} else {
			out += fmt.Sprintf("The plugin request the following grants (%v):\n\n", len(d.GrantRequests))
			for _, v := range d.GrantRequests {
				out += fmt.Sprintf("- **Name**: %v\n", v.Grant)
				out += fmt.Sprintf("  - **Description**: %v\n", v.Description)
			}
		}

		return StripNewline(out)
	}
	fm["Settings"] = func() string {
		out := ""

		if len(d.Settings) == 0 {
			out += "The plugin does not have any settings."
		} else {
			out += fmt.Sprintf("The plugin has the follow settings (%v):\n\n", len(d.Settings))
			for _, v := range d.Settings {
				out += fmt.Sprintf("- **Name**: %v\n", v.Name)
				if len(v.Type) > 0 {
					out += fmt.Sprintf("  - **Type**: %v\n", v.Type)
				} else {
					out += fmt.Sprintf("  - **Type**: %v\n", "input")
				}
				if len(v.Description.Text) > 0 {
					out += fmt.Sprintf("  - **Description**: %v\n", v.Description.Text)
					if len(v.Description.URL) > 0 {
						out += fmt.Sprintf("    - **URL**: %v\n", v.Description.URL)
					}
				}

				out += fmt.Sprintf("  - **Hidden**: %v\n", v.Hide)

				if v.Default != nil {
					out += fmt.Sprintf("  - **Has Default**: %v\n", true)
				}
			}
		}

		return StripNewline(out)
	}
	fm["FuncMap"] = func() string {
		out := ""
		fmr := d.FuncMap
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			return fmt.Sprintf("Error generating request: %v\n", err.Error())
		}
		if fmr == nil {
			out += "The plugin does not have a FuncMap.\n"
		} else {
			fm := fmr(req)
			out += fmt.Sprintf("The plugin has the follow FuncMap items (%v):\n\n", len(fm))

			arr := make([]string, 0)
			for k := range fm {
				arr = append(arr, k)
			}

			sort.Strings(arr)

			for _, v := range arr {
				out += fmt.Sprintf("  - {{%v}}\n", v)
			}
		}

		return StripNewline(out)
	}
	fm["Assets"] = func() string {
		out := ""
		if len(d.Assets) == 0 {
			out += "The plugin does not inject any assets."
		} else {
			out += fmt.Sprintf("The plugin injects the following assets (%v):\n\n", len(d.Assets))
			for _, v := range d.Assets {
				out += fmt.Sprintf("  - **Type:** %v\n", v.Filetype)
				out += fmt.Sprintf("    - **Location:** %v\n", v.Location)
				if v.Auth != "" {
					out += fmt.Sprintf("    - **Auth Type:** %v\n", v.Auth)
				}
				if v.External {
					out += fmt.Sprintf("    - **External:** %v\n", v.External)
				}
				if v.Inline {
					out += fmt.Sprintf("    - **Inline:** %v\n", v.Inline)
				}
				if v.SkipExistCheck {
					out += fmt.Sprintf("    - **Skip Exist Check:** %v\n", v.SkipExistCheck)
				}
				if len(v.Path) > 0 {
					out += fmt.Sprintf("    - **Path:** %v\n", v.Path)
				}
				if len(v.Content) > 0 {
					out += fmt.Sprintf("    - **Has Content:** %v\n", true)
				}
				if len(v.TagName) > 0 {
					out += fmt.Sprintf("    - **Tag Name:** %v\n", v.TagName)
				}
				if v.ClosingTag {
					out += fmt.Sprintf("    - **Closing Tag:** %v\n", v.ClosingTag)
				}
				if len(v.LayoutOnly) > 0 {
					arr := make([]string, 0)
					for _, layout := range v.LayoutOnly {
						arr = append(arr, string(layout))
					}
					out += fmt.Sprintf("    - **Layout Only:** %v\n", strings.Join(arr, ","))
				}
				if len(v.Attributes) > 0 {
					out += fmt.Sprintf("    - **Attributes (%v):** %v\n", len(v.Attributes), v.Auth)
					for _, attr := range v.Attributes {
						out += fmt.Sprintf("      - **Name:** %v | **Value:** %v\n", attr.Name, attr.Value)
					}
				}
				if len(v.Replace) > 0 {
					out += fmt.Sprintf("    - **Replace (%v):**\n", len(v.Replace))
					for _, repl := range v.Replace {
						out += fmt.Sprintf("      - **Find:** %v | **Replace:** %v\n", repl.Find, repl.Replace)
					}
				}
			}
		}

		return StripNewline(out)
	}
	fm["EmbeddedFS"] = func() string {
		out := ""
		if d.EmbeddedFS == nil {
			out += "The plugin does not have any embedded files.\n"
		} else {
			out += "The plugin has embedded files.\n"
		}

		return StripNewline(out)
	}
	fm["Routes"] = func() string {
		out := ""
		if len(d.Routes) == 0 {
			out += "The plugin does not have any routes.\n"
		} else {
			out += fmt.Sprintf("The plugin has the following routes (%v):\n", len(d.Routes))
			for _, v := range d.Routes {
				out += fmt.Sprintf("  - **Method:** %v | **Path:** %v\n", v.Method, v.Path)
			}
		}

		return StripNewline(out)
	}
	fm["Middleware"] = func() string {
		out := ""
		if len(d.Middleware) == 0 {
			out += "The plugin does not have any middleware.\n"
		} else {
			out += fmt.Sprintf("The plugin has middleware (%v).\n", len(d.Middleware))
		}

		return StripNewline(out)
	}
	fm["Example"] = func() template.HTML {
		out := ""
		if len(d.Example) == 0 {
			out += "There is no example usage for the plugin."
		} else {
			out += "```go\n"
			out += d.Example
			out += "```\n"
		}

		return template.HTML(StripNewline(out))
	}

	return fm
}

// StripNewline strips newlines from the end of a string.
func StripNewline(s string) string {
	out := s
	out = strings.TrimSuffix(out, "\r\n")
	out = strings.TrimSuffix(out, "\n")
	return out
}

// Output will write the plugin doc using the template.
func (d *Doc) Output(filename string) {
	vars := make(map[string]interface{})
	vars["PackageName"] = d.PackageName
	vars["PackagePath"] = d.PackagePath
	vars["PackageDescription"] = d.PackageDescription
	vars["PackageVersion"] = d.PackageVersion
	vars["Logger"] = d.Logger
	vars["StorageSystem"] = d.StorageSystem
	vars["Router"] = d.Router
	vars["TemplateEngine"] = d.TemplateEngine
	vars["SessionManager"] = d.SessionManager

	// Write temporarily to a buffer pool.
	buf := bufpool.Get()
	defer bufpool.Put(buf)

	templatePath := "template.tmpl"
	tmpl, err := template.New(path.Base(templatePath)).Funcs(d.funcMap()).ParseFS(assets, templatePath)
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = tmpl.Execute(buf, vars)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Output the readme.
	err = os.WriteFile(filename, buf.Bytes(), 0644)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
