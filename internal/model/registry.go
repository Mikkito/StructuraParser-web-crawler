package model

var handlers = make(map[string]BlockHandler)

func RegisterHandler(handler BlockHandler) {
	handlers[handler.Type()] = handler
}

func GetHandler(blockType string) (BlockHandler, bool) {
	h, ok := handlers[blockType]
	return h, ok
}

func GetAllHandlers() map[string]BlockHandler {
	return handlers
}
