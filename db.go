package main

import (
	"fmt"
	"log"
	"encoding/json"
	"github.com/boltdb/bolt"
	"os/exec"
	"os"
	"strings"
	"strconv"
)

func setupDB() (*bolt.DB, error) {
	db, err := bolt.Open("swapi.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		// todo：从这里开始
		// 第一处：这里是建数据库的操作，已经修改了，不用做更改，不过需要记一下，后面要读下面的表
		// 到这里结束
		_, err = root.CreateBucketIfNotExists([]byte("Planets"))
		_, err = root.CreateBucketIfNotExists([]byte("Species"))
		_, err = root.CreateBucketIfNotExists([]byte("Vehicle"))
		_, err = root.CreateBucketIfNotExists([]byte("Starship"))
		_, err = root.CreateBucketIfNotExists([]byte("People"))
		_, err = root.CreateBucketIfNotExists([]byte("Film"))
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}
	fmt.Println("DB Setup Done")
	return db, nil
}

// todo：从这里开始
// 第二处：这里是往数据库中添加内容的操作，实际上不用做任何修改
// jsonStr表示要插入的json字符串，id为该字符串对应的id，bucketName为数据库中的桶的名字，可以直接看做一张表
func addToBucket(db *bolt.DB, jsonStr string, id string, bucketName string) error {

	planetBytes := []byte(jsonStr)

	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte(bucketName)).Put([]byte(id), []byte(planetBytes))
		if err != nil {
			return fmt.Errorf("could not insert %s: %v", bucketName, err)
		}

		return nil
	})
	fmt.Println("Added " + bucketName)
	return err
}

// 到这里结束

