package v1

import (
	"minitest/db"
	"minitest/shape"
	"minitest/utils"

	"github.com/graphql-go/graphql"
)

/*
Querry:
	{
		area(shape_id:1)
		perimeter(shape_id:2)
	}
*/
func ShapeScheme() *graphql.Schema {
	fields := graphql.Fields{
		"area": &graphql.Field{
			Type:        graphql.String,
			Description: "Get Shape Area by ID",
			Args: graphql.FieldConfigArgument{
				"shape_id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["shape_id"].(int)
				if !ok {
					return "NA", utils.ErrBadShapeId
				}
				shapeInfo, err := db.ShapeTblInstance().Get(int64(id))
				if err != nil {
					return "NA", err
				}
				shapeObj, err := shape.CreateShape(shapeInfo.Shape, shapeInfo.Edges)
				if err != nil {
					return "NA", err
				}
				return shapeObj.Area(), nil
			},
		},
		"perimeter": &graphql.Field{
			Type:        graphql.String,
			Description: "Get Shape Perimeter by ID",
			Args: graphql.FieldConfigArgument{
				"shape_id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["shape_id"].(int)
				if !ok {
					return "NA", utils.ErrBadShapeId
				}
				shapeInfo, err := db.ShapeTblInstance().Get(int64(id))
				if err != nil {
					return "NA", err
				}
				shapeObj, err := shape.CreateShape(shapeInfo.Shape, shapeInfo.Edges)
				if err != nil {
					return "NA", err
				}
				return shapeObj.Perimeter(), nil
			},
		},
	}

	shapeQuery := graphql.ObjectConfig{
		Name:   "ShapeCal",
		Fields: fields,
	}

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(shapeQuery),
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		utils.ErrorLog("failed to create new schema, error: %v", err)
		return nil
	}
	return &schema
}
