package gstrtmp

/*
#cgo pkg-config: gstreamer-1.0 gstreamer-app-1.0
#include "gst.h"
*/
import "C"

func init() {
	go C.gst_rtmp_start_mainloop()
}

type Pipeline struct {
}
