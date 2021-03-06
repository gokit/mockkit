package static


import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

type fileData struct{
  path string
  root string
  data []byte
}

var (
  assets = map[string][]string{
    
      ".tml": []string{  // all .tml assets.
        
          "impl-only.tml",
        
          "impl.tml",
        
          "mock.tml",
        
      },
    
  }

  assetFiles = map[string]fileData{
    
      
        "impl-only.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\xbd\x6e\xdc\x3c\x10\xec\x05\xe8\x1d\xa6\xb8\xc2\x6e\xe8\xaf\xfe\x00\x17\x29\x62\xe0\x80\x24\x30\x02\xbf\x00\x8f\x5a\x49\xc4\x51\xa4\x42\x52\xe7\x1c\x16\x7a\xf7\x40\x14\x75\x7f\xd6\x21\x4e\x29\xee\xec\xec\xec\xcc\x8a\x79\xb3\xb5\x91\x7c\x2d\x15\xfd\x90\x1d\xe1\xff\x67\x88\xeb\x97\x71\x2c\x0b\xe6\xcd\xab\x54\x7b\xd9\x9c\x41\xaf\xfb\x46\xe4\xb7\x04\x29\x8b\xa7\x27\x30\x5f\x37\x8f\xe3\xb6\xeb\x0d\x2a\xaa\xb5\xa5\x00\x09\x43\x8d\xc3\xce\x38\xb5\x47\x88\x7e\x50\x11\xef\xad\x56\x2d\x02\xf9\x83\x56\x14\x10\x5b\x42\x47\xb1\x75\x55\xc8\x6d\x55\x62\xde\x1d\x53\xe9\xe3\x00\xe8\xe5\x1b\x32\x40\x86\xa0\x1b\x2b\x77\x86\x50\x6b\x32\x55\x10\x78\x6b\x29\x10\xea\xc1\xaa\xa8\x9d\x0d\x90\x9e\xa0\xa4\x31\x99\xf8\xbd\x25\x9b\xa8\x3d\x85\x9e\x54\xd4\x87\x4b\xb0\xab\x53\x2d\x6b\xd5\x21\x77\x8a\xb2\x88\xc7\x7e\x4d\x4e\xda\x77\x86\x73\x59\x00\x00\x33\xbc\xb4\x0d\x41\x7c\xcf\x7b\x4d\x76\xcd\x15\x31\x37\xbd\x0c\x56\xa5\xa1\x0f\xcc\xe2\xeb\xef\x2f\xbe\x19\x3a\xb2\xf1\x9b\x0e\x11\xd1\x0f\x84\x4b\xf7\xc7\xf1\x11\x33\xee\x27\xc5\xc1\xdb\xbb\xa8\x9b\x21\x5b\x7b\x70\x7b\xaa\xb0\x73\xce\x9c\x94\x91\xad\x92\x9c\x14\xe0\xba\xd2\x39\xd6\xc5\xec\xae\x37\x34\x69\x0b\x39\x8e\xcd\xcd\xfe\xe2\x8c\x9d\x63\x44\xed\xfc\x0a\xee\x1c\x9b\x48\x23\xb6\x31\x59\x7b\x8f\x76\xb2\x55\xdc\xf8\x35\xe5\x8b\xc1\x56\xe4\x2d\xc9\xd8\x8a\xb2\x98\x2c\xc4\xc3\xa4\xf1\x0e\xc5\xe3\x79\x95\xcf\x5a\xcd\xac\x6b\x58\x17\xe9\xd7\x20\x0d\xc4\x9b\x8b\xd2\xcc\xc6\x07\xfc\x87\x85\xe7\x2f\x51\x30\x93\xad\xc6\x91\x67\x36\xea\xfa\x78\x84\x98\x7b\x26\x44\x98\x1a\x97\xab\xd0\x57\xab\x2e\xa9\x3d\x27\xde\x35\xc4\x4b\xbe\x9b\x65\x95\x0b\xc2\xd3\x09\x90\x09\xf4\x79\x7e\xe6\x8f\xda\xa6\x7f\xfe\x5f\xe7\xfa\x44\xb2\x4a\x97\x0d\x49\x87\x77\x3a\xc3\x3f\x01\x00\x00\xff\xff\x7c\x87\x0a\x25\x8d\x04\x00\x00"),
          path: "impl-only.tml",
          root: "impl-only.tml",
        },
      
        "impl.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\x41\x8b\xdb\x30\x10\x85\xef\x06\xff\x87\x77\xd8\x43\x72\xd1\xde\x0b\x7b\xe8\x65\x21\xd0\xf6\x50\xfa\x07\x14\x79\x1c\x8b\x95\x25\xa3\x91\xb3\x04\xe1\xff\x5e\x24\x2b\x89\xe3\x75\xa0\x3d\xc6\x33\xf3\xbd\x99\xf7\x94\x18\x5f\x0e\x36\x90\x6f\xa5\xa2\x5f\xb2\x27\x7c\x7b\x83\x78\xfc\x32\x4d\x75\x55\x57\xaf\xaf\x88\xf1\xb1\x32\x4d\x87\x7e\x30\x68\xa8\xd5\x96\x18\x12\x86\x4e\x0e\x47\xe3\xd4\x07\x38\xf8\x51\x05\x7c\x76\x5a\x75\x60\xf2\x67\xad\x88\x11\x3a\x42\x4f\xa1\x73\x0d\x97\xb1\x26\x93\x8f\x97\x5c\xfa\x2a\x00\x7d\xfd\x0d\xc9\x90\xcc\xfa\x64\xe5\xd1\x10\x5a\x4d\xa6\x61\x81\x3f\x1d\x31\xa1\x1d\xad\x0a\xda\x59\x86\xf4\x04\x25\x8d\x29\xe0\xcf\x8e\x6c\x46\x7b\xe2\x81\x54\xd0\xe7\x65\xb3\x6b\x73\xad\xec\xaa\xb9\x4c\x8a\xba\x0a\x97\x61\x6b\x9d\x7c\xef\xdc\x1e\xeb\x0a\x00\x62\x84\x97\xf6\x44\x10\x3f\xcb\x5d\xc9\xae\xb9\x22\xe6\xa1\xf7\xd1\xaa\x2c\xba\x8b\x51\x7c\xf7\xa7\xb1\x27\x1b\x7e\x68\x0e\x08\x7e\xa4\x69\xda\x23\x15\x7e\x53\x18\xbd\x5d\x7e\x5e\x61\x0e\xf6\xec\x3e\xa8\xc1\xd1\x39\x73\xd3\x26\xdb\x64\xc1\x1c\xd1\xf6\x2e\x73\x70\x57\x3b\xfb\xc1\x50\xd2\xe7\x62\xf8\xcb\xea\x42\x71\xef\x9d\x83\x42\xeb\xfc\x46\xdf\x3d\x18\x91\x25\x0e\x21\x9b\xf7\x0c\x9b\x8c\x13\x2b\x47\x52\x82\x18\x6d\x43\xde\x92\x0c\x9d\xa8\xab\x64\x12\x76\x69\xc7\x27\x88\xfd\xfd\x94\xa7\x66\xc6\xa8\x5b\x58\x17\xa8\x1f\xc2\x05\xbb\xb5\xb1\xfb\x79\xf4\xab\xdd\x31\x92\x6d\xa6\x29\xce\x80\x79\xba\xb4\x25\x45\x4e\xbd\xd7\x6c\xf5\xc3\x39\xd7\x64\xde\x32\x6a\xab\xe3\x7d\x95\xfe\x02\x78\x8b\x99\x0c\xd3\xbf\xf3\x6f\x27\x2c\x50\xe9\xbf\xfb\xbf\xba\x3e\x43\x36\x71\xc5\x90\xfc\xb8\x6e\x4f\xed\x6f\x00\x00\x00\xff\xff\x8c\xb3\x35\x58\x30\x04\x00\x00"),
          path: "impl.tml",
          root: "impl.tml",
        },
      
        "mock.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x57\xdf\x8b\xe3\x36\x10\x7e\x5f\xd8\xff\x61\xee\xe8\x83\x5d\x52\x67\xfb\x7a\x25\x0f\x47\x9b\x85\x42\x6f\x29\xdd\x83\x3e\x2c\x4b\x51\xec\x71\x2c\xd6\x96\x8c\x24\x27\x9b\x0a\xff\xef\x65\x24\x39\xfe\x11\x3b\xec\xb5\x14\xda\xf3\x4b\xf0\x68\x34\x33\xfa\xf4\xcd\x37\xce\xed\x8d\xb5\xdf\xfc\x2c\x0c\xaa\x9c\xa5\xf8\xc0\x2a\x84\x0f\x1b\x48\xc6\x96\xb6\xbd\xbd\x21\x47\x50\x4c\xec\x11\x92\x4f\x68\x0a\x99\x69\x67\x5f\xaf\xc1\xbf\xfe\xc8\xca\xf2\x5e\x2a\x6b\x13\xda\xd3\xb6\x90\x61\xce\x05\x6a\x60\x60\x4e\x35\xc2\xb1\xe0\x69\x01\x85\x2c\x33\x0d\x15\x1a\xf6\x5d\x86\x86\xf1\x52\x03\xdb\xc9\xc6\x80\x29\x10\xf6\xfc\xc0\xc5\x1e\x52\x56\x92\x59\x6b\x99\x72\x66\x30\x73\x49\x8e\xdc\x14\xce\x69\x5a\x6f\xdb\x26\xe7\x9c\x51\x4c\xa1\x0b\x99\x25\xb7\x37\x2e\xe9\x42\x69\xda\xa8\x26\x35\xf6\xf6\x06\x00\xe0\xf7\x02\x05\x18\x5e\x61\xf2\x99\x57\xe8\x6d\x8f\x86\x29\x33\x35\x6e\x45\x36\x34\x79\xe3\x7a\x0d\x3f\x85\x83\xc8\x1c\x6a\x26\x78\x0a\x3c\x07\xdd\xa4\x05\xc8\x34\x6d\x94\x4e\xbc\xe3\xaf\xb4\xf4\x68\x58\xfa\x02\x4f\xcf\xbb\x93\xc1\x81\x79\xab\x94\x54\xc0\xbb\x63\xd9\x76\x10\xfd\xa3\xda\x37\x15\x0a\x03\x07\x56\x36\xd8\x45\xeb\x2f\xe3\xa3\xda\xfb\x9b\xf0\xe6\x94\xd5\xdc\xb0\x92\xff\x89\xd0\x9d\xd6\xda\x64\xfb\xfa\xf9\x54\x63\xef\x85\x22\x6b\x87\x49\x7e\x43\xd3\x28\xb1\x94\xc2\xaf\x7e\x79\x16\x08\x69\x5c\x26\x22\x0a\x33\x69\xd1\x9d\x47\x83\x0a\x61\x8d\x6a\x70\x9d\xb3\x52\x23\x21\x57\x2b\x79\xe0\x19\x2a\x90\xa6\x40\x15\xcd\xdf\x60\x0c\xec\x1c\xc6\x85\xf6\x95\x43\x45\x19\xc0\x14\x5c\x03\xbe\x72\x6d\x88\x4e\xec\x02\xc0\xbc\x11\x29\x44\xd5\x12\x3d\xe2\x49\xa1\x91\xab\x64\xd1\x79\x27\x65\x09\xf6\xea\xb5\xd0\xc3\x73\x78\xa7\x30\x2f\x31\x35\x5c\x8a\xc4\xa5\xd8\x96\x48\x19\xa2\x0a\x93\x39\x4c\x57\x1e\x83\x85\x35\x82\x2d\xb6\x7d\x02\x7a\x3c\xa4\xe0\xc0\xec\x57\x46\x17\x72\x2e\x2a\xf8\x52\x18\x77\x45\x83\xe5\x70\x5f\xd6\x26\x93\x56\xfb\x24\xd3\x97\xf9\xc6\xe6\x55\xed\x0f\x43\x76\xdf\x5f\xe7\x9e\x75\xc1\xaa\x20\x1a\xb9\x54\xa1\x8f\xa7\xc1\x81\x69\xc8\x39\x92\x40\xf8\x98\xac\x2c\xe5\x51\xc3\x49\x36\x1d\x2b\xfa\x34\x8c\x50\xa4\x9e\x73\xc1\x4d\x81\x1a\x81\xee\xd5\x9b\x8d\x3c\xef\xc8\x4b\x7c\xe5\xbb\x12\xc1\xa0\xe3\x43\x27\x0d\x0b\x87\x1b\x49\xc3\xbc\xe4\xf9\x95\x70\x0f\x3d\x2b\x34\x3c\x3d\xcf\x73\xe4\x02\xfe\xeb\x7a\xda\xeb\xd4\x00\xd4\x37\x4b\xdf\x00\xe0\x0b\xef\x8e\xfa\x14\x17\xbe\xbd\x74\x20\x04\xe2\x3e\x7f\x64\x6d\xd2\x75\xc1\x2f\x5c\x1b\xc7\x15\xa2\xfc\x42\x57\xae\x00\x49\xc8\xe2\xae\x1b\x0e\x4c\x39\x31\x5f\x6c\x1e\xcf\x34\x72\xf5\x6e\x89\xd3\xe2\x8d\x57\xd9\x07\x79\x8c\xe2\xd1\xaa\x57\xe5\xcd\xd0\xb9\xdb\xbf\xd0\x78\xc1\x73\x56\xb0\x36\xb0\x7c\x41\x7d\xfd\xb9\x6c\x44\xe6\x7a\xdc\xdb\x08\xdd\x3f\x56\x50\x4b\xad\x89\x55\x74\x22\x9a\x98\x3e\x39\xe1\x9a\xcc\x51\xc3\x8e\x74\x60\xb8\x39\x99\x68\x8d\xaf\x78\xda\xd7\xbe\x8a\x4d\x68\xd6\xe1\xca\xb2\x4a\x0f\x9f\xeb\x38\x8c\xea\x99\xf3\xb9\x48\x39\xc6\x69\x92\x65\x30\xd2\x26\xb1\xfb\x95\xe5\x8d\x7e\x44\xce\x6d\x74\x2b\xe3\x8d\x3b\x85\xec\xe5\x42\xe8\xda\x09\xab\xb6\x1e\xbb\x09\xa9\x78\x1e\x60\x1d\x60\x1d\x34\xd1\xef\x5b\x81\xe0\xe5\x24\xe2\xc4\xc1\xf1\x5d\x27\x0f\x78\x8c\xde\x0b\xe9\xc7\x0f\xcd\x1c\x85\xba\x96\x82\x24\x89\x32\xbc\x8f\x2f\x05\x76\x56\x5e\x1f\x05\xa7\xf9\xf5\xb5\x0a\x6c\x38\xde\xbf\x29\xb1\xe1\xf5\x9e\x64\x8e\xaa\x5d\x96\x30\x6b\x43\xc7\x0c\xcd\xff\x1b\xa5\xf6\x50\xbe\x45\xab\xad\xe5\x39\x08\x69\xb0\xaa\xcd\x09\xa2\xe9\xa1\x63\xbf\xf5\x12\x8a\xf0\x91\xf8\x25\x52\x4e\x8e\x19\xe6\xa8\x3c\xf4\xf1\x58\xf5\x50\x29\xa7\x94\x98\xca\x03\xaa\x28\xfe\xc1\x59\xde\x6d\xa8\xcb\x60\xa2\x77\x46\xb1\xd4\xfd\x13\xa9\xd8\x0b\x46\xfe\x7b\x79\x05\xdf\xdf\xdd\xdd\xc5\x73\x8e\x1b\xff\xfb\xf4\x41\x35\xc2\xb5\xb9\x93\x8a\xc8\x19\xc3\x87\xd2\xf3\x5b\xc4\x0a\xd5\x1b\x84\xc9\x45\x1d\x69\x4e\xff\x72\x55\x70\x1c\x0e\x4b\x13\x02\x36\xc0\xea\x1a\x45\x16\x2d\xba\xac\x42\xf8\x10\xae\xa5\xb8\xff\xe1\xe9\xe9\x88\xe7\x59\x17\xe8\x45\xee\x9a\x38\x36\xfa\x2e\x1e\x1d\xf7\x7e\xd2\xb3\x83\x2d\xe7\xe6\xc4\x52\x8f\x66\xd2\x99\xbe\x03\x67\xe2\xce\xdf\x88\x3c\xfc\x5b\x74\x7d\xb2\xfe\x33\x6c\xc2\x18\x99\xab\x7c\x3a\x2a\xfe\x0a\x00\x00\xff\xff\x8e\x87\x21\xc1\xa2\x0f\x00\x00"),
          path: "mock.tml",
          root: "mock.tml",
        },
      
    
  }
)

