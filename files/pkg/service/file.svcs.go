package service

import (
	"common/grpcs/file/pb"
	"context"
	"files-microservice/pkg/repo"
	"io"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FileSvcs struct {
	pb.UnimplementedFileServiceServer
	filerepo *repo.FileRepo
}

func NewFileSvcs() *FileSvcs {
	return &FileSvcs{}
}

func (f *FileSvcs) Upload(stream pb.FileService_UploadServer) error {

	file := NewFile()

	defer file.Close()

	for {
		req, err := stream.Recv()
		if file.FilePath == "" {
			file.SetFile(req.GetFileName(), "client_files")
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}

		chunk := req.GetChunk()
		if err := file.Write(chunk); err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}

	}

	_file := &pb.File{
		Path:         file.FilePath,
		OriginalName: file.fileName,
	}

	if err := f.filerepo.SaveFile(stream.Context(), _file); err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	return stream.SendAndClose(&pb.FileUploadRes{File: _file})
}

func (f *FileSvcs) GetById(ctx context.Context, req *pb.GetByIdReq) (*pb.GetByIdRes, error) {
	id := req.GetId()

	result, err := f.filerepo.GetById(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.GetByIdRes{
		File: result,
	}, nil
}

func (f *FileSvcs) ListFiles(req *pb.ListFilesReq, stream pb.FileService_ListFilesServer) error {
	files, err := f.filerepo.Listfiles(stream.Context())
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, file := range files {
		if err := stream.Send(&pb.ListFilesRes{
			File: file,
		}); err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	}

	return nil
}
