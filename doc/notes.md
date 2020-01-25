#  planned command line usage
It is possible that it would be more elegant using a configuration file like
`$HOME/.grename/config.toml`, that might look like this:

``` 
[global]
library = "/home/peter/grename-lib/"
keep-duplicates = false
hash-algorithm = blake3

[database]
driver = mysql
host = localhost
# ...

[tags]
ignore = ["old", "new"]

[import]
  [import.images]
  mimetypes = ["image/jpeg", "image/raw"]
  template = ".MimeMediaType/.MimeSubType/.ExifModel/.ExifYear..."

  [import.videos]
  mimetypes = ["video"]

```
## import 
imports (moves or copies) all media from source into a destination
```
grename import <source-path> <destinatiion-path>
```

Parameters
- mode=copy,move: decide, if source files should be removed (default)
- keep-duplicates: decide, if source files that already exist, should be kept (default:false)

## find
find all items by a specific query string

```
grename find <destination-path>
```
Parameters:
- --add-tag: adds tags to a find result (perhaps an extra command: modify)
- --remove-tag: removes tags from a find result
- --search-files: decides whether to search in tags only or also in files table

## cleanup
compare an existing database to an existing destination path and update existing or remove all non-existing items
```
grename cleanup <destination-path>
```

## serve
start a web server
```
grename serve <destination-path>
```

Parameters:
- --watch: directories to watch and auto-import with the given settings


# Database

- files
    - id (guid)
    - mimetype (string)
    - hash (varchar, blake2)
- tags
    - id
    - file_id
    - category (smallint, 0=CUSTOM, 1=EXIF, 2=IPTC, 3=ID3, ...)
    - qualifier (0=NONE, 1=EXIF-Make, 2=EXIF-Model)
    - tag
    


# Structs
- MetaData
    - File
        - Extension
        - Size
        - Atime
        - Mtime
    - Exif
        - CameraModel
        - ...
    - Id3
        - Album
        - Artist
        - ...
        
# Behaviour
    - Build filename by template
    - if file exists
        - Build hash
        - If hash equal, ignore / delete source
        - if hash not equal, append index
    - Copy or move file 


# gorm reference

```

// db.FirstOrCreate(&user, User{Name: "non_existing"})

//  if err != nil {
//    panic("failed to connect database")
//  }
//  defer Db.Close()
//
//  // Migrate the schema
//  Db.AutoMigrate(&Product{})
//
//  // Create
//  Db.Create(&Product{Code: "L1212", Price: 1000})
//
//  // Read
//  var product Product
//  Db.First(&product, 1) // find product with id 1
//  Db.First(&product, "code = ?", "L1212") // find product with code l1212
//
//  // Update - update product's price to 2000
//  Db.Model(&product).Update("Price", 2000)
//
//  // Delete - delete product
//  Db.Delete(&product)
	// db, err := gorm.Open("sqlite3", "test.db")
	//  if err != nil {
	//    panic("failed to connect database")
	//  }
	//  defer db.Close()
	//
	//  // Migrate the schema
	//  db.AutoMigrate(&Product{})
	//
	//  // Create
	//  db.Create(&Product{Code: "L1212", Price: 1000})
	//
	//  // Read
	//  var product Product
	//  db.First(&product, 1) // find product with id 1
	//  db.First(&product, "code = ?", "L1212") // find product with code l1212
	//
	//  // Update - update product's price to 2000
	//  db.Model(&product).Update("Price", 2000)
	//
	//  // Delete - delete product
	//  db.Delete(&product)

```
