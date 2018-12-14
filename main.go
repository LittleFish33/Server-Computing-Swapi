package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/boltdb/bolt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

// 这边是认证功能的实现，先不用理
func createToken(key string, m map[string]interface{}) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	for index, val := range m {
		claims[index] = val
	}

	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(key))
	return tokenString
}

// 这边是认证功能的实现，先不用理
func parseToken(tokenString string, key string) (interface{}, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println(err)
		return "", false
	}
}

func main() {

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// todo：从这里开始
			// 第一步：添加查询语句，下面有两个例子，planet和species，补上people，films，starships，vehicles
			"planets": &graphql.Field{
				Type: createPlanetType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"]
					v, _ := id.(int)
					log.Printf("fetching planet with id: %d", v)
					return fetchPlanetByiD(v)
				},
			},
			"species": &graphql.Field{
				Type: createSpeciesType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["id"]
					v, _ := id.(int)
					log.Printf("fetching species with id: %d", v)
					return fetchSpeciesByPostID(v)
				},
			},

			"people": &graphql.Field{
				Type: createPeopleType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["id"]
					v, _ := id.(int)
					log.Printf("fetching people with id: %d", v)
					return fetchPeopleByiD(v)
				},
			},
			"films": &graphql.Field{
				Type: createFilmType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["id"]
					v, _ := id.(int)
					log.Printf("fetching films with id: %d", v)
					return fetchFilmByiD(v)
				},
			},
			"starships": &graphql.Field{
				Type: createStarshipType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["id"]
					v, _ := id.(int)
					log.Printf("fetching starships with id: %d", v)
					return fetchStarshipByiD(v)
				},
			},
			"vehicles": &graphql.Field{
				Type: createVehicleType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["id"]
					v, _ := id.(int)
					log.Printf("fetching Vehicles with id: %d", v)
					return fetchVehicleByiD(v)
				},
			},
			// 到这里结束
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	handler := gqlhandler.New(&gqlhandler.Config{
		Schema: &schema,
	})
	http.Handle("/graphql", handler)
	log.Println("Server started at http://localhost:3000/graphql")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// todo：从这里开始
