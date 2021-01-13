/*
	Package main - transpiled by c2go version: v0.26.0 Erbium 2020-03-17

	If you have found any issues, please raise an issue at:
	https://github.com/elliotchance/c2go/
*/

// Warning (RecordDecl):  /usr/include/x86_64-linux-gnu/bits/libio.h:149 : could not lookup type definition for : _IO_FILE
// Warning (FieldDecl):  /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:173 : Error : name of FieldDecl is empty
// Warning (IndirectFieldDecl):  /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:175 : could not parse &{94913359793864 {/usr/include/x86_64-linux-gnu/bits/thread-shared-types.h 175 0 42 0 } col:42 true __wseq unsigned long long [0xc0000c3380 0xc0000c3480]}
// Warning (IndirectFieldDecl):  /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:180 : could not parse &{94913359793944 {/usr/include/x86_64-linux-gnu/bits/thread-shared-types.h 180 0 7 0 } col:7 true __wseq32 struct (anonymous struct at /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:176:5)':'struct __pthread_cond_s::(anonymous at /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:176:5) [0xc0000c3540 0xc0000c3600]}
// Warning (FieldDecl):  /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:182 : Error : name of FieldDecl is empty
// Warning (IndirectFieldDecl):  /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:184 : could not parse &{94913359795008 {/usr/include/x86_64-linux-gnu/bits/thread-shared-types.h 184 0 42 0 } col:42 true __g1_start unsigned long long [0xc0000c3840 0xc0000c3900]}
// Warning (IndirectFieldDecl):  /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:189 : could not parse &{94913359795088 {/usr/include/x86_64-linux-gnu/bits/thread-shared-types.h 189 0 7 0 } col:7 true __g1_start32 struct (anonymous struct at /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:185:5)':'struct __pthread_cond_s::(anonymous at /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:185:5) [0xc0000c3a00 0xc0000c3b00]}
// Warning (RecordDecl):  /usr/include/x86_64-linux-gnu/bits/pthreadtypes.h:75 : could not determine the size of type `union pthread_cond_t` for that reason: Cannot determine sizeof : |union pthread_cond_t|. err = Cannot determine sizeof : |struct __pthread_cond_s|. err = Cannot determine sizeof : |__pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD173D17E|. err = error in array size
// Error (RecordDecl):  /usr/include/x86_64-linux-gnu/bits/pthreadtypes.h:75 : Cannot determine sizeof : |union pthread_cond_t|. err = Cannot determine sizeof : |struct __pthread_cond_s|. err = Cannot determine sizeof : |__pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD173D17E|. err = error in array size

package gosalsa

import (
	"errors"
	"unsafe"

	"github.com/elliotchance/c2go/noarch"
)

type size_t uint32
type __u_char uint8
type __u_short uint16
type __u_int uint32
type __u_long uint32
type __int8_t int8
type __uint8_t uint8
type __int16_t int16
type __uint16_t uint16
type __int32_t int32
type __uint32_t uint32
type __int64_t int32
type __uint64_t uint32
type __quad_t int32
type __u_quad_t uint32
type __intmax_t int32
type __uintmax_t uint32
type __dev_t uint32
type __uid_t uint32
type __gid_t uint32
type __ino_t uint32
type __ino64_t uint32
type __mode_t uint32
type __nlink_t uint32
type __off_t int32
type __off64_t int32
type __pid_t int32
type __fsid_t struct {
	__val [2]int32
}
type __clock_t int32
type __rlim_t uint32
type __rlim64_t uint32
type __id_t uint32
type __time_t int32
type __useconds_t uint32
type __suseconds_t int32
type __daddr_t int32
type __key_t int32
type __clockid_t int32
type __timer_t unsafe.Pointer
type __blksize_t int32
type __blkcnt_t int32
type __blkcnt64_t int32
type __fsblkcnt_t uint32
type __fsblkcnt64_t uint32
type __fsfilcnt_t uint32
type __fsfilcnt64_t uint32
type __fsword_t int32
type __ssize_t int32
type __syscall_slong_t int32
type __syscall_ulong_t uint32
type __loff_t __off64_t
type __caddr_t *byte
type __intptr_t int32
type __socklen_t uint32
type __sig_atomic_t int32
type __FILE _IO_FILE
type FILE _IO_FILE
type BSunionSatSSusrSincludeSx86_64TlinuxTgnuSbitsStypesS__mbstate_tPhD16D3E struct{ memory unsafe.Pointer }

