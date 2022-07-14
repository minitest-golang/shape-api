package db

// A Mock DB for unittest
type SQLMock struct {
}

func (db *SQLMock) GetName() string {
	return "mockdb"
}

func (db *SQLMock) DbInit() error {
	return nil
}

func (db *SQLMock) DbClose() {}

func (db *SQLMock) IsRowExisted(tableName, condition string) bool {
	return true
}

func (db *SQLMock) UpdateTableColumn(statement string) error {
	return nil
}

type MockDBUserTbl struct {
	DummyError error
}

func (tbl *MockDBUserTbl) Create(username, password string) error {
	return tbl.DummyError
}

// Search user by username
func (tbl *MockDBUserTbl) SearchByUsername(username string) (*UserTbl, error) {
	return &UserTbl{
		Username: username,
		Password: "key",
	}, tbl.DummyError
}

func (tbl *MockDBUserTbl) SearchByUsernamePassword(username, password string) (*UserTbl, error) {
	return &UserTbl{
		Username: username,
		Password: "key",
	}, tbl.DummyError
}

func (tbl *MockDBUserTbl) Login(username, password string) error {
	return tbl.DummyError
}

type MockShapeTbl struct {
	DummyError error
}

func (tbl *MockShapeTbl) Create(username, shape string, edges []string) (int64, error) {
	return 123, tbl.DummyError
}

func (tbl *MockShapeTbl) GetAll(username string) ([]ShapeRow, error) {
	return []ShapeRow{
		{
			ShapeId: 123,
			Shape:   "square",
			Edges:   []string{"3"},
		},
	}, tbl.DummyError
}

func (tbl *MockShapeTbl) Get(shapeId int64) (*ShapeRow, error) {
	return &ShapeRow{
		ShapeId: 123,
		Shape:   "square",
		Edges:   []string{"3"},
	}, tbl.DummyError
}

func (tbl *MockShapeTbl) Update(shapeId int64, edges []string) error {
	return nil
}
func (tbl *MockShapeTbl) Delete(shapeId int64) error {
	return tbl.DummyError
}

var mockUserTbl *MockDBUserTbl = &MockDBUserTbl{}
var mockShapeTbl *MockShapeTbl = &MockShapeTbl{}
