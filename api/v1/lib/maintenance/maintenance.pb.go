// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: maintenance/maintenance.proto

/*
	Package maintenance is a generated protocol buffer package.

	It is generated from these files:
		maintenance/maintenance.proto

	It has these top-level messages:
		Window
		Schedule
		ClusterStatus
*/
package maintenance

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import mesos "github.com/mesos/mesos-go/api/v1/lib"
import mesos_allocator "github.com/mesos/mesos-go/api/v1/lib/allocator"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// *
// A set of machines scheduled to go into maintenance
// in the same `unavailability`.
type Window struct {
	// Machines affected by this maintenance window.
	MachineIDs []mesos.MachineID `protobuf:"bytes,1,rep,name=machine_ids,json=machineIds" json:"machine_ids"`
	// Interval during which this set of machines is expected to be down.
	Unavailability mesos.Unavailability `protobuf:"bytes,2,req,name=unavailability" json:"unavailability"`
}

func (m *Window) Reset()                    { *m = Window{} }
func (*Window) ProtoMessage()               {}
func (*Window) Descriptor() ([]byte, []int) { return fileDescriptorMaintenance, []int{0} }

func (m *Window) GetMachineIDs() []mesos.MachineID {
	if m != nil {
		return m.MachineIDs
	}
	return nil
}

func (m *Window) GetUnavailability() mesos.Unavailability {
	if m != nil {
		return m.Unavailability
	}
	return mesos.Unavailability{}
}

// *
// A list of maintenance windows.
// For example, this may represent a rolling restart of agents.
type Schedule struct {
	Windows []Window `protobuf:"bytes,1,rep,name=windows" json:"windows"`
}

func (m *Schedule) Reset()                    { *m = Schedule{} }
func (*Schedule) ProtoMessage()               {}
func (*Schedule) Descriptor() ([]byte, []int) { return fileDescriptorMaintenance, []int{1} }

func (m *Schedule) GetWindows() []Window {
	if m != nil {
		return m.Windows
	}
	return nil
}

// *
// Represents the maintenance status of each machine in the cluster.
// The lists correspond to the `MachineInfo.Mode` enumeration.
type ClusterStatus struct {
	DrainingMachines []ClusterStatus_DrainingMachine `protobuf:"bytes,1,rep,name=draining_machines,json=drainingMachines" json:"draining_machines"`
	DownMachines     []mesos.MachineID               `protobuf:"bytes,2,rep,name=down_machines,json=downMachines" json:"down_machines"`
}

func (m *ClusterStatus) Reset()                    { *m = ClusterStatus{} }
func (*ClusterStatus) ProtoMessage()               {}
func (*ClusterStatus) Descriptor() ([]byte, []int) { return fileDescriptorMaintenance, []int{2} }

func (m *ClusterStatus) GetDrainingMachines() []ClusterStatus_DrainingMachine {
	if m != nil {
		return m.DrainingMachines
	}
	return nil
}

func (m *ClusterStatus) GetDownMachines() []mesos.MachineID {
	if m != nil {
		return m.DownMachines
	}
	return nil
}

type ClusterStatus_DrainingMachine struct {
	ID mesos.MachineID `protobuf:"bytes,1,req,name=id" json:"id"`
	// A list of the most recent responses to inverse offers from frameworks
	// running on this draining machine.
	Statuses []mesos_allocator.InverseOfferStatus `protobuf:"bytes,2,rep,name=statuses" json:"statuses"`
}

func (m *ClusterStatus_DrainingMachine) Reset()      { *m = ClusterStatus_DrainingMachine{} }
func (*ClusterStatus_DrainingMachine) ProtoMessage() {}
func (*ClusterStatus_DrainingMachine) Descriptor() ([]byte, []int) {
	return fileDescriptorMaintenance, []int{2, 0}
}

func (m *ClusterStatus_DrainingMachine) GetID() mesos.MachineID {
	if m != nil {
		return m.ID
	}
	return mesos.MachineID{}
}

func (m *ClusterStatus_DrainingMachine) GetStatuses() []mesos_allocator.InverseOfferStatus {
	if m != nil {
		return m.Statuses
	}
	return nil
}

