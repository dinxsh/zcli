package repository

import (
	"context"

	"github.com/zeropsio/zcli/src/entity"
	"github.com/zeropsio/zcli/src/errorsx"
	"github.com/zeropsio/zcli/src/zeropsRestApiClient"
	"github.com/zeropsio/zerops-go/dto/input/path"
	"github.com/zeropsio/zerops-go/dto/output"
	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

func GetServiceByIdOrName(
	ctx context.Context,
	restApiClient *zeropsRestApiClient.Handler,
	projectId uuid.ProjectId,
	serviceIdOrName string,
) (*entity.Service, error) {
	service, err := GetServiceById(ctx, restApiClient, uuid.ServiceStackId(serviceIdOrName))
	if err != nil {
		if errorsx.IsUserError(err) {
			service, err = GetServiceByName(ctx, restApiClient, projectId, types.String(serviceIdOrName))
			if err != nil {
				return nil, err
			}
		}
	}

	return service, err
}

func GetServiceById(
	ctx context.Context,
	restApiClient *zeropsRestApiClient.Handler,
	serviceId uuid.ServiceStackId,
) (*entity.Service, error) {
	serviceResponse, err := restApiClient.GetServiceStack(ctx, path.ServiceStackId{Id: serviceId})
	if err != nil {
		return nil, err
	}

	serviceOutput, err := serviceResponse.Output()
	if err != nil {
		return nil, err
	}

	service := serviceFromApiOutput(serviceOutput)
	return &service, nil
}

func GetServiceByName(
	ctx context.Context,
	restApiClient *zeropsRestApiClient.Handler,
	projectId uuid.ProjectId,
	serviceName types.String,
) (*entity.Service, error) {
	serviceResponse, err := restApiClient.GetServiceStackByName(ctx, path.GetServiceStackByName{
		ProjectId: projectId,
		Name:      serviceName,
	})
	if err != nil {
		return nil, err
	}

	serviceOutput, err := serviceResponse.Output()
	if err != nil {
		return nil, err
	}

	service := serviceFromApiOutput(serviceOutput)
	return &service, nil
}

func GetNonSystemServicesByProject(
	ctx context.Context,
	restApiClient *zeropsRestApiClient.Handler,
	project entity.Project,
) ([]entity.Service, error) {
	servicesResponse, err := restApiClient.GetServiceStackByProject(ctx, project.ID, project.ClientId)
	if err != nil {
		return nil, err
	}

	servicesOutput, err := servicesResponse.Output()
	if err != nil {
		return nil, err
	}

	services := make([]entity.Service, 0, len(servicesOutput.Items))
	for _, service := range servicesOutput.Items {
		if !service.IsSystem {
			services = append(services, serviceFromEsSearch(service))
		}
	}

	return services, nil
}

func serviceFromEsSearch(esServiceStack zeropsRestApiClient.EsServiceStack) entity.Service {
	return entity.Service{
		ID:                  esServiceStack.Id,
		ClientId:            esServiceStack.ClientId,
		Name:                esServiceStack.Name,
		Status:              esServiceStack.Status,
		ServiceTypeId:       esServiceStack.ServiceStackTypeId,
		ServiceTypeCategory: esServiceStack.ServiceStackTypeInfo.ServiceStackTypeCategory,
	}
}

func serviceFromApiOutput(service output.ServiceStack) entity.Service {
	return entity.Service{
		ID:                          service.Id,
		ClientId:                    service.Project.ClientId,
		Name:                        service.Name,
		Status:                      service.Status,
		ServiceTypeId:               service.ServiceStackTypeId,
		ServiceTypeCategory:         service.ServiceStackTypeInfo.ServiceStackTypeCategory,
		ServiceStackTypeVersionName: service.ServiceStackTypeInfo.ServiceStackTypeVersionName,
	}
}
