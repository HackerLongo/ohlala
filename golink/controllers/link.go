package controllers

import (
    "github.com/QLeelulu/goku"
    "github.com/QLeelulu/ohlala/golink/filters"
    "github.com/QLeelulu/ohlala/golink/forms"
    "github.com/QLeelulu/ohlala/golink/models"
    "strconv"
)

var _ = goku.Controller("link").
    // 
    Get("index", func(ctx *goku.HttpContext) goku.ActionResulter {

    return ctx.View(nil)
}).
    /**
     * 查看一个链接的评论
     */
    Get("show", func(ctx *goku.HttpContext) goku.ActionResulter {

    linkId, _ := strconv.ParseInt(ctx.RouteData.Params["id"], 10, 64)
    link, _ := models.Link_GetById(linkId)

    if link == nil {
        ctx.ViewData["errorMsg"] = "内容不存在"
        return ctx.Render("error", nil)
    }

    // links := models.Link_ByUser(link.Id, 1, 10)

    // ctx.ViewData["Comments"] = links
    return ctx.View(link)
}).

    /**
     * 提交链接的表单页面
     */
    Get("submit", func(ctx *goku.HttpContext) goku.ActionResulter {

    ctx.ViewData["Values"] = map[string]string{
        "title":   ctx.Get("t"),
        "context": ctx.Get("u"),
    }
    return ctx.View(nil)

}).Filters(filters.NewRequireLoginFilter()).

    /**
     * 提交一个链接并保存到数据库
     */
    Post("submit", func(ctx *goku.HttpContext) goku.ActionResulter {

    f := forms.CreateLinkSubmitForm()
    f.FillByRequest(ctx.Request)

    success, errorMsgs := models.Link_SaveForm(f, (ctx.Data["user"].(*models.User)).Id)

    if success {
        return ctx.Redirect("/")
    } else {
        ctx.ViewData["Errors"] = errorMsgs
        ctx.ViewData["Values"] = f.Values()
    }
    return ctx.View(nil)

}).Filters(filters.NewRequireLoginFilter()).

    /**
     * 添加评论
     */
    Post("comment", func(ctx *goku.HttpContext) goku.ActionResulter {

    return ctx.View(nil)

}).Filters(filters.NewRequireLoginFilter())