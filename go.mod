module github.com/silencily/sparktime

go 1.12

require (
	github.com/go-kivik/couchdb v0.0.0-20190520192200-57eca2c0435d
	github.com/go-kivik/kivik v0.0.0-20190509113000-22b54eda3db0
	github.com/kataras/golog v0.0.0-20180321173939-03be10146386
	github.com/kataras/iris v0.0.0-20190526035200-63c26dc97890
	github.com/mojocn/base64Captcha v0.0.0-20190509175000-87c9c59224d8
	github.com/robfig/cron/v3 v3.0.0-rc1
	golang.org/x/net v0.0.0-20181114220301-adae6a3d119a
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace (
	golang.org/x/crypto v0.0.0-20181112202954-3d3f9f413869 => github.com/golang/crypto v0.0.0-20180803061100-56440b844dfe
	golang.org/x/image v0.0.0-20190501045829-6d32002ffd75 => github.com/golang/image v0.0.0-20190615100300-92942e4437e2
	golang.org/x/net v0.0.0-20180906233101-161cd47e91fd => github.com/golang/net v0.0.0-20181005061600-48a9fcba44db
	golang.org/x/net v0.0.0-20181023162649-9b4f9f5ad519 => github.com/golang/net v0.0.0-20181005061600-48a9fcba44db
	golang.org/x/net v0.0.0-20181114220301-adae6a3d119a => github.com/golang/net v0.0.0-20181005061600-48a9fcba44db
	golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f => github.com/golang/sync v0.0.0-20190423061100-112230192c58
	golang.org/x/sys v0.0.0-20180909124046-d0be0721c37e => github.com/golang/sys v0.0.0-20190621125100-bf70e4678053
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 => github.com/natefinch/lumberjack v2.0.0+incompatible
)
