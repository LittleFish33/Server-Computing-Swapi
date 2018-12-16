# SWAPI

## 负责数据库的人

### 运行说明：

将`db.go` 第117-120行去掉注释：

```go
// 如果要测试db.go类，去掉下面的注释
func main() {
	testdb()
}
```

控制台输入：`go run db.go`

![](https://littlefish33.cn/image/temp/237.png)

目前的程序会往数据库中的Planet桶和Species桶添加两条数据，然后将打印桶里所有的数据

**boltDB**数据库的桶就相当于数据结构中的map，挺适合我们的项目的，数据的编号正好作为map里的key，json作为map里的value，如果编号一样，即key一样插入两次会自动覆盖

### 怎么做：

负责`db.go`的人只需要关注`db.go` 文件，我在该文件中留了三处todo，实际上只有一处需要添加代码，另外两处只是为了说明，直接`ctrl+f` 找一下就好了，具体操作在代码里

### 测试说明

我已经给出了打印出所有桶的内容和打印桶里某个对象的内容的代码，稍加修改就可以测试了

### 负责网络连接的人

控制台输入：`go run struct.go db.go main.go`

![](https://littlefish33.cn/image/temp/238.png)

运行之后可能需要等一会，直到控制台输出

`2018/12/11 20:46:12 Server started at http://localhost:3000/graphql`

### 怎么做：

负责网络连接的人需要关注`main.go`和`struct.go` 文件，不过需要添加代码的只有`main.go`文件，我在该文件中留了三处todo，直接`ctrl+f` 找一下就好了，具体操作在代码里

### 测试说明：

我用的是`graphql` ，现在已经往Planet和Species添加了一条记录，可以使用下面的链接查看：

```
// 下面是两个现在可用的例子
// graphql 可以不用返回整个planet的json字符串，而是取出其中特定的元素，如第一个例子只取出了Name，而第二个例子取出了Name和RotationPeriod
http://localhost:3000/graphql?query={planet(id:2){Name,}}
http://localhost:3000/graphql?query={planet(id:2){Name,RotationPeriod,}}
http://localhost:3000/graphql?query={species(id:2){Name,Classification,}}
```

具体效果如下

![](https://littlefish33.cn/image/temp/239.png)



![](https://littlefish33.cn/image/temp/240.png)



## url

```
http://localhost:3000/graphql?query={planets(id:2){Name,RotationPeriod,OrbitalPeriod,Diameter,Climate,Gravity,Terrain,SurfaceWater,Population,ResidentURLs,FilmURLs,Created,Edited,URL,}}

http://localhost:3000/graphql?query={Species(id:2){Name,Classification,Designation,AverageHeight,SkinColors,HairColors,EyeColors,AverageLifespan,Homeworld,Language,PeopleURLs,FilmURLs,Created,Edited,URL,}}

http://localhost:3000/graphql?query={vehicles(id:14){Name,Model,Manufacturer,CostInCredits,Length,MaxAtmospheringSpeed,Crew,Passengers,CargoCapacity,Consumables,VehicleClass,PilotURLs,FilmURLs,Created,Edited,URL,}}

http://localhost:3000/graphql?query={starships(id:2){Name,Model,Manufacturer,CostInCredits,Length,MaxAtmospheringSpeed,Crew,Passengers,CargoCapacity,Consumables,HyperdriveRating,MGLT,StarshipClass,PilotURLs,FilmURLs,Created,Edited,URL,}}

http://localhost:3000/graphql?query={people(id:2){Name,Height,Mass,HairColor,SkinColor,EyeColor,BirthYear,Gender,Homeworld,FilmURLs,SpeciesURLs,VehicleURLs,StarshipURLs,Created,Edited,URL,}}

http://localhost:3000/graphql?query={films(id:2){Title,EpisodeID,OpeningCrawl,Director,Producer,CharacterURLs,PlanetURLs,StarshipURLs,VehicleURLs,SpeciesURLs,Created,Edited,URL,}}

```











