package bass
/*
#cgo linux,386 LDFLAGS: -Wl,-rpath=${SRCDIR}/lib/linux_386 -L${SRCDIR}/lib/linux_386 -Wl,-Rpath=${SRCDIR}/lib/linux_386
#cgo linux,amd64 LDFLAGS: -Wl,-rpath=${SRCDIR}/lib/linux_amd64 -L${SRCDIR}/lib/linux_amd64
#cgo windows,386 LDFLAGS: -L${SRCDIR}/lib/windows_386
#cgo windows,amd64 LDFLAGS: -L${SRCDIR}/lib/windows_amd64
#cgo darwin LDFLAGS: -L/opt/homebrew/var/bass
#cgo LDFLAGS: -lbass

#cgo CFLAGS: -I${SRCDIR}/include
*/
import "C"