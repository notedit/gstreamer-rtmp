package gstrtmp

/*
#cgo pkg-config: gstreamer-1.0 gstreamer-app-1.0

#include "gst.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func init() {
	go C.gst_rtmp_start_mainloop()
}

type Pipeline struct {
	Pipeline *C.GstElement
}

func CreatePipeline(rtmpUrl string) *Pipeline {

	pipelineStr := "appsrc is-live=true do-timestamp=true name=src ! h264parse ! video/x-h264,stream-format=(string)avc ! flvmux ! rtmpsink location='%s live=1'"

	pipelineStr = fmt.Sprintf(pipelineStr, rtmpUrl)

	pipelineStrUnsafe := C.CString(pipelineStr)
	defer C.free(unsafe.Pointer(pipelineStrUnsafe))
	return &Pipeline{Pipeline: C.gst_rtmp_create_pipeline(pipelineStrUnsafe)}
}

func (p *Pipeline) Start() {
	C.gst_rtmp_start_pipeline(p.Pipeline)
}

func (p *Pipeline) Stop() {
	C.gst_rtmp_stop_pipeline(p.Pipeline)
}

// Push pushes a buffer on the appsrc of the GStreamer Pipeline
func (p *Pipeline) Push(buffer []byte) {
	b := C.CBytes(buffer)
	defer C.free(unsafe.Pointer(b))
	C.gst_rtmp_push_buffer(p.Pipeline, b, C.int(1), C.int(len(buffer)))
}
