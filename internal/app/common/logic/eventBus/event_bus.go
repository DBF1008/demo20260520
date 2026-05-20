package eventBus

import (
	"github.com/asaskevich/EventBus"
	"github.com/tiger1103/gfast/v3/internal/app/common/service"
)

func init() {
	service.RegisterEventBus(EventBus.New())
}
