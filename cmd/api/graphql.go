package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"server/models"
	"strings"

	"github.com/graphql-go/graphql"
)

var members []*models.Member

var fields = graphql.Fields{

	"member": &graphql.Field{
		Type:        memberType,
		Description: "Get member by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(int)

			if ok {
				for _, member := range members {
					if member.ID == id {
						return member, nil
					}
				}
			}
			return nil, nil
		},
	},
	"list": &graphql.Field{
		Type:        graphql.NewList(memberType),
		Description: "Get all member",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return members, nil
		},
	},
	"search": &graphql.Field{
		Type:        graphql.NewList(memberType),
		Description: "Search member by name",
		Args: graphql.FieldConfigArgument{
			"nameContain": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var theList []*models.Member
			search, ok := p.Args["nameContain"].(string)

			if ok {
				for _, currentMember := range members {
					if strings.Contains(currentMember.Name, search) {
						theList = append(theList, currentMember)
					}
				}
			}

			return theList, nil
		},
	},
}

var memberType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Member",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"nickname": &graphql.Field{
				Type: graphql.String,
			},
			"channel": &graphql.Field{
				Type: graphql.String,
			},
			"height": &graphql.Field{
				Type: graphql.Int,
			},
			"debut_date": &graphql.Field{
				Type: graphql.DateTime,
			},
			"birth_date": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

func (app *Application) memberGraphQL(w http.ResponseWriter, r *http.Request) {
	members, _ = app.Models.DB.All()

	q, _ := io.ReadAll(r.Body)
	query := string(q)

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		app.errorJSON(w, errors.New("failed to create schema"))
		log.Println(err)
		return
	}

	params := graphql.Params{Schema: schema, RequestString: query }

	resp := graphql.Do(params)

	if len(resp.Errors) > 0 {
		app.errorJSON(w, errors.New(fmt.Sprintf("failed: %+v", resp.Errors)))
	}

	j,_ := json.MarshalIndent(resp,"","\t")
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(j)

}