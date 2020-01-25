package log

func NewFileWrapper(filename)  {
    return &FileWrapper {
        filename: filename
    }
}

type FileWrapper struct {
    filename string
}
