package daemon

import (
	"fmt"

	"github.com/ubuntu/zsys"
	"github.com/ubuntu/zsys/internal/authorizer"
	"github.com/ubuntu/zsys/internal/i18n"
	"github.com/ubuntu/zsys/internal/log"
)

// MachineShow returns information about the machine id passed in argument
func (s *Server) MachineShow(req *zsys.MachineShowRequest, stream zsys.Zsys_MachineShowServer) error {
	if err := s.authorizer.IsAllowedFromContext(stream.Context(), authorizer.ActionAlwaysAllowed); err != nil {
		return err
	}

	fullInfo := req.GetFull()

	m, err := s.Machines.GetMachine(req.GetMachineId())
	if err != nil {
		return err
	}

	log.Infof(stream.Context(), i18n.G("Retrieving information for machine %s"), m.ID)

	machineInfo, err := m.Info(fullInfo)
	if err != nil {
		return fmt.Errorf(i18n.G("couldn't fetch matching information: %v"), err)
	}

	stream.Send(&zsys.MachineShowResponse{
		Reply: &zsys.MachineShowResponse_MachineInfo{
			MachineInfo: machineInfo,
		},
	})

	return nil

}
