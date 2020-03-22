package opencl

// #cgo !darwin LDFLAGS: -lOpenCL
// #cgo darwin LDFLAGS: -framework OpenCL
//
// #define CL_USE_DEPRECATED_OPENCL_1_2_APIS 1
// #ifdef __APPLE__// #include <OpenCL/opencl.h>
// #else
// #include <CL/opencl.h>
// #endif
import "C"

import "unsafe"

// Enqueues a command to fill a buffer object with a pattern of a
// given pattern size.

func (q *CommandQueue) EnqueueFillBuffer(buffer *MemObject, pattern unsafe.Pointer, patternSize, offset, size int, eventWaitList []*Event) (*Event, error) {
	var event C.cl_event
	err := toError(C.clEnqueueFillBuffer(q.clQueue, buffer.clMem, pattern, C.size_t(patternSize), C.size_t(offset), C.size_t(size), C.cl_uint(len(eventWaitList)), eventListPtr(eventWaitList), &event))
	return newEvent(event), err
}

// A synchronization point that enqueues a barrier operation.
func (q *CommandQueue) EnqueueBarrierWithWaitList(eventWaitList []*Event) (*Event, error) {
	var event C.cl_event
	err := toError(C.clEnqueueBarrierWithWaitList(q.clQueue, C.cl_uint(len(eventWaitList)), eventListPtr(eventWaitList), &event))
	return newEvent(event), err
}

// Enqueues a marker command which waits for either a list of events to complete, or all previously enqueued commands to complete.
func (q *CommandQueue) EnqueueMarkerWithWaitList(eventWaitList []*Event) (*Event, error) {
	var event C.cl_event
	err := toError(C.clEnqueueMarkerWithWaitList(q.clQueue, C.cl_uint(len(eventWaitList)), eventListPtr(eventWaitList), &event))
	return newEvent(event), err
}