func (unionVar *BSunionSatSSusrSincludeSx86_64TlinuxTgnuSbitsStypesS__mbstate_tPhD16D3E) copy() BSunionSatSSusrSincludeSx86_64TlinuxTgnuSbitsStypesS__mbstate_tPhD16D3E {
	var buffer [8]byte
	for i := range buffer {
		buffer[i] = (*((*[8]byte)(unionVar.memory)))[i]
	}
	var newUnion BSunionSatSSusrSincludeSx86_64TlinuxTgnuSbitsStypesS__mbstate_tPhD16D3E
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *BSunionSatSSusrSincludeSx86_64TlinuxTgnuSbitsStypesS__mbstate_tPhD16D3E) __wch() *uint32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*uint32)(unionVar.memory)
}
func (unionVar *BSunionSatSSusrSincludeSx86_64TlinuxTgnuSbitsStypesS__mbstate_tPhD16D3E) __wchb() *[4]byte {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[4]byte)(unionVar.memory)
}

type __mbstate_t struct {
	__count int32
	__value BSunionSatSSusrSincludeSx86_64TlinuxTgnuSbitsStypesS__mbstate_tPhD16D3E
}
type _G_fpos_t struct {
	__pos   __off_t
	__state int64
}
type _G_fpos64_t struct {
	__pos   __off64_t
	__state int64
}
type va_list int64
type __gnuc_va_list int64
type _IO_jump_t struct {
}
type _IO_lock_t interface{}
type _IO_marker struct {
	_next *_IO_marker
	_sbuf *_IO_FILE
	_pos  int32
}
type __codecvt_result int32

const ( // Warning (RecordDecl):  /usr/include/x86_64-linux-gnu/bits/libio.h:149 : could not lookup type definition for : _IO_FILE
	__codecvt_ok      __codecvt_result = 0
	__codecvt_partial                  = 1
	__codecvt_error                    = 2
	__codecvt_noconv                   = 3
)

type _IO_FILE struct {
	_flags          int32
	_IO_read_ptr    *byte
	_IO_read_end    *byte
	_IO_read_base   *byte
	_IO_write_base  *byte
	_IO_write_ptr   *byte
	_IO_write_end   *byte
	_IO_buf_base    *byte
	_IO_buf_end     *byte
	_IO_save_base   *byte
	_IO_backup_base *byte
	_IO_save_end    *byte
	_markers        *_IO_marker
	_chain          *_IO_FILE
	_fileno         int32
	_flags2         int32
	_old_offset     __off_t
	_cur_column     uint16
	_vtable_offset  int8
	_shortbuf       [1]byte
	_lock           *_IO_lock_t
	_offset         __off64_t
	__pad1          unsafe.Pointer
	__pad2          unsafe.Pointer
	__pad3          unsafe.Pointer
	__pad4          unsafe.Pointer
	__pad5          uint32
	_mode           int32
	_unused2        [20]byte
}
type _IO_FILE_plus struct {
}
type __io_read_fn func(unsafe.Pointer, *byte, uint32) __ssize_t
type __io_write_fn func(unsafe.Pointer, *byte, uint32) __ssize_t
type __io_seek_fn func(unsafe.Pointer, *__off64_t, int32) int32
type __io_close_fn func(unsafe.Pointer) int32
type cookie_read_function_t __io_read_fn
type cookie_write_function_t __io_write_fn
type cookie_seek_function_t __io_seek_fn
type cookie_close_function_t __io_close_fn
type _IO_cookie_io_functions_t struct {
	read  __io_read_fn
	write __io_write_fn
	seek  __io_seek_fn
	close __io_close_fn
}
type cookie_io_functions_t _IO_cookie_io_functions_t
type _IO_cookie_file struct {
}
type off_t __off_t
type off64_t __off64_t
type ssize_t __ssize_t
type fpos_t _G_fpos_t
type fpos64_t _G_fpos64_t

