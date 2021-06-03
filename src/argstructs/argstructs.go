package argstructs

type ServerArgs struct {
	Port int
}

type ImageHandlerArgs struct {
	Shapes int
	AllowShapeCountQP bool
	MaxShapeCountQP int
}

type QueryParameters struct {
	Shapes int
}