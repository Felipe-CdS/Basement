// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package activity_views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"nugu.dev/basement/pkg/models"
	"strconv"
	"time"
)

func EditDailyLogModal(d time.Time, tags []models.Tag) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<dialog id=\"editDialog\" x-ref=\"editDialog\" class=\"relative self-center place-self-center py-5 px-6 w-4/12 rounded backdrop:bg-black backdrop:bg-opacity-40 backdrop:backdrop-blur-[2px]\" x-on:keyup.escape=\"modalOpen = false; document.body.classList.remove(&#39;no-scroll&#39;);\"><div class=\"flex justify-between items-center mb-3 space-x-9\"><h1 class=\"inline text-xl font-semibold\">New entry (")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(formatDate(d, time.RFC1123))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/activity_views/modal.partial.templ`, Line: 19, Col: 44}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(")</h1><button x-on:click=\"$refs.editDialog.close()\"><img class=\"size-6\" src=\"/assets/img/x.svg\"></button></div><form id=\"log-form\" hx-post=\"/log/create\" hx-target=\"body\" class=\"flex flex-col items-center\"><input class=\"hidden\" type=\"text\" name=\"date\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(formatDate(d, time.DateOnly))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/activity_views/modal.partial.templ`, Line: 35, Col: 40}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"grid grid-cols-2 w-full\"><label for=\"start\">Start:</label> <input class=\"justify-self-end w-fit\" type=\"time\" id=\"start-input\" name=\"start\" required> <label for=\"start\">End:</label> <input class=\"justify-self-end w-fit\" type=\"time\" id=\"start-input\" name=\"end\" required> <label for=\"start\">Title:</label> <input type=\"text\" id=\"title-input\" name=\"title\" autocomplete=\"off\" placeholder=\"Placeholder...\" class=\"col-span-2 px-2 mb-1 w-full border border-black\"> <label for=\"description\">Description:</label> <textarea id=\"description-input\" name=\"description\" rows=\"8\" autocomplete=\"off\" placeholder=\"Placeholder...\" class=\"col-span-2 px-2 mb-1 w-full border border-black\"></textarea> <label for=\"tags\">Tags:</label><div class=\"grid grid-cols-4 col-span-2 mb-3 w-full\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, t := range tags {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-full\"><input type=\"checkbox\" class=\"inline mr-0.5\" name=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("check-%s", strconv.Itoa(t.ID)))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/activity_views/modal.partial.templ`, Line: 67, Col: 58}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"> <label class=\"inline\" for=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(t.ID))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/activity_views/modal.partial.templ`, Line: 69, Col: 53}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(t.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/activity_views/modal.partial.templ`, Line: 70, Col: 16}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</label></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></form><button form=\"log-form\" type=\"submit\" class=\"p-2 mt-3 w-full font-bold rounded border border-black hover:text-white hover:bg-black hover:bg-opacity-90\">Save</button></dialog>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
