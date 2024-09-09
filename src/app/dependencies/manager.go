package dependencies

const (
	productionEnvironment = "production"
)

type Manager interface {
}

func NewManager() Manager {
	return Production{}
}