//==============================================================================

// FilesFor returns all files that use the provided extension, returning a
// empty/nil slice if none is found.
func FilesFor(ext string) []string {
  return assets[ext]
}

// MustFindFile calls FindFile to retrieve file reader with path else panics.
func MustFindFile(path string, doGzip bool) (io.Reader, int64) {
  reader, size, err := FindFile(path, doGzip)
  if err != nil {
    panic(err)
  }

  return reader, size
}

// FindDecompressedGzippedFile returns a io.Reader by seeking the giving file path if it exists.
// It returns an uncompressed file.
func FindDecompressedGzippedFile(path string) (io.Reader, int64, error){
	return FindFile(path, true)
}

// MustFindDecompressedGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindDecompressedGzippedFile(path string) (io.Reader, int64){
	reader, size, err := FindDecompressedGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindGzippedFile returns a io.Reader by seeking the giving file path if it exists.
// It returns an uncompressed file.
func FindGzippedFile(path string) (io.Reader, int64, error){
	return FindFile(path, false)
}

// MustFindGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindGzippedFile(path string) (io.Reader, int64){
	reader, size, err := FindGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFile returns a io.Reader by seeking the giving file path if it exists.
func FindFile(path string, doGzip bool) (io.Reader, int64, error){
	reader, size, err := FindFileReader(path)
	if err != nil {
		return nil, size, err
	}

	if !doGzip {
		return reader, size, nil
	}

  gzr, err := gzip.NewReader(reader)
	return gzr, size, err
}

