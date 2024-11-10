// Code generated by linux/mkall.go generatePtracePair("arm", "arm64"). DO NOT EDIT.

//go:build linux && (arm || arm64)

package unix

import "unsafe"

// PtraceRegsArm is the registers used by arm binaries.
type PtraceRegsArm struct {
	Uregs [18]uint32
}

// PtraceGetRegsArm fetches the registers used by arm binaries.
func PtraceGetRegsArm(pid int, regsout *PtraceRegsArm) error {
	return ptracePtr(PTRACE_GETREGS, pid, 0, unsafe.Pointer(regsout))
}

// PtraceSetRegsArm sets the registers used by arm binaries.
func PtraceSetRegsArm(pid int, regs *PtraceRegsArm) error {
	return ptracePtr(PTRACE_SETREGS, pid, 0, unsafe.Pointer(regs))
}

// PtraceRegsArm64 is the registers used by arm64 binaries.
type PtraceRegsArm64 struct {
	Regs   [31]uint64
	Sp     uint64
	Pc     uint64
	Pstate uint64
}

// PtraceGetRegsArm64 fetches the registers used by arm64 binaries.
func PtraceGetRegsArm64(pid int, regsout *PtraceRegsArm64) error {
	return ptracePtr(PTRACE_GETREGS, pid, 0, unsafe.Pointer(regsout))
}

// PtraceSetRegsArm64 sets the registers used by arm64 binaries.
func PtraceSetRegsArm64(pid int, regs *PtraceRegsArm64) error {
	return ptracePtr(PTRACE_SETREGS, pid, 0, unsafe.Pointer(regs))
}