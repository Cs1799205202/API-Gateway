// Code generated by thriftgo (0.2.12). DO NOT EDIT.

package upper

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type NormalRequest struct {
	Message string `thrift:"message,1" frugal:"1,default,string" json:"message"`
}

func NewNormalRequest() *NormalRequest {
	return &NormalRequest{}
}

func (p *NormalRequest) InitDefault() {
	*p = NormalRequest{}
}

func (p *NormalRequest) GetMessage() (v string) {
	return p.Message
}
func (p *NormalRequest) SetMessage(val string) {
	p.Message = val
}

var fieldIDToName_NormalRequest = map[int16]string{
	1: "message",
}

func (p *NormalRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_NormalRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *NormalRequest) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Message = v
	}
	return nil
}

func (p *NormalRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("NormalRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *NormalRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Message); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *NormalRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("NormalRequest(%+v)", *p)
}

func (p *NormalRequest) DeepEqual(ano *NormalRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Message) {
		return false
	}
	return true
}

func (p *NormalRequest) Field1DeepEqual(src string) bool {

	if strings.Compare(p.Message, src) != 0 {
		return false
	}
	return true
}

type UpperResponse struct {
	Result_ string `thrift:"result,1" frugal:"1,default,string" json:"result"`
}

func NewUpperResponse() *UpperResponse {
	return &UpperResponse{}
}

func (p *UpperResponse) InitDefault() {
	*p = UpperResponse{}
}

func (p *UpperResponse) GetResult_() (v string) {
	return p.Result_
}
func (p *UpperResponse) SetResult_(val string) {
	p.Result_ = val
}

var fieldIDToName_UpperResponse = map[int16]string{
	1: "result",
}

func (p *UpperResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UpperResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UpperResponse) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Result_ = v
	}
	return nil
}

func (p *UpperResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("UpperResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UpperResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("result", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Result_); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *UpperResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UpperResponse(%+v)", *p)
}

func (p *UpperResponse) DeepEqual(ano *UpperResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Result_) {
		return false
	}
	return true
}

func (p *UpperResponse) Field1DeepEqual(src string) bool {

	if strings.Compare(p.Result_, src) != 0 {
		return false
	}
	return true
}

type UpperService interface {
	Toupper(ctx context.Context, req *NormalRequest) (r *UpperResponse, err error)
}

type UpperServiceClient struct {
	c thrift.TClient
}

func NewUpperServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *UpperServiceClient {
	return &UpperServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewUpperServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *UpperServiceClient {
	return &UpperServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewUpperServiceClient(c thrift.TClient) *UpperServiceClient {
	return &UpperServiceClient{
		c: c,
	}
}

func (p *UpperServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *UpperServiceClient) Toupper(ctx context.Context, req *NormalRequest) (r *UpperResponse, err error) {
	var _args UpperServiceToupperArgs
	_args.Req = req
	var _result UpperServiceToupperResult
	if err = p.Client_().Call(ctx, "toupper", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type UpperServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      UpperService
}

func (p *UpperServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *UpperServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *UpperServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewUpperServiceProcessor(handler UpperService) *UpperServiceProcessor {
	self := &UpperServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("toupper", &upperServiceProcessorToupper{handler: handler})
	return self
}
func (p *UpperServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type upperServiceProcessorToupper struct {
	handler UpperService
}

func (p *upperServiceProcessorToupper) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := UpperServiceToupperArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("toupper", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := UpperServiceToupperResult{}
	var retval *UpperResponse
	if retval, err2 = p.handler.Toupper(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing toupper: "+err2.Error())
		oprot.WriteMessageBegin("toupper", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("toupper", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type UpperServiceToupperArgs struct {
	Req *NormalRequest `thrift:"req,1" frugal:"1,default,NormalRequest" json:"req"`
}

func NewUpperServiceToupperArgs() *UpperServiceToupperArgs {
	return &UpperServiceToupperArgs{}
}

func (p *UpperServiceToupperArgs) InitDefault() {
	*p = UpperServiceToupperArgs{}
}

var UpperServiceToupperArgs_Req_DEFAULT *NormalRequest

func (p *UpperServiceToupperArgs) GetReq() (v *NormalRequest) {
	if !p.IsSetReq() {
		return UpperServiceToupperArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *UpperServiceToupperArgs) SetReq(val *NormalRequest) {
	p.Req = val
}

var fieldIDToName_UpperServiceToupperArgs = map[int16]string{
	1: "req",
}

func (p *UpperServiceToupperArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpperServiceToupperArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UpperServiceToupperArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UpperServiceToupperArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Req = NewNormalRequest()
	if err := p.Req.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *UpperServiceToupperArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("toupper_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UpperServiceToupperArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *UpperServiceToupperArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UpperServiceToupperArgs(%+v)", *p)
}

func (p *UpperServiceToupperArgs) DeepEqual(ano *UpperServiceToupperArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *UpperServiceToupperArgs) Field1DeepEqual(src *NormalRequest) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type UpperServiceToupperResult struct {
	Success *UpperResponse `thrift:"success,0,optional" frugal:"0,optional,UpperResponse" json:"success,omitempty"`
}

func NewUpperServiceToupperResult() *UpperServiceToupperResult {
	return &UpperServiceToupperResult{}
}

func (p *UpperServiceToupperResult) InitDefault() {
	*p = UpperServiceToupperResult{}
}

var UpperServiceToupperResult_Success_DEFAULT *UpperResponse

func (p *UpperServiceToupperResult) GetSuccess() (v *UpperResponse) {
	if !p.IsSetSuccess() {
		return UpperServiceToupperResult_Success_DEFAULT
	}
	return p.Success
}
func (p *UpperServiceToupperResult) SetSuccess(x interface{}) {
	p.Success = x.(*UpperResponse)
}

var fieldIDToName_UpperServiceToupperResult = map[int16]string{
	0: "success",
}

func (p *UpperServiceToupperResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpperServiceToupperResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UpperServiceToupperResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UpperServiceToupperResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewUpperResponse()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *UpperServiceToupperResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("toupper_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UpperServiceToupperResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *UpperServiceToupperResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UpperServiceToupperResult(%+v)", *p)
}

func (p *UpperServiceToupperResult) DeepEqual(ano *UpperServiceToupperResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *UpperServiceToupperResult) Field0DeepEqual(src *UpperResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}
