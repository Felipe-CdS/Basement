// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package activity_views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"nugu.dev/basement/pkg/models"
	"strings"
	"time"
)

func DetailedLog(d time.Time, log []models.Activity) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"detailed-log-section\"><div class=\"flex justify-between items-center\"><h2 class=\"heading-primary\">/ Day Stats - ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(formatDate(d, time.RFC1123))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/activity_views/detailed_log.partial.templ`, Line: 13, Col: 47}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><button @click=\"$refs.editDialog.showModal()\" class=\"py-0.5 px-2 border border-white\">Add new entry</button></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(log) != 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p>No info for this day...</p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, l := range log {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(l.Title)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/activity_views/detailed_log.partial.templ`, Line: 25, Col: 18}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h2 class=\"mt-3 text-lg font-semibold\">#1) 10:20AM - 11:10AM | Studied about inotify in C </h2><p class=\"my-2 w-full text-justify indent-10\">The counter is now working and the cookie bugs seem fixed. The solution isn't the prettier one but it works. Studied a little bit about htmx events too and added it to the project. Don't know I'm doing it right though.</p><p class=\"mb-3\">Tags: Programming | Sapos</p><hr class=\"mb-0.5 hr-divider-dashed\"><hr class=\"hr-divider-dashed\"><h2 class=\"mt-3 text-lg font-semibold\">#2) 11:20 - 12:10 - Chapter 5</h2><p class=\"my-3 w-full text-justify indent-10\"></p><p class=\"mb-3\">Tags: Reading | Tao </p><hr class=\"mb-0.5 hr-divider-dashed\"><hr class=\"hr-divider-dashed\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = EditDailyLogModal(d).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func formatDate(d time.Time, timeFormat string) string {
	holder := d.Format(timeFormat)
	return strings.TrimSuffix(holder, " 00:00:00 UTC")
}
