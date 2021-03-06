// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"google.golang.org/grpc/status"

	app "github.com/ibm/the-mesh-for-data/manager/apis/app/v1alpha1"
	"github.com/ibm/the-mesh-for-data/manager/controllers/app/modules"
	"github.com/ibm/the-mesh-for-data/manager/controllers/utils"
	pb "github.com/ibm/the-mesh-for-data/pkg/connectors/protobuf"
	pc "github.com/ibm/the-mesh-for-data/pkg/policy-compiler/policy-compiler"
)

// ConstructApplicationContext constructs ApplicationContext structure to send to Policy Compiler
func ConstructApplicationContext(datasetID string, input *app.M4DApplication, operationType pb.AccessOperation_AccessType) *pb.ApplicationContext {
	return &pb.ApplicationContext{
		AppInfo: &pb.ApplicationDetails{
			Purpose:             input.Spec.AppInfo.Purpose,
			ProcessingGeography: input.Spec.AppInfo.ProcessingGeography,
			Role:                string(input.Spec.AppInfo.Role),
		},
		AppId: utils.CreateAppIdentifier(input),
		Datasets: []*pb.DatasetContext{{
			Dataset: &pb.DatasetIdentifier{
				DatasetId: datasetID,
			},
			Operation: &pb.AccessOperation{
				Type:        operationType,
				Destination: input.Spec.AppInfo.ProcessingGeography,
			},
		}},
	}
}

// LookupPolicyDecisions provides a list of governance actions for the given dataset and the given operation
func LookupPolicyDecisions(datasetID string, policyCompiler pc.IPolicyCompiler, req *modules.DataInfo, input *app.M4DApplication, op pb.AccessOperation_AccessType) {
	// call external policy manager to get governance instructions for this operation
	appContext := ConstructApplicationContext(datasetID, input, op)
	var flow app.ModuleFlow
	switch op {
	case pb.AccessOperation_READ:
		flow = app.Read
	case pb.AccessOperation_COPY:
		flow = app.Copy
	case pb.AccessOperation_WRITE:
		flow = app.Write
	}

	if flow == app.Copy {
		// Special use-case: in case of different geographies and required transformations, require a copy that will handle the transformations
		// TODO: get this information from policy compiler
		if req.DataDetails.Geo != input.Spec.AppInfo.ProcessingGeography && len(req.Actions[app.Read].EnforcementActions) > 0 && req.Actions[app.Read].Allowed {
			req.Actions[flow] = modules.Transformations{
				Allowed:            true,
				Required:           true,
				EnforcementActions: req.Actions[app.Read].EnforcementActions,
			}
			req.Actions[app.Read] = modules.Transformations{
				Allowed:            true,
				EnforcementActions: make([]pb.EnforcementAction, 0),
			}
			return
		}
	}

	pcresponse, err := policyCompiler.GetPoliciesDecisions(appContext)
	if err != nil {
		errStatus, _ := status.FromError(err)
		req.Actions[flow] = modules.Transformations{
			Allowed:            false,
			Message:            errStatus.Message(),
			Reason:             utils.DetermineCause(err, "PolicyManagerService"),
			EnforcementActions: make([]pb.EnforcementAction, 0),
		}
	} else {
		for _, datasetDecision := range pcresponse.GetDatasetDecisions() {
			if datasetDecision.GetDataset().GetDatasetId() != datasetID {
				continue // not our data set
			}
			var actions []pb.EnforcementAction
			operationDecisions := datasetDecision.GetDecisions()
			for _, operationDecision := range operationDecisions {
				enforcementActions := operationDecision.GetEnforcementActions()
				for _, action := range enforcementActions {
					if utils.IsDenied(action.GetName()) {
						var msg, reason string
						if operationDecision.Operation.Type == pb.AccessOperation_READ {
							msg = "Governance policies forbid access to the data"
							reason = "AccessDenied"
						} else {
							msg = "Copy of the data is required but can not be done according to the governance policies."
							reason = "CopyDenied"
						}
						req.Actions[flow] = modules.Transformations{
							Allowed:            false,
							Message:            msg,
							Reason:             reason,
							EnforcementActions: make([]pb.EnforcementAction, 0),
						}
						return
					}
					// Check if this is a real action (i.e. not Allow)
					if utils.IsAction(action.GetName()) {
						actions = append(actions, *action.DeepCopy())
					}
				}
			}
			req.Actions[flow] = modules.Transformations{
				Allowed:            true,
				EnforcementActions: actions,
			}
		}
	}
}
