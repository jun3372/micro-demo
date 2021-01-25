module member

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.4.3
	github.com/jun3372/micro-demo v0.0.0-20210124040303-b67f76d9b839
	github.com/micro/micro/v3 v3.0.4
	google.golang.org/protobuf v1.25.0
	gorm.io/gorm v1.20.11
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
