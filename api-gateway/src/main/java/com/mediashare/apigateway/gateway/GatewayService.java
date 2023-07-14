package com.mediashare.apigateway.gateway;

import com.mediashare.apigateway.externalService.ExternalService;

public interface GatewayService {
    String getForwardingAddress(String service);
    void registerService(ExternalService externalService);
    String encodeAddress(String address);
}
