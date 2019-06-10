module github.com/silencily/sparktime

require (
	github.com/go-kivik/couchdb v0.0.0-20190520192200-57eca2c0435d
	github.com/go-kivik/kivik v0.0.0-20190509113000-22b54eda3db0
	github.com/kataras/iris v0.0.0-20190526035200-63c26dc97890
	golang.org/x/net v0.0.0-20181114220301-adae6a3d119a
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace (
	golang.org/x/crypto v0.0.0-20181112202954-3d3f9f413869 => github.com/golang/crypto v0.0.0-20180803061100-56440b844dfe
	golang.org/x/net v0.0.0-20181023162649-9b4f9f5ad519 => github.com/golang/net v0.0.0-20181005061600-48a9fcba44db
	golang.org/x/net v0.0.0-20181114220301-adae6a3d119a => github.com/golang/net v0.0.0-20181005061600-48a9fcba44db
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 => github.com/natefinch/lumberjack v2.0.0+incompatible
)
