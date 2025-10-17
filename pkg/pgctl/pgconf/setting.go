package pgconf

import (
	"cmp"
	"fmt"
	"strconv"

	"github.com/innoai-tech/postgres-operator/pkg/units"
)

type Setting struct {
	// CPU db cpu requests
	CPU int `flag:",omitzero"`
	// MEM db mem requests
	MEM units.BinarySize `flag:",omitzero"`
	// MaxConnections db max connections
	MaxConnections int `flag:",omitzero"`
	// ApplicationType db which application
	ApplicationType ApplicationType `flag:",omitzero"`
	// DiskType disk type
	DiskType DiskType `flag:",omitzero"`
}

func (s *Setting) SetDefaults() {
	if s.CPU == 0 {
		s.CPU = 2
	}

	if s.MEM == 0 {
		s.MEM = 4 * units.GiB
	}

	if s.ApplicationType.IsZero() {
		s.ApplicationType = APPLICATION_TYPE__MIXED
	}

	if s.DiskType.IsZero() {
		s.DiskType = DISK_TYPE__SSD
	}
}

func (s *Setting) ToPgConf(pgVersion string) map[string]string {
	// https://github.com/le0pard/pgtune/blob/master/src/features/configuration/configurationSlice.js

	pgVer, _ := strconv.Atoi(cmp.Or(pgVersion, "16"))

	settings := map[string]string{
		"max_connections":              fromInt(s.maxConnections()),
		"shared_buffers":               fromBinarySize(s.sharedBuffers()),
		"effective_cache_size":         fromBinarySize(s.effectiveCacheSize()),
		"maintenance_work_mem":         fromBinarySize(s.maintenanceWorkMem()),
		"checkpoint_completion_target": fromFloat(s.checkpointCompletionTarget()),
		"wal_buffers":                  fromBinarySize(s.walBuffers()),
		"default_statistics_target":    fromInt(s.defaultStatisticsTarget()),
		"random_page_cost":             fromFloat(s.randomPageCost()),
		"effective_io_concurrency":     fromInt(s.effectiveIoConcurrency()),
		"huge_pages":                   s.hugePages(),
		"min_wal_size":                 fromBinarySize(s.minWalSize()),
		"max_wal_size":                 fromBinarySize(s.maxWalSize()),
	}

	if s.CPU >= 4 {
		settings["max_worker_processes"] = fromInt(s.maxWorkerProcesses())

		workersPerGather := s.CPU / 2
		if s.ApplicationType != APPLICATION_TYPE__DATA_WAREHOUSE && workersPerGather > 4 {
			workersPerGather = 4 // no clear evidence, that each new worker will provide big benefit for each noew core
		}
		settings["max_parallel_workers_per_gather"] = fromInt(workersPerGather)

		if pgVer >= 10 {
			settings["max_parallel_workers"] = fromInt(s.CPU)

			if pgVer > 11 {
				parallelMaintenanceWorkers := s.CPU / 2
				if parallelMaintenanceWorkers > 4 {
					parallelMaintenanceWorkers = 4 // no clear evidence, that each new worker will provide big benefit for each noew core
				}
				settings["max_parallel_maintenance_workers"] = fromInt(parallelMaintenanceWorkers)
			}
		}
	}

	settings["work_mem"] = fromBinarySize(s.workMem())

	s.patchWithWalLevel(settings)

	return settings
}

func fromInt(i int) string {
	return fmt.Sprintf("%d", i)
}

func fromFloat(i float64) string {
	return fmt.Sprintf("%v", i)
}

func fromBinarySize(bs units.BinarySize) string {
	return fmt.Sprintf("%dMB", int(bs/units.MiB))
}

// https://github.com/le0pard/pgtune/blob/master/src/features/configuration/configurationSlice.js#L107
func (s *Setting) maxConnections() int {
	if s.MaxConnections != 0 {
		return s.MaxConnections
	}

	switch s.ApplicationType {
	case APPLICATION_TYPE__DATA_WAREHOUSE:
		return s.CPU * 10
	case APPLICATION_TYPE__DESKTOP:
		return s.CPU / 2 * 10
	default:
		return s.CPU * 100
	}
}

// https://github.com/le0pard/pgtune/blob/master/src/features/configuration/configurationSlice.js#L127
func (s *Setting) sharedBuffers() units.BinarySize {
	if s.ApplicationType == APPLICATION_TYPE__DESKTOP {
		return s.MEM / 16
	}
	return s.MEM / 4
}

// https://github.com/le0pard/pgtune/blob/master/src/features/configuration/configurationSlice.js#L148
func (s *Setting) effectiveCacheSize() units.BinarySize {
	return s.MEM * 3 / 4
}

// https://github.com/le0pard/pgtune/blob/master/src/features/configuration/configurationSlice.js#L160C14-L160C38
func (s *Setting) maintenanceWorkMem() units.BinarySize {
	maintenanceWorkMem := s.MEM * 16
	if s.ApplicationType == APPLICATION_TYPE__DATA_WAREHOUSE {
		maintenanceWorkMem = s.MEM / 8
	}
	// Cap maintenance RAM at 2GB on servers with lots of memory
	if memoryLimit := 2 * units.GiB; maintenanceWorkMem > memoryLimit {
		maintenanceWorkMem = memoryLimit
	}
	return maintenanceWorkMem
}

