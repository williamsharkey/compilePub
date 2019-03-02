// Derived from Inferno utils/6l/obj.c and utils/6l/span.c
// https://bitbucket.org/inferno-os/inferno-os/src/default/utils/6l/obj.c
// https://bitbucket.org/inferno-os/inferno-os/src/default/utils/6l/span.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package ld

import (
	"github.com/williamsharkey/compilePub/internal/objabi"
	"github.com/williamsharkey/compilePub/internal/sys"
	"github.com/williamsharkey/compilePub/link/internal/sym"
	"log"
)

func linknew(arch *sys.Arch) *Link {
	ctxt := &Link{
		Syms:         sym.NewSymbols(),
		Out:          &OutBuf{arch: arch},
		Arch:         arch,
		LibraryByPkg: make(map[string]*sym.Library),
	}

	if objabi.GOARCH != arch.Name {
		log.Fatalf("invalid objabi.GOARCH %s (want %s)", objabi.GOARCH, arch.Name)
	}

	AtExit(func() {
		if nerrors > 0 && ctxt.Out.f != nil {
			ctxt.Out.f.Close()
			mayberemoveoutfile()
		}
	})

	return ctxt
}

// computeTLSOffset records the thread-local storage offset.
func (ctxt *Link) computeTLSOffset() {
	switch ctxt.HeadType {
	default:
		log.Fatalf("unknown thread-local storage offset for %v", ctxt.HeadType)

	case objabi.Hplan9, objabi.Hwindows:
		break

		/*
		 * ELF uses TLS offset negative from FS.
		 * Translate 0(FS) and 8(FS) into -16(FS) and -8(FS).
		 * Known to low-level assembly in package runtime and runtime/cgo.
		 */
	case objabi.Hlinux,
		objabi.Hfreebsd,
		objabi.Hnetbsd,
		objabi.Hopenbsd,
		objabi.Hdragonfly,
		objabi.Hsolaris:
		if objabi.GOOS == "android" {
			switch ctxt.Arch.Family {
			case sys.AMD64:
				// Android/amd64 constant - offset from 0(FS) to our TLS slot.
				// Explained in src/runtime/cgo/gcc_android_*.c
				ctxt.Tlsoffset = 0x1d0
			case sys.I386:
				// Android/386 constant - offset from 0(GS) to our TLS slot.
				ctxt.Tlsoffset = 0xf8
			default:
				ctxt.Tlsoffset = -1 * ctxt.Arch.PtrSize
			}
		} else {
			ctxt.Tlsoffset = -1 * ctxt.Arch.PtrSize
		}

	case objabi.Hnacl:
		switch ctxt.Arch.Family {
		default:
			log.Fatalf("unknown thread-local storage offset for nacl/%s", ctxt.Arch.Name)

		case sys.ARM:
			ctxt.Tlsoffset = 0

		case sys.AMD64:
			ctxt.Tlsoffset = 0

		case sys.I386:
			ctxt.Tlsoffset = -8
		}

		/*
		 * OS X system constants - offset from 0(GS) to our TLS.
		 * Explained in src/runtime/cgo/gcc_darwin_*.c.
		 */
	case objabi.Hdarwin:
		switch ctxt.Arch.Family {
		default:
			log.Fatalf("unknown thread-local storage offset for darwin/%s", ctxt.Arch.Name)

		case sys.ARM:
			ctxt.Tlsoffset = 0 // dummy value, not needed

		case sys.AMD64:
			ctxt.Tlsoffset = 0x8a0

		case sys.ARM64:
			ctxt.Tlsoffset = 0 // dummy value, not needed

		case sys.I386:
			ctxt.Tlsoffset = 0x468
		}
	}

}