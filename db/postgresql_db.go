package db

import (
	"database/sql"
	"fmt"
	"minitest/utils"
	"os"
	"strings"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var (
	dbHost     = ""
	dbPort     = "5432"
	dbUsername = ""
	dbPassword = ""
	dbName     = ""
	pqDb       *sql.DB
)

// This is an implementation of DbModel interface for PostgreSQL
type PostgreSQL struct {
}

func (db *PostgreSQL) GetName() string {
	return "postgresql"
}

func (db *PostgreSQL) DbInit() error {
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUsername = os.Getenv("POSTGRES_USER")
	dbPassword = os.Getenv("POSTGRES_PASSWORD")
	dbName = os.Getenv("POSTGRES_DB")

	psqlconn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUsername,
		dbPassword,
		dbName)

	utils.InfoLog("%s", psqlconn)

	// Open DB
	var err error
	pqDb, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return err
	}

	// Check db
	err = pqDb.Ping()
	if err != nil {
		pqDb.Close()
		return err
	}
	utils.InfoLog("Connected DB successfully!")
	return nil
}

func (db *PostgreSQL) DbClose() {
	if pqDb != nil {
		pqDb.Close()
		utils.InfoLog("closed DB!")
	}
}

func (db *PostgreSQL) IsRowExisted(tableName, condition string) bool {
	var count int
	cmd := fmt.Sprintf("SELECT 1 FROM %s WHERE %s LIMIT 1", tableName, condition)
	err := pqDb.QueryRow(cmd).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (db *PostgreSQL) UpdateTableColumn(statement string) error {
	res, err := pqDb.Exec(statement)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

type PqUserTbl struct {
}

func (tbl *PqUserTbl) Create(username, password string) error {
	if DB.IsRowExisted("user_tbl", fmt.Sprintf("username = '%s'", username)) {
		return utils.ErrDBUserExisted
	}
	statement := `INSERT into "user_tbl"("username", "login_key") values($1, $2)`
	_, err := pqDb.Exec(statement, username, password)
	if err != nil {
		return err
	}
	return nil
}

// User search wraper
func (tbl *PqUserTbl) search(condition string) (*UserTbl, error) {
	username := ""
	loginKey := ""
	err := pqDb.QueryRow(fmt.Sprintf(
		`SELECT username, login_key from "user_tbl" where %s;`,
		condition)).Scan(&username, &loginKey)
	if err != nil {
		return nil, err
	}
	return &UserTbl{
		Username: username,
		Password: loginKey,
	}, nil
}

// Search user by username
func (tbl *PqUserTbl) SearchByUsername(username string) (*UserTbl, error) {
	return tbl.search(fmt.Sprintf("username = '%s'", username))
}

// Search user by username and password
// This function can be used for Login
func (tbl *PqUserTbl) SearchByUsernamePassword(username, password string) (*UserTbl, error) {
	return tbl.search(fmt.Sprintf("username = '%s' AND login_key = '%s'", username, password))
}

// User Login
func (tbl *PqUserTbl) Login(username, password string) error {
	_, err := tbl.SearchByUsernamePassword(username, password)
	return err
}

type PqShapeTbl struct {
}

func (tbl *PqShapeTbl) Create(username, shape string, edges []string) (int64, error) {
	statement := `INSERT into "shape_tbl"("username", "shape", "edges") values($1, $2, $3) RETURNING shape_id`

	shape_id := int64(0)
	err := pqDb.QueryRow(statement, username, shape, pq.Array(edges)).Scan(&shape_id)
	if err != nil {
		return 0, err
	}
	return shape_id, nil
}

func (tbl *PqShapeTbl) GetAll(username string) ([]ShapeRow, error) {
	cmd := fmt.Sprintf(`SELECT shape_id, shape, edges from "shape_tbl" where username='%s';`, username)
	rows, err := pqDb.Query(cmd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	shapes := []ShapeRow{}
	for rows.Next() {
		shape_id := int64(0)
		shape := ""
		edges := []string{}
		err = rows.Scan(
			&shape_id,
			&shape,
			pq.Array(&edges),
		)
		if err == nil {
			shapes = append(shapes, ShapeRow{
				ShapeId: shape_id,
				Shape:   shape,
				Edges:   edges,
			})
		} else {
			utils.ErrorLog("DB Error: %s!", err.Error())
		}
	}
	return shapes, nil
}

func (tbl *PqShapeTbl) Get(shapeId int64) (*ShapeRow, error) {
	shape := ""
	edges := []string{}
	err := pqDb.QueryRow(fmt.Sprintf(
		`SELECT shape, edges from "shape_tbl" where shape_id=%d;`,
		shapeId)).Scan(&shape, pq.Array(&edges))
	if err != nil {
		return nil, err
	}
	return &ShapeRow{
		Shape: shape,
		Edges: edges,
	}, nil
}

func (tbl *PqShapeTbl) Update(shapeId int64, edges []string) error {
	params := make([]string, 0, len(edges))
	for i := range edges {
		params = append(params, fmt.Sprintf("$%v", i+1))
	}
	statement := fmt.Sprintf("UPDATE shape_tbl SET edges = ARRAY[%s]", strings.Join(params, ", "))
	edgesAny := []any{}
	for _, e := range edges {
		edgesAny = append(edgesAny, e)
	}
	_, err := pqDb.Exec(statement, edgesAny...)
	return err
}

func (tbl *PqShapeTbl) Delete(shapeId int64) error {
	statement := fmt.Sprintf(`DELETE FROM "shape_tbl" WHERE shape_id=%d;`, shapeId)
	_, err := pqDb.Exec(statement)
	return err
}
