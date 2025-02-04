// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package activity_views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"nugu.dev/basement/pkg/models"
	layouts_view "nugu.dev/basement/views/layouts"
	"time"
)

func Log(calendarLog []models.ActivityDayOverview, partialLog templ.Component, stats models.StatsOverview, loggedUser bool) templ.Component {
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
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1 class=\"text-3xl font-semibold text-center\">/ Habits Log / </h1><p class=\"text-center\">At least sometimes I need to do something...</p><div class=\"grid grid-cols-5 mt-3 w-full border border-gray-100\"><p class=\"col-span-5 p-0.5 text-center text-black bg-gray-100\">Category</p><button hx-get=\"/log\" hx-target=\"body\" hx-push-url=\"true\" class=\"p-0.5 hover:underline\">Overall</button> <button hx-get=\"/log?tab=Reading\" hx-target=\"body\" hx-push-url=\"true\" class=\"p-0.5 hover:underline\">Reading</button> <button hx-get=\"/log?tab=Programming\" hx-target=\"body\" hx-push-url=\"true\" class=\"p-0.5 hover:underline\">Programming</button> <button hx-get=\"/log?tab=Studying\" hx-target=\"body\" hx-push-url=\"true\" class=\"p-0.5 hover:underline\">Studying</button> <button class=\"p-0.5 hover:underline\">Other Tags</button></div><div class=\"flex pb-2 mt-5\"><div class=\"grid grid-cols-1 gap-1 mr-2\" style=\"grid-template-rows: repeat(7, 0.75rem);\"><span class=\"text-sm leading-[.9rem]\">S</span> <span class=\"text-sm leading-[.9rem]\">M</span> <span class=\"text-sm leading-[.9rem]\">T</span> <span class=\"text-sm leading-[.9rem]\">W</span> <span class=\"text-sm leading-[.9rem]\">T</span> <span class=\"text-sm leading-[.9rem]\">F</span> <span class=\"text-sm leading-[.9rem]\">S</span></div><div class=\"grid overflow-y-scroll grid-flow-col gap-1 pb-2 select-none\" style=\"grid-template-columns: repeat(53, 1fr);\n\t\t\t\tgrid-template-rows: repeat(7, 0.75rem);\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, d := range calendarLog {
				var templ_7745c5c3_Var3 = []any{fmt.Sprintf("%s rounded-sm first:row-span-4 first:self-end hover:border hover:border-red-500 size-3",
					getCalendarDayBg(d.TotalSec))}
				templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var3...)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button hx-get=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("log/%s?partial=true", d.Date.Format(time.DateOnly)))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/activity_views/log.page.templ`, Line: 78, Col: 79}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#detailed-log-section\" hx-swap=\"outerHTML\" class=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var3).String())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/activity_views/log.page.templ`, Line: 1, Col: 0}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></button>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><div class=\"grid grid-cols-3 w-full\"><span class=\"text-center\">Daily Average: ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", stats.DailyAverage))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/activity_views/log.page.templ`, Line: 89, Col: 58}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" hours</span> <span class=\"text-center\">Current Streak: ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", stats.CurrentStreak))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/activity_views/log.page.templ`, Line: 92, Col: 60}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" days</span> <span class=\"text-center\">Longest Streak: ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", stats.LongestStreak))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/activity_views/log.page.templ`, Line: 95, Col: 60}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" days</span></div><hr class=\"my-5 w-full\"><div x-data=\"{ modalOpen: false }\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = partialLog.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><script src=\"/assets/js/htmx-2.0.4.min.js\"></script> <script src=\"/assets/js/alpine-3.14.8.min.js\"></script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts_view.StaticHome().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func NoLogSelected() templ.Component {
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
		templ_7745c5c3_Var9 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var9 == nil {
			templ_7745c5c3_Var9 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"detailed-log-section\"><h2 class=\"heading-primary\">/ Day Stats</h2><div><p>Select a day on the heatmap for more info...</p></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func calcDate(selectedYear int, i int) string {

	r := time.Date(selectedYear, 1, 1, 0, 0, 0, 0, time.UTC)
	dur := time.Duration(i) * time.Duration(24) * time.Hour
	r = r.Add(dur)

	return r.Format(time.DateOnly)
}

func getCalendarDayBg(total int) string {

	if total >= 1 && total < 3000 {
		return "bg-red-500"
	}
	if total >= 3000 && total < 7200 {
		return "bg-green-800"
	}
	if total >= 7200 && total < 10800 {
		return "bg-green-700"
	}
	if total >= 10800 && total < 14400 {
		return "bg-green-600"
	}
	if total >= 14400 && total < 18000 {
		return "bg-green-400"
	}
	if total >= 18000 {
		return "bg-green-300"
	}
	return "bg-gray-800"
}
