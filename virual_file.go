package utils

import "io"

//VirtualFile 在内存中虚拟文件操作
type VirtualFile struct {
	data []byte
}

//Write  io.Write 接口实现
func (vf *VirtualFile) Write(p []byte) (int, error) {
	if vf.data == nil {
		vf.data = []byte{}
	}
	vf.data = append(vf.data, p...)
	return len(p), nil
}

// io.Reader 接口实现
func (vf *VirtualFile) Read(p []byte) (n int, err error) {
	l := copy(p, vf.data)
	if l > 0 {
		vf.data = vf.data[l:]
		return l, nil
	}

	return 0, io.EOF

}

//Seek ...
func (vf *VirtualFile) Seek(offset int64, whence int) (int64, error) {

	return 0, nil
}

// // ReadSeeker is the interface that groups the basic Read and Seek methods.
// type ReadSeeker interface {
// 	Reader
// 	Seeker
// }
