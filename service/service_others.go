// +build !windows

package service

func RunAsService(handler func()) bool {
	return false
}