// MustFindFileReader returns bytes.Reader for path else panics.
func MustFindFileReader(path string) (*bytes.Reader, int64){
	reader, size, err := FindFileReader(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFileReader returns a io.Reader by seeking the giving file path if it exists.
func FindFileReader(path string) (*bytes.Reader, int64, error){
  item, ok := assetFiles[path]
  if !ok {
    return nil,0, fmt.Errorf("File %q not found in file system", path)
  }

  return bytes.NewReader(item.data), int64(len(item.data)), nil
}

// MustReadFile calls ReadFile to retrieve file content with path else panics.
func MustReadFile(path string, doGzip bool) string {
  body, err := ReadFile(path, doGzip)
  if err != nil {
    panic(err)
  }

  return body
}

// ReadFile attempts to return the underline data associated with the given path
// if it exists else returns an error.
func ReadFile(path string, doGzip bool) (string, error){
  body, err := ReadFileByte(path, doGzip)
  return string(body), err
}

// MustReadFileByte calls ReadFile to retrieve file content with path else panics.
func MustReadFileByte(path string, doGzip bool) []byte {
  body, err := ReadFileByte(path, doGzip)
  if err != nil {
    panic(err)
  }

  return body
}

// ReadFileByte attempts to return the underline data associated with the given path
// if it exists else returns an error.
func ReadFileByte(path string, doGzip bool) ([]byte, error){
  reader, _, err := FindFile(path, doGzip)
  if err != nil {
    return nil, err
  }

  if closer, ok := reader.(io.Closer); ok {
    defer closer.Close()
  }

  var bu bytes.Buffer

  _, err = io.Copy(&bu, reader);
  if err != nil && err != io.EOF {
   return nil, fmt.Errorf("File %q failed to be read: %+q", path, err)
  }

  return bu.Bytes(), nil
}