func init() {
	proto.RegisterType((*Window)(nil), "mesos.maintenance.Window")
	proto.RegisterType((*Schedule)(nil), "mesos.maintenance.Schedule")
	proto.RegisterType((*ClusterStatus)(nil), "mesos.maintenance.ClusterStatus")
	proto.RegisterType((*ClusterStatus_DrainingMachine)(nil), "mesos.maintenance.ClusterStatus.DrainingMachine")
}
func (this *Window) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Window)
	if !ok {
		that2, ok := that.(Window)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Window")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Window but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Window but is not nil && this == nil")
	}
	if len(this.MachineIDs) != len(that1.MachineIDs) {
		return fmt.Errorf("MachineIDs this(%v) Not Equal that(%v)", len(this.MachineIDs), len(that1.MachineIDs))
	}
	for i := range this.MachineIDs {
		if !this.MachineIDs[i].Equal(&that1.MachineIDs[i]) {
			return fmt.Errorf("MachineIDs this[%v](%v) Not Equal that[%v](%v)", i, this.MachineIDs[i], i, that1.MachineIDs[i])
		}
	}
	if !this.Unavailability.Equal(&that1.Unavailability) {
		return fmt.Errorf("Unavailability this(%v) Not Equal that(%v)", this.Unavailability, that1.Unavailability)
	}
	return nil
}
func (this *Window) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Window)
	if !ok {
		that2, ok := that.(Window)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.MachineIDs) != len(that1.MachineIDs) {
		return false
	}
	for i := range this.MachineIDs {
		if !this.MachineIDs[i].Equal(&that1.MachineIDs[i]) {
			return false
		}
	}
	if !this.Unavailability.Equal(&that1.Unavailability) {
		return false
	}
	return true
}
func (this *Schedule) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Schedule)
	if !ok {
		that2, ok := that.(Schedule)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Schedule")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Schedule but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Schedule but is not nil && this == nil")
	}
	if len(this.Windows) != len(that1.Windows) {
		return fmt.Errorf("Windows this(%v) Not Equal that(%v)", len(this.Windows), len(that1.Windows))
	}
	for i := range this.Windows {
		if !this.Windows[i].Equal(&that1.Windows[i]) {
			return fmt.Errorf("Windows this[%v](%v) Not Equal that[%v](%v)", i, this.Windows[i], i, that1.Windows[i])
		}
	}
	return nil
}
func (this *Schedule) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Schedule)
	if !ok {
		that2, ok := that.(Schedule)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Windows) != len(that1.Windows) {
		return false
	}
	for i := range this.Windows {
		if !this.Windows[i].Equal(&that1.Windows[i]) {
			return false
		}
	}
	return true
}
func (this *ClusterStatus) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ClusterStatus)
	if !ok {
		that2, ok := that.(ClusterStatus)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ClusterStatus")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ClusterStatus but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ClusterStatus but is not nil && this == nil")
	}
	if len(this.DrainingMachines) != len(that1.DrainingMachines) {
		return fmt.Errorf("DrainingMachines this(%v) Not Equal that(%v)", len(this.DrainingMachines), len(that1.DrainingMachines))
	}
	for i := range this.DrainingMachines {
		if !this.DrainingMachines[i].Equal(&that1.DrainingMachines[i]) {
			return fmt.Errorf("DrainingMachines this[%v](%v) Not Equal that[%v](%v)", i, this.DrainingMachines[i], i, that1.DrainingMachines[i])
		}
	}
	if len(this.DownMachines) != len(that1.DownMachines) {
		return fmt.Errorf("DownMachines this(%v) Not Equal that(%v)", len(this.DownMachines), len(that1.DownMachines))
	}
	for i := range this.DownMachines {
		if !this.DownMachines[i].Equal(&that1.DownMachines[i]) {
			return fmt.Errorf("DownMachines this[%v](%v) Not Equal that[%v](%v)", i, this.DownMachines[i], i, that1.DownMachines[i])
		}
	}
	return nil
}
func (this *ClusterStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClusterStatus)
	if !ok {
		that2, ok := that.(ClusterStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.DrainingMachines) != len(that1.DrainingMachines) {
		return false
	}
	for i := range this.DrainingMachines {
		if !this.DrainingMachines[i].Equal(&that1.DrainingMachines[i]) {
			return false
		}
	}
	if len(this.DownMachines) != len(that1.DownMachines) {
		return false
	}
	for i := range this.DownMachines {
		if !this.DownMachines[i].Equal(&that1.DownMachines[i]) {
			return false
		}
	}
	return true
}
func (this *ClusterStatus_DrainingMachine) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ClusterStatus_DrainingMachine)
	if !ok {
		that2, ok := that.(ClusterStatus_DrainingMachine)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ClusterStatus_DrainingMachine")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ClusterStatus_DrainingMachine but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ClusterStatus_DrainingMachine but is not nil && this == nil")
	}
	if !this.ID.Equal(&that1.ID) {
		return fmt.Errorf("ID this(%v) Not Equal that(%v)", this.ID, that1.ID)
	}
	if len(this.Statuses) != len(that1.Statuses) {
		return fmt.Errorf("Statuses this(%v) Not Equal that(%v)", len(this.Statuses), len(that1.Statuses))
	}
	for i := range this.Statuses {
		if !this.Statuses[i].Equal(&that1.Statuses[i]) {
			return fmt.Errorf("Statuses this[%v](%v) Not Equal that[%v](%v)", i, this.Statuses[i], i, that1.Statuses[i])
		}
	}
	return nil
}
func (this *ClusterStatus_DrainingMachine) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClusterStatus_DrainingMachine)
	if !ok {
		that2, ok := that.(ClusterStatus_DrainingMachine)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ID.Equal(&that1.ID) {
		return false
	}
	if len(this.Statuses) != len(that1.Statuses) {
		return false
	}
	for i := range this.Statuses {
		if !this.Statuses[i].Equal(&that1.Statuses[i]) {
			return false
		}
	}
	return true
}
func (this *Window) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&maintenance.Window{")
	if this.MachineIDs != nil {
		vs := make([]*mesos.MachineID, len(this.MachineIDs))
		for i := range vs {
			vs[i] = &this.MachineIDs[i]
		}
		s = append(s, "MachineIDs: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "Unavailability: "+strings.Replace(this.Unavailability.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Schedule) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&maintenance.Schedule{")
	if this.Windows != nil {
		vs := make([]*Window, len(this.Windows))
		for i := range vs {
			vs[i] = &this.Windows[i]
		}
		s = append(s, "Windows: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ClusterStatus) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&maintenance.ClusterStatus{")
	if this.DrainingMachines != nil {
		vs := make([]*ClusterStatus_DrainingMachine, len(this.DrainingMachines))
		for i := range vs {
			vs[i] = &this.DrainingMachines[i]
		}
		s = append(s, "DrainingMachines: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	if this.DownMachines != nil {
		vs := make([]*mesos.MachineID, len(this.DownMachines))
		for i := range vs {
			vs[i] = &this.DownMachines[i]
		}
		s = append(s, "DownMachines: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ClusterStatus_DrainingMachine) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&maintenance.ClusterStatus_DrainingMachine{")
	s = append(s, "ID: "+strings.Replace(this.ID.GoString(), `&`, ``, 1)+",\n")
	if this.Statuses != nil {
		vs := make([]*mesos_allocator.InverseOfferStatus, len(this.Statuses))
		for i := range vs {
			vs[i] = &this.Statuses[i]
		}
		s = append(s, "Statuses: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMaintenance(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Window) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Window) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.MachineIDs) > 0 {
		for _, msg := range m.MachineIDs {
			dAtA[i] = 0xa
			i++
			i = encodeVarintMaintenance(dAtA, i, uint64(msg.ProtoSize()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintMaintenance(dAtA, i, uint64(m.Unavailability.ProtoSize()))
	n1, err := m.Unavailability.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func (m *Schedule) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Schedule) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Windows) > 0 {
		for _, msg := range m.Windows {
			dAtA[i] = 0xa
			i++
			i = encodeVarintMaintenance(dAtA, i, uint64(msg.ProtoSize()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ClusterStatus) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.DrainingMachines) > 0 {
		for _, msg := range m.DrainingMachines {
			dAtA[i] = 0xa
			i++
			i = encodeVarintMaintenance(dAtA, i, uint64(msg.ProtoSize()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.DownMachines) > 0 {
		for _, msg := range m.DownMachines {
			dAtA[i] = 0x12
			i++
			i = encodeVarintMaintenance(dAtA, i, uint64(msg.ProtoSize()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ClusterStatus_DrainingMachine) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterStatus_DrainingMachine) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintMaintenance(dAtA, i, uint64(m.ID.ProtoSize()))
	n2, err := m.ID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.Statuses) > 0 {
		for _, msg := range m.Statuses {
			dAtA[i] = 0x12
			i++
			i = encodeVarintMaintenance(dAtA, i, uint64(msg.ProtoSize()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintMaintenance(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedWindow(r randyMaintenance, easy bool) *Window {
	this := &Window{}
	if r.Intn(10) != 0 {
		v1 := r.Intn(5)
		this.MachineIDs = make([]mesos.MachineID, v1)
		for i := 0; i < v1; i++ {
			v2 := mesos.NewPopulatedMachineID(r, easy)
			this.MachineIDs[i] = *v2
		}
	}
	v3 := mesos.NewPopulatedUnavailability(r, easy)
	this.Unavailability = *v3
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedSchedule(r randyMaintenance, easy bool) *Schedule {
	this := &Schedule{}
	if r.Intn(10) != 0 {
		v4 := r.Intn(5)
		this.Windows = make([]Window, v4)
		for i := 0; i < v4; i++ {
			v5 := NewPopulatedWindow(r, easy)
			this.Windows[i] = *v5
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedClusterStatus(r randyMaintenance, easy bool) *ClusterStatus {
	this := &ClusterStatus{}
	if r.Intn(10) != 0 {
		v6 := r.Intn(5)
		this.DrainingMachines = make([]ClusterStatus_DrainingMachine, v6)
		for i := 0; i < v6; i++ {
			v7 := NewPopulatedClusterStatus_DrainingMachine(r, easy)
			this.DrainingMachines[i] = *v7
		}
	}
	if r.Intn(10) != 0 {
		v8 := r.Intn(5)
		this.DownMachines = make([]mesos.MachineID, v8)
		for i := 0; i < v8; i++ {
			v9 := mesos.NewPopulatedMachineID(r, easy)
			this.DownMachines[i] = *v9
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedClusterStatus_DrainingMachine(r randyMaintenance, easy bool) *ClusterStatus_DrainingMachine {
	this := &ClusterStatus_DrainingMachine{}
	v10 := mesos.NewPopulatedMachineID(r, easy)
	this.ID = *v10
	if r.Intn(10) != 0 {
		v11 := r.Intn(5)
		this.Statuses = make([]mesos_allocator.InverseOfferStatus, v11)
		for i := 0; i < v11; i++ {
			v12 := mesos_allocator.NewPopulatedInverseOfferStatus(r, easy)
			this.Statuses[i] = *v12
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyMaintenance interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMaintenance(r randyMaintenance) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringMaintenance(r randyMaintenance) string {
	v13 := r.Intn(100)
	tmps := make([]rune, v13)
	for i := 0; i < v13; i++ {
		tmps[i] = randUTF8RuneMaintenance(r)
	}
	return string(tmps)
}
func randUnrecognizedMaintenance(r randyMaintenance, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldMaintenance(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldMaintenance(dAtA []byte, r randyMaintenance, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateMaintenance(dAtA, uint64(key))
		v14 := r.Int63()
		if r.Intn(2) == 0 {
			v14 *= -1
		}
		dAtA = encodeVarintPopulateMaintenance(dAtA, uint64(v14))
	case 1:
		dAtA = encodeVarintPopulateMaintenance(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateMaintenance(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateMaintenance(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateMaintenance(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateMaintenance(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Window) ProtoSize() (n int) {
	var l int
	_ = l
	if len(m.MachineIDs) > 0 {
		for _, e := range m.MachineIDs {
			l = e.ProtoSize()
			n += 1 + l + sovMaintenance(uint64(l))
		}
	}
	l = m.Unavailability.ProtoSize()
	n += 1 + l + sovMaintenance(uint64(l))
	return n
}

func (m *Schedule) ProtoSize() (n int) {
	var l int
	_ = l
	if len(m.Windows) > 0 {
		for _, e := range m.Windows {
			l = e.ProtoSize()
			n += 1 + l + sovMaintenance(uint64(l))
		}
	}
	return n
}

func (m *ClusterStatus) ProtoSize() (n int) {
	var l int
	_ = l
	if len(m.DrainingMachines) > 0 {
		for _, e := range m.DrainingMachines {
			l = e.ProtoSize()
			n += 1 + l + sovMaintenance(uint64(l))
		}
	}
	if len(m.DownMachines) > 0 {
		for _, e := range m.DownMachines {
			l = e.ProtoSize()
			n += 1 + l + sovMaintenance(uint64(l))
		}
	}
	return n
}

func (m *ClusterStatus_DrainingMachine) ProtoSize() (n int) {
	var l int
	_ = l
	l = m.ID.ProtoSize()
	n += 1 + l + sovMaintenance(uint64(l))
	if len(m.Statuses) > 0 {
		for _, e := range m.Statuses {
			l = e.ProtoSize()
			n += 1 + l + sovMaintenance(uint64(l))
		}
	}
	return n
}

func sovMaintenance(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMaintenance(x uint64) (n int) {
	return sovMaintenance(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Window) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Window{`,
		`MachineIDs:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.MachineIDs), "MachineID", "mesos.MachineID", 1), `&`, ``, 1) + `,`,
		`Unavailability:` + strings.Replace(strings.Replace(this.Unavailability.String(), "Unavailability", "mesos.Unavailability", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Schedule) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Schedule{`,
		`Windows:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Windows), "Window", "Window", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ClusterStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ClusterStatus{`,
		`DrainingMachines:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.DrainingMachines), "ClusterStatus_DrainingMachine", "ClusterStatus_DrainingMachine", 1), `&`, ``, 1) + `,`,
		`DownMachines:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.DownMachines), "MachineID", "mesos.MachineID", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ClusterStatus_DrainingMachine) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ClusterStatus_DrainingMachine{`,
		`ID:` + strings.Replace(strings.Replace(this.ID.String(), "MachineID", "mesos.MachineID", 1), `&`, ``, 1) + `,`,
		`Statuses:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Statuses), "InverseOfferStatus", "mesos_allocator.InverseOfferStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMaintenance(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Window) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMaintenance
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Window: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Window: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MachineIDs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaintenance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMaintenance
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MachineIDs = append(m.MachineIDs, mesos.MachineID{})
			if err := m.MachineIDs[len(m.MachineIDs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unavailability", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaintenance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMaintenance
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Unavailability.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipMaintenance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMaintenance
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return proto.NewRequiredNotSetError("unavailability")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Schedule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMaintenance
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Schedule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Schedule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Windows", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaintenance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMaintenance
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Windows = append(m.Windows, Window{})
			if err := m.Windows[len(m.Windows)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMaintenance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMaintenance
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClusterStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMaintenance
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClusterStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DrainingMachines", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaintenance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMaintenance
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DrainingMachines = append(m.DrainingMachines, ClusterStatus_DrainingMachine{})
			if err := m.DrainingMachines[len(m.DrainingMachines)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DownMachines", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaintenance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMaintenance
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DownMachines = append(m.DownMachines, mesos.MachineID{})
			if err := m.DownMachines[len(m.DownMachines)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMaintenance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMaintenance
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClusterStatus_DrainingMachine) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMaintenance
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DrainingMachine: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DrainingMachine: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaintenance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMaintenance
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Statuses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMaintenance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMaintenance
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Statuses = append(m.Statuses, mesos_allocator.InverseOfferStatus{})
			if err := m.Statuses[len(m.Statuses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMaintenance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMaintenance
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return proto.NewRequiredNotSetError("id")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMaintenance(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMaintenance
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMaintenance
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMaintenance
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMaintenance
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMaintenance
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMaintenance(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMaintenance = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMaintenance   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("maintenance/maintenance.proto", fileDescriptorMaintenance) }

var fileDescriptorMaintenance = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x31, 0x6f, 0x13, 0x31,
	0x14, 0xc7, 0xcf, 0x07, 0x82, 0xca, 0xa1, 0xd0, 0x5a, 0x42, 0x0a, 0x91, 0x70, 0xab, 0xb0, 0x74,
	0xa0, 0x77, 0xa5, 0x1b, 0x42, 0x62, 0x48, 0x93, 0xe1, 0x06, 0x84, 0x94, 0x0a, 0x21, 0xb1, 0x54,
	0xce, 0xd9, 0xb9, 0x58, 0xba, 0xd8, 0xd5, 0xd9, 0x97, 0x88, 0x8d, 0x89, 0x99, 0x81, 0x0f, 0xc0,
	0xc8, 0x47, 0x40, 0x4c, 0x8c, 0x1d, 0x3b, 0x32, 0x55, 0x9c, 0xbb, 0x30, 0x76, 0x64, 0x44, 0xb5,
	0x7d, 0xc9, 0xa5, 0x05, 0xa9, 0x4b, 0xf4, 0xfc, 0xde, 0xfb, 0xfd, 0xdf, 0xff, 0xbd, 0x1c, 0x7c,
	0x3c, 0x25, 0x5c, 0x68, 0x26, 0x88, 0x48, 0x59, 0xdc, 0x88, 0xa3, 0xe3, 0x42, 0x6a, 0x89, 0x36,
	0xa7, 0x4c, 0x49, 0x15, 0x35, 0x0a, 0x9d, 0xbd, 0x8c, 0xeb, 0x49, 0x39, 0x8a, 0x52, 0x39, 0x8d,
	0x6d, 0xd5, 0xfd, 0xee, 0x66, 0x32, 0x26, 0xc7, 0x3c, 0x9e, 0x3d, 0x8b, 0x73, 0x3e, 0x72, 0x39,
	0x27, 0xd2, 0x79, 0x79, 0x23, 0x82, 0xe4, 0xb9, 0x4c, 0x89, 0x96, 0xc5, 0x32, 0xf2, 0xfc, 0x6e,
	0x83, 0xcf, 0x64, 0x26, 0x63, 0x9b, 0x1e, 0x95, 0x63, 0xfb, 0xb2, 0x0f, 0x1b, 0xb9, 0xf6, 0xee,
	0x67, 0x00, 0xef, 0xbc, 0xe5, 0x82, 0xca, 0x39, 0x1a, 0xc0, 0xd6, 0x94, 0xa4, 0x13, 0x2e, 0xd8,
	0x11, 0xa7, 0xaa, 0x0d, 0xb6, 0x6f, 0xed, 0xb4, 0xf6, 0x37, 0x22, 0x67, 0xee, 0x95, 0xab, 0x24,
	0xfd, 0x1e, 0x3a, 0x39, 0xdb, 0x0a, 0xcc, 0xd9, 0x16, 0x5c, 0xa4, 0xd4, 0x10, 0x7a, 0x30, 0xa1,
	0x0a, 0x1d, 0xc0, 0xfb, 0xa5, 0x20, 0x33, 0xc2, 0x73, 0x32, 0xe2, 0x39, 0xd7, 0xef, 0xdb, 0xe1,
	0x76, 0xb8, 0xd3, 0xda, 0x7f, 0xe8, 0x95, 0xde, 0xac, 0x14, 0x7b, 0xb7, 0x2f, 0xe5, 0x86, 0x57,
	0x90, 0xee, 0x00, 0xae, 0x1d, 0xa6, 0x13, 0x46, 0xcb, 0x9c, 0xa1, 0xe7, 0xf0, 0xee, 0xdc, 0x3a,
	0xac, 0x3d, 0x3d, 0x8a, 0xae, 0x1d, 0x3a, 0x72, 0x3b, 0x78, 0xb5, 0xba, 0xbf, 0xfb, 0x3d, 0x84,
	0xeb, 0x07, 0x79, 0xa9, 0x34, 0x2b, 0x0e, 0x35, 0xd1, 0xa5, 0x42, 0x29, 0xdc, 0xa4, 0x05, 0xe1,
	0x82, 0x8b, 0xec, 0xc8, 0x9b, 0xae, 0x65, 0xf7, 0xfe, 0x21, 0xbb, 0x02, 0x47, 0x7d, 0x4f, 0xfa,
	0xcd, 0xfd, 0xb4, 0x0d, 0xba, 0x9a, 0x56, 0xe8, 0x05, 0x5c, 0xa7, 0x72, 0x2e, 0x96, 0x03, 0xc2,
	0xff, 0xdc, 0xd2, 0x09, 0xdc, 0xbb, 0x6c, 0xae, 0xe1, 0xce, 0x47, 0x00, 0x1f, 0x5c, 0x19, 0x84,
	0x9e, 0xc2, 0x90, 0xd3, 0x36, 0xb0, 0x77, 0xbc, 0xae, 0x02, 0xfd, 0x3f, 0x12, 0x26, 0xfd, 0x61,
	0xc8, 0x29, 0x1a, 0xc0, 0x35, 0x65, 0x0d, 0x2f, 0x26, 0x3f, 0xf1, 0xcc, 0xf2, 0x63, 0x49, 0xc4,
	0x8c, 0x15, 0x8a, 0xbd, 0x1e, 0x8f, 0xeb, 0xed, 0xbc, 0x99, 0x05, 0xda, 0x4b, 0x4e, 0x2b, 0x1c,
	0xfc, 0xac, 0x70, 0xf0, 0xab, 0xc2, 0xe0, 0xa2, 0xc2, 0xe0, 0x4f, 0x85, 0xc1, 0x07, 0x83, 0xc1,
	0x57, 0x83, 0xc1, 0x37, 0x83, 0xc1, 0x0f, 0x83, 0xc1, 0x89, 0xc1, 0xe0, 0xd4, 0x60, 0xf0, 0xdb,
	0xe0, 0xe0, 0xc2, 0x60, 0xf0, 0xe9, 0x1c, 0x07, 0x5f, 0xce, 0x31, 0x78, 0xd7, 0x6a, 0x9c, 0xf1,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xff, 0xb9, 0xd9, 0x39, 0x03, 0x00, 0x00,
}
