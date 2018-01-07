package nous

type Nous interface {
	Store(Knowledge) error
	Retain() error
	Research() error
}