var stdin *noarch.File

var stdout *noarch.File

var stderr *noarch.File

type obstack struct {
}
type wchar_t int32

const P_ALL int32 = 0
const P_PID int32 = 1
const P_PGID int32 = 2

type idtype_t int32
type _Float32 float32
type _Float64 float64
type _Float32x float64
type _Float64x float64
type div_t struct {
	quot int32
	rem  int32
}
type ldiv_t struct {
	quot int32
	rem  int32
}
type lldiv_t struct {
	quot int64
	rem  int64
}
type __locale_t int32
type locale_t __locale_t
type u_char __u_char
type u_short __u_short
type u_int __u_int
type u_long __u_long
type quad_t __quad_t
type u_quad_t __u_quad_t
type fsid_t __fsid_t
type loff_t __loff_t
type ino_t __ino_t
type ino64_t __ino64_t
type dev_t __dev_t
type gid_t __gid_t
type mode_t __mode_t
type nlink_t __nlink_t
type uid_t __uid_t
type pid_t __pid_t
type id_t __id_t
type daddr_t __daddr_t
type caddr_t __caddr_t
type key_t __key_t
type clock_t __clock_t
type clockid_t __clockid_t
type time_t __time_t
type timer_t __timer_t
type useconds_t __useconds_t
type suseconds_t __suseconds_t
type ulong uint32
type ushort uint16
type uint uint32
type int8_t int8
type int16_t int16
type int32_t int32
type int64_t int64
type u_int8_t uint8
type u_int16_t uint16
type u_int32_t uint32
type u_int64_t uint32
type register_t int32

// __uint16_identity - transpiled function from  /usr/include/x86_64-linux-gnu/bits/uintn-identity.h:32
func __uint16_identity(__x uint16) uint16 {
	return uint16(__x)
}

// __uint32_identity - transpiled function from  /usr/include/x86_64-linux-gnu/bits/uintn-identity.h:38
func __uint32_identity(__x uint32) uint32 {
	return uint32(__x)
}

// __uint64_identity - transpiled function from  /usr/include/x86_64-linux-gnu/bits/uintn-identity.h:44
func __uint64_identity(__x uint64) uint64 {
	return uint64(__x)
}

type __sigset_t struct {
	__val [16]uint32
}
type sigset_t __sigset_t
type timeval struct {
	tv_sec  __time_t
	tv_usec __suseconds_t
}
type timespec struct {
	tv_sec  __time_t
	tv_nsec __syscall_slong_t
}
type __fd_mask int32
type fd_set struct {
	fds_bits [16]__fd_mask
}
type fd_mask __fd_mask
type blksize_t __blksize_t
type blkcnt_t __blkcnt_t
type fsblkcnt_t __fsblkcnt_t
type fsfilcnt_t __fsfilcnt_t
type blkcnt64_t __blkcnt64_t
type fsblkcnt64_t __fsblkcnt64_t
type fsfilcnt64_t __fsfilcnt64_t
type __pthread_rwlock_arch_t struct {
	__readers       uint32
	__writers       uint32
	__wrphase_futex uint32
	__writers_futex uint32
	__pad3          uint32
	__pad4          uint32
	__cur_writer    int32
	__shared        int32
	__rwelision     int8
	__pad1          [7]uint8
	__pad2          uint32
	__flags         uint32
}
type __pthread_internal_list struct {
	__prev *__pthread_internal_list
	__next *__pthread_internal_list
}
type __pthread_list_t __pthread_internal_list
type __pthread_mutex_s struct {
	__lock    int32
	__count   uint32
	__owner   int32
	__nusers  uint32
	__kind    int32
	__spins   int16
	__elision int16
	__list    __pthread_list_t
}
type BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD176D5E struct {
	__low  uint32
	__high uint32
}
type __pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD173D17E struct{ memory unsafe.Pointer }

