//go:generate mockgen -source demo.go -destination demo_mock.go -package demo
package demo

type HTTP interface {
	Get() string
}

func GetString(h HTTP) string {
	return h.Get()
}
