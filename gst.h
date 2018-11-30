#ifndef GST_H
#define GST_H

#include <glib.h>
#include <gst/gst.h>
#include <stdint.h>
#include <stdlib.h>

GstElement *gst_rtmp_create_pipeline(char *pipeline);
void gst_rtmp_start_pipeline(GstElement *pipeline);
void gst_rtmp_stop_pipeline(GstElement *pipeline);
void gst_rtmp_push_buffer(GstElement *pipeline, void *buffer, int mtype, int len);
void gst_rtmp_start_mainloop(void);

#endif