func (unionVar *__pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD173D17E) copy() __pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD173D17E {
	var buffer [16]byte
	for i := range buffer {
		buffer[i] = (*((*[16]byte)(unionVar.memory)))[i]
	}
	var newUnion __pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD173D17E
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *__pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD173D17E) __wseq() *uint64 {
	if unionVar.memory == nil {
		var buffer [16]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*uint64)(unionVar.memory)
}
func (unionVar *__pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD173D17E) __wseq32() *BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD176D5E {
	if unionVar.memory == nil {
		var buffer [16]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD176D5E)(unionVar.memory)
}

type BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD185D5E struct {
	__low  uint32
	__high uint32
}
type __pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD182D17E struct{ memory unsafe.Pointer }

func (unionVar *__pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD182D17E) copy() __pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD182D17E {
	var buffer [16]byte
	for i := range buffer {
		buffer[i] = (*((*[16]byte)(unionVar.memory)))[i]
	}
	var newUnion __pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD182D17E
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *__pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD182D17E) __g1_start() *uint64 {
	if unionVar.memory == nil {
		var buffer [16]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*uint64)(unionVar.memory)
}
func (unionVar *__pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD182D17E) __g1_start32() *BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD185D5E {
	if unionVar.memory == nil {
		var buffer [16]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*BSstructSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD185D5E)(unionVar.memory)
}

