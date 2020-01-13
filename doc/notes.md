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
