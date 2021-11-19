package middleware

import (
	"bytes"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Recovery(ctx *gin.Context) {
	defer func(c *gin.Context) {
		if err := recover(); err != nil {
			// zlog.Warn(c, err)
			// 请求url
			path, raw := c.Request.URL.Path, c.Request.URL.RawQuery
			if raw != "" {
				path = path + "?" + raw
			}
			// 请求报文
			body, _ := ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

			// zlog.NewEntry(zlog.SugaredLogger).WithFields(zlog.Fields{
			// 	"logId":     zlog.GetLogID(c),
			// 	"uri":       path,
			// 	"refer":     c.Request.Referer(),
			// 	"cookie":    c.Request.Cookies(),
			// 	"client_ip": c.ClientIP(),
			// 	"module":    env.AppName,
			// 	"ua":        c.Request.UserAgent(),
			// 	"host":      c.Request.Host,
			// }).Info("Panic[recover]")
			// base.RenderJsonAbort(c, component.ErrServer)
		}
	}(ctx)
	ctx.Next()
}