// 第二步：添加createXXXType函数，下面有两个例子，createPlanetType和createSpeciesType，补上people，films，starships，vehicles的createXXXType函数
// Fields 里面是返回的json字符串里面的属性，替换为struct.go 里面的名字就可以了，
// Type有两种，graphql.String和graphql.NewList(graphql.String)，照着struct.go里面的改就行了
func createPlanetType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Planet",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"OrbitalPeriod": &graphql.Field{
				Type: graphql.String,
			},
			"RotationPeriod": &graphql.Field{
				Type: graphql.String,
			},
			"Diameter": &graphql.Field{
				Type: graphql.String,
			},
			"Climate": &graphql.Field{
				Type: graphql.String,
			},
			"Gravity": &graphql.Field{
				Type: graphql.String,
			},
			"Terrain": &graphql.Field{
				Type: graphql.String,
			},
			"SurfaceWater": &graphql.Field{
				Type: graphql.String,
			},
			"Population": &graphql.Field{
				Type: graphql.String,
			},
			"ResidentURLs": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"FilmURLs": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"Created": &graphql.Field{
				Type: graphql.String,
			},
			"Edited": &graphql.Field{
				Type: graphql.String,
			},
			"URL": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func createSpeciesType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Species",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Classification": &graphql.Field{
				Type: graphql.String,
			},
			"Designation": &graphql.Field{
				Type: graphql.String,
			},
			"AverageHeight": &graphql.Field{
				Type: graphql.String,
			},
			"SkinColors": &graphql.Field{
				Type: graphql.String,
			},
			"HairColors": &graphql.Field{
				Type: graphql.String,
			},
			"EyeColors": &graphql.Field{
				Type: graphql.String,
			},
			"AverageLifespan": &graphql.Field{
				Type: graphql.String,
			},
			"Homeworld": &graphql.Field{
				Type: graphql.String,
			},
			"Language": &graphql.Field{
				Type: graphql.String,
			},
			"PeopleURLs": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"FilmURLs": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"Created": &graphql.Field{
				Type: graphql.String,
			},
			"Edited": &graphql.Field{
				Type: graphql.String,
			},
			"URL": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func createPeopleType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Person",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Height": &graphql.Field{
				Type: graphql.String,
			},
			"Mass": &graphql.Field{
				Type: graphql.String,
			},
			"HairColor": &graphql.Field{
				Type: graphql.String,
			},
			"SkinColor": &graphql.Field{
				Type: graphql.String,
			},
			"EyeColor": &graphql.Field{
				Type: graphql.String,
			},
			"BirthYear": &graphql.Field{
				Type: graphql.String,
			},
			"Gender": &graphql.Field{
				Type: graphql.String,
			},
			"Homeworld": &graphql.Field{
				Type: graphql.String,
			},
			"FilmURLs": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"SpeciesURLs": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"VehicleURLs": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"StarshipURLs": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"Created": &graphql.Field{
				Type: graphql.String,
			},
			"Edited": &graphql.Field{
				Type: graphql.String,
			},
			"URL": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func createFilmType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Film",
		Fields: graphql.Fields{
			"Title": &graphql.Field{
				Type: graphql.String,
			},
			"EpisodeID": &graphql.Field{
				Type: graphql.Int,
			},
			"OpeningCrawl": &graphql.Field{
				Type: graphql.String,
			},
			"Director": &graphql.Field{
				Type: graphql.String,
			},
			"Producer": &graphql.Field{
				Type: graphql.String,
			},
			"CharacterURLs": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"PlanetURLs": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"StarshipURLs": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"VehicleURLs": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"SpeciesURLs": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"Created": &graphql.Field{
				Type: graphql.String,
			},
			"Edited": &graphql.Field{
				Type: graphql.String,
			},
			"URL": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func createStarshipType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Starship",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Model": &graphql.Field{
				Type: graphql.String,
			},
			"Manufacturer": &graphql.Field{
				Type: graphql.String,
			},
			"CostInCredits": &graphql.Field{
				Type: graphql.String,
			},
			"Length": &graphql.Field{
				Type: graphql.String,
			},
			"MaxAtmospheringSpeed": &graphql.Field{
				Type: graphql.String,
			},
			"Crew": &graphql.Field{
				Type: graphql.String,
			},
			"Passengers": &graphql.Field{
				Type: graphql.String,
			},
			"CargoCapacity": &graphql.Field{
				Type: graphql.String,
			},
			"Consumables": &graphql.Field{
				Type: graphql.String,
			},
			"HyperdriveRating": &graphql.Field{
				Type: graphql.String,
			},
			"MGLT": &graphql.Field{
				Type: graphql.String,
			},
			"StarshipClass": &graphql.Field{
				Type: graphql.String,
			},
			"PilotURLs": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"FilmURLs": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"Created": &graphql.Field{
				Type: graphql.String,
			},
			"Edited": &graphql.Field{
				Type: graphql.String,
			},
			"URL": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func createVehicleType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Vehicle",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Model": &graphql.Field{
				Type: graphql.String,
			},
			"Manufacturer": &graphql.Field{
				Type: graphql.String,
			},
			"CostInCredits": &graphql.Field{
				Type: graphql.String,
			},
			"Length": &graphql.Field{
				Type: graphql.String,
			},
			"MaxAtmospheringSpeed": &graphql.Field{
				Type: graphql.String,
			},
			"Crew": &graphql.Field{
				Type: graphql.String,
			},
			"Passengers": &graphql.Field{
				Type: graphql.String,
			},
			"CargoCapacity": &graphql.Field{
				Type: graphql.String,
			},
			"Consumables": &graphql.Field{
				Type: graphql.String,
			},
			"VehicleClass": &graphql.Field{
				Type: graphql.String,
			},
			"PilotURLs": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"FilmURLs": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"Created": &graphql.Field{
				Type: graphql.String,
			},
			"Edited": &graphql.Field{
				Type: graphql.String,
			},
			"URL": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

// 到这里结束

// todo：从这里开始
// 第三步：添加fetchXXXByiD函数，下面有两个例子，fetchPlanetByiD和fetchSpeciesByPostID，补上people，films，starships，vehicles的cfetchXXXByiD函数
// 超简单，替换一下变量名而已

func fetchPlanetByiD(id int) (*Planet, error) {
	result := Planet{}
	db, _ := setupDB()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("Planets"))
		v := string(b.Get([]byte(strconv.Itoa(id))))
		err := json.Unmarshal([]byte(v), &result)
		if err != nil {
			return fmt.Errorf("could not Unmarshal json string: %v", err)
		}

		return nil
	})
	db.Close()
	return &result, nil
}

func fetchSpeciesByPostID(id int) (*Species, error) {
	result := Species{}
	db, _ := setupDB()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("Species"))
		v := string(b.Get([]byte(strconv.Itoa(id))))
		err := json.Unmarshal([]byte(v), &result)
		if err != nil {
			return fmt.Errorf("could not Unmarshal json string: %v", err)
		}

		return nil
	})
	db.Close()
	return &result, nil
}

func fetchPeopleByiD(id int) (*Person, error) {
	result := Person{}
	db, _ := setupDB()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("People"))
		v := string(b.Get([]byte(strconv.Itoa(id))))
		err := json.Unmarshal([]byte(v), &result)
		if err != nil {
			return fmt.Errorf("could not Unmarshal json string: %v", err)
		}

		return nil
	})
	db.Close()
	return &result, nil
}

func fetchFilmByiD(id int) (*Film, error) {
	result := Film{}
	db, _ := setupDB()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("Films"))
		v := string(b.Get([]byte(strconv.Itoa(id))))
		err := json.Unmarshal([]byte(v), &result)
		if err != nil {
			return fmt.Errorf("could not Unmarshal json string: %v", err)
		}

		return nil
	})
	db.Close()
	return &result, nil
}

func fetchStarshipByiD(id int) (*Starship, error) {
	result := Starship{}
	db, _ := setupDB()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("Starships"))
		v := string(b.Get([]byte(strconv.Itoa(id))))
		err := json.Unmarshal([]byte(v), &result)
		if err != nil {
			return fmt.Errorf("could not Unmarshal json string: %v", err)
		}

		return nil
	})
	db.Close()
	return &result, nil
}

func fetchVehicleByiD(id int) (*Vehicle, error) {
	result := Vehicle{}
	db, _ := setupDB()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("Vehicles"))
		v := string(b.Get([]byte(strconv.Itoa(id))))
		err := json.Unmarshal([]byte(v), &result)
		if err != nil {
			return fmt.Errorf("could not Unmarshal json string: %v", err)
		}

		return nil
	})
	db.Close()
	return &result, nil
}

// 到这里结束
