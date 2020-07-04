module github.com/klische/gopenpgp/v2

go 1.12

require (
	github.com/ProtonMail/go-mime v0.0.0-20190923161245-9b5a4261663a
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.4.0
	golang.org/x/crypto v0.0.0-20190923035154-9ee001bba392
)

replace golang.org/x/crypto => github.com/klische/crypto v0.0.0-20200704001232-4a9868ba58c3
