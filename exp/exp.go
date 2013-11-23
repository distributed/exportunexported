package exp

type Exported struct {
	Field int
}

type shell struct {
	Exported
	error
}

func GetLeStuff() (Exported, error) {
	val := shell{Exported{42}, nil}
	return val.Exported, val.error
}
