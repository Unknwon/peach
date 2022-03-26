// Copyright 2015 ASoulDocs. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package oldroutes

import (
	"fmt"
	"strings"

	"github.com/unknwon/com"

	"github.com/asoul-sig/asouldocs/pkg/context"
	"github.com/asoul-sig/asouldocs/pkg/setting"

	"github.com/asoul-sig/asouldocs/oldmodels"
)

func Pages(ctx *context.Context) {
	toc := oldmodels.Tocs[ctx.Locale.Language()]
	if toc == nil {
		toc = oldmodels.Tocs[setting.Docs.Langs[0]]
	}

	pageName := strings.ToLower(strings.TrimSuffix(ctx.Req.URL.Path[1:], ".html"))
	for i := range toc.Pages {
		if toc.Pages[i].Name == pageName {
			page := toc.Pages[i]
			langVer := toc.Lang
			if !com.IsFile(page.FileName) {
				ctx.Data["IsShowingDefault"] = true
				langVer = setting.Docs.Langs[0]
				page = oldmodels.Tocs[langVer].Pages[i]
			}
			if !setting.ProdMode {
				_ = page.ReloadContent()
			}

			ctx.Data["Title"] = page.Title
			ctx.Data["Content"] = fmt.Sprintf(`<script type="text/javascript" src="/%s/%s?=%d"></script>`, langVer, page.DocumentPath+".js", page.LastBuildTime)
			ctx.Data["Pages"] = toc.Pages

			renderEditPage(ctx, page.DocumentPath)
			ctx.HTML(200, "docs")
			return
		}
	}

	NotFound(ctx)
}

func NotFound(ctx *context.Context) {
	ctx.Data["Title"] = "404"
	ctx.HTML(404, "404")
}