// todo：从这里开始
// 第三处，在这里添加往数据库里加入数据的函数
func testdb() {
	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 下面是两个往数据库中添加内容的例子
	// 照着改就可以了，planetJSONStr为https://swapi.co/api/planets/2/返回的json字符串压缩后的结果，json字符串压缩可以直接百度
	// ID对应编号，很明显
	// 可以考虑用脚本，否则一个一个打太麻烦了
	
	const (
		planet = 1
		species = 2
		vehicle = 3
		starship = 4
		people = 5
		film  = 6
	)
	var tabText = map[int]string{
		planet:  "Planets",
		species: "Species",
		vehicle: "Vehicle",
		starship: "Starship",
		people: "People",
		film:  "Film",
	}
	var nameText = map[int]string{
		planet:  "planets",
		species: "species",
		vehicle: "vehicles",
		starship: "starships",
		people: "people",
		film:  "films",
	}
	var curl []byte
	var cmd *exec.Cmd
	for i := 1; i < 100; i++ {
		for k := 1; k <= 6; k++ {
			// fmt.Println("https://swapi.co/api/" + nameText[k] + "/" + string(i) + "/")
			// fmt.Println(string(i))
			ID := strconv.Itoa(i)
			cmd = exec.Command("curl", ("https://swapi.co/api/" + nameText[k] + "/" + ID + "/"))
			if curl, err = cmd.Output(); err != nil {
			    fmt.Println(err)
			    os.Exit(1)
			}

			// fmt.Println(("https://swapi.co/api/" + nameText[k] + "/" + ID + "/"))
			JSONStr := strings.TrimRight(string(curl), "\n")
			fmt.Println(JSONStr)
			fmt.Println(ID)

			err = addToBucket(db, JSONStr, ID, tabText[k])
			if err != nil {
				log.Fatal(err)
			}
		}
	}
// {
	// planetJSONStr := `{"name":"Alderaan","rotation_period":"24","orbital_period":"364","diameter":"12500","climate":"temperate","gravity":"1 standard","terrain":"grasslands, mountains","surface_water":"40","population":"2000000000","residents":["https://swapi.co/api/people/5/","https://swapi.co/api/people/68/","https://swapi.co/api/people/81/"],"films":["https://swapi.co/api/films/6/","https://swapi.co/api/films/1/"],"created":"2014-12-10T11:35:48.479000Z","edited":"2014-12-20T20:58:18.420000Z","url":"https://swapi.co/api/planets/2/"}`
	// planetID := "2"
	// err = addToBucket(db, planetJSONStr, planetID, "Planets")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// speciesJSONStr := `{"name":"Wookiee","classification":"mammal","designation":"sentient","average_height":"210","skin_colors":"gray","hair_colors":"black, brown","eye_colors":"blue, green, yellow, brown, golden, red","average_lifespan":"400","homeworld":"https://swapi.co/api/planets/14/","language":"Shyriiwook","people":["https://swapi.co/api/people/13/","https://swapi.co/api/people/80/"],"films":["https://swapi.co/api/films/2/","https://swapi.co/api/films/7/","https://swapi.co/api/films/6/","https://swapi.co/api/films/3/","https://swapi.co/api/films/1/"],"created":"2014-12-10T16:44:31.486000Z","edited":"2015-01-30T21:23:03.074598Z","url":"https://swapi.co/api/species/3/"}`
	// speciesID := "2"
	// err = addToBucket(db, speciesJSONStr, speciesID, "Species")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// vehicleJSONStr := `{"name":"Sand Crawler","model":"Digger Crawler","manufacturer":"Corellia Mining Corporation","cost_in_credits":"150000","length":"36.8","max_atmosphering_speed":"30","crew":"46","passengers":"30","cargo_capacity":"50000","consumables":"2 months","vehicle_class":"wheeled","pilots":[],"films":["https://swapi.co/api/films/5/","https://swapi.co/api/films/1/"],"created":"2014-12-10T15:36:25.724000Z","edited":"2014-12-22T18:21:15.523587Z","url":"https://swapi.co/api/vehicles/4/"}`
	// vehicleID := "4"
	// err = addToBucket(db, vehicleJSONStr, vehicleID, "Vehicle")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// starshipJSONStr := `{"name":"Death Star","model":"DS-1 Orbital Battle Station","manufacturer":"Imperial Department of Military Research, Sienar Fleet Systems","cost_in_credits":"1000000000000","length":"120000","max_atmosphering_speed":"n/a","crew":"342953","passengers":"843342","cargo_capacity":"1000000000000","consumables":"3 years","hyperdrive_rating":"4.0","MGLT":"10","starship_class":"Deep Space Mobile Battlestation","pilots":[],"films":["https://swapi.co/api/films/1/"],"created":"2014-12-10T16:36:50.509000Z","edited":"2014-12-22T17:35:44.452589Z","url":"https://swapi.co/api/starships/9/"}`
	// starshipID := "9"
	// err = addToBucket(db, starshipJSONStr, starshipID, "Starship")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// personJSONStr := `{"name":"Luke Skywalker","height":"172","mass":"77","hair_color":"blond","skin_color":"fair","eye_color":"blue","birth_year":"19BBY","gender":"male","homeworld":"https://swapi.co/api/planets/1/","films":["https://swapi.co/api/films/2/","https://swapi.co/api/films/6/","https://swapi.co/api/films/3/","https://swapi.co/api/films/1/","https://swapi.co/api/films/7/"],"species":["https://swapi.co/api/species/1/"],"vehicles":["https://swapi.co/api/vehicles/14/","https://swapi.co/api/vehicles/30/"],"starships":["https://swapi.co/api/starships/12/","https://swapi.co/api/starships/22/"],"created":"2014-12-09T13:50:51.644000Z","edited":"2014-12-20T21:17:56.891000Z","url":"https://swapi.co/api/people/1/"}`
	// personID := "1"
	// err = addToBucket(db, personJSONStr, personID, "Person")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// filmJSONStr := `{"title":"A New Hope","episode_id":4,"opening_crawl":"It is a period of civil war.\r\nRebel spaceships, striking\r\nfrom a hidden base, have won\r\ntheir first victory against\r\nthe evil Galactic Empire.\r\n\r\nDuring the battle, Rebel\r\nspies managed to steal secret\r\nplans to the Empire's\r\nultimate weapon, the DEATH\r\nSTAR, an armored space\r\nstation with enough power\r\nto destroy an entire planet.\r\n\r\nPursued by the Empire's\r\nsinister agents, Princess\r\nLeia races home aboard her\r\nstarship, custodian of the\r\nstolen plans that can save her\r\npeople and restore\r\nfreedom to the galaxy....","director":"George Lucas","producer":"Gary Kurtz, Rick McCallum","release_date":"1977-05-25","characters":["https://swapi.co/api/people/1/","https://swapi.co/api/people/2/","https://swapi.co/api/people/3/","https://swapi.co/api/people/4/","https://swapi.co/api/people/5/","https://swapi.co/api/people/6/","https://swapi.co/api/people/7/","https://swapi.co/api/people/8/","https://swapi.co/api/people/9/","https://swapi.co/api/people/10/","https://swapi.co/api/people/12/","https://swapi.co/api/people/13/","https://swapi.co/api/people/14/","https://swapi.co/api/people/15/","https://swapi.co/api/people/16/","https://swapi.co/api/people/18/","https://swapi.co/api/people/19/","https://swapi.co/api/people/81/"],"planets":["https://swapi.co/api/planets/2/","https://swapi.co/api/planets/3/","https://swapi.co/api/planets/1/"],"starships":["https://swapi.co/api/starships/2/","https://swapi.co/api/starships/3/","https://swapi.co/api/starships/5/","https://swapi.co/api/starships/9/","https://swapi.co/api/starships/10/","https://swapi.co/api/starships/11/","https://swapi.co/api/starships/12/","https://swapi.co/api/starships/13/"],"vehicles":["https://swapi.co/api/vehicles/4/","https://swapi.co/api/vehicles/6/","https://swapi.co/api/vehicles/7/","https://swapi.co/api/vehicles/8/"],"species":["https://swapi.co/api/species/5/","https://swapi.co/api/species/3/","https://swapi.co/api/species/2/","https://swapi.co/api/species/1/","https://swapi.co/api/species/4/"],"created":"2014-12-10T14:23:31.880000Z","edited":"2015-04-11T09:46:52.774897Z","url":"https://swapi.co/api/films/1/"}`
	// filmID := "1"
	// err = addToBucket(db, filmJSONStr, filmID, "Film")
	// if err != nil {
	// 	log.Fatal(err)
	// }
// }
	
	// 下面的代码会打印出Planets桶和Species桶里的内容，替换为其他的桶的名字可以帮助你查看你的代码是否正确
	// 目前Planets桶和Species桶里默认各有一条内容
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("Planets"))
		b.ForEach(func(k, v []byte) error {
			fmt.Println(string(k), string(v))
			return nil
		})
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("Species"))
		b.ForEach(func(k, v []byte) error {
			fmt.Println(string(k), string(v))
			return nil
		})
		return nil
	})

	// 下面的函数可以用于查询Planets里编号为2的星球的名字，稍作修改就可以测试其他的内容
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("Planets"))
		v := string(b.Get([]byte("2")))
		fmt.Printf(v)
		var planet Planet
		err := json.Unmarshal([]byte(v), &planet)
		if err != nil {
			return fmt.Errorf("could not Unmarshal json string: %v", err)
		}
		fmt.Printf("\nThe name of the planet 2 is: %s\n", planet.Name)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}

// 到这里结束

// 如果要测试db.go类，去掉下面的注释
func main() {
	testdb()
}
