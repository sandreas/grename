package metadata

import (
	"fmt"
	"github.com/dsoprea/go-exif"
	log "github.com/dsoprea/go-logging"
	"io/ioutil"
	"os"
	"time"
)

type Exif struct {
	Make  string // VALUE=[Canon]
	Model string // VALUE=[Canon EOS 5D Mark III]
	//Orientation int // VALUE=[1]
	//XResolution // VALUE=[72/1]
	//YResolution // VALUE=[72/1]
	//ResolutionUnit // VALUE=[2]
	// 2006-01-02T15:04:05-0700
	DateTime time.Time // VALUE=[2017:12:02 08:18:50]
}
func (e *Exif) Ss() string {
	return e.DateTime.Format("05")
}
func (e *Exif) Mm() string {
	return e.DateTime.Format("04")
}
func (e *Exif) Hh() string {
	return e.DateTime.Format("15")
}

func (e *Exif) YYYY() string {
	return e.DateTime.Format("2006")
}

func (e *Exif) MM() string {
	return e.DateTime.Format("01")
}


func (e *Exif) DD() string {
	return e.DateTime.Format("02")
}



func exifReadFromFile(filepathArgument string) (*Exif, error) {
	println("import executed")
	// filepathArgument := ""
	f, err := os.Open(filepathArgument)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	exifData, err := exif.SearchAndExtractExif(data)
	if err != nil {
		return nil, err
	}

	im := exif.NewIfdMappingWithStandard()
	ti := exif.NewTagIndex()

	exifContainer := new(Exif)
	visitor := func(fqIfdPath string, ifdIndex int, tagId uint16, tagType exif.TagType, valueContext exif.ValueContext) (err error) {
		ifdPath, err := im.StripPathPhraseIndices(fqIfdPath)
		log.PanicIf(err)

		it, err := ti.Get(ifdPath, tagId)
		if err != nil {
			if log.Is(err, exif.ErrTagNotFound) {
				fmt.Printf("WARNING: Unknown tag: [%s] (%04x)\n", ifdPath, tagId)
				return err
			} else {
				return err
			}
		}

		valueString := ""
		if tagType.Type() == exif.TypeUndefined {
			value, err := exif.UndefinedValue(ifdPath, tagId, valueContext, tagType.ByteOrder())
			if log.Is(err, exif.ErrUnhandledUnknownTypedTag) {
				valueString = "!UNDEFINED!"
			} else if err != nil {
				panic(err)
			} else {
				valueString = fmt.Sprintf("%v", value)
			}
		} else {
			valueString, err = tagType.ResolveAsString(valueContext, true)
			if err != nil {
				return err
			}
		}

		switch it.Name {
		case "Make":
			exifContainer.Make = valueString
		case "Model":
			exifContainer.Model = valueString
		case "DateTime":
			// 2017:12:02 08:18:50
			dateTime, err := time.Parse("2006:01:02 03:04:05", valueString)
			if err != nil {
				return err
			}
			exifContainer.DateTime = dateTime
		}

		// fmt.Printf("FQ-IFD-PATH=[%s] ID=(0x%04x) NAME=[%s] COUNT=(%d) TYPE=[%s] VALUE=[%s]\n", fqIfdPath, tagId, it.Name, valueContext.UnitCount, tagType.Name(), valueString)
		return nil
	}

	_, err = exif.Visit(exif.IfdStandard, im, ti, exifData, visitor)

	return exifContainer, err
}