type __pthread_cond_s struct {
	__g_refs       [2]uint32
	__g_size       [2]uint32
	__g1_orig_size uint32
	__wrefs        uint32
	__g_signals    [2]uint32
}
type pthread_t uint32
type pthread_mutexattr_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_mutexattr_t) copy() pthread_mutexattr_t {
	var buffer [8]byte
	for i := range buffer {
		buffer[i] = (*((*[8]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_mutexattr_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_mutexattr_t) __size() *[4]byte {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[4]byte)(unionVar.memory)
}
func (unionVar *pthread_mutexattr_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type pthread_condattr_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_condattr_t) copy() pthread_condattr_t {
	var buffer [8]byte
	for i := range buffer {
		buffer[i] = (*((*[8]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_condattr_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_condattr_t) __size() *[4]byte {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[4]byte)(unionVar.memory)
}
func (unionVar *pthread_condattr_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type pthread_key_t uint32
type pthread_once_t int32
type pthread_attr_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_attr_t) copy() pthread_attr_t {
	var buffer [56]byte
	for i := range buffer {
		buffer[i] = (*((*[56]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_attr_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_attr_t) __size() *[56]byte {
	if unionVar.memory == nil {
		var buffer [56]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[56]byte)(unionVar.memory)
}
func (unionVar *pthread_attr_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [56]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type pthread_mutex_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_mutex_t) copy() pthread_mutex_t {
	var buffer [40]byte
	for i := range buffer {
		buffer[i] = (*((*[40]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_mutex_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_mutex_t) __data() *__pthread_mutex_s {
	if unionVar.memory == nil {
		var buffer [40]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*__pthread_mutex_s)(unionVar.memory)
}
func (unionVar *pthread_mutex_t) __size() *[40]byte {
	if unionVar.memory == nil {
		var buffer [40]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[40]byte)(unionVar.memory)
}
func (unionVar *pthread_mutex_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [40]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type pthread_rwlock_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_rwlock_t) copy() pthread_rwlock_t {
	var buffer [56]byte
	for i := range buffer {
		buffer[i] = (*((*[56]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_rwlock_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_rwlock_t) __data() *__pthread_rwlock_arch_t {
	if unionVar.memory == nil {
		var buffer [56]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*__pthread_rwlock_arch_t)(unionVar.memory)
}
func (unionVar *pthread_rwlock_t) __size() *[56]byte {
	if unionVar.memory == nil {
		var buffer [56]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[56]byte)(unionVar.memory)
}
func (unionVar *pthread_rwlock_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [56]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type pthread_rwlockattr_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_rwlockattr_t) copy() pthread_rwlockattr_t {
	var buffer [8]byte
	for i := range buffer {
		buffer[i] = (*((*[8]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_rwlockattr_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_rwlockattr_t) __size() *[8]byte {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[8]byte)(unionVar.memory)
}
func (unionVar *pthread_rwlockattr_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type pthread_spinlock_t int32
type pthread_barrier_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_barrier_t) copy() pthread_barrier_t {
	var buffer [32]byte
	for i := range buffer {
		buffer[i] = (*((*[32]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_barrier_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_barrier_t) __size() *[32]byte {
	if unionVar.memory == nil {
		var buffer [32]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[32]byte)(unionVar.memory)
}
func (unionVar *pthread_barrier_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [32]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type pthread_barrierattr_t struct{ memory unsafe.Pointer }

func (unionVar *pthread_barrierattr_t) copy() pthread_barrierattr_t {
	var buffer [8]byte
	for i := range buffer {
		buffer[i] = (*((*[8]byte)(unionVar.memory)))[i]
	}
	var newUnion pthread_barrierattr_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *pthread_barrierattr_t) __size() *[4]byte {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[4]byte)(unionVar.memory)
}
func (unionVar *pthread_barrierattr_t) __align() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}

type random_data struct {
	fptr      *int32
	rptr      *int32
	state     *int32
	rand_type int32
	rand_deg  int32
	rand_sep  int32
	end_ptr   *int32
}
type drand48_data struct {
	__x     [3]uint16
	__old_x [3]uint16
	__c     uint16
	__init  uint16
	__a     uint64
}
type __compar_fn_t func(unsafe.Pointer, unsafe.Pointer) int32
type comparison_fn_t __compar_fn_t
type __compar_d_fn_t func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int32
type uint8_t uint8
type uint16_t uint16
type uint32_t uint32
type uint64_t uint64
type int_least8_t int8
type int_least16_t int16
type int_least32_t int32
type int_least64_t int32
type uint_least8_t uint8
type uint_least16_t uint16
type uint_least32_t uint32
type uint_least64_t uint32
type int_fast8_t int8
type int_fast16_t int32
type int_fast32_t int32
type int_fast64_t int32
type uint_fast8_t uint8
type uint_fast16_t uint32
type uint_fast32_t uint32
type uint_fast64_t uint32
type intptr_t int32
type uintptr_t uint32
type intmax_t __intmax_t
type uintmax_t __uintmax_t

var // Warning (FieldDecl):  /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:173 : Error : name of FieldDecl is empty
// Warning (IndirectFieldDecl):  /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:175 : could not parse &{94913359793864 {/usr/include/x86_64-linux-gnu/bits/thread-shared-types.h 175 0 42 0 } col:42 true __wseq unsigned long long [0xc0000c3380 0xc0000c3480]}
// Warning (IndirectFieldDecl):  /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:180 : could not parse &{94913359793944 {/usr/include/x86_64-linux-gnu/bits/thread-shared-types.h 180 0 7 0 } col:7 true __wseq32 struct (anonymous struct at /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:176:5)':'struct __pthread_cond_s::(anonymous at /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:176:5) [0xc0000c3540 0xc0000c3600]}
// Warning (FieldDecl):  /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:182 : Error : name of FieldDecl is empty
// Warning (IndirectFieldDecl):  /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:184 : could not parse &{94913359795008 {/usr/include/x86_64-linux-gnu/bits/thread-shared-types.h 184 0 42 0 } col:42 true __g1_start unsigned long long [0xc0000c3840 0xc0000c3900]}
// Warning (IndirectFieldDecl):  /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:189 : could not parse &{94913359795088 {/usr/include/x86_64-linux-gnu/bits/thread-shared-types.h 189 0 7 0 } col:7 true __g1_start32 struct (anonymous struct at /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:185:5)':'struct __pthread_cond_s::(anonymous at /usr/include/x86_64-linux-gnu/bits/thread-shared-types.h:185:5) [0xc0000c3a00 0xc0000c3b00]}
// Warning (RecordDecl):  /usr/include/x86_64-linux-gnu/bits/pthreadtypes.h:75 : could not determine the size of type `union pthread_cond_t` for that reason: Cannot determine sizeof : |union pthread_cond_t|. err = Cannot determine sizeof : |struct __pthread_cond_s|. err = Cannot determine sizeof : |__pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD173D17E|. err = error in array size
// Error (RecordDecl):  /usr/include/x86_64-linux-gnu/bits/pthreadtypes.h:75 : Cannot determine sizeof : |union pthread_cond_t|. err = Cannot determine sizeof : |struct __pthread_cond_s|. err = Cannot determine sizeof : |__pthread_cond_sDDBSatSSusrSincludeSx86_64TlinuxTgnuSbitsSthreadTsharedTtypesPhD173D17E|. err = error in array size
key []uint8 = []uint8{uint8(int32(247)), uint8(int32(155)), uint8(int32(247)), uint8(int32(51)), uint8(int32(242)), uint8(int32(63)), uint8(int32(157)), uint8(int32(122)), uint8(int32(242)), uint8(int32(162)), uint8(int32(145)), uint8(int32(203)), uint8(int32(77)), uint8(int32(203)), uint8(int32(94)), uint8(int32(73)), uint8(int32(99)), uint8(int32(230)), uint8(int32(168)), uint8(int32(113)), uint8(int32(224)), uint8(int32(81)), uint8(int32(44)), uint8(int32(225)), uint8(int32(70)), uint8(int32(189)), uint8(int32(3)), uint8(int32(103)), uint8(int32(147)), uint8(int32(86)), uint8(int32(164)), uint8(int32(115))}
var iv []uint8 = []uint8{uint8(int32(188)), uint8(int32(121)), uint8(int32(204)), uint8(int32(117)), uint8(int32(145)), uint8(int32(180)), uint8(int32(141)), uint8(int32(112))}
var state [][]uint8 = [][]uint8{[]uint8{uint8('e'), uint8('x'), uint8('p'), uint8('a')}, []uint8{uint8('n'), uint8('d'), uint8(' '), uint8('3')}, []uint8{uint8('2'), uint8('-'), uint8('b'), uint8('y')}, []uint8{uint8('t'), uint8('e'), uint8(' '), uint8('k')}}

// crypt - transpiled function from  /root/gbsalsa.c:31
/* Cookie Run-specific constants. */ //
/* Salsa20-specific constants. */ //
/* Implements the Salsa20 cipher. */ //
//
func crypt(buf *uint8, len uint32) {
	var exp []uint8 = make([]uint8, 64, 64)
	var n []uint8 = (&[16]uint8{uint8(int32(0))})[:]
	var i uint32
	var j uint32
	var k uint32
	var x []uint32 = make([]uint32, 16, 16)
	for i = uint32(int32(0)); i < uint32((uint32((uint32(int32(8)))))); i++ {
		*((*uint8)(func() unsafe.Pointer {
			tempVar := &n[0]
			return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint32((uint32((i))))))*unsafe.Sizeof(*tempVar))
		}())) = *((*uint8)(func() unsafe.Pointer {
			tempVar := &iv[0]
			return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint32((uint32((i))))))*unsafe.Sizeof(*tempVar))
		}()))
	}
	for i = uint32(int32(0)); i < len; i++ {
		if i%uint32((uint32((uint32(int32(64)))))) == uint32((uint32((uint32(int32(0)))))) {
			*(*uint32)(unsafe.Pointer(((*uint8)(func() unsafe.Pointer {
				tempVar := &n[0]
				return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(8))*unsafe.Sizeof(*tempVar))
			}())))) = i / uint32((uint32((uint32(int32(64))))))
			for j = uint32(int32(0)); j < uint32((uint32((uint32(int32(64)))))); j += uint32((uint32((uint32(int32(20)))))) {
				for k = uint32(int32(0)); k < uint32((uint32((uint32(int32(4)))))); k++ {
					*((*uint8)(func() unsafe.Pointer {
						tempVar := &exp[0]
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint32((uint32((j + k))))))*unsafe.Sizeof(*tempVar))
					}())) = *((*uint8)(func() unsafe.Pointer {
						tempVar := &state[j/uint32((uint32((uint32(int32(20))))))][0]
						return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint32((uint32((k))))))*unsafe.Sizeof(*tempVar))
					}()))
				}
			}
			for j = uint32(int32(0)); j < uint32((uint32((uint32(int32(16)))))); j++ {
				*((*uint8)(func() unsafe.Pointer {
					tempVar := &exp[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint32(int32(4))+uint32((uint32((j))))))*unsafe.Sizeof(*tempVar))
				}())) = *((*uint8)(func() unsafe.Pointer {
					tempVar := &key[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint32((uint32((j))))))*unsafe.Sizeof(*tempVar))
				}()))
				*((*uint8)(func() unsafe.Pointer {
					tempVar := &exp[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint32(int32(44))+uint32((uint32((j))))))*unsafe.Sizeof(*tempVar))
				}())) = *((*uint8)(func() unsafe.Pointer {
					tempVar := &key[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint32((uint32((j + uint32((uint32((uint32(int32(16))))))))))))*unsafe.Sizeof(*tempVar))
				}()))
				*((*uint8)(func() unsafe.Pointer {
					tempVar := &exp[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint32(int32(24))+uint32((uint32((j))))))*unsafe.Sizeof(*tempVar))
				}())) = *((*uint8)(func() unsafe.Pointer {
					tempVar := &n[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint32((uint32((j))))))*unsafe.Sizeof(*tempVar))
				}()))
			}
			for j = uint32(int32(0)); j < uint32((uint32((uint32(int32(16)))))); j++ {
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint32((uint32((j))))))*unsafe.Sizeof(*tempVar))
				}())) = *((*uint32)(func() unsafe.Pointer {
					tempVar := ((*uint32)(unsafe.Pointer(&exp[0])))
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint32((uint32((j))))))*unsafe.Sizeof(*tempVar))
				}()))
			}
			for j = uint32(int32(0)); j < uint32((uint32((uint32(int32(10)))))); j++ {
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(4))*unsafe.Sizeof(*tempVar))
				}())) ^= (*&x[0]+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(12))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(7)) | (*&x[0]+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(12))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(7))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(8))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(4))*unsafe.Sizeof(*tempVar))
				}()))+*&x[0])<<uint64(int32(9)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(4))*unsafe.Sizeof(*tempVar))
				}()))+*&x[0])>>uint64(int32(32)-int32(9))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(12))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(8))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(4))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(13)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(8))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(4))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(13))
				*&x[0] ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(12))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(8))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(18)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(12))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(8))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(18))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(9))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(5))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(1))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(7)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(5))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(1))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(7))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(13))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(9))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(5))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(9)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(9))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(5))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(9))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(1))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(13))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(9))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(13)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(13))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(9))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(13))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(5))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(1))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(13))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(18)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(1))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(13))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(18))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(14))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(10))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(6))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(7)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(10))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(6))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(7))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(2))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(14))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(10))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(9)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(14))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(10))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(9))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(6))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(2))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(14))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(13)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(2))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(14))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(13))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(10))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(6))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(2))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(18)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(6))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(2))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(18))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(3))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(15))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(11))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(7)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(15))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(11))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(7))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(7))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(3))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(15))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(9)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(3))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(15))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(9))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(11))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(7))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(3))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(13)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(7))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(3))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(13))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(15))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(11))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(7))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(18)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(11))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(7))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(18))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(1))*unsafe.Sizeof(*tempVar))
				}())) ^= (*&x[0]+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(3))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(7)) | (*&x[0]+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(3))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(7))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(2))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(1))*unsafe.Sizeof(*tempVar))
				}()))+*&x[0])<<uint64(int32(9)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(1))*unsafe.Sizeof(*tempVar))
				}()))+*&x[0])>>uint64(int32(32)-int32(9))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(3))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(2))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(1))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(13)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(2))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(1))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(13))
				*&x[0] ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(3))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(2))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(18)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(3))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(2))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(18))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(6))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(5))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(4))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(7)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(5))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(4))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(7))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(7))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(6))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(5))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(9)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(6))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(5))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(9))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(4))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(7))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(6))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(13)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(7))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(6))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(13))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(5))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(4))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(7))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(18)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(4))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(7))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(18))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(11))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(10))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(9))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(7)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(10))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(9))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(7))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(8))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(11))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(10))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(9)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(11))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(10))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(9))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(9))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(8))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(11))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(13)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(8))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(11))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(13))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(10))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(9))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(8))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(18)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(9))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(8))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(18))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(12))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(15))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(14))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(7)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(15))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(14))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(7))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(13))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(12))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(15))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(9)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(12))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(15))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(9))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(14))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(13))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(12))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(13)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(13))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(12))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(13))
				*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(15))*unsafe.Sizeof(*tempVar))
				}())) ^= (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(14))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(13))*unsafe.Sizeof(*tempVar))
				}())))<<uint64(int32(18)) | (*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(14))*unsafe.Sizeof(*tempVar))
				}()))+*((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(13))*unsafe.Sizeof(*tempVar))
				}())))>>uint64(int32(32)-int32(18))
			}
			for j = uint32(int32(0)); j < uint32((uint32((uint32(int32(16)))))); j++ {
				*((*uint32)(func() unsafe.Pointer {
					tempVar := ((*uint32)(unsafe.Pointer(&exp[0])))
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint32((uint32((j))))))*unsafe.Sizeof(*tempVar))
				}())) += *((*uint32)(func() unsafe.Pointer {
					tempVar := &x[0]
					return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint32((uint32((j))))))*unsafe.Sizeof(*tempVar))
				}()))
			}
		}
		*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + (uintptr)(int32(uint32((uint32((i))))))*unsafe.Sizeof(*buf)))) ^= uint8((uint8((uint8(uint8((uint8((uint8(int32(uint8((uint8((*((*uint8)(func() unsafe.Pointer {
			tempVar := &exp[0]
			return unsafe.Pointer(uintptr(unsafe.Pointer(tempVar)) + (uintptr)(int32(uint32((uint32((i % uint32((uint32((uint32(int32(64))))))))))))*unsafe.Sizeof(*tempVar))
		}())))))))))))))))))
	}
}

