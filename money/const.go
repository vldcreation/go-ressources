package money

type Version string

func (v Version) String() string {
	return string(v)
}

const (
	V1 Version = "v1"
)
