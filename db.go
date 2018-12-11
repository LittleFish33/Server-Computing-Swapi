package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
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
		_, err = root.CreateBucketIfNotExists([]byte("Person"))
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
	planetJSONStr := `{"name":"Alderaan","rotation_period":"24","orbital_period":"364","diameter":"12500","climate":"temperate","gravity":"1 standard","terrain":"grasslands, mountains","surface_water":"40","population":"2000000000","residents":["https://swapi.co/api/people/5/","https://swapi.co/api/people/68/","https://swapi.co/api/people/81/"],"films":["https://swapi.co/api/films/6/","https://swapi.co/api/films/1/"],"created":"2014-12-10T11:35:48.479000Z","edited":"2014-12-20T20:58:18.420000Z","url":"https://swapi.co/api/planets/2/"}`
	planetID := "2"
	err = addToBucket(db, planetJSONStr, planetID, "Planets")
	if err != nil {
		log.Fatal(err)
	}

	speciesJSONStr := `{"name":"Wookiee","classification":"mammal","designation":"sentient","average_height":"210","skin_colors":"gray","hair_colors":"black, brown","eye_colors":"blue, green, yellow, brown, golden, red","average_lifespan":"400","homeworld":"https://swapi.co/api/planets/14/","language":"Shyriiwook","people":["https://swapi.co/api/people/13/","https://swapi.co/api/people/80/"],"films":["https://swapi.co/api/films/2/","https://swapi.co/api/films/7/","https://swapi.co/api/films/6/","https://swapi.co/api/films/3/","https://swapi.co/api/films/1/"],"created":"2014-12-10T16:44:31.486000Z","edited":"2015-01-30T21:23:03.074598Z","url":"https://swapi.co/api/species/3/"}`
	speciesID := "2"
	err = addToBucket(db, speciesJSONStr, speciesID, "Species")
	if err != nil {
		log.Fatal(err)
	}

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
	// db.View(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("DB")).Bucket([]byte("Planets"))
	// 	v := string(b.Get([]byte("2")))
	// 	fmt.Printf(v)
	// 	var planet Planet
	// 	err := json.Unmarshal([]byte(v), &planet)
	// 	if err != nil {
	// 		return fmt.Errorf("could not Unmarshal json string: %v", err)
	// 	}
	// 	fmt.Printf("The name of the planet 2 is: %s\n", planet.Name)
	// 	return nil
	// })

	if err != nil {
		log.Fatal(err)
	}

}

// 到这里结束

// 如果要测试db.go类，去掉下面的注释
// func main() {
// 	testdb()
// }
