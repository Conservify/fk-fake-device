package main

import (
	"context"

	"github.com/golang/protobuf/proto"

	pb "github.com/fieldkit/app-protocol"
)

func handleQueryCapabilities(ctx context.Context, rw replyWriter) (err error) {
	reply := &pb.WireMessageReply{
		Type: pb.ReplyType_REPLY_CAPABILITIES,
		Capabilities: &pb.Capabilities{
			Version: 0x1,
			Name:    "FieldKit Station",
			Sensors: []*pb.SensorCapabilities{
				&pb.SensorCapabilities{
					Id:            0,
					Name:          "Conductivity",
					UnitOfMeasure: "µS/cm",
					Frequency:     60,
				},
				&pb.SensorCapabilities{
					Id:            1,
					Name:          "Temperature",
					UnitOfMeasure: "C",
					Frequency:     60,
				},
				&pb.SensorCapabilities{
					Id:            2,
					Name:          "Depth",
					UnitOfMeasure: "m",
					Frequency:     60,
				},
				&pb.SensorCapabilities{
					Id:            3,
					Name:          "Hydrophone",
					UnitOfMeasure: "",
					Frequency:     0,
				},
			},
		},
	}

	_, err = rw.WriteReply(reply)

	return
}

func handleQueryStatus(ctx context.Context, rw replyWriter) (err error) {
	reply := &pb.WireMessageReply{
		Type:   pb.ReplyType_REPLY_STATUS,
		Status: &pb.DeviceStatus{},
	}

	_, err = rw.WriteReply(reply)

	return
}

func handleQueryFiles(ctx context.Context, rw replyWriter) (err error) {
	reply := &pb.WireMessageReply{
		Type: pb.ReplyType_REPLY_FILES,
		Files: &pb.Files{
			Files: []*pb.File{
				&pb.File{},
				&pb.File{},
				&pb.File{},
				&pb.File{},
			},
		},
	}

	_, err = rw.WriteReply(reply)

	return
}

func handleDownloadFile(ctx context.Context, rw replyWriter) (err error) {
	size := 0
	required := 10 * 1024 * 1024
	body := proto.NewBuffer(make([]byte, 0))

	for size < required {
		reply := &pb.WireMessageReply{
			Type: pb.ReplyType_REPLY_DOWNLOAD_FILE,
			FileData: &pb.FileData{
				Size: uint32(size),
			},
		}

		body.EncodeMessage(reply)

		size = len(body.Bytes())
	}

	rw.Prepare(size)

	reply := &pb.WireMessageReply{
		Type: pb.ReplyType_REPLY_DOWNLOAD_FILE,
		FileData: &pb.FileData{
			Size: uint32(len(body.Bytes())),
		},
	}

	rw.WriteReply(reply)
	rw.WriteBytes(body.Bytes())

	return
}
