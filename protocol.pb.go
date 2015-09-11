// Code generated by protoc-gen-gogo.
// source: protocol.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	protocol.proto

It has these top-level messages:
	Request
	ReqBatch
	Response
	ResBatch
	InfoReq
	InfoRes
	DBInfo
	OpenReq
	OpenRes
	EditReq
	EditRes
	PutReq
	PutRes
	IncReq
	IncRes
	GetReq
	GetRes
	ResSeries
	ResPoint
	MetricsReq
	MetricsRes
*/
package main

import proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Request struct {
	InfoReq *InfoReq `protobuf:"bytes,1,opt,name=infoReq" json:"infoReq,omitempty"`
	OpenReq *OpenReq `protobuf:"bytes,2,opt,name=openReq" json:"openReq,omitempty"`
	EditReq *EditReq `protobuf:"bytes,3,opt,name=editReq" json:"editReq,omitempty"`
	PutReq  *PutReq  `protobuf:"bytes,4,opt,name=putReq" json:"putReq,omitempty"`
	IncReq  *IncReq  `protobuf:"bytes,5,opt,name=incReq" json:"incReq,omitempty"`
	GetReq  *GetReq  `protobuf:"bytes,6,opt,name=getReq" json:"getReq,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetInfoReq() *InfoReq {
	if m != nil {
		return m.InfoReq
	}
	return nil
}

func (m *Request) GetOpenReq() *OpenReq {
	if m != nil {
		return m.OpenReq
	}
	return nil
}

func (m *Request) GetEditReq() *EditReq {
	if m != nil {
		return m.EditReq
	}
	return nil
}

func (m *Request) GetPutReq() *PutReq {
	if m != nil {
		return m.PutReq
	}
	return nil
}

func (m *Request) GetIncReq() *IncReq {
	if m != nil {
		return m.IncReq
	}
	return nil
}

func (m *Request) GetGetReq() *GetReq {
	if m != nil {
		return m.GetReq
	}
	return nil
}

type ReqBatch struct {
	Batch []*Request `protobuf:"bytes,1,rep,name=batch" json:"batch,omitempty"`
}

func (m *ReqBatch) Reset()         { *m = ReqBatch{} }
func (m *ReqBatch) String() string { return proto.CompactTextString(m) }
func (*ReqBatch) ProtoMessage()    {}

func (m *ReqBatch) GetBatch() []*Request {
	if m != nil {
		return m.Batch
	}
	return nil
}

type Response struct {
	InfoRes *InfoRes `protobuf:"bytes,1,opt,name=infoRes" json:"infoRes,omitempty"`
	OpenRes *OpenRes `protobuf:"bytes,2,opt,name=openRes" json:"openRes,omitempty"`
	EditRes *EditRes `protobuf:"bytes,3,opt,name=editRes" json:"editRes,omitempty"`
	PutRes  *PutRes  `protobuf:"bytes,4,opt,name=putRes" json:"putRes,omitempty"`
	IncRes  *IncRes  `protobuf:"bytes,5,opt,name=incRes" json:"incRes,omitempty"`
	GetRes  *GetRes  `protobuf:"bytes,6,opt,name=getRes" json:"getRes,omitempty"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetInfoRes() *InfoRes {
	if m != nil {
		return m.InfoRes
	}
	return nil
}

func (m *Response) GetOpenRes() *OpenRes {
	if m != nil {
		return m.OpenRes
	}
	return nil
}

func (m *Response) GetEditRes() *EditRes {
	if m != nil {
		return m.EditRes
	}
	return nil
}

func (m *Response) GetPutRes() *PutRes {
	if m != nil {
		return m.PutRes
	}
	return nil
}

func (m *Response) GetIncRes() *IncRes {
	if m != nil {
		return m.IncRes
	}
	return nil
}

func (m *Response) GetGetRes() *GetRes {
	if m != nil {
		return m.GetRes
	}
	return nil
}

type ResBatch struct {
	Batch []*Response `protobuf:"bytes,1,rep,name=batch" json:"batch,omitempty"`
}

func (m *ResBatch) Reset()         { *m = ResBatch{} }
func (m *ResBatch) String() string { return proto.CompactTextString(m) }
func (*ResBatch) ProtoMessage()    {}

func (m *ResBatch) GetBatch() []*Response {
	if m != nil {
		return m.Batch
	}
	return nil
}

type InfoReq struct {
}

func (m *InfoReq) Reset()         { *m = InfoReq{} }
func (m *InfoReq) String() string { return proto.CompactTextString(m) }
func (*InfoReq) ProtoMessage()    {}

type InfoRes struct {
	Databases []*DBInfo `protobuf:"bytes,1,rep,name=databases" json:"databases,omitempty"`
}

func (m *InfoRes) Reset()         { *m = InfoRes{} }
func (m *InfoRes) String() string { return proto.CompactTextString(m) }
func (*InfoRes) ProtoMessage()    {}

func (m *InfoRes) GetDatabases() []*DBInfo {
	if m != nil {
		return m.Databases
	}
	return nil
}

type DBInfo struct {
	Database   string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Resolution uint32 `protobuf:"varint,2,opt,name=resolution,proto3" json:"resolution,omitempty"`
	Retention  uint32 `protobuf:"varint,3,opt,name=retention,proto3" json:"retention,omitempty"`
}

func (m *DBInfo) Reset()         { *m = DBInfo{} }
func (m *DBInfo) String() string { return proto.CompactTextString(m) }
func (*DBInfo) ProtoMessage()    {}

type OpenReq struct {
	Database    string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Resolution  uint32 `protobuf:"varint,2,opt,name=resolution,proto3" json:"resolution,omitempty"`
	Retention   uint32 `protobuf:"varint,3,opt,name=retention,proto3" json:"retention,omitempty"`
	EpochTime   uint32 `protobuf:"varint,4,opt,name=epochTime,proto3" json:"epochTime,omitempty"`
	MaxROEpochs uint32 `protobuf:"varint,5,opt,name=maxROEpochs,proto3" json:"maxROEpochs,omitempty"`
	MaxRWEpochs uint32 `protobuf:"varint,6,opt,name=maxRWEpochs,proto3" json:"maxRWEpochs,omitempty"`
}

func (m *OpenReq) Reset()         { *m = OpenReq{} }
func (m *OpenReq) String() string { return proto.CompactTextString(m) }
func (*OpenReq) ProtoMessage()    {}

type OpenRes struct {
}

func (m *OpenRes) Reset()         { *m = OpenRes{} }
func (m *OpenRes) String() string { return proto.CompactTextString(m) }
func (*OpenRes) ProtoMessage()    {}

type EditReq struct {
	Database    string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Retention   uint32 `protobuf:"varint,2,opt,name=retention,proto3" json:"retention,omitempty"`
	MaxROEpochs uint32 `protobuf:"varint,3,opt,name=maxROEpochs,proto3" json:"maxROEpochs,omitempty"`
	MaxRWEpochs uint32 `protobuf:"varint,4,opt,name=maxRWEpochs,proto3" json:"maxRWEpochs,omitempty"`
}

func (m *EditReq) Reset()         { *m = EditReq{} }
func (m *EditReq) String() string { return proto.CompactTextString(m) }
func (*EditReq) ProtoMessage()    {}

type EditRes struct {
}

func (m *EditRes) Reset()         { *m = EditRes{} }
func (m *EditRes) String() string { return proto.CompactTextString(m) }
func (*EditRes) ProtoMessage()    {}

type PutReq struct {
	Database  string   `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Timestamp uint32   `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Value     float64  `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	Count     uint32   `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	Fields    []string `protobuf:"bytes,5,rep,name=fields" json:"fields,omitempty"`
}

func (m *PutReq) Reset()         { *m = PutReq{} }
func (m *PutReq) String() string { return proto.CompactTextString(m) }
func (*PutReq) ProtoMessage()    {}

type PutRes struct {
}

func (m *PutRes) Reset()         { *m = PutRes{} }
func (m *PutRes) String() string { return proto.CompactTextString(m) }
func (*PutRes) ProtoMessage()    {}

type IncReq struct {
	Database  string   `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Timestamp uint32   `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Value     float64  `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	Count     uint32   `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	Fields    []string `protobuf:"bytes,5,rep,name=fields" json:"fields,omitempty"`
}

func (m *IncReq) Reset()         { *m = IncReq{} }
func (m *IncReq) String() string { return proto.CompactTextString(m) }
func (*IncReq) ProtoMessage()    {}

type IncRes struct {
}

func (m *IncRes) Reset()         { *m = IncRes{} }
func (m *IncRes) String() string { return proto.CompactTextString(m) }
func (*IncRes) ProtoMessage()    {}

type GetReq struct {
	Database   string   `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	StartTime  uint32   `protobuf:"varint,2,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime    uint32   `protobuf:"varint,3,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Fields     []string `protobuf:"bytes,4,rep,name=fields" json:"fields,omitempty"`
	GroupBy    []bool   `protobuf:"varint,5,rep,packed,name=groupBy" json:"groupBy,omitempty"`
	Resolution uint32   `protobuf:"varint,6,opt,name=resolution,proto3" json:"resolution,omitempty"`
}

func (m *GetReq) Reset()         { *m = GetReq{} }
func (m *GetReq) String() string { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()    {}

type GetRes struct {
	Groups []*ResSeries `protobuf:"bytes,2,rep,name=groups" json:"groups,omitempty"`
}

func (m *GetRes) Reset()         { *m = GetRes{} }
func (m *GetRes) String() string { return proto.CompactTextString(m) }
func (*GetRes) ProtoMessage()    {}

func (m *GetRes) GetGroups() []*ResSeries {
	if m != nil {
		return m.Groups
	}
	return nil
}

type ResSeries struct {
	Fields []string    `protobuf:"bytes,1,rep,name=fields" json:"fields,omitempty"`
	Points []*ResPoint `protobuf:"bytes,2,rep,name=points" json:"points,omitempty"`
}

func (m *ResSeries) Reset()         { *m = ResSeries{} }
func (m *ResSeries) String() string { return proto.CompactTextString(m) }
func (*ResSeries) ProtoMessage()    {}

func (m *ResSeries) GetPoints() []*ResPoint {
	if m != nil {
		return m.Points
	}
	return nil
}

type ResPoint struct {
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	Count uint32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *ResPoint) Reset()         { *m = ResPoint{} }
func (m *ResPoint) String() string { return proto.CompactTextString(m) }
func (*ResPoint) ProtoMessage()    {}

type MetricsReq struct {
}

func (m *MetricsReq) Reset()         { *m = MetricsReq{} }
func (m *MetricsReq) String() string { return proto.CompactTextString(m) }
func (*MetricsReq) ProtoMessage()    {}

type MetricsRes struct {
}

func (m *MetricsRes) Reset()         { *m = MetricsRes{} }
func (m *MetricsRes) String() string { return proto.CompactTextString(m) }
func (*MetricsRes) ProtoMessage()    {}