// https://github.com/le0pard/pgtune/blob/master/src/features/configuration/configurationSlice.js#L160C14-L160C38
func (s *Setting) minWalSize() units.BinarySize {
	x := 1024
	switch s.ApplicationType {
	case APPLICATION_TYPE__DESKTOP:
		x = x / 10
	case APPLICATION_TYPE__OLTP:
		x = x * 2
	case APPLICATION_TYPE__DATA_WAREHOUSE:
		x = x * 4
	default:
	}
	return units.BinarySize(x) * units.MiB
}

func (s *Setting) maxWalSize() units.BinarySize {
	x := 4096
	switch s.ApplicationType {
	case APPLICATION_TYPE__DESKTOP:
		x = x / 2
	case APPLICATION_TYPE__OLTP:
		x = x * 2
	case APPLICATION_TYPE__DATA_WAREHOUSE:
		x = x * 4
	default:
	}
	return units.BinarySize(x) * units.MiB
}

// https://github.com/le0pard/pgtune/blob/master/src/features/configuration/configurationSlice.js#L209
func (s *Setting) checkpointCompletionTarget() float64 {
	return 0.9 // based on https://github.com/postgres/postgres/commit/bbcc4eb2
}

// https://github.com/le0pard/pgtune/blob/master/src/features/configuration/configurationSlice.js#L214
func (s *Setting) walBuffers() units.BinarySize {
	walBuffers := s.sharedBuffers() * 3 / 100

	// Follow auto-tuning guideline for wal_buffers added in 9.1, where it's
	// set to 3% of shared_buffers up to a maximum of 16MB.
	maxWalBuffer := 16 * units.MiB
	if walBuffers > maxWalBuffer {
		walBuffers = maxWalBuffer
	}

	// It's nice of wal_buffers is an even 16MB if it's near that number. Since
	// that is a common case on Windows, where shared_buffers is clipped to 512MB,
	// round upwards in that situation
	if walBufferNearValue := 14 * units.MiB; walBufferNearValue > walBuffers && walBuffers < maxWalBuffer {
		walBuffers = maxWalBuffer
	}

	if minWalBuffers := 32 * units.KiB; walBuffers < minWalBuffers {
		walBuffers = minWalBuffers
	}

	return walBuffers
}

// https://github.com/le0pard/pgtune/blob/master/src/features/configuration/configurationSlice.js#L236
func (s *Setting) defaultStatisticsTarget() int {
	if s.ApplicationType == APPLICATION_TYPE__DATA_WAREHOUSE {
		return 500
	}
	return 100
}

// https://github.com/le0pard/pgtune/blob/master/src/features/configuration/configurationSlice.js#L248C14-L248C34
func (s *Setting) randomPageCost() float64 {
	switch s.DiskType {
	case DISK_TYPE__HDD:
		return 4
	default:
		return 1.1
	}
}

// https://github.com/le0pard/pgtune/blob/master/src/features/configuration/configurationSlice.js#L256
func (s *Setting) effectiveIoConcurrency() int {
	switch s.DiskType {
	case DISK_TYPE__HDD:
		return 2
	case DISK_TYPE__SSD:
		return 200
	case DISK_TYPE__SAN:
		return 300
	default:
		return 200
	}
}

func (v *Setting) maxWorkerProcesses() int {
	if v.CPU > 0 {
		return v.CPU
	}
	return 1
}

// https://github.com/le0pard/pgtune/blob/master/src/features/configuration/configurationSlice.js#L318
func (s *Setting) workMem() units.BinarySize {
	workMem := (s.MEM - s.sharedBuffers()) / units.BinarySize((s.maxConnections()+s.maxWorkerProcesses())*3)

	switch s.ApplicationType {
	case APPLICATION_TYPE__DATA_WAREHOUSE:
		workMem = workMem / 2
	case APPLICATION_TYPE__DESKTOP:
		workMem = workMem / 6
	default:
	}
	if minWorkMem := 64 * units.KiB; workMem < minWorkMem {
		workMem = minWorkMem
	}
	return workMem
}

// https://github.com/le0pard/pgtune/blob/master/src/features/configuration/configurationSlice.js#L121
func (s *Setting) hugePages() string {
	if s.MEM > 32*units.GiB {
		return "try"
	}
	return "off"
}

// https://github.com/le0pard/pgtune/blob/master/src/features/configuration/configurationSlice.js#L381
func (s *Setting) patchWithWalLevel(settings map[string]string) {
	if s.ApplicationType == APPLICATION_TYPE__DESKTOP {
		settings["wal_level"] = "minimal"
		settings["max_wal_senders"] = "0"
	}
}

// +gengo:enum
type ApplicationType uint8

const (
	APPLICATION_TYPE_UNKNOWN ApplicationType = iota
	APPLICATION_TYPE__WEB
	APPLICATION_TYPE__OLTP
	APPLICATION_TYPE__DATA_WAREHOUSE
	APPLICATION_TYPE__DESKTOP
	APPLICATION_TYPE__MIXED
)

// +gengo:enum
type DiskType uint8

const (
	DISK_TYPE_UNKNOWN DiskType = iota
	DISK_TYPE__SSD
	DISK_TYPE__HDD
	DISK_TYPE__SAN
)
