package taskrepo

// import (
// 	"context"

// 	"github.com/Muha113/task-service/pkg/apiproto"
// 	"google.golang.org/grpc"
// )

// type TaskRepoClient struct{}

// func (trc *TaskRepoClient) AddNewEmployee(ctx context.Context, in *apiproto.Employee, opts ...grpc.CallOption) (*apiproto.ResponseEmployee, error) {
// 	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
// 	if err != nil {
// 		return nil, err
// 	}

// }

// func (trc *TaskRepoClient) UpdateEmployee(ctx context.Context, in *apiproto.Employee, opts ...grpc.CallOption) (*apiproto.ResponseEmployee, error)

// func (trc *TaskRepoClient) FindEmployee(ctx context.Context, in *apiproto.Employee, opts ...grpc.CallOption) (*apiproto.ResponseEmployee, error)

// func (trc *TaskRepoClient) UpdateEmployeeWithFormData(ctx context.Context, in *apiproto.Employee, opts ...grpc.CallOption) (*apiproto.ResponseEmployee, error)

// func (trc *TaskRepoClient) DeleteEmployee(ctx context.Context, in *apiproto.Employee, opts ...grpc.CallOption) (*apiproto.ResponseEmployee, error)

// func (trc *TaskRepoClient) AddNewCompany(ctx context.Context, in *apiproto.Company, opts ...grpc.CallOption) (*apiproto.ResponseCompany, error)

// func (trc *TaskRepoClient) UpdateCompany(ctx context.Context, in *apiproto.Company, opts ...grpc.CallOption) (*apiproto.ResponseCompany, error)

// func (trc *TaskRepoClient) FindCompany(ctx context.Context, in *apiproto.Company, opts ...grpc.CallOption) (*apiproto.ResponseCompany, error)

// func (trc *TaskRepoClient) UpdateCompanyWithFormData(ctx context.Context, in *apiproto.Company, opts ...grpc.CallOption) (*apiproto.ResponseCompany, error)

// func (trc *TaskRepoClient) DeleteCompany(ctx context.Context, in *apiproto.Company, opts ...grpc.CallOption) (*apiproto.ResponseCompany, error)

// func (trc *TaskRepoClient) GetCompanyEmployeesList(ctx context.Context, in *apiproto.Company, opts ...grpc.CallOption) (*apiproto.ResponseEmployee, error)
