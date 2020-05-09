// +build linux
// from https://github.com/nsf/termbox-go/

// package tty - utilities to work with terminals.
package tty

import "syscall"
import "os"
import "unsafe"

type sc_Termios syscall.Termios

const (
	sc_ECHO   = syscall.ECHO
	sc_ICANON = syscall.ICANON
	sc_VMIN   = syscall.VMIN

	sc_TCGETS = syscall.TCGETS
	sc_TCSETS = syscall.TCSETS
)

var (
	orig_tios sc_Termios
	out       *os.File
)

// disable input buffering
// similar  exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
// do not display entered characters on the screen
// similar  exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
func RawMode() {
	var err error

	out, err = os.OpenFile("/dev/tty", syscall.O_WRONLY, 0)
	if err != nil {
		Error("RawMode(): ", err)
		return
	}

	err = tcgetattr(out.Fd(), &orig_tios)
	if err != nil {
		Error("RawMode(): ", err)
		return
	}

	tios := orig_tios
	tios.Lflag &^= sc_ECHO | sc_ICANON
	tios.Cc[sc_VMIN] = 1

	err = tcsetattr(out.Fd(), &tios)
	if err != nil {
		Error("RawMode(): ", err)
	}
}

// returns the default mode
func OrigMode() {
	if out != nil {
		err := tcsetattr(out.Fd(), &orig_tios)
		if err != nil {
			Error("OrigMode(): ", err)
		}
	}
}

func tcsetattr(fd uintptr, termios *sc_Termios) error {
	r, _, e := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(sc_TCSETS), uintptr(unsafe.Pointer(termios)))
	if r != 0 {
		return os.NewSyscallError("SYS_IOCTL", e)
	}
	return nil
}

func tcgetattr(fd uintptr, termios *sc_Termios) error {
	r, _, e := syscall.Syscall(syscall.SYS_IOCTL,
		fd, uintptr(sc_TCGETS), uintptr(unsafe.Pointer(termios)))
	if r != 0 {
		return os.NewSyscallError("SYS_IOCTL", e)
	}
	return nil
}
