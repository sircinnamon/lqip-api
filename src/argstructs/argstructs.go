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
}

type QueryParameters struct {
	Shapes int
	Mode int
}