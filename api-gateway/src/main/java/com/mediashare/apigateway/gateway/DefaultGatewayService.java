package com.mediashare.apigateway.gateway;

import com.mediashare.apigateway.externalService.ExternalService;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.Map;

@Service
public class DefaultGatewayService implements GatewayService {
    private Map<String, ExternalService> serviceCache;

    public DefaultGatewayService() {
        this.serviceCache = new HashMap<>();
//        Load forwarding addresses into memory
    }

    public String getForwardingAddress(String service) {
        return "";
    }

    public void registerService(ExternalService externalService) {}

    public String encodeAddress(String address) {
//        URLEncoder.encode(address, StandardCharsets.UTF_8.toString());
        return "";
    }
}
