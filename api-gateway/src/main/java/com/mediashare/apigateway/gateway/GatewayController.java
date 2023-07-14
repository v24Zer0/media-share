package com.mediashare.apigateway.gateway;

import jakarta.servlet.http.HttpServletRequest;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpEntity;
import org.springframework.http.RequestEntity;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

@RestController
@RequestMapping("/")
public class GatewayController {
    private final GatewayService gatewayService;

    @Autowired
    public GatewayController(GatewayService gatewayService) {
        this.gatewayService = gatewayService;
    }

    @PostMapping("register")
    public String registerService() {
//        this.gatewayService.registerService();
        return "register service";
    }

    @RequestMapping("/image/**")
    public ResponseEntity<byte[]> handleImageRequest(HttpServletRequest request) {
        String path = request.getServletPath();
        String params = request.getQueryString();

        RestTemplate template = new RestTemplate();

//        RequestEntity<Void> ent = RequestEntity.post().headers().build();
        return template.getForEntity("http://localhost:8081/image/post_01", byte[].class);
    }

    @RequestMapping("**")
    public Object handleRequest(HttpServletRequest request) {
        return request.getRequestURL() + " : " + request.getQueryString();
    }
}