// main - transpiled function from  /root/gbsalsa.c:106
func CryptFile(filename string) error {
	filenameAsArgs := []string{"gbsalsa", filename}
	argv__multiarray := [][]byte{}
	argv__array := []*byte{}
	for _, argvSingle := range filenameAsArgs {
		argv__multiarray = append(argv__multiarray, append([]byte(argvSingle), 0))
	}
	for _, argvSingle := range argv__multiarray {
		argv__array = append(argv__array, &argvSingle[0])
	}
	argv := *(***byte)(unsafe.Pointer(&argv__array))
	var f *noarch.File
	var len uint32
	var buf *byte
	if (func() *noarch.File {
		tempVar := noarch.Fopen(*((**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + (uintptr)(int32(1))*unsafe.Sizeof(*argv)))), (&[]byte("rb+\x00")[0]))
		f = tempVar
		return tempVar
	}()) == nil {
		return errors.New("Error")
	}
	noarch.Fseek(f, int32(0), int32(2))
	len = uint32((uint32((uint32(noarch.Ftell(f))))))
	noarch.Rewind(f)
	buf = (*byte)(noarch.Malloc(int32(uint32((uint32((uint32(len))))))))
	noarch.Fread(unsafe.Pointer(buf), int32(int32(1)), int32(uint32(len)), f)
	crypt((*uint8)(unsafe.Pointer(buf)), uint32(len))
	noarch.Rewind(f)
	noarch.Fwrite(buf, int32(int32(1)), int32(uint32(len)), f)
	noarch.Fclose(f)
	noarch.Free(unsafe.Pointer(buf))
	return nil
}
func init() {
	stdin = noarch.Stdin
	stdout = noarch.Stdout
	stderr = noarch.Stderr
}
