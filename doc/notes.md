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
