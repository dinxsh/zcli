package vpn

import "context"

func (h *Handler) vpnStatusStatus(ctx context.Context) {
	h.lock.Lock()
	defer h.lock.Unlock()

	data := h.storage.Data()

	if data.ProjectId != "" {
		if !h.isVpnAlive(data.ServerIp) {
			err := h.startVpn(
				ctx,
				data.GrpcApiAddress,
				data.GrpcVpnAddress,
				data.Token,
				data.ProjectId,
				data.UserId,
				data.Mtu,
				data.CaCertificate,
			)
			if err != nil {
				h.logger.Error(err)
			}
		}
	}
}
