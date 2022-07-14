package db

type DbModel interface {
	GetName() string
	DbInit() error
	DbClose()
	IsRowExisted(string, string) bool
	UpdateTableColumn(statement string) error
}

type iUserTbl interface {
	Create(username, password string) error
	SearchByUsername(username string) (*UserTbl, error)
	SearchByUsernamePassword(username, password string) (*UserTbl, error)
	Login(username, password string) error
}

type ShapeRow struct {
	ShapeId int64    `json:"shape_id,omitempty"`
	Shape   string   `json:"shape,omitempty"`
	Edges   []string `json:"edges,omitempty"`
}

type iShapeTbl interface {
	Create(username, shape string, edges []string) (int64, error)
	GetAll(username string) ([]ShapeRow, error)
	Get(shapeId int64) (*ShapeRow, error)
	Update(shapeId int64, edges []string) error
	Delete(shapeId int64) error
}

func NewDbDriver(name string) DbModel {
	if name == "postgresql" {
		return &PostgreSQL{}
	} else if name == "mockdb" {
		return &SQLMock{}
	}
	return nil
}

func UserTblInstance() iUserTbl {
	if DB.GetName() == "postgresql" {
		return &PqUserTbl{}
	} else if DB.GetName() == "mockdb" {
		return mockUserTbl
	}
	return nil
}

func ShapeTblInstance() iShapeTbl {
	if DB.GetName() == "postgresql" {
		return &PqShapeTbl{}
	} else if DB.GetName() == "mockdb" {
		return mockShapeTbl
	}
	return nil
}

var DB DbModel = nil
