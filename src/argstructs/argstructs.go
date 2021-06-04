package argstructs

type ServerArgs struct {
	Port int
	AllowAsync bool
	AsyncCacheExpiry int
	AsyncCacheGC int
}

type ImageHandlerArgs struct {
	Shapes int
	AllowShapeCountQP bool
	MaxShapeCountQP int
	Mode int
	AllowedModeQPs string
	Blur int
	AllowBlurQP bool
}

type QueryParameters struct {
	Shapes int
	Mode int
	Blur int
}