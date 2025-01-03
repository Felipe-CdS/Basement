// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package layouts_view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func StaticHome() templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"mx-auto w-full bg-center bg-cover lg:w-10/12 xl:w-7/12 min-h-48\" style=\"background-image: url(assets/img/kyubey.png);\"></div><h1 class=\"mx-auto mt-2 text-4xl heading-primary font-kr\">누구세요?</h1><div id=\"content\" class=\"flex justify-center self-center mt-10 w-full lg:w-full xl:w-7/12\"><div id=\"side\" class=\"overflow-hidden w-1/3 border-r lg:w-2/12 mr-1/12\"><ul><span>Static Content</span><li class=\"list-none hover:text-gray-500\"><a href=\"/reads\">Reads</a></li><li class=\"list-none hover:text-gray-500\"><a href=\"/bookmarks\">Bookmarks</a></li><li class=\"list-none hover:text-gray-500\"><a href=\"/log\">Log</a></li></ul><hr class=\"my-5 w-11/12\"><ul><span>JS Required pages</span><li class=\"list-none hover:text-gray-500\"><a href=\"/gallery\">Gallery</a></li><li class=\"list-none hover:text-gray-500\"><a href=\"/activities\">Activities</a></li></ul><hr class=\"my-5 w-11/12\"><ul class=\"space-y-5\"><li class=\"list-none\"><audio class=\"w-9/12\" id=\"noise\" loop controls src=\"/assets/audio/noise.mp3\"></audio></li><li class=\"list-none\"><audio class=\"w-9/12\" id=\"noise\" loop controls src=\"/assets/audio/rain.mp3\"></audio></li></ul></div><div class=\"flex overflow-hidden flex-col px-10 w-2/3\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
