module github.com/terhitormanen/cmd

go 1.16

retract (
	v1.1.1
	v1.1.0 // v1.1.0-1.1.1 are failed releases
)

require (
	github.com/agtorre/gocolorize v1.0.0
	github.com/fsnotify/fsnotify v1.5.1
	github.com/jessevdk/go-flags v1.4.0
	github.com/mattn/go-colorable v0.1.12
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.1
	github.com/terhitormanen/config v1.0.1
	github.com/terhitormanen/log15 v2.11.20+incompatible
	github.com/terhitormanen/revel v1.1.1
	golang.org/x/tools v0.1.10
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/stack.v0 v0.0.0-20141108040640-9b43fcefddd0